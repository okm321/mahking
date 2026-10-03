# ==============================================================================
# ALB
# ==============================================================================
resource "aws_lb" "main" {
  name               = var.alb_name
  load_balancer_type = "application"
  internal           = false
  subnets            = var.subnet_ids
  security_groups    = [var.security_group_id]

  # 個人プロジェクトなのでdestroyを妨げない
  enable_deletion_protection = false

  tags = {
    Name = var.alb_name
  }
}

# ==============================================================================
# Target Group
# ==============================================================================
resource "aws_lb_target_group" "main" {
  name     = var.target_group_name
  vpc_id   = var.vpc_id
  port     = var.app_port
  protocol = "HTTP"

  # Fargate（awsvpc）のタスクはIPで登録されるのでinstanceは使えない
  target_type = "ip"

  health_check {
    path                = "/healthz"
    matcher             = "200"
    interval            = 30
    healthy_threshold   = 2
    unhealthy_threshold = 3
  }

  # 入れ替え時に古いタスクを抜くまでの待ち時間（既定300秒は長い）
  deregistration_delay = 30

  tags = {
    Name = var.target_group_name
  }
}

# ==============================================================================
# Listener
# ==============================================================================
resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.main.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn   = var.certificate_arn

  # TLS1.3と1.2だけを許可するポリシー。AWSが推奨するpost-quantum対応版で、
  # マネジメントコンソールでHTTPSリスナーを作ったときの既定値でもある
  ssl_policy = "ELBSecurityPolicy-TLS13-1-2-Res-PQ-2025-09"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.arn
  }
}

# ==============================================================================
# DNS (Cloudflare)
# ==============================================================================
resource "cloudflare_dns_record" "api" {
  zone_id = var.cloudflare_zone_id
  name    = var.domain_name
  type    = "CNAME"
  content = aws_lb.main.dns_name

  # ttlは必須。ALBのIPはAWS側で変わるのでproxyは通さず直接引かせる
  ttl     = 60
  proxied = false

  comment = "ALB for ${var.domain_name}"
}
