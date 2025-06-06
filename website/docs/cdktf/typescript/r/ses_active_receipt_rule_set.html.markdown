---
subcategory: "SES (Simple Email)"
layout: "aws"
page_title: "AWS: aws_ses_active_receipt_rule_set"
description: |-
  Provides a resource to designate the active SES receipt rule set
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ses_active_receipt_rule_set

Provides a resource to designate the active SES receipt rule set

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SesActiveReceiptRuleSet } from "./.gen/providers/aws/ses-active-receipt-rule-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SesActiveReceiptRuleSet(this, "main", {
      ruleSetName: "primary-rules",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `ruleSetName` - (Required) The name of the rule set

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The SES receipt rule set name.
* `arn` - The SES receipt rule set ARN.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import active SES receipt rule sets using the rule set name. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SesActiveReceiptRuleSet } from "./.gen/providers/aws/ses-active-receipt-rule-set";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SesActiveReceiptRuleSet.generateConfigForImport(
      this,
      "myRuleSet",
      "my_rule_set_name"
    );
  }
}

```

Using `terraform import`, import active SES receipt rule sets using the rule set name. For example:

```console
% terraform import aws_ses_active_receipt_rule_set.my_rule_set my_rule_set_name
```

<!-- cache-key: cdktf-0.20.8 input-8aa116a29d3ebaa80524e2c31efd1f8f2d61078377ae336e09f4241a5b71237e -->