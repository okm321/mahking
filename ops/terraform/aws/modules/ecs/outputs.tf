output "cluster_name" {
  description = "ECSクラスター名"
  value       = aws_ecs_cluster.main.name
}

output "service_name" {
  description = "ECSサービス名"
  value       = aws_ecs_service.api.name
}

output "task_definition_arn" {
  description = "タスク定義のARN（リビジョン込み）"
  value       = aws_ecs_task_definition.api.arn
}

output "log_group_name" {
  description = "コンテナログのロググループ名"
  value       = aws_cloudwatch_log_group.api.name
}
