# ==============================================================================
# ALB
#
# ルールはインラインのingress/egressブロックではなく個別リソースで書く
# （provider 6.xの推奨。インラインと個別リソースの併用は競合する）
# ==============================================================================
resource "aws_security_group" "alb" {
  name   = var.alb_sg_name
  vpc_id = var.vpc_id

  tags = {
    Name = var.alb_sg_name
  }
}

resource "aws_vpc_security_group_ingress_rule" "alb_https" {
  security_group_id = aws_security_group.alb.id
  description       = "HTTPS from internet"

  ip_protocol = "tcp"
  from_port   = 443
  to_port     = 443
  cidr_ipv4   = "0.0.0.0/0"
}

# aws_security_groupをTerraformで作ると既定のegress（全許可）が削除されるので明示する
resource "aws_vpc_security_group_egress_rule" "alb_all" {
  security_group_id = aws_security_group.alb.id
  description       = "All outbound"

  ip_protocol = "-1"
  cidr_ipv4   = "0.0.0.0/0"
}

# ==============================================================================
# ECS Task
# ==============================================================================
resource "aws_security_group" "ecs" {
  name   = var.ecs_sg_name
  vpc_id = var.vpc_id

  tags = {
    Name = var.ecs_sg_name
  }
}

# 送信元はALBのSGそのもの。タスクのIPが変わっても追従する
resource "aws_vpc_security_group_ingress_rule" "ecs_from_alb" {
  security_group_id = aws_security_group.ecs.id
  description       = "App port from ALB"

  ip_protocol                  = "tcp"
  from_port                    = var.app_port
  to_port                      = var.app_port
  referenced_security_group_id = aws_security_group.alb.id
}

# 既定のegressが削除されるので明示する。ECR pullとAurora接続に必要
resource "aws_vpc_security_group_egress_rule" "ecs_all" {
  security_group_id = aws_security_group.ecs.id
  description       = "All outbound"

  ip_protocol = "-1"
  cidr_ipv4   = "0.0.0.0/0"
}

# ==============================================================================
# Aurora
#
# egressは作らない（DBから外に出る通信は不要）
# ==============================================================================
resource "aws_security_group" "aurora" {
  name   = var.aurora_sg_name
  vpc_id = var.vpc_id

  tags = {
    Name = var.aurora_sg_name
  }
}

resource "aws_vpc_security_group_ingress_rule" "aurora_from_ecs" {
  security_group_id = aws_security_group.aurora.id
  description       = "PostgreSQL from ECS task"

  ip_protocol                  = "tcp"
  from_port                    = 5432
  to_port                      = 5432
  referenced_security_group_id = aws_security_group.ecs.id
}
