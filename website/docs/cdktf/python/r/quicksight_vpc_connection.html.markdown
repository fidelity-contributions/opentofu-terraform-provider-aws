---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_vpc_connection"
description: |-
  Terraform resource for managing an AWS QuickSight VPC Connection.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_vpc_connection

Terraform resource for managing an AWS QuickSight VPC Connection.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iam_role import IamRole
from imports.aws.quicksight_vpc_connection import QuicksightVpcConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        vpc_connection_role = IamRole(self, "vpc_connection_role",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "quicksight.amazonaws.com"
                        }
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            inline_policy=[IamRoleInlinePolicy(
                name="QuickSightVPCConnectionRolePolicy",
                policy=Token.as_string(
                    Fn.jsonencode({
                        "Statement": [{
                            "Action": ["ec2:CreateNetworkInterface", "ec2:ModifyNetworkInterfaceAttribute", "ec2:DeleteNetworkInterface", "ec2:DescribeSubnets", "ec2:DescribeSecurityGroups"
                            ],
                            "Effect": "Allow",
                            "Resource": ["*"]
                        }
                        ],
                        "Version": "2012-10-17"
                    }))
            )
            ]
        )
        QuicksightVpcConnection(self, "example",
            name="Example Connection",
            role_arn=vpc_connection_role.arn,
            security_group_ids=["sg-00000000000000000"],
            subnet_ids=["subnet-00000000000000000", "subnet-00000000000000001"],
            vpc_connection_id="example-connection-id"
        )
```

## Argument Reference

The following arguments are required:

* `vpc_connection_id` - (Required) The ID of the VPC connection.
* `name` - (Required) The display name for the VPC connection.
* `role_arn` - (Required) The IAM role to associate with the VPC connection.
* `security_group_ids` - (Required) A list of security group IDs for the VPC connection.
* `subnet_ids` - (Required) A list of subnet IDs for the VPC connection.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `aws_account_id` - (Optional) AWS account ID.
* `dns_resolvers` - (Optional) A list of IP addresses of DNS resolver endpoints for the VPC connection.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the VPC connection.
* `availability_status` - The availability status of the VPC connection. Valid values are `AVAILABLE`, `UNAVAILABLE` or `PARTIALLY_AVAILABLE`.
* `id` - A comma-delimited string joining AWS account ID and VPC connection ID.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `5m`)
* `update` - (Default `5m`)
* `delete` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import QuickSight VPC connection using the AWS account ID and VPC connection ID separated by commas (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_vpc_connection import QuicksightVpcConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightVpcConnection.generate_config_for_import(self, "example", "123456789012,example")
```

Using `terraform import`, import QuickSight VPC connection using the AWS account ID and VPC connection ID separated by commas (`,`). For example:

```console
% terraform import aws_quicksight_vpc_connection.example 123456789012,example
```

<!-- cache-key: cdktf-0.20.8 input-cfc45e8dd6e4b46dbe662279ff59b4941da91d878b04659cbd8d9938b2a64099 -->