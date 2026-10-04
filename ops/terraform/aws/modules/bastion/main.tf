# ==============================================================================
# AMI
#
# Amazon Linux 2023（arm64）。SSMエージェントが最初から入っているのでuser_dataは不要
# ==============================================================================
data "aws_ssm_parameter" "al2023_arm64" {
  name = "/aws/service/ami-amazon-linux-latest/al2023-ami-kernel-default-arm64"
}

# ==============================================================================
# Security Group
#
# ingressは作らない。SSMはbastionから外向き443でAWSに繋ぐ接続なので
# インバウンドの穴は不要（SSHも使わない）
# ==============================================================================
resource "aws_security_group" "bastion" {
  name   = var.name
  vpc_id = var.vpc_id

  tags = {
    Name = var.name
  }
}

# aws_security_groupをTerraformで作ると既定のegress（全許可）が削除されるので明示する
resource "aws_vpc_security_group_egress_rule" "bastion_all" {
  security_group_id = aws_security_group.bastion.id
  description       = "All outbound"

  ip_protocol = "-1"
  cidr_ipv4   = "0.0.0.0/0"
}

# Aurora側に「5432 ← bastion」を足す（段2のecs → auroraと同じ書き方）
resource "aws_vpc_security_group_ingress_rule" "aurora_from_bastion" {
  security_group_id = var.aurora_security_group_id
  description       = "PostgreSQL from bastion"

  ip_protocol                  = "tcp"
  from_port                    = 5432
  to_port                      = 5432
  referenced_security_group_id = aws_security_group.bastion.id
}

# ==============================================================================
# IAM
# ==============================================================================
data "aws_iam_policy_document" "ec2_assume" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "bastion" {
  name               = var.name
  assume_role_policy = data.aws_iam_policy_document.ec2_assume.json
}

# Session Managerでの接続に必要な権限一式
resource "aws_iam_role_policy_attachment" "ssm" {
  role       = aws_iam_role.bastion.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

resource "aws_iam_instance_profile" "bastion" {
  name = var.name
  role = aws_iam_role.bastion.name
}

# ==============================================================================
# Instance
#
# 普段は停止しておき、使うときだけ起動する（ops/db の make prd-tunnel）。
# 起動/停止の状態はTerraformの管理外なので、apply直後は起動している。
# 使い終わったら make prd-tunnel-stop で止める
# ==============================================================================
resource "aws_instance" "bastion" {
  ami           = data.aws_ssm_parameter.al2023_arm64.value
  instance_type = var.instance_type
  subnet_id     = var.subnet_ids[0]

  vpc_security_group_ids = [aws_security_group.bastion.id]
  iam_instance_profile   = aws_iam_instance_profile.bastion.name

  # NAT Gatewayが無いので、パブリックIPが無いとSSMエージェントがAWSに届かない
  associate_public_ip_address = true

  # IMDSv2必須
  metadata_options {
    http_tokens = "required"
  }

  root_block_device {
    volume_type = "gp3"
    volume_size = 8
  }

  # AMIの新版が出るたびにplanがインスタンスの作り直しを出すのを防ぐ。
  # 更新したいときは意図的にtaintする
  lifecycle {
    ignore_changes = [ami]
  }

  tags = {
    Name = var.name
  }
}
