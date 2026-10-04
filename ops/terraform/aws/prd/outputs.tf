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

output "alb_dns_name" {
  description = "ALBのDNS名"
  value       = module.alb.alb_dns_name
}

output "ecs_cluster_name" {
  description = "ECSクラスター名"
  value       = module.ecs.cluster_name
}

output "ecs_service_name" {
  description = "ECSサービス名"
  value       = module.ecs.service_name
}

output "aurora_endpoint" {
  description = "Auroraのwriterエンドポイント"
  value       = module.aurora.cluster_endpoint
}

output "db_password_ssm_name" {
  description = "DBパスワードを置いたSSMパラメータ名"
  value       = module.aurora.password_ssm_parameter_name
}

output "bastion_instance_id" {
  description = "bastion EC2のインスタンスID"
  value       = module.bastion.instance_id
}
