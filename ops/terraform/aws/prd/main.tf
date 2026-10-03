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
