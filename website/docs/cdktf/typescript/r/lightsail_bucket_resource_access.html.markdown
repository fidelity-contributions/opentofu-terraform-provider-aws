---
subcategory: "Lightsail"
layout: "aws"
page_title: "AWS: aws_lightsail_bucket_resource_access"
description: |-
  Manages access permissions between Lightsail resources and buckets.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lightsail_bucket_resource_access

Manages a Lightsail bucket resource access. Use this resource to grant a Lightsail resource (such as an instance) access to a specific bucket.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LightsailBucket } from "./.gen/providers/aws/lightsail-bucket";
import { LightsailBucketResourceAccess } from "./.gen/providers/aws/lightsail-bucket-resource-access";
import { LightsailInstance } from "./.gen/providers/aws/lightsail-instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new LightsailBucket(this, "example", {
      bundleId: "small_1_0",
      name: "example-bucket",
    });
    const awsLightsailInstanceExample = new LightsailInstance(
      this,
      "example_1",
      {
        availabilityZone: "us-east-1b",
        blueprintId: "amazon_linux_2",
        bundleId: "nano_3_0",
        name: "example-instance",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLightsailInstanceExample.overrideLogicalId("example");
    const awsLightsailBucketResourceAccessExample =
      new LightsailBucketResourceAccess(this, "example_2", {
        bucketName: example.id,
        resourceName: Token.asString(awsLightsailInstanceExample.id),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLightsailBucketResourceAccessExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `bucketName` - (Required) Name of the bucket to grant access to.
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resourceName` - (Required) Name of the resource to grant bucket access.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Combination of attributes separated by a `,` to create a unique id: `bucketName`,`resourceName`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_lightsail_bucket_resource_access` using the `id` attribute. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LightsailBucketResourceAccess } from "./.gen/providers/aws/lightsail-bucket-resource-access";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    LightsailBucketResourceAccess.generateConfigForImport(
      this,
      "example",
      "example-bucket,example-instance"
    );
  }
}

```

Using `terraform import`, import `aws_lightsail_bucket_resource_access` using the `id` attribute. For example:

```console
% terraform import aws_lightsail_bucket_resource_access.example example-bucket,example-instance
```

<!-- cache-key: cdktf-0.20.8 input-c1a99be77c2d4629f7044ffd71aaec61cf560a59cf05514edcd95dc93b81da11 -->