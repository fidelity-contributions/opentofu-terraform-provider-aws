---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_ipv6_cidr_block_association"
description: |-
  Associate additional IPv6 CIDR blocks with a VPC
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_ipv6_cidr_block_association

Provides a resource to associate additional IPv6 CIDR blocks with a VPC.

The `aws_vpc_ipv6_cidr_block_association` resource allows IPv6 CIDR blocks to be added to the VPC.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Vpc } from "./.gen/providers/aws/vpc";
import { VpcIpv6CidrBlockAssociation } from "./.gen/providers/aws/vpc-ipv6-cidr-block-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const test = new Vpc(this, "test", {
      cidrBlock: "10.0.0.0/16",
    });
    const awsVpcIpv6CidrBlockAssociationTest = new VpcIpv6CidrBlockAssociation(
      this,
      "test_1",
      {
        ipv6IpamPoolId: Token.asString(awsVpcIpamPoolTest.id),
        vpcId: test.id,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsVpcIpv6CidrBlockAssociationTest.overrideLogicalId("test");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `assignGeneratedIpv6CidrBlock` - (Optional) Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IPv6 addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6IpamPoolId`, `ipv6Pool`, `ipv6CidrBlock` and `ipv6NetmaskLength`.
* `ipv6CidrBlock` - (Optional) The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6NetmaskLength`. This parameter is required if `ipv6NetmaskLength` is not set and the IPAM pool does not have `allocation_default_netmask` set. Conflicts with `assignGeneratedIpv6CidrBlock`.
* `ipv6IpamPoolId` - (Optional) The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Conflict with `assignGeneratedIpv6CidrBlock` and `ipv6Pool`.
* `ipv6NetmaskLength` - (Optional) The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6IpamPoolId`. This parameter is optional if the IPAM pool has `allocation_default_netmask` set, otherwise it or `ipv6CidrBlock` are required. Conflicts with `ipv6CidrBlock`.
* `ipv6Pool` - (Optional) The  ID of an IPv6 address pool from which to allocate the IPv6 CIDR block. Conflicts with `assignGeneratedIpv6CidrBlock` and `ipv6IpamPoolId`.
* `vpcId` - (Required) The ID of the VPC to make the association with.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `10m`)

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the VPC CIDR association.
* `ipSource` - The source that allocated the IP address space. Values: `amazon`, `byoip`, `none`.
* `ipv6AddressAttribute` - Public IPv6 addresses are those advertised on the internet from AWS. Private IP addresses are not and cannot be advertised on the internet from AWS. Values: `public`, `private`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_vpc_ipv6_cidr_block_association` using the VPC CIDR association ID and optionally the IPv6 IPAM pool ID and netmask length. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcIpv6CidrBlockAssociation } from "./.gen/providers/aws/vpc-ipv6-cidr-block-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcIpv6CidrBlockAssociation.generateConfigForImport(
      this,
      "example",
      "vpc-cidr-assoc-0754129087e149dcd"
    );
  }
}

```

or

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcIpv6CidrBlockAssociation } from "./.gen/providers/aws/vpc-ipv6-cidr-block-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcIpv6CidrBlockAssociation.generateConfigForImport(
      this,
      "example",
      "vpc-cidr-assoc-0754129087e149dcd,ipam-pool-0611d1d6bbc05ce60"
    );
  }
}

```

or

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcIpv6CidrBlockAssociation } from "./.gen/providers/aws/vpc-ipv6-cidr-block-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcIpv6CidrBlockAssociation.generateConfigForImport(
      this,
      "example",
      "vpc-cidr-assoc-0754129087e149dcd,ipam-pool-0611d1d6bbc05ce60,56"
    );
  }
}

```

Using `terraform import`, import `aws_vpc_ipv6_cidr_block_association` using the VPC CIDR association ID and optionally the IPv6 IPAM pool ID and netmask length. For example:

```console
% terraform import aws_vpc_ipv6_cidr_block_association.example vpc-cidr-assoc-0754129087e149dcd
```

or

```console
% terraform import aws_vpc_ipv6_cidr_block_association.example vpc-cidr-assoc-0754129087e149dcd,ipam-pool-0611d1d6bbc05ce60
```

or

```console
% terraform import aws_vpc_ipv6_cidr_block_association.example vpc-cidr-assoc-0754129087e149dcd,ipam-pool-0611d1d6bbc05ce60,56
```

<!-- cache-key: cdktf-0.20.8 input-ff4215c2e1ad5349d2ed3c9c2d4481876038ce488ee913140fe879292a1ae050 -->