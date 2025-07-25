---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_block_public_access_options"
description: |-
  Terraform resource for managing AWS VPC Block Public Access Options in a region.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_block_public_access_options

Terraform resource for managing an AWS VPC Block Public Access Options.

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
from imports.aws.vpc_block_public_access_options import VpcBlockPublicAccessOptions
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpcBlockPublicAccessOptions(self, "example",
            internet_gateway_block_mode="block-bidirectional"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `internet_gateway_block_mode` - (Required) Block mode. Needs to be one of `block-bidirectional`, `block-ingress`, `off`. If this resource is deleted, then this value will be set to `off` in the AWS account and region.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `aws_account_id` - The AWS account id to which these options apply.
* `aws_region` - The AWS region to which these options apply.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import VPC Block Public Access Options using the `aws_region`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.vpc_block_public_access_options import VpcBlockPublicAccessOptions
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpcBlockPublicAccessOptions.generate_config_for_import(self, "example", "us-east-1")
```

Using `terraform import`, import VPC Block Public Access Options using the `aws_region`. For example:

```console
% terraform import aws_vpc_block_public_access_options.example us-east-1
```

<!-- cache-key: cdktf-0.20.8 input-778eb60c05437279155c72de8a11f65ba0b3c437e2fb01c387fe7434a6e64d9a -->