---
subcategory: "Lightsail"
layout: "aws"
page_title: "AWS: aws_lightsail_disk"
description: |-
  Manages a Lightsail block storage disk.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lightsail_disk

Manages a Lightsail disk. Use this resource to create additional block storage that can be attached to Lightsail instances for extra storage capacity.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_availability_zones import DataAwsAvailabilityZones
from imports.aws.lightsail_disk import LightsailDisk
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        available = DataAwsAvailabilityZones(self, "available",
            filter=[DataAwsAvailabilityZonesFilter(
                name="opt-in-status",
                values=["opt-in-not-required"]
            )
            ],
            state="available"
        )
        LightsailDisk(self, "example",
            availability_zone=Token.as_string(Fn.lookup_nested(available.names, ["0"])),
            name="example-disk",
            size_in_gb=8
        )
```

## Argument Reference

The following arguments are required:

* `availability_zone` - (Required) Availability Zone in which to create the disk.
* `name` - (Required) Name of the disk. Must begin with an alphabetic character and contain only alphanumeric characters, underscores, hyphens, and dots.
* `size_in_gb` - (Required) Size of the disk in GB.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `tags` - (Optional) Map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the disk.
* `created_at` - Date and time when the disk was created.
* `id` - Name of the disk (matches `name`).
* `support_code` - Support code for the disk. Include this code in your email to support when you have questions about a disk in Lightsail.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_lightsail_disk` using the name attribute. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lightsail_disk import LightsailDisk
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        LightsailDisk.generate_config_for_import(self, "example", "example-disk")
```

Using `terraform import`, import `aws_lightsail_disk` using the name attribute. For example:

```console
% terraform import aws_lightsail_disk.example example-disk
```

<!-- cache-key: cdktf-0.20.8 input-b2b00c0d047812c476fdbbe150146dd440fcb6bce85e6462675d33d1795e7240 -->