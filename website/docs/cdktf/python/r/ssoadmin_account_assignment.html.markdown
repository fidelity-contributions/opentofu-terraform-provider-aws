---
subcategory: "SSO Admin"
layout: "aws"
page_title: "AWS: aws_ssoadmin_account_assignment"
description: |-
  Manages a Single Sign-On (SSO) Account Assignment
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssoadmin_account_assignment

Provides a Single Sign-On (SSO) Account Assignment resource

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_identitystore_group import DataAwsIdentitystoreGroup
from imports.aws.data_aws_ssoadmin_instances import DataAwsSsoadminInstances
from imports.aws.data_aws_ssoadmin_permission_set import DataAwsSsoadminPermissionSet
from imports.aws.ssoadmin_account_assignment import SsoadminAccountAssignment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsSsoadminInstances(self, "example")
        data_aws_ssoadmin_permission_set_example =
        DataAwsSsoadminPermissionSet(self, "example_1",
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.arns), ["0"])),
            name="AWSReadOnlyAccess"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_ssoadmin_permission_set_example.override_logical_id("example")
        data_aws_identitystore_group_example = DataAwsIdentitystoreGroup(self, "example_2",
            alternate_identifier=DataAwsIdentitystoreGroupAlternateIdentifier(
                unique_attribute=DataAwsIdentitystoreGroupAlternateIdentifierUniqueAttribute(
                    attribute_path="DisplayName",
                    attribute_value="ExampleGroup"
                )
            ),
            identity_store_id=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.identity_store_ids), ["0"]))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_identitystore_group_example.override_logical_id("example")
        aws_ssoadmin_account_assignment_example = SsoadminAccountAssignment(self, "example_3",
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.arns), ["0"])),
            permission_set_arn=Token.as_string(data_aws_ssoadmin_permission_set_example.arn),
            principal_id=Token.as_string(data_aws_identitystore_group_example.group_id),
            principal_type="GROUP",
            target_id="123456789012",
            target_type="AWS_ACCOUNT"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssoadmin_account_assignment_example.override_logical_id("example")
```

### With Managed Policy Attachment

~> Because destruction of a managed policy attachment resource also re-provisions the associated permission set to all accounts, explicitly indicating the dependency with the account assignment resource via the [`depends_on` meta argument](https://developer.hashicorp.com/terraform/language/meta-arguments/depends_on) is necessary to ensure proper deletion order when these resources are used together.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ssoadmin_instances import DataAwsSsoadminInstances
from imports.aws.identitystore_group import IdentitystoreGroup
from imports.aws.ssoadmin_account_assignment import SsoadminAccountAssignment
from imports.aws.ssoadmin_managed_policy_attachment import SsoadminManagedPolicyAttachment
from imports.aws.ssoadmin_permission_set import SsoadminPermissionSet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsSsoadminInstances(self, "example")
        aws_identitystore_group_example = IdentitystoreGroup(self, "example_1",
            description="Admin Group",
            display_name="Admin",
            identity_store_id=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.identity_store_ids), ["0"]))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_identitystore_group_example.override_logical_id("example")
        aws_ssoadmin_permission_set_example = SsoadminPermissionSet(self, "example_2",
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.arns), ["0"])),
            name="Example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssoadmin_permission_set_example.override_logical_id("example")
        SsoadminAccountAssignment(self, "account_assignment",
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.arns), ["0"])),
            permission_set_arn=Token.as_string(aws_ssoadmin_permission_set_example.arn),
            principal_id=Token.as_string(aws_identitystore_group_example.group_id),
            principal_type="GROUP",
            target_id="123456789012",
            target_type="AWS_ACCOUNT"
        )
        aws_ssoadmin_managed_policy_attachment_example =
        SsoadminManagedPolicyAttachment(self, "example_4",
            depends_on=[aws_ssoadmin_account_assignment_example],
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(example.arns), ["0"])),
            managed_policy_arn="arn:aws:iam::aws:policy/AlexaForBusinessDeviceSetup",
            permission_set_arn=Token.as_string(aws_ssoadmin_permission_set_example.arn)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssoadmin_managed_policy_attachment_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `instance_arn` - (Required, Forces new resource) The Amazon Resource Name (ARN) of the SSO Instance.
* `permission_set_arn` - (Required, Forces new resource) The Amazon Resource Name (ARN) of the Permission Set that the admin wants to grant the principal access to.
* `principal_id` - (Required, Forces new resource) An identifier for an object in SSO, such as a user or group. PrincipalIds are GUIDs (For example, `f81d4fae-7dec-11d0-a765-00a0c91e6bf6`).
* `principal_type` - (Required, Forces new resource) The entity type for which the assignment will be created. Valid values: `USER`, `GROUP`.
* `target_id` - (Required, Forces new resource) An AWS account identifier, typically a 10-12 digit string.
* `target_type` - (Optional, Forces new resource) The entity type for which the assignment will be created. Valid values: `AWS_ACCOUNT`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The identifier of the Account Assignment i.e., `principal_id`, `principal_type`, `target_id`, `target_type`, `permission_set_arn`, `instance_arn` separated by commas (`,`).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `5m`)
- `delete` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SSO Account Assignments using the `principal_id`, `principal_type`, `target_id`, `target_type`, `permission_set_arn`, `instance_arn` separated by commas (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ssoadmin_account_assignment import SsoadminAccountAssignment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SsoadminAccountAssignment.generate_config_for_import(self, "example", "f81d4fae-7dec-11d0-a765-00a0c91e6bf6,GROUP,1234567890,AWS_ACCOUNT,arn:aws:sso:::permissionSet/ssoins-0123456789abcdef/ps-0123456789abcdef,arn:aws:sso:::instance/ssoins-0123456789abcdef")
```

Using `terraform import`, import SSO Account Assignments using the `principal_id`, `principal_type`, `target_id`, `target_type`, `permission_set_arn`, `instance_arn` separated by commas (`,`). For example:

```console
% terraform import aws_ssoadmin_account_assignment.example f81d4fae-7dec-11d0-a765-00a0c91e6bf6,GROUP,1234567890,AWS_ACCOUNT,arn:aws:sso:::permissionSet/ssoins-0123456789abcdef/ps-0123456789abcdef,arn:aws:sso:::instance/ssoins-0123456789abcdef
```

<!-- cache-key: cdktf-0.20.8 input-92515b633334bcc4a3ebf73668c0382588fb749b26743442c083c4cf3575380b -->