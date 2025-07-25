---
subcategory: "Verified Permissions"
layout: "aws"
page_title: "AWS: aws_verifiedpermissions_policy_store"
description: |-
  Terraform data source for managing an AWS Verified Permissions Policy Store.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_verifiedpermissions_policy_store

Terraform data source for managing an AWS Verified Permissions Policy Store.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_verifiedpermissions_policy_store import DataAwsVerifiedpermissionsPolicyStore
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsVerifiedpermissionsPolicyStore(self, "example",
            id="example"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `id` - (Required) The ID of the Policy Store.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - The ARN of the Policy Store.
* `created_date` - The date the Policy Store was created.
* `last_updated_date` - The date the Policy Store was last updated.
* `tags` - Map of key-value pairs associated with the policy store.
* `validation_settings` - Validation settings for the policy store.

<!-- cache-key: cdktf-0.20.8 input-f2dd75e291998052773be62746a1a8848d12aba1362273b47dcf6d9a933e3154 -->