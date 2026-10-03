# cloudflareはhashicorp名前空間に無いので、モジュール側でもsourceを宣言する
terraform {
  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
    }
  }
}
