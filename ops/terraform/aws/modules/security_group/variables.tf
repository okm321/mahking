variable "vpc_id" {
  description = "Security Groupを作成するVPCのID"
  type        = string
}

variable "alb_sg_name" {
  description = "ALB用Security Group名"
  type        = string
}

variable "ecs_sg_name" {
  description = "ECSタスク用Security Group名"
  type        = string
}

variable "aurora_sg_name" {
  description = "Aurora用Security Group名"
  type        = string
}

variable "app_port" {
  description = "ECSタスクが待ち受けるポート（Goサーバーの既定値）"
  type        = number
}
