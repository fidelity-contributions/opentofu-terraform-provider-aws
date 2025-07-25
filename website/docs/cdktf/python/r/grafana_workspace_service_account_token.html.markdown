---
subcategory: "Managed Grafana"
layout: "aws"
page_title: "AWS: aws_grafana_workspace_service_account_token"
description: |-
  Terraform resource for managing an Amazon Managed Grafana Workspace Service Account Token.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_grafana_workspace_service_account_token

-> **Note:** You cannot update a service account token. If you change any attribute, Terraform
will delete the current and create a new one.

Read about Service Accounts Tokens in the [Amazon Managed Grafana user guide](https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html#service-account-tokens).

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.grafana_workspace_service_account import GrafanaWorkspaceServiceAccount
from imports.aws.grafana_workspace_service_account_token import GrafanaWorkspaceServiceAccountToken
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = GrafanaWorkspaceServiceAccount(self, "example",
            grafana_role="ADMIN",
            name="example-admin",
            workspace_id=Token.as_string(aws_grafana_workspace_example.id)
        )
        aws_grafana_workspace_service_account_token_example =
        GrafanaWorkspaceServiceAccountToken(self, "example_1",
            name="example-key",
            seconds_to_live=3600,
            service_account_id=example.service_account_id,
            workspace_id=Token.as_string(aws_grafana_workspace_example.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_grafana_workspace_service_account_token_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) A name for the token to create. The name must be unique within the workspace.
* `seconds_to_live` - (Required) Sets how long the token will be valid, in seconds. You can set the time up to 30 days in the future.
* `service_account_id` - (Required) The ID of the service account for which to create a token.
* `workspace_id` - (Required) The Grafana workspace with which the service account token is associated.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `service_account_token_id` - Identifier of the service account token in the given Grafana workspace.
* `created_at` - Specifies when the service account token was created.
* `expires_at` - Specifies when the service account token will expire.
* `key` - The key for the service account token. Used when making calls to the Grafana HTTP APIs to authenticate and authorize the requests.

<!-- cache-key: cdktf-0.20.8 input-98e0cea224ce9705a24bc19d210ea99d30c310bd089a192f43e8923fb39afe42 -->