---
subcategory: "Transit Gateway"
layout: "aws"
page_title: "AWS: aws_ec2_transit_gateway_route_table"
description: |-
  Manages an EC2 Transit Gateway Route Table
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ec2_transit_gateway_route_table

Manages an EC2 Transit Gateway Route Table.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Ec2TransitGatewayRouteTable } from "./.gen/providers/aws/ec2-transit-gateway-route-table";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Ec2TransitGatewayRouteTable(this, "example", {
      transitGatewayId: Token.asString(awsEc2TransitGatewayExample.id),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `transitGatewayId` - (Required) Identifier of EC2 Transit Gateway.
* `tags` - (Optional) Key-value tags for the EC2 Transit Gateway Route Table. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - EC2 Transit Gateway Route Table Amazon Resource Name (ARN).
* `defaultAssociationRouteTable` - Boolean whether this is the default association route table for the EC2 Transit Gateway.
* `defaultPropagationRouteTable` - Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
* `id` - EC2 Transit Gateway Route Table identifier
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_ec2_transit_gateway_route_table` using the EC2 Transit Gateway Route Table identifier. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Ec2TransitGatewayRouteTable } from "./.gen/providers/aws/ec2-transit-gateway-route-table";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Ec2TransitGatewayRouteTable.generateConfigForImport(
      this,
      "example",
      "tgw-rtb-12345678"
    );
  }
}

```

Using `terraform import`, import `aws_ec2_transit_gateway_route_table` using the EC2 Transit Gateway Route Table identifier. For example:

```console
% terraform import aws_ec2_transit_gateway_route_table.example tgw-rtb-12345678
```

<!-- cache-key: cdktf-0.20.8 input-c7779520f790084f4f1fe55c59c983ab3dde17731aecaedbd5334d42da97141c -->