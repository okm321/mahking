locals {
  # ============================================================================
  # Common
  # ============================================================================
  env         = "prd"
  aws_profile = "mahking"
  region      = "ap-northeast-1"

  # ============================================================================
  # VPC
  # ============================================================================
  vpc = {
    name     = "mahking-${local.env}-vpc"
    cidr     = "10.0.0.0/16"
    igw_name = "mahking-${local.env}-igw"
  }

  # publicサブネット（ALBとFargateタスクを配置）
  public_subnets = {
    a = {
      az   = "${local.region}a"
      cidr = "10.0.1.0/24"
      name = "mahking-${local.env}-public-a"
    }
    c = {
      az   = "${local.region}c"
      cidr = "10.0.2.0/24"
      name = "mahking-${local.env}-public-c"
    }
  }

  # privateサブネット（Auroraを配置）
  private_subnets = {
    a = {
      az   = "${local.region}a"
      cidr = "10.0.11.0/24"
      name = "mahking-${local.env}-private-a"
    }
    c = {
      az   = "${local.region}c"
      cidr = "10.0.12.0/24"
      name = "mahking-${local.env}-private-c"
    }
  }
}
