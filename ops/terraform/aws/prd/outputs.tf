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
