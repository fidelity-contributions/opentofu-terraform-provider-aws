---
subcategory: "GuardDuty"
layout: "aws"
page_title: "AWS: aws_guardduty_organization_configuration_feature"
description: |-
  Provides a resource to manage an Amazon GuardDuty organization configuration feature
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_guardduty_organization_configuration_feature

Provides a resource to manage a single Amazon GuardDuty [organization configuration feature](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-features-activation-model.html#guardduty-features).

~> **NOTE:** Deleting this resource does not disable the organization configuration feature, the resource is simply removed from state instead.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { GuarddutyDetector } from "./.gen/providers/aws/guardduty-detector";
import { GuarddutyOrganizationConfigurationFeature } from "./.gen/providers/aws/guardduty-organization-configuration-feature";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new GuarddutyDetector(this, "example", {
      enable: true,
    });
    new GuarddutyOrganizationConfigurationFeature(
      this,
      "eks_runtime_monitoring",
      {
        additionalConfiguration: [
          {
            autoEnable: "NEW",
            name: "EKS_ADDON_MANAGEMENT",
          },
        ],
        autoEnable: "ALL",
        detectorId: example.id,
        name: "EKS_RUNTIME_MONITORING",
      }
    );
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `autoEnable` - (Required) The status of the feature that is configured for the member accounts within the organization. Valid values: `NEW`, `ALL`, `NONE`.
* `detectorId` - (Required) The ID of the detector that configures the delegated administrator.
* `name` - (Required) The name of the feature that will be configured for the organization. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`. Only one of two features `EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING` can be added, adding both features will cause an error. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorFeatureConfiguration.html) for the current list of supported values.
* `additionalConfiguration` - (Optional) Additional feature configuration block for features `EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING`. See [below](#additional-configuration).

### Additional Configuration

The `additionalConfiguration` block supports the following:

* `autoEnable` - (Required) The status of the additional configuration that will be configured for the organization. Valid values: `NEW`, `ALL`, `NONE`.
* `name` - (Required) The name of the additional configuration for a feature that will be configured for the organization. Valid values: `EKS_ADDON_MANAGEMENT`, `ECS_FARGATE_AGENT_MANAGEMENT`, `EC2_AGENT_MANAGEMENT`. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorAdditionalConfiguration.html) for the current list of supported values.

## Attribute Reference

This resource exports no additional attributes.

<!-- cache-key: cdktf-0.20.8 input-e8c9b0e73d4df5f41de6f642b9098ea031b4fd42b90910dc23c01f5f701b9f0f -->