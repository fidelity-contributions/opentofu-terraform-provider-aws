---
subcategory: "Redshift Serverless"
layout: "aws"
page_title: "AWS: aws_redshiftserverless_resource_policy"
description: |-
  Provides a Redshift Serverless Resource Policy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_redshiftserverless_resource_policy

Creates a new Amazon Redshift Serverless Resource Policy.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RedshiftserverlessResourcePolicy } from "./.gen/providers/aws/redshiftserverless-resource-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new RedshiftserverlessResourcePolicy(this, "example", {
      policy: Token.asString(
        Fn.jsonencode({
          Statement: [
            {
              Action: ["redshift-serverless:RestoreFromSnapshot"],
              Effect: "Allow",
              Principal: {
                AWS: ["12345678901"],
              },
              Sid: "",
            },
          ],
          Version: "2012-10-17",
        })
      ),
      resourceArn: Token.asString(awsRedshiftserverlessSnapshotExample.arn),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resourceArn` - (Required) The Amazon Resource Name (ARN) of the account to create or update a resource policy for.
* `policy` - (Required) The policy to create or update. For example, the following policy grants a user authorization to restore a snapshot.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Amazon Resource Name (ARN) of the account to create or update a resource policy for.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Redshift Serverless Resource Policies using the `resourceArn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RedshiftserverlessResourcePolicy } from "./.gen/providers/aws/redshiftserverless-resource-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    RedshiftserverlessResourcePolicy.generateConfigForImport(
      this,
      "example",
      "example"
    );
  }
}

```

Using `terraform import`, import Redshift Serverless Resource Policies using the `resourceArn`. For example:

```console
% terraform import aws_redshiftserverless_resource_policy.example example
```

<!-- cache-key: cdktf-0.20.8 input-8fcb02f320a67b1d265579b4032bb742e9deacc0f5ede169163fe050c0de77f6 -->