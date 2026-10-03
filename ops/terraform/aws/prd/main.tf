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

# ==============================================================================
# ALB
# ==============================================================================
module "alb" {
  source = "../modules/alb"

  alb_name           = local.alb_name
  target_group_name  = local.target_group_name
  vpc_id             = module.vpc.vpc_id
  subnet_ids         = module.vpc.public_subnet_ids
  security_group_id  = module.security_group.alb_sg_id
  app_port           = local.app_port
  certificate_arn    = module.acm.certificate_arn
  domain_name        = local.api_domain
  cloudflare_zone_id = local.cloudflare_zone_id
}

# ==============================================================================
# ECS
# ==============================================================================
module "ecs" {
  source = "../modules/ecs"

  region             = local.region
  log_group_name     = local.log_group_name
  log_retention_days = local.log_retention_days
  cluster_name       = local.ecs_cluster_name
  task_family        = local.task_family
  task_cpu           = local.task_cpu
  task_memory        = local.task_memory
  image              = "${module.ecr.repository_url}:${local.image_tag}"
  app_port           = local.app_port
  environment        = local.container_environment
  desired_count      = local.desired_count
  subnet_ids         = module.vpc.public_subnet_ids
  security_group_id  = module.security_group.ecs_sg_id
  target_group_arn   = module.alb.target_group_arn

  # ターゲットグループのARNだけではリスナーの作成を待てない。
  # リスナーが無い状態でサービスを作るとALBへの登録に失敗するのでモジュールごと待つ
  depends_on = [module.alb]
}
