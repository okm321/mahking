# ==============================================================================
# Outputs
# ==============================================================================

output "vpc_id" {
  description = "VPCのID"
  value       = module.vpc.vpc_id
}

output "vpc_cidr" {
  description = "VPCのCIDRブロック"
  value       = module.vpc.vpc_cidr
}

output "public_subnet_ids" {
  description = "publicサブネットのIDリスト"
  value       = module.vpc.public_subnet_ids
}

output "private_subnet_ids" {
  description = "privateサブネットのIDリスト"
  value       = module.vpc.private_subnet_ids
}

output "alb_sg_id" {
  description = "ALB用Security GroupのID"
  value       = module.security_group.alb_sg_id
}

output "ecs_sg_id" {
  description = "ECSタスク用Security GroupのID"
  value       = module.security_group.ecs_sg_id
}

output "aurora_sg_id" {
  description = "Aurora用Security GroupのID"
  value       = module.security_group.aurora_sg_id
}

output "ecr_repository_url" {
  description = "ECRリポジトリのURL"
  value       = module.ecr.repository_url
}

output "certificate_arn" {
  description = "api.mahking.appのACM証明書ARN"
  value       = module.acm.certificate_arn
}
