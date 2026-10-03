output "vpc_id" {
  description = "VPCのID"
  value       = aws_vpc.main.id
}

output "vpc_cidr" {
  description = "VPCのCIDRブロック"
  value       = aws_vpc.main.cidr_block
}

output "public_subnet_ids" {
  description = "publicサブネットのIDリスト"
  value       = [for s in aws_subnet.public : s.id]
}

output "private_subnet_ids" {
  description = "privateサブネットのIDリスト"
  value       = [for s in aws_subnet.private : s.id]
}

output "igw_id" {
  description = "Internet GatewayのID"
  value       = aws_internet_gateway.main.id
}
