---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_group_membership"
description: |-
  Manages a Resource QuickSight Group Membership.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_group_membership

Resource for managing QuickSight Group Membership

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightGroupMembership } from "./.gen/providers/aws/quicksight-group-membership";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new QuicksightGroupMembership(this, "example", {
      groupName: "all-access-users",
      memberName: "john_smith",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `groupName` - (Required) The name of the group in which the member will be added.
* `memberName` - (Required) The name of the member to add to the group.
* `awsAccountId` - (Optional) The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
* `namespace` - (Required) The namespace that you want the user to be a part of. Defaults to `default`.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import QuickSight Group membership using the AWS account ID, namespace, group name and member name separated by `/`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightGroupMembership } from "./.gen/providers/aws/quicksight-group-membership";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    QuicksightGroupMembership.generateConfigForImport(
      this,
      "example",
      "123456789123/default/all-access-users/john_smith"
    );
  }
}

```

Using `terraform import`, import QuickSight Group membership using the AWS account ID, namespace, group name and member name separated by `/`. For example:

```console
% terraform import aws_quicksight_group_membership.example 123456789123/default/all-access-users/john_smith
```

<!-- cache-key: cdktf-0.20.8 input-a650dc7bdc278adb5f9bc2243685571b0dcb59b014e77ccfa668410220af52d1 -->