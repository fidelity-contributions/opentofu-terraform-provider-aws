---
subcategory: "Redshift Serverless"
layout: "aws"
page_title: "AWS: aws_redshiftserverless_workgroup"
description: |-
  Terraform data source for managing an AWS Redshift Serverless Workgroup.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_redshiftserverless_workgroup

Terraform data source for managing an AWS Redshift Serverless Workgroup.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_redshiftserverless_workgroup import DataAwsRedshiftserverlessWorkgroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsRedshiftserverlessWorkgroup(self, "example",
            workgroup_name=Token.as_string(aws_redshiftserverless_workgroup_example.workgroup_name)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `workgroup_name` - (Required) The name of the workgroup associated with the database.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the Redshift Serverless Workgroup.
* `id` - The Redshift Workgroup Name.
* `endpoint` - The endpoint that is created from the workgroup. See `Endpoint` below.
* `enhanced_vpc_routing` - The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
* `publicly_accessible` - A value that specifies whether the workgroup can be accessed from a public network.
* `security_group_ids` - An array of security group IDs to associate with the workgroup.
* `subnet_ids` - An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
* `track_name` - The name of the track for the workgroup.
* `workgroup_id` - The Redshift Workgroup ID.

### Endpoint

* `address` - The DNS address of the VPC endpoint.
* `port` - The port that Amazon Redshift Serverless listens on.
* `vpc_endpoint` - The VPC endpoint or the Redshift Serverless workgroup. See `VPC Endpoint` below.

#### VPC Endpoint

* `vpc_endpoint_id` - The DNS address of the VPC endpoint.
* `vpc_id` - The port that Amazon Redshift Serverless listens on.
* `network_interface` - The network interfaces of the endpoint.. See `Network Interface` below.

##### Network Interface

* `availability_zone` - The availability Zone.
* `network_interface_id` - The unique identifier of the network interface.
* `private_ip_address` - The IPv4 address of the network interface within the subnet.
* `subnet_id` - The unique identifier of the subnet.

<!-- cache-key: cdktf-0.20.8 input-694cd55672591c32871dd2095116e69edec959f297502c13cb90376417ca136b -->