---
subcategory: "EFS (Elastic File System)"
layout: "aws"
page_title: "AWS: aws_efs_backup_policy"
description: |-
  Provides an Elastic File System (EFS) Backup Policy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_efs_backup_policy

Provides an Elastic File System (EFS) Backup Policy resource.
Backup policies turn automatic backups on or off for an existing file system.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.efs_backup_policy import EfsBackupPolicy
from imports.aws.efs_file_system import EfsFileSystem
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        fs = EfsFileSystem(self, "fs",
            creation_token="my-product"
        )
        EfsBackupPolicy(self, "policy",
            backup_policy=EfsBackupPolicyBackupPolicy(
                status="ENABLED"
            ),
            file_system_id=fs.id
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `file_system_id` - (Required) The ID of the EFS file system.
* `backup_policy` - (Required) A backup_policy object (documented below).

### Backup Policy Arguments

`backup_policy` supports the following arguments:

* `status` - (Required) A status of the backup policy. Valid values: `ENABLED`, `DISABLED`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID that identifies the file system (e.g., fs-ccfc0d65).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the EFS backup policies using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.efs_backup_policy import EfsBackupPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EfsBackupPolicy.generate_config_for_import(self, "example", "fs-6fa144c6")
```

Using `terraform import`, import the EFS backup policies using the `id`. For example:

```console
% terraform import aws_efs_backup_policy.example fs-6fa144c6
```

<!-- cache-key: cdktf-0.20.8 input-edb2131594081169c813eced70a5fa90a495f6ca171799a3ea9849321846452a -->