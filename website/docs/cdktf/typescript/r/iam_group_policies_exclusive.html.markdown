---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "AWS: aws_iam_group_policies_exclusive"
description: |-
  Terraform resource for maintaining exclusive management of inline policies assigned to an AWS IAM (Identity & Access Management) group.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_iam_group_policies_exclusive

Terraform resource for maintaining exclusive management of inline policies assigned to an AWS IAM (Identity & Access Management) group.

!> This resource takes exclusive ownership over inline policies assigned to a group. This includes removal of inline policies which are not explicitly configured. To prevent persistent drift, ensure any `aws_iam_group_policy` resources managed alongside this resource are included in the `policyNames` argument.

~> Destruction of this resource means Terraform will no longer manage reconciliation of the configured inline policy assignments. It __will not__ delete the configured policies from the group.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { IamGroupPoliciesExclusive } from "./.gen/providers/aws/iam-group-policies-exclusive";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new IamGroupPoliciesExclusive(this, "example", {
      groupName: Token.asString(awsIamGroupExample.name),
      policyNames: [Token.asString(awsIamGroupPolicyExample.name)],
    });
  }
}

```

### Disallow Inline Policies

To automatically remove any configured inline policies, set the `policyNames` argument to an empty list.

~> This will not __prevent__ inline policies from being assigned to a group via Terraform (or any other interface). This resource enables bringing inline policy assignments into a configured state, however, this reconciliation happens only when `apply` is proactively run.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { IamGroupPoliciesExclusive } from "./.gen/providers/aws/iam-group-policies-exclusive";
interface MyConfig {
  policyNames: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new IamGroupPoliciesExclusive(this, "example", {
      groupName: Token.asString(awsIamGroupExample.name),
      policyNames: config.policyNames,
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `groupName` - (Required) IAM group name.
* `policyNames` - (Required) A list of inline policy names to be assigned to the group. Policies attached to this group but not configured in this argument will be removed.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to exclusively manage inline policy assignments using the `groupName`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { IamGroupPoliciesExclusive } from "./.gen/providers/aws/iam-group-policies-exclusive";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    IamGroupPoliciesExclusive.generateConfigForImport(
      this,
      "example",
      "MyGroup"
    );
  }
}

```

Using `terraform import`, import exclusive management of inline policy assignments using the `groupName`. For example:

```console
% terraform import aws_iam_group_policies_exclusive.example MyGroup
```

<!-- cache-key: cdktf-0.20.8 input-6d54235023a98a7e42e32eb54bff87275a9863c025f3cbd76e2eea91b5953669 -->