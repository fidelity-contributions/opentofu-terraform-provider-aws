---
subcategory: "App Runner"
layout: "aws"
page_title: "AWS: aws_apprunner_auto_scaling_configuration_version"
description: |-
  Manages an App Runner AutoScaling Configuration Version.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_apprunner_auto_scaling_configuration_version

Manages an App Runner AutoScaling Configuration Version.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApprunnerAutoScalingConfigurationVersion } from "./.gen/providers/aws/apprunner-auto-scaling-configuration-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ApprunnerAutoScalingConfigurationVersion(this, "example", {
      autoScalingConfigurationName: "example",
      maxConcurrency: 50,
      maxSize: 10,
      minSize: 2,
      tags: {
        Name: "example-apprunner-autoscaling",
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `autoScalingConfigurationName` - (Required, Forces new resource) Name of the auto scaling configuration.
* `maxConcurrency` - (Optional, Forces new resource) Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
* `maxSize` - (Optional, Forces new resource) Maximal number of instances that App Runner provisions for your service.
* `minSize` - (Optional, Forces new resource) Minimal number of instances that App Runner provisions for your service.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of this auto scaling configuration version.
* `autoScalingConfigurationRevision` - The revision of this auto scaling configuration.
* `latest` - Whether the auto scaling configuration has the highest `autoScalingConfigurationRevision` among all configurations that share the same `autoScalingConfigurationName`.
* `status` - Current state of the auto scaling configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import App Runner AutoScaling Configuration Versions using the `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApprunnerAutoScalingConfigurationVersion } from "./.gen/providers/aws/apprunner-auto-scaling-configuration-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ApprunnerAutoScalingConfigurationVersion.generateConfigForImport(
      this,
      "example",
      "arn:aws:apprunner:us-east-1:1234567890:autoscalingconfiguration/example/1/69bdfe0115224b0db49398b7beb68e0f"
    );
  }
}

```

Using `terraform import`, import App Runner AutoScaling Configuration Versions using the `arn`. For example:

```console
% terraform import aws_apprunner_auto_scaling_configuration_version.example "arn:aws:apprunner:us-east-1:1234567890:autoscalingconfiguration/example/1/69bdfe0115224b0db49398b7beb68e0f
```

<!-- cache-key: cdktf-0.20.8 input-f0d0fec8a60dc01e31b8e2009293a581be38d645b2f272894706bd7997e3761e -->