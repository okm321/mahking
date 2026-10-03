output "alb_arn" {
  description = "ALBのARN"
  value       = aws_lb.main.arn
}

output "alb_dns_name" {
  description = "ALBのDNS名（CNAMEの宛先）"
  value       = aws_lb.main.dns_name
}

output "target_group_arn" {
  description = "ECSサービスが登録するターゲットグループのARN"
  value       = aws_lb_target_group.main.arn
}
