---
subcategory: "App Runner"
layout: "aws"
page_title: "AWS: aws_apprunner_vpc_connector"
description: |-
  Manages an App Runner VPC Connector.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_apprunner_vpc_connector

Manages an App Runner VPC Connector.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApprunnerVpcConnector } from "./.gen/providers/aws/apprunner-vpc-connector";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ApprunnerVpcConnector(this, "connector", {
      securityGroups: ["sg1", "sg2"],
      subnets: ["subnet1", "subnet2"],
      vpcConnectorName: "name",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `vpcConnectorName` - (Required) Name for the VPC connector.
* `subnets` (Required) List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
* `securityGroups` - List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of VPC connector.
* `status` - Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
* `vpcConnectorRevision` - The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import App Runner vpc connector using the `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApprunnerVpcConnector } from "./.gen/providers/aws/apprunner-vpc-connector";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ApprunnerVpcConnector.generateConfigForImport(
      this,
      "example",
      "arn:aws:apprunner:us-east-1:1234567890:vpcconnector/example/1/0a03292a89764e5882c41d8f991c82fe"
    );
  }
}

```

Using `terraform import`, import App Runner vpc connector using the `arn`. For example:

```console
% terraform import aws_apprunner_vpc_connector.example arn:aws:apprunner:us-east-1:1234567890:vpcconnector/example/1/0a03292a89764e5882c41d8f991c82fe
```

<!-- cache-key: cdktf-0.20.8 input-54aba0cae40e42c5f052998da0f78a26b4253418a6efb074f00766c7c88196d9 -->