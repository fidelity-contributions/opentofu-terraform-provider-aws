---
subcategory: "GuardDuty"
layout: "aws"
page_title: "AWS: aws_guardduty_detector_feature"
description: |-
  Provides a resource to manage an Amazon GuardDuty detector feature
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_guardduty_detector_feature

Provides a resource to manage a single Amazon GuardDuty [detector feature](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-features-activation-model.html#guardduty-features).

~> **NOTE:** Deleting this resource does not disable the detector feature, the resource in simply removed from state instead.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.guardduty_detector import GuarddutyDetector
from imports.aws.guardduty_detector_feature import GuarddutyDetectorFeature
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = GuarddutyDetector(self, "example",
            enable=True
        )
        GuarddutyDetectorFeature(self, "s3_protection",
            detector_id=example.id,
            name="S3_DATA_EVENTS",
            status="ENABLED"
        )
```

## Extended Threat Detection for EKS

To enable GuardDuty [Extended Threat Detection](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-extended-threat-detection.html) for EKS, you need at least one of these features enabled: [EKS Protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html) or [Runtime Monitoring](https://docs.aws.amazon.com/guardduty/latest/ug/runtime-monitoring-configuration.html). For maximum detection coverage, enabling both is recommended to enhance detection capabilities.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.guardduty_detector import GuarddutyDetector
from imports.aws.guardduty_detector_feature import GuarddutyDetectorFeature
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = GuarddutyDetector(self, "example",
            enable=True
        )
        GuarddutyDetectorFeature(self, "eks_protection",
            detector_id=example.id,
            name="EKS_AUDIT_LOGS",
            status="ENABLED"
        )
        GuarddutyDetectorFeature(self, "eks_runtime_monitoring",
            additional_configuration=[GuarddutyDetectorFeatureAdditionalConfiguration(
                name="EKS_ADDON_MANAGEMENT",
                status="ENABLED"
            )
            ],
            detector_id=example.id,
            name="EKS_RUNTIME_MONITORING",
            status="ENABLED"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `detector_id` - (Required) Amazon GuardDuty detector ID.
* `name` - (Required) The name of the detector feature. Valid values: `S3_DATA_EVENTS`, `EKS_AUDIT_LOGS`, `EBS_MALWARE_PROTECTION`, `RDS_LOGIN_EVENTS`, `EKS_RUNTIME_MONITORING`, `LAMBDA_NETWORK_LOGS`, `RUNTIME_MONITORING`. Only one of two features `EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING` can be added, adding both features will cause an error. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorFeatureConfiguration.html) for the current list of supported values.
* `status` - (Required) The status of the detector feature. Valid values: `ENABLED`, `DISABLED`.
* `additional_configuration` - (Optional) Additional feature configuration block for features`EKS_RUNTIME_MONITORING` or `RUNTIME_MONITORING`. See [below](#additional-configuration).

### Additional Configuration

The `additional_configuration` block supports the following:

* `name` - (Required) The name of the additional configuration for a feature. Valid values: `EKS_ADDON_MANAGEMENT`, `ECS_FARGATE_AGENT_MANAGEMENT`, `EC2_AGENT_MANAGEMENT`. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorAdditionalConfiguration.html) for the current list of supported values.
* `status` - (Required) The status of the additional configuration. Valid values: `ENABLED`, `DISABLED`.

## Attribute Reference

This resource exports no additional attributes.

<!-- cache-key: cdktf-0.20.8 input-3b9e54dcc31c404d29b0506166e291ff042e0b69fd0be6ecf897393b7825035d -->