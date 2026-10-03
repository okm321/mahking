output "cluster_endpoint" {
  description = "writerエンドポイント"
  value       = aws_rds_cluster.main.endpoint
}

output "reader_endpoint" {
  description = "readerエンドポイント"
  value       = aws_rds_cluster.main.reader_endpoint
}

output "database_name" {
  description = "データベース名"
  value       = aws_rds_cluster.main.database_name
}

output "master_username" {
  description = "マスターユーザー名"
  value       = aws_rds_cluster.main.master_username
}

output "password_ssm_parameter_arn" {
  description = "パスワードを置いたSSMパラメータのARN"
  value       = aws_ssm_parameter.master_password.arn
}

output "password_ssm_parameter_name" {
  description = "パスワードを置いたSSMパラメータ名"
  value       = aws_ssm_parameter.master_password.name
}
