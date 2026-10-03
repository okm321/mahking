output "repository_url" {
  description = "ECRリポジトリのURL（docker pushの宛先）"
  value       = aws_ecr_repository.main.repository_url
}

output "repository_arn" {
  description = "ECRリポジトリのARN"
  value       = aws_ecr_repository.main.arn
}
