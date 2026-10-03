# ==============================================================================
# DB Subnet Group
# ==============================================================================
resource "aws_db_subnet_group" "main" {
  name       = var.cluster_identifier
  subnet_ids = var.subnet_ids

  tags = {
    Name = var.cluster_identifier
  }
}

# ==============================================================================
# Master Password
#
# 公開リポジトリなのでパスワードはコードに書かずTerraformで生成する。
# DSNに記号が入ると扱いが面倒なのでspecialは無効
# ==============================================================================
resource "random_password" "master" {
  length  = 32
  special = false
}

# パラメータは既定のKMSキー（aws/ssm）で暗号化する。
# 既定キーならECSの実行ロールにkms:Decryptは不要
resource "aws_ssm_parameter" "master_password" {
  name  = var.password_ssm_name
  type  = "SecureString"
  value = random_password.master.result

  tags = {
    Name = var.password_ssm_name
  }
}

# ==============================================================================
# Cluster
#
# Serverless v2はengine_modeをprovisionedにして
# serverlessv2_scaling_configurationを付ける（engine_mode = "serverless"は v1）
# ==============================================================================
resource "aws_rds_cluster" "main" {
  cluster_identifier = var.cluster_identifier
  engine             = "aurora-postgresql"
  engine_mode        = "provisioned"
  engine_version     = var.engine_version

  database_name   = var.database_name
  master_username = var.master_username
  master_password = random_password.master.result

  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [var.security_group_id]

  # 既定のKMSキー（aws/rds）で暗号化する
  storage_encrypted = true

  backup_retention_period = 7

  # JSTの3:00-3:30（UTC指定）
  preferred_backup_window = "18:00-18:30"

  # 誤ってdestroyされないよう削除保護を有効にする。消すときはfalseにしてapplyしてから
  # 消す。最終スナップショットは取らない（消すときは意図的に消す前提）
  skip_final_snapshot = true
  deletion_protection = true

  apply_immediately = true

  # min_capacity = 0 で接続が無い間はインスタンスを自動停止する。
  # 復帰には15秒程度かかるのでアプリ側はリトライ前提
  serverlessv2_scaling_configuration {
    min_capacity             = var.min_capacity
    max_capacity             = var.max_capacity
    seconds_until_auto_pause = var.seconds_until_auto_pause
  }

  tags = {
    Name = var.cluster_identifier
  }
}

# ==============================================================================
# Instance
#
# writer1台のみ。可用性が要るようになったらreaderを足す
# ==============================================================================
resource "aws_rds_cluster_instance" "writer" {
  identifier         = "${var.cluster_identifier}-1"
  cluster_identifier = aws_rds_cluster.main.id
  instance_class     = "db.serverless"

  engine         = aws_rds_cluster.main.engine
  engine_version = aws_rds_cluster.main.engine_version

  # privateサブネットに置くのでパブリックIPは付けない
  publicly_accessible = false

  tags = {
    Name = "${var.cluster_identifier}-1"
  }
}
