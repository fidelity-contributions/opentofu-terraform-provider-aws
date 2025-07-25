---
subcategory: "Secrets Manager"
layout: "aws"
page_title: "AWS: aws_secretsmanager_secret_policy"
description: |-
  Provides a resource to manage AWS Secrets Manager secret policy
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_secretsmanager_secret_policy

Provides a resource to manage AWS Secrets Manager secret policy.

## Example Usage

### Basic

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { SecretsmanagerSecret } from "./.gen/providers/aws/secretsmanager-secret";
import { SecretsmanagerSecretPolicy } from "./.gen/providers/aws/secretsmanager-secret-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new SecretsmanagerSecret(this, "example", {
      name: "example",
    });
    const dataAwsIamPolicyDocumentExample = new DataAwsIamPolicyDocument(
      this,
      "example_1",
      {
        statement: [
          {
            actions: ["secretsmanager:GetSecretValue"],
            effect: "Allow",
            principals: [
              {
                identifiers: ["arn:aws:iam::123456789012:root"],
                type: "AWS",
              },
            ],
            resources: ["*"],
            sid: "EnableAnotherAWSAccountToReadTheSecret",
          },
        ],
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsIamPolicyDocumentExample.overrideLogicalId("example");
    const awsSecretsmanagerSecretPolicyExample = new SecretsmanagerSecretPolicy(
      this,
      "example_2",
      {
        policy: Token.asString(dataAwsIamPolicyDocumentExample.json),
        secretArn: example.arn,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsSecretsmanagerSecretPolicyExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `policy` - (Required) Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://learn.hashicorp.com/terraform/aws/iam-policy). Unlike `aws_secretsmanager_secret`, where `policy` can be set to `"{}"` to delete the policy, `"{}"` is not a valid policy since `policy` is required.
* `secretArn` - (Required) Secret ARN.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `blockPublicPolicy` - (Optional) Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Amazon Resource Name (ARN) of the secret.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_secretsmanager_secret_policy` using the secret Amazon Resource Name (ARN). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecretsmanagerSecretPolicy } from "./.gen/providers/aws/secretsmanager-secret-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SecretsmanagerSecretPolicy.generateConfigForImport(
      this,
      "example",
      "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456"
    );
  }
}

```

Using `terraform import`, import `aws_secretsmanager_secret_policy` using the secret Amazon Resource Name (ARN). For example:

```console
% terraform import aws_secretsmanager_secret_policy.example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
```

<!-- cache-key: cdktf-0.20.8 input-556fd793c140149ab481605100c351f6af6bf69d7c4002848917dd8455a9d4f2 -->