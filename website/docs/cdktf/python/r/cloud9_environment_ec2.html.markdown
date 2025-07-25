---
subcategory: "Cloud9"
layout: "aws"
page_title: "AWS: aws_cloud9_environment_ec2"
description: |-
  Provides a Cloud9 EC2 Development Environment.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cloud9_environment_ec2

Provides a Cloud9 EC2 Development Environment.

## Example Usage

Basic usage:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloud9_environment_ec2 import Cloud9EnvironmentEc2
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Cloud9EnvironmentEc2(self, "example",
            image_id="amazonlinux-2023-x86_64",
            instance_type="t2.micro",
            name="example-env"
        )
```

Get the URL of the Cloud9 environment after creation:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformOutput, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloud9_environment_ec2 import Cloud9EnvironmentEc2
from imports.aws.data_aws_instance import DataAwsInstance
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, imageId, name):
        super().__init__(scope, name)
        example = Cloud9EnvironmentEc2(self, "example",
            instance_type="t2.micro",
            image_id=image_id,
            name=name
        )
        DataAwsInstance(self, "cloud9_instance",
            filter=[DataAwsInstanceFilter(
                name="tag:aws:cloud9:environment",
                values=[example.id]
            )
            ]
        )
        TerraformOutput(self, "cloud9_url",
            value="https://${" + region.value + "}.console.aws.amazon.com/cloud9/ide/${" + example.id + "}"
        )
```

Allocate a static IP to the Cloud9 environment:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformOutput, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloud9_environment_ec2 import Cloud9EnvironmentEc2
from imports.aws.data_aws_instance import DataAwsInstance
from imports.aws.eip import Eip
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, imageId, name):
        super().__init__(scope, name)
        example = Cloud9EnvironmentEc2(self, "example",
            instance_type="t2.micro",
            image_id=image_id,
            name=name
        )
        cloud9_instance = DataAwsInstance(self, "cloud9_instance",
            filter=[DataAwsInstanceFilter(
                name="tag:aws:cloud9:environment",
                values=[example.id]
            )
            ]
        )
        cloud9_eip = Eip(self, "cloud9_eip",
            domain="vpc",
            instance=Token.as_string(cloud9_instance.id)
        )
        TerraformOutput(self, "cloud9_public_ip",
            value=cloud9_eip.public_ip
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the environment.
* `instance_type` - (Required) The type of instance to connect to the environment, e.g., `t2.micro`.
* `image_id` - (Required) The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
    * `amazonlinux-2-x86_64`
    * `amazonlinux-2023-x86_64`
    * `ubuntu-18.04-x86_64`
    * `ubuntu-22.04-x86_64`
    * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
    * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2023-x86_64`
    * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
    * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-22.04-x86_64`
* `automatic_stop_time_minutes` - (Optional) The number of minutes until the running instance is shut down after the environment has last been used.
* `connection_type` - (Optional) The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
* `description` - (Optional) The description of the environment.
* `owner_arn` - (Optional) The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
* `subnet_id` - (Optional) The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the environment.
* `arn` - The ARN of the environment.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `type` - The type of the environment (e.g., `ssh` or `ec2`).

<!-- cache-key: cdktf-0.20.8 input-ce0b830510ee642d37cdce16200104f8b4e81b55f37b9d8f815c04640abd56da -->