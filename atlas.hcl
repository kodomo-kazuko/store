data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "postgres", // | postgres | sqlite | sqlserver
  ]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://postgres:91200290@127.0.0.1:5432/postgres?search_path=store_dev&sslmode=disable&TimeZone=Asia/Ulaanbaatar"
  url = "postgres://postgres:91200290@127.0.0.1:5432/postgres?search_path=store&sslmode=disable&TimeZone=Asia/Ulaanbaatar"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}