variable "domain_name" {
  description = "証明書を発行するドメイン名"
  type        = string
}

variable "cloudflare_zone_id" {
  description = "DNS検証レコードを作るCloudflareのゾーンID"
  type        = string
}
