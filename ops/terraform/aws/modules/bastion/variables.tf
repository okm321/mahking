variable "name" {
  description = "bastionの名前（EC2・Security Group・IAMロールの名前に使う）"
  type        = string
}

variable "vpc_id" {
  description = "Security Groupを作成するVPCのID"
  type        = string
}

variable "subnet_ids" {
  description = "bastionを配置するサブネットのIDリスト（public、先頭だけ使う）"
  type        = list(string)
}

variable "aurora_security_group_id" {
  description = "5432のingressを足すAurora用Security GroupのID"
  type        = string
}

variable "instance_type" {
  description = "bastionのインスタンスタイプ（arm64）"
  type        = string
}
