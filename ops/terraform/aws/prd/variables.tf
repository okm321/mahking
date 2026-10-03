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

  # ============================================================================
  # Security Group
  # ============================================================================
  security_groups = {
    alb    = "mahking-${local.env}-sg-alb"
    ecs    = "mahking-${local.env}-sg-ecs"
    aurora = "mahking-${local.env}-sg-aurora"
  }

  # Goサーバーの待ち受けポート（go/config/config.goの既定値、DockerfileのEXPOSEと一致）
  app_port = 8080

  # ============================================================================
  # ECR
  # ============================================================================
  ecr_repository_name = "mahking"

  # ============================================================================
  # ACM / Cloudflare
  # ============================================================================
  api_domain = "api.mahking.app"

  # mahking.appのゾーンID
  cloudflare_zone_id = "6e0f20b2766a8f0268487504db6ef237"
}

variable "cloudflare_api_token" {
  description = "Cloudflare API token（mahking.app の DNS 編集用）"
  type        = string
  sensitive   = true
}
