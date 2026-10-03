# ==============================================================================
# ECR
# ==============================================================================
resource "aws_ecr_repository" "main" {
  name = var.repository_name

  # タグはコミットSHAで付けるので上書きを許さない
  image_tag_mutability = "IMMUTABLE"

  # 個人プロジェクトなのでイメージが残っていてもdestroyできるようにする
  force_delete = true

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = var.repository_name
  }
}

# タグ付きイメージは最新10個だけ残す
resource "aws_ecr_lifecycle_policy" "main" {
  repository = aws_ecr_repository.main.name

  policy = jsonencode({
    rules = [
      {
        rulePriority = 1
        description  = "Keep last 10 tagged images"
        selection = {
          tagStatus      = "tagged"
          tagPatternList = ["*"]
          countType      = "imageCountMoreThan"
          countNumber    = 10
        }
        action = {
          type = "expire"
        }
      }
    ]
  })
}
