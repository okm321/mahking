output "certificate_arn" {
  description = "発行完了したACM証明書のARN"
  value       = aws_acm_certificate_validation.main.certificate_arn
}
