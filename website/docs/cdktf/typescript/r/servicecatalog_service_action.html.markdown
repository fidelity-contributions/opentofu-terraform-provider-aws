---
subcategory: "Service Catalog"
layout: "aws"
page_title: "AWS: aws_servicecatalog_service_action"
description: |-
  Manages a Service Catalog Service Action
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_servicecatalog_service_action

Manages a Service Catalog self-service action.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ServicecatalogServiceAction } from "./.gen/providers/aws/servicecatalog-service-action";
interface MyConfig {
  version: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new ServicecatalogServiceAction(this, "example", {
      definition: {
        name: "AWS-RestartEC2Instance",
        version: config.version,
      },
      description: "Motor generator unit",
      name: "MGU",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `definition` - (Required) Self-service action definition configuration block. Detailed below.
* `name` - (Required) Self-service action name.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `acceptLanguage` - (Optional) Language code. Valid values are `en` (English), `jp` (Japanese), and `zh` (Chinese). Default is `en`.
* `description` - (Optional) Self-service action description.

### `definition`

The `definition` configuration block supports the following attributes:

* `assumeRole` - (Optional) ARN of the role that performs the self-service actions on your behalf. For example, `arn:aws:iam::12345678910:role/ActionRole`. To reuse the provisioned product launch role, set to `LAUNCH_ROLE`.
* `name` - (Required) Name of the SSM document. For example, `AWS-RestartEC2Instance`. If you are using a shared SSM document, you must provide the ARN instead of the name.
* `parameters` - (Optional) List of parameters in JSON format. For example: `[{\"Name\":\"InstanceId\",\"Type\":\"TARGET\"}]` or `[{\"Name\":\"InstanceId\",\"Type\":\"TEXT_VALUE\"}]`.
* `type` - (Optional) Service action definition type. Valid value is `SSM_AUTOMATION`. Default is `SSM_AUTOMATION`.
* `version` - (Required) SSM document version. For example, `1`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Identifier of the service action.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `3m`)
- `read` - (Default `10m`)
- `update` - (Default `3m`)
- `delete` - (Default `3m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_servicecatalog_service_action` using the service action ID. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ServicecatalogServiceAction } from "./.gen/providers/aws/servicecatalog-service-action";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ServicecatalogServiceAction.generateConfigForImport(
      this,
      "example",
      "act-f1w12eperfslh"
    );
  }
}

```

Using `terraform import`, import `aws_servicecatalog_service_action` using the service action ID. For example:

```console
% terraform import aws_servicecatalog_service_action.example act-f1w12eperfslh
```

<!-- cache-key: cdktf-0.20.8 input-cb614aaabe9ca54761d73f5f074ab64ad92a185d91392b7f0f7903fc491cf422 -->