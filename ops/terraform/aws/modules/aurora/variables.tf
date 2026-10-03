variable "cluster_identifier" {
  description = "Auroraクラスターの識別子（DBインスタンスとサブネットグループの名前にも使う）"
  type        = string
}

variable "subnet_ids" {
  description = "DBサブネットグループに入れるサブネットのIDリスト（private）"
  type        = list(string)
}

variable "security_group_id" {
  description = "クラスターに付けるSecurity GroupのID"
  type        = string
}

variable "engine_version" {
  description = "Aurora PostgreSQLのエンジンバージョン"
  type        = string
}

variable "database_name" {
  description = "クラスター作成時に作るデータベース名"
  type        = string
}

variable "master_username" {
  description = "マスターユーザー名"
  type        = string
}

variable "min_capacity" {
  description = "Serverless v2の最小ACU（0で自動停止）"
  type        = number
}

variable "max_capacity" {
  description = "Serverless v2の最大ACU"
  type        = number
}

variable "seconds_until_auto_pause" {
  description = "接続が無くなってから自動停止するまでの秒数（300〜86400）"
  type        = number
}

variable "password_ssm_name" {
  description = "マスターユーザーのパスワードを置くSSMパラメータ名"
  type        = string
}


