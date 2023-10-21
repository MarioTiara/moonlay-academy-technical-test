data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "postgres", // | postgres | sqlite
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}