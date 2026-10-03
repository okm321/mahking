variable "region" {
  description = "ログの出力先リージョン（awslogsドライバに渡す）"
  type        = string
}

variable "log_group_name" {
  description = "CloudWatch Logsのロググループ名"
  type        = string
}

variable "log_retention_days" {
  description = "ログの保持日数"
  type        = number
}

variable "cluster_name" {
  description = "ECSクラスター名"
  type        = string
}

variable "task_family" {
  description = "タスク定義のfamily名（サービス名とIAMロール名にも使う）"
  type        = string
}

variable "task_cpu" {
  description = "タスクに割り当てるCPUユニット"
  type        = number
}

variable "task_memory" {
  description = "タスクに割り当てるメモリ（MiB）"
  type        = number
}

variable "image" {
  description = "起動するコンテナイメージ（タグ込み）"
  type        = string
}

variable "app_port" {
  description = "コンテナの待ち受けポート"
  type        = number
}

variable "environment" {
  description = "コンテナに渡す環境変数"
  type        = map(string)
}

variable "secrets" {
  description = "コンテナに渡す秘密の環境変数（環境変数名 => SSMパラメータのARN）"
  type        = map(string)
  default     = {}
}

variable "desired_count" {
  description = "起動するタスク数"
  type        = number
}

variable "subnet_ids" {
  description = "タスクを配置するサブネットのIDリスト（public）"
  type        = list(string)
}

variable "security_group_id" {
  description = "タスクに付けるSecurity GroupのID"
  type        = string
}

variable "target_group_arn" {
  description = "タスクを登録するALBターゲットグループのARN"
  type        = string
}
