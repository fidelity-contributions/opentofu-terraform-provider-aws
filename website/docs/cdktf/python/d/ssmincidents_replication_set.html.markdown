---
subcategory: "SSM Incident Manager Incidents"
layout: "aws"
page_title: "AWS: aws_ssmincidents_replication_set"
description: |-
  Terraform data source for managing an incident replication set in AWS Systems Manager Incident Manager.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ssmincidents_replication_set

~> **NOTE:** The AWS Region specified by a Terraform provider must always be one of the Regions specified for the replication set.

Use this Terraform data source to manage a replication set in AWS Systems Manager Incident Manager.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ssmincidents_replication_set import DataAwsSsmincidentsReplicationSet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSsmincidentsReplicationSet(self, "example")
```

## Argument Reference

This data source does not support any arguments.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the replication set.
* `created_by` - The ARN of the user who created the replication set.
* `deletion_protected` - If `true`, the last remaining Region in a replication set can’t be deleted.
* `last_modified_by` - The ARN of the user who last modified the replication set.
* `region` - (**Deprecated**) The replication set's Regions. Use `regions` instead.
* `regions` - The replication set's Regions.
* `status` - The overall status of a replication set.
    * Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
* `tags` - All tags applied to the replication set.

The `regions` configuration block exports the following attributes for each Region:

* `name` - The name of the Region.
* `kms_key_arn` - The ARN of the AWS Key Management Service (AWS KMS) encryption key.
* `status` - The current status of the Region.
    * Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
* `status_message` - More information about the status of a Region.

<!-- cache-key: cdktf-0.20.8 input-3bb6e56e2ec19fe4bf2b06809872b6724ec4e65d6c90ed72df2afd490eea6995 -->