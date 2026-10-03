output "alb_sg_id" {
  description = "ALB用Security GroupのID"
  value       = aws_security_group.alb.id
}

output "ecs_sg_id" {
  description = "ECSタスク用Security GroupのID"
  value       = aws_security_group.ecs.id
}

output "aurora_sg_id" {
  description = "Aurora用Security GroupのID"
  value       = aws_security_group.aurora.id
}
