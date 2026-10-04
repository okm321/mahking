output "instance_id" {
  description = "bastion EC2のインスタンスID（start-instances / start-sessionの--target）"
  value       = aws_instance.bastion.id
}

output "security_group_id" {
  description = "bastion用Security GroupのID"
  value       = aws_security_group.bastion.id
}
