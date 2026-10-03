# https://claude.ai/artifact/3a7raZy6ha4A6utfsdN6F7
# ==============================================================================
# VPC
# ==============================================================================
resource "aws_vpc" "main" {
  cidr_block = var.vpc_cidr

  # 10.0.0.2 のAWSのDNSサーバーが応答するか。falseだとAuroraのホスト名を引けない
  enable_dns_support = true

  # public IPにpublic DNS名を付けるか。将来のinterface VPC endpoint用（既定false）
  enable_dns_hostnames = true

  tags = {
    Name = var.vpc_name
  }
}

# ==============================================================================
# Internet Gateway
# ==============================================================================
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = var.igw_name
  }
}

# ==============================================================================
# Subnets
#
# NAT Gatewayは置かない（月$30超。タスクのpublic IPv4なら月$3.6）。Fargateは
# publicに置いてSecurity Groupで守る。privateはAurora用。
# ==============================================================================
resource "aws_subnet" "public" {
  for_each = var.public_subnets

  vpc_id                  = aws_vpc.main.id
  availability_zone       = each.value.az
  cidr_block              = each.value.cidr
  map_public_ip_on_launch = true

  tags = {
    Name = each.value.name
    Tier = "public"
  }
}

resource "aws_subnet" "private" {
  for_each = var.private_subnets

  vpc_id            = aws_vpc.main.id
  availability_zone = each.value.az
  cidr_block        = each.value.cidr

  tags = {
    Name = each.value.name
    Tier = "private"
  }
}

# ==============================================================================
# Route Tables
# ==============================================================================
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "${var.vpc_name}-public-rt"
  }
}

# privateサブネットはVPC内のローカルルートのみ（外向きのルートは持たない）
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "${var.vpc_name}-private-rt"
  }
}

resource "aws_route_table_association" "public" {
  for_each = aws_subnet.public

  subnet_id      = each.value.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "private" {
  for_each = aws_subnet.private

  subnet_id      = each.value.id
  route_table_id = aws_route_table.private.id
}
