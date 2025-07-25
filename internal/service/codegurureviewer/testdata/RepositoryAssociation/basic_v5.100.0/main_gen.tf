# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_codegurureviewer_repository_association" "test" {
  repository {
    codecommit {
      name = aws_codecommit_repository.test.repository_name
    }
  }
}

# testAccRepositoryAssociation_codecommit_repository

resource "aws_codecommit_repository" "test" {
  repository_name = var.rName
  description     = "This is a test description"
  lifecycle {
    ignore_changes = [
      tags["codeguru-reviewer"]
    ]
  }
}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.100.0"
    }
  }
}

provider "aws" {}
