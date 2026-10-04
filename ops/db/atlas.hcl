variable "db_password" {
  type    = string
  default = ""
}

env "local" {
  src = "file://schema"
  // Define the URL of the database which is managed
  // in this environment.
  url = "postgres://postgres:password@127.0.0.1:15432/postgres?search_path=mahking_local&sslmode=disable"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "docker://postgres/16/dev?search_path=public"

  migration {
    dir = "file://migrations"
  }
}

env "dev" {
  src = "file://schema"
  url = "postgres://app_user:${var.db_password}@localhost:15432/mahking_dev?search_path=public&sslmode=disable"
  dev = "docker://postgres/18/dev?search_path=public"

  migration {
    dir = "file://migrations"
  }
}

env "prd" {
  src = "file://schema"
  url = "postgres://mahking:${var.db_password}@localhost:15433/mahking?search_path=public&sslmode=require"
  dev = "docker://postgres/18/dev?search_path=public"

  migration {
    dir = "file://migrations"
  }
}
