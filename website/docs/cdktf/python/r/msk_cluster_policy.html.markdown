---
subcategory: "Managed Streaming for Kafka"
layout: "aws"
page_title: "AWS: aws_msk_cluster_policy"
description: |-
  Terraform resource for managing an AWS Managed Streaming for Kafka Cluster Policy.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_msk_cluster_policy

Terraform resource for managing an AWS Managed Streaming for Kafka Cluster Policy.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_caller_identity import DataAwsCallerIdentity
from imports.aws.data_aws_partition import DataAwsPartition
from imports.aws.msk_cluster_policy import MskClusterPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        current = DataAwsCallerIdentity(self, "current")
        data_aws_partition_current = DataAwsPartition(self, "current_1")
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_partition_current.override_logical_id("current")
        MskClusterPolicy(self, "example",
            cluster_arn=Token.as_string(aws_msk_cluster_example.arn),
            policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": ["kafka:Describe*", "kafka:Get*", "kafka:CreateVpcConnection", "kafka:GetBootstrapBrokers"
                        ],
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:${" + data_aws_partition_current.partition + "}:iam::${" + current.account_id + "}:root"
                        },
                        "Resource": aws_msk_cluster_example.arn,
                        "Sid": "ExampleMskClusterPolicy"
                    }
                    ],
                    "Version": "2012-10-17"
                }))
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `cluster_arn` - (Required) The Amazon Resource Name (ARN) that uniquely identifies the cluster.
* `policy` - (Required) Resource policy for cluster.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Same as `cluster_arn`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Managed Streaming for Kafka Cluster Policy using the `cluster_arn. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.msk_cluster_policy import MskClusterPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        MskClusterPolicy.generate_config_for_import(self, "example", "arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3")
```

Using `terraform import`, import Managed Streaming for Kafka Cluster Policy using the `cluster_arn`. For example:

```console
% terraform import aws_msk_cluster_policy.example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
```

<!-- cache-key: cdktf-0.20.8 input-c73751e80c7b773eb93bc038a018a40b85795e72756c0938b467faea9801db7b -->