# ==============================================================================
# CloudWatch Logs
# ==============================================================================
resource "aws_cloudwatch_log_group" "api" {
  name              = var.log_group_name
  retention_in_days = var.log_retention_days

  tags = {
    Name = var.log_group_name
  }
}

# ==============================================================================
# IAM
# ==============================================================================
data "aws_iam_policy_document" "ecs_tasks_assume" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

# ECS側がタスクを起動するときに使うロール
resource "aws_iam_role" "execution" {
  name               = "${var.task_family}-exec"
  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume.json
}

# ECRからのpullとCloudWatch Logsへの書き込みに必要
resource "aws_iam_role_policy_attachment" "execution" {
  role       = aws_iam_role.execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# アプリのコードが使うロール。将来S3等を使うときにここにポリシーを足す
# （今のGoアプリはAWSのAPIを呼ばないのでポリシーは付けない）
resource "aws_iam_role" "task" {
  name               = "${var.task_family}-task"
  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume.json
}

# ==============================================================================
# Cluster
# ==============================================================================
resource "aws_ecs_cluster" "main" {
  name = var.cluster_name

  tags = {
    Name = var.cluster_name
  }
}

# ==============================================================================
# Task Definition
# ==============================================================================
resource "aws_ecs_task_definition" "api" {
  family                   = var.task_family
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = var.task_cpu
  memory                   = var.task_memory
  execution_role_arn       = aws_iam_role.execution.arn
  task_role_arn            = aws_iam_role.task.arn

  # go/DockerfileがGOARCH=amd64でビルドしているのでX86_64
  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }

  container_definitions = jsonencode([
    {
      name      = "api"
      image     = var.image
      essential = true

      portMappings = [
        {
          containerPort = var.app_port
          protocol      = "tcp"
        }
      ]

      environment = [for k, v in var.environment : { name = k, value = v }]

      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.api.name
          awslogs-region        = var.region
          awslogs-stream-prefix = "api"
        }
      }
    }
  ])

  tags = {
    Name = var.task_family
  }
}

# ==============================================================================
# Service
# ==============================================================================
resource "aws_ecs_service" "api" {
  name            = var.task_family
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.api.arn
  launch_type     = "FARGATE"
  desired_count   = var.desired_count

  # NAT Gatewayが無い構成なのでpublicサブネットに置く。
  # assign_public_ipがfalseだとECRからイメージをpullできないので必須
  network_configuration {
    subnets          = var.subnet_ids
    security_groups  = [var.security_group_id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = var.target_group_arn
    container_name   = "api"
    container_port   = var.app_port
  }

  # 起動直後のヘルスチェック失敗でタスクが落とされないための猶予
  health_check_grace_period_seconds = 60

  # 入れ替え中も常に1タスクは残す
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200

  tags = {
    Name = var.task_family
  }
}
