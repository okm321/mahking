variable "alb_name" {
  description = "ALB名"
  type        = string
}

variable "target_group_name" {
  description = "ターゲットグループ名"
  type        = string
}

variable "vpc_id" {
  description = "ターゲットグループを作成するVPCのID"
  type        = string
}

variable "subnet_ids" {
  description = "ALBを配置するサブネットのIDリスト（public）"
  type        = list(string)
}

variable "security_group_id" {
  description = "ALBに付けるSecurity GroupのID"
  type        = string
}

variable "app_port" {
  description = "転送先コンテナの待ち受けポート"
  type        = number
}

variable "certificate_arn" {
  description = "HTTPSリスナーに付けるACM証明書のARN"
  type        = string
}

variable "domain_name" {
  description = "ALBに向けるドメイン名"
  type        = string
}

variable "cloudflare_zone_id" {
  description = "CNAMEを作るCloudflareのゾーンID"
  type        = string
}
