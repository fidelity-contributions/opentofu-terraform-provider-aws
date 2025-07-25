---
subcategory: "S3 Control"
layout: "aws"
page_title: "AWS: aws_s3control_object_lambda_access_point_policy"
description: |-
  Provides a resource to manage an S3 Object Lambda Access Point resource policy.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_s3control_object_lambda_access_point_policy

Provides a resource to manage an S3 Object Lambda Access Point resource policy.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { S3AccessPoint } from "./.gen/providers/aws/s3-access-point";
import { S3Bucket } from "./.gen/providers/aws/s3-bucket";
import { S3ControlObjectLambdaAccessPoint } from "./.gen/providers/aws/s3-control-object-lambda-access-point";
import { S3ControlObjectLambdaAccessPointPolicy } from "./.gen/providers/aws/s3-control-object-lambda-access-point-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new S3Bucket(this, "example", {
      bucket: "example",
    });
    const awsS3AccessPointExample = new S3AccessPoint(this, "example_1", {
      bucket: example.id,
      name: "example",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3AccessPointExample.overrideLogicalId("example");
    const awsS3ControlObjectLambdaAccessPointExample =
      new S3ControlObjectLambdaAccessPoint(this, "example_2", {
        configuration: {
          supportingAccessPoint: Token.asString(awsS3AccessPointExample.arn),
          transformationConfiguration: [
            {
              actions: ["GetObject"],
              contentTransformation: {
                awsLambda: {
                  functionArn: Token.asString(awsLambdaFunctionExample.arn),
                },
              },
            },
          ],
        },
        name: "example",
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3ControlObjectLambdaAccessPointExample.overrideLogicalId("example");
    const awsS3ControlObjectLambdaAccessPointPolicyExample =
      new S3ControlObjectLambdaAccessPointPolicy(this, "example_3", {
        name: Token.asString(awsS3ControlObjectLambdaAccessPointExample.name),
        policy: Token.asString(
          Fn.jsonencode({
            Statement: [
              {
                Action: "s3-object-lambda:GetObject",
                Effect: "Allow",
                Principal: {
                  AWS: current.accountId,
                },
                Resource: awsS3ControlObjectLambdaAccessPointExample.arn,
              },
            ],
            Version: "2008-10-17",
          })
        ),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3ControlObjectLambdaAccessPointPolicyExample.overrideLogicalId(
      "example"
    );
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `accountId` - (Optional) The AWS account ID for the account that owns the Object Lambda Access Point. Defaults to automatically determined account ID of the Terraform AWS provider.
* `name` - (Required) The name of the Object Lambda Access Point.
* `policy` - (Required) The Object Lambda Access Point resource policy document.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `hasPublicAccessPolicy` - Indicates whether this access point currently has a policy that allows public access.
* `id` - The AWS account ID and access point name separated by a colon (`:`).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Object Lambda Access Point policies using the `accountId` and `name`, separated by a colon (`:`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { S3ControlObjectLambdaAccessPointPolicy } from "./.gen/providers/aws/s3-control-object-lambda-access-point-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    S3ControlObjectLambdaAccessPointPolicy.generateConfigForImport(
      this,
      "example",
      "123456789012:example"
    );
  }
}

```

Using `terraform import`, import Object Lambda Access Point policies using the `accountId` and `name`, separated by a colon (`:`). For example:

```console
% terraform import aws_s3control_object_lambda_access_point_policy.example 123456789012:example
```

<!-- cache-key: cdktf-0.20.8 input-a36fb9dd4ec39c19d3d515d38f4252c1a8c7065d0c694219f84e91dac801697c -->