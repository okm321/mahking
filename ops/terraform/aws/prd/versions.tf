terraform {
  required_version = ">= 1.10.5"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
  }

  # tfstateはS3で管理。use_lockfileはTerraform 1.10以降のS3ネイティブロック
  # （DynamoDBによるロックは公式でdeprecatedなので使わない）
  backend "s3" {
    bucket       = "mahking-tfstate"
    key          = "aws/prd/terraform.tfstate"
    region       = "ap-northeast-1"
    profile      = "mahking"
    use_lockfile = true
    encrypt      = true
  }
}

provider "aws" {
  region  = local.region
  profile = local.aws_profile

  # 全リソースに共通タグを付与
  default_tags {
    tags = {
      Project   = "mahking"
      Env       = local.env
      ManagedBy = "terraform"
    }
  }
}
