---
subcategory: "VPC IPAM (IP Address Manager)"
layout: "aws"
page_title: "AWS: aws_vpc_ipam_preview_next_cidr"
description: |-
  Previews a CIDR from an IPAM address pool.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_ipam_preview_next_cidr

Previews a CIDR from an IPAM address pool. Only works for private IPv4.

## Example Usage

Basic usage:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_region import DataAwsRegion
from imports.aws.vpc_ipam import VpcIpam
from imports.aws.vpc_ipam_pool import VpcIpamPool
from imports.aws.vpc_ipam_pool_cidr import VpcIpamPoolCidr
from imports.aws.vpc_ipam_preview_next_cidr import VpcIpamPreviewNextCidr
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        current = DataAwsRegion(self, "current")
        example = VpcIpam(self, "example",
            operating_regions=[VpcIpamOperatingRegions(
                region_name=Token.as_string(current.region)
            )
            ]
        )
        aws_vpc_ipam_pool_example = VpcIpamPool(self, "example_2",
            address_family="ipv4",
            ipam_scope_id=example.private_default_scope_id,
            locale=Token.as_string(current.region)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_vpc_ipam_pool_example.override_logical_id("example")
        aws_vpc_ipam_pool_cidr_example = VpcIpamPoolCidr(self, "example_3",
            cidr="172.20.0.0/16",
            ipam_pool_id=Token.as_string(aws_vpc_ipam_pool_example.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_vpc_ipam_pool_cidr_example.override_logical_id("example")
        aws_vpc_ipam_preview_next_cidr_example = VpcIpamPreviewNextCidr(self, "example_4",
            depends_on=[aws_vpc_ipam_pool_cidr_example],
            disallowed_cidrs=["172.2.0.0/32"],
            ipam_pool_id=Token.as_string(aws_vpc_ipam_pool_example.id),
            netmask_length=28
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_vpc_ipam_preview_next_cidr_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `disallowed_cidrs` - (Optional) Exclude a particular CIDR range from being returned by the pool.
* `ipam_pool_id` - (Required) The ID of the pool to which you want to assign a CIDR.
* `netmask_length` - (Optional) The netmask length of the CIDR you would like to preview from the IPAM pool.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `cidr` - The previewed CIDR from the pool.
* `id` - The ID of the preview.

<!-- cache-key: cdktf-0.20.8 input-a9d469cc28f2c124149dea8d55d016430e288e984381bdb07bdbbc72e3cf54ba -->