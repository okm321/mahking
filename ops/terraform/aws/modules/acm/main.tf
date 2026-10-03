# ==============================================================================
# ACM Certificate
# ==============================================================================
resource "aws_acm_certificate" "main" {
  domain_name       = var.domain_name
  validation_method = "DNS"

  # 証明書の差し替え時に先に新しいものを作る（ALBから参照中でも入れ替えられる）
  lifecycle {
    create_before_destroy = true
  }

  tags = {
    Name = var.domain_name
  }
}

# ==============================================================================
# DNS Validation (Cloudflare)
#
# ACMが要求する検証用CNAME。ACMはFQDNを末尾ドット付きで返すが、
# cloudflare_dns_recordのname/contentは末尾ドットなしで書くのでtrimsuffixする
# ==============================================================================
resource "cloudflare_dns_record" "validation" {
  for_each = {
    for dvo in aws_acm_certificate.main.domain_validation_options : dvo.domain_name => dvo
  }

  zone_id = var.cloudflare_zone_id
  name    = trimsuffix(each.value.resource_record_name, ".")
  type    = each.value.resource_record_type
  content = trimsuffix(each.value.resource_record_value, ".")

  # ttlは必須。検証レコードはCloudflareのproxyを通さない
  ttl     = 60
  proxied = false

  comment = "ACM validation for ${each.value.domain_name}"
}

# レコードが引けて証明書がISSUEDになるまで待つ
resource "aws_acm_certificate_validation" "main" {
  certificate_arn         = aws_acm_certificate.main.arn
  validation_record_fqdns = [for r in cloudflare_dns_record.validation : r.name]
}
