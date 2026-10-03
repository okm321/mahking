# ==============================================================================
# VPC
# ==============================================================================
module "vpc" {
  source = "../modules/vpc"

  vpc_name        = local.vpc.name
  vpc_cidr        = local.vpc.cidr
  igw_name        = local.vpc.igw_name
  public_subnets  = local.public_subnets
  private_subnets = local.private_subnets
}

# ==============================================================================
# Security Group
# ==============================================================================
module "security_group" {
  source = "../modules/security_group"

  vpc_id         = module.vpc.vpc_id
  alb_sg_name    = local.security_groups.alb
  ecs_sg_name    = local.security_groups.ecs
  aurora_sg_name = local.security_groups.aurora
  app_port       = local.app_port
}

# ==============================================================================
# ECR
# ==============================================================================
module "ecr" {
  source = "../modules/ecr"

  repository_name = local.ecr_repository_name
}

# ==============================================================================
# ACM
# ==============================================================================
module "acm" {
  source = "../modules/acm"

  domain_name        = local.api_domain
  cloudflare_zone_id = local.cloudflare_zone_id
}
