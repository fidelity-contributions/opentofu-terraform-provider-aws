---
subcategory: "CloudFormation"
layout: "aws"
page_title: "AWS: aws_cloudformation_stack"
description: |-
    Provides metadata of a CloudFormation stack (e.g., outputs)
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_cloudformation_stack

The CloudFormation Stack data source allows access to stack
outputs and other useful data including the template body.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_cloudformation_stack import DataAwsCloudformationStack
from imports.aws.instance import Instance
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        network = DataAwsCloudformationStack(self, "network",
            name="my-network-stack"
        )
        Instance(self, "web",
            ami="ami-abb07bcb",
            instance_type="t2.micro",
            subnet_id=Token.as_string(
                Fn.lookup_nested(network.outputs, ["\"SubnetId\""])),
            tags={
                "Name": "HelloWorld"
            }
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the stack

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `capabilities` - List of capabilities
* `description` - Description of the stack
* `disable_rollback` - Whether the rollback of the stack is disabled when stack creation fails
* `notification_arns` - List of SNS topic ARNs to publish stack related events
* `outputs` - Map of outputs from the stack.
* `parameters` - Map of parameters that specify input parameters for the stack.
* `tags` - Map of tags associated with this stack.
* `template_body` - Structure containing the template body.
* `iam_role_arn` - ARN of the IAM role used to create the stack.
* `timeout_in_minutes` - Amount of time that can pass before the stack status becomes `CREATE_FAILED`

<!-- cache-key: cdktf-0.20.8 input-8ebecf636f61d7fcc417fb6bdfa67d5217e91adf031e8dc9be82f962fe19e429 -->