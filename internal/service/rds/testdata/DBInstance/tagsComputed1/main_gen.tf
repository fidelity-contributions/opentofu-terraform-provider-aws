# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "null" {}

ephemeral "aws_secretsmanager_random_password" "test" {
  password_length     = 20
  exclude_punctuation = true
}

resource "aws_db_instance" "test" {
  identifier          = var.rName
  allocated_storage   = 10
  engine              = data.aws_rds_orderable_db_instance.test.engine
  engine_version      = data.aws_rds_orderable_db_instance.test.engine_version
  instance_class      = data.aws_rds_orderable_db_instance.test.instance_class
  skip_final_snapshot = true
  password_wo         = ephemeral.aws_secretsmanager_random_password.test.random_password
  password_wo_version = 1
  username            = "tfacctest"

  tags = {
    (var.unknownTagKey) = null_resource.test.id
  }
}

# testAccInstanceConfig_orderableClassMySQL

data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine         = data.aws_rds_engine_version.default.engine
  engine_version = data.aws_rds_engine_version.default.version
  license_model  = "general-public-license"
  storage_type   = "standard"

  preferred_instance_classes = ["db.t4g.micro"]
}

resource "null_resource" "test" {}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}

variable "unknownTagKey" {
  type     = string
  nullable = false
}
