---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_connections"
description: |-
  Provides details about existing Network Manager connections.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_networkmanager_connections

Provides details about existing Network Manager connections.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_networkmanager_connections import DataAwsNetworkmanagerConnections
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsNetworkmanagerConnections(self, "example",
            global_network_id=global_network_id.string_value,
            tags={
                "Env": "test"
            }
        )
```

## Argument Reference

This data source supports the following arguments:

* `device_id` - (Optional) ID of the device of the connections to retrieve.
* `global_network_id` - (Required) ID of the Global Network of the connections to retrieve.
* `tags` - (Optional) Restricts the list to the connections with these tags.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `ids` - IDs of the connections.

<!-- cache-key: cdktf-0.20.8 input-e8236f73dfe9cabafd035bbcf99a5830e8112d9e04ce070c41bbcf9f78debba1 -->