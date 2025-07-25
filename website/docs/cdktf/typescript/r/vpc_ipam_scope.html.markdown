---
subcategory: "VPC IPAM (IP Address Manager)"
layout: "aws"
page_title: "AWS: aws_vpc_ipam_scope"
description: |-
  Creates a scope for AWS IPAM.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_ipam_scope

Creates a scope for AWS IPAM.

## Example Usage

Basic usage:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
import { VpcIpam } from "./.gen/providers/aws/vpc-ipam";
import { VpcIpamScope } from "./.gen/providers/aws/vpc-ipam-scope";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const current = new DataAwsRegion(this, "current", {});
    const example = new VpcIpam(this, "example", {
      operatingRegions: [
        {
          regionName: Token.asString(current.region),
        },
      ],
    });
    const awsVpcIpamScopeExample = new VpcIpamScope(this, "example_2", {
      description: "Another Scope",
      ipamId: example.id,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsVpcIpamScopeExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `ipamId` - The ID of the IPAM for which you're creating this scope.
* `description` - (Optional) A description for the scope you're creating.
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the scope.
* `id` - The ID of the IPAM Scope.
* `ipamArn` - The ARN of the IPAM for which you're creating this scope.
* `isDefault` - Defines if the scope is the default scope or not.
* `poolCount` - The number of pools in the scope.
* `type` - The type of the scope.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import IPAMs using the `scope_id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcIpamScope } from "./.gen/providers/aws/vpc-ipam-scope";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcIpamScope.generateConfigForImport(
      this,
      "example",
      "ipam-scope-0513c69f283d11dfb"
    );
  }
}

```

Using `terraform import`, import IPAMs using the `scope_id`. For example:

```console
% terraform import aws_vpc_ipam_scope.example ipam-scope-0513c69f283d11dfb
```

<!-- cache-key: cdktf-0.20.8 input-f8a3b356b43d37fc9f51ab8fd7668de845677242cade22679fec7aa92160e259 -->