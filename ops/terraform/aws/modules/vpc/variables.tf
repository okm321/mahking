variable "vpc_name" {
  description = "VPC名"
  type        = string
}

variable "vpc_cidr" {
  description = "VPCのCIDRブロック"
  type        = string
}

variable "igw_name" {
  description = "Internet Gateway名"
  type        = string
}

variable "public_subnets" {
  description = "publicサブネットの定義（キーは任意の識別子）"
  type = map(object({
    az   = string
    cidr = string
    name = string
  }))
}

variable "private_subnets" {
  description = "privateサブネットの定義（キーは任意の識別子）"
  type = map(object({
    az   = string
    cidr = string
    name = string
  }))
}
