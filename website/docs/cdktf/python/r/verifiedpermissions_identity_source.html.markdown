---
subcategory: "Verified Permissions"
layout: "aws"
page_title: "AWS: aws_verifiedpermissions_identity_source"
description: |-
  Terraform resource for managing an AWS Verified Permissions Identity Source.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_verifiedpermissions_identity_source

Terraform resource for managing an AWS Verified Permissions Identity Source.

## Example Usage

### Cognito User Pool Configuration Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cognito_user_pool import CognitoUserPool
from imports.aws.cognito_user_pool_client import CognitoUserPoolClient
from imports.aws.verifiedpermissions_identity_source import VerifiedpermissionsIdentitySource
from imports.aws.verifiedpermissions_policy_store import VerifiedpermissionsPolicyStore
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = CognitoUserPool(self, "example",
            name="example"
        )
        aws_cognito_user_pool_client_example = CognitoUserPoolClient(self, "example_1",
            explicit_auth_flows=["ADMIN_NO_SRP_AUTH"],
            name="example",
            user_pool_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_cognito_user_pool_client_example.override_logical_id("example")
        aws_verifiedpermissions_policy_store_example =
        VerifiedpermissionsPolicyStore(self, "example_2",
            validation_settings=[VerifiedpermissionsPolicyStoreValidationSettings(
                mode="STRICT"
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_verifiedpermissions_policy_store_example.override_logical_id("example")
        aws_verifiedpermissions_identity_source_example =
        VerifiedpermissionsIdentitySource(self, "example_3",
            configuration=[VerifiedpermissionsIdentitySourceConfiguration(
                cognito_user_pool_configuration=[VerifiedpermissionsIdentitySourceConfigurationCognitoUserPoolConfiguration(
                    client_ids=[Token.as_string(aws_cognito_user_pool_client_example.id)],
                    user_pool_arn=example.arn
                )
                ]
            )
            ],
            policy_store_id=Token.as_string(aws_verifiedpermissions_policy_store_example.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_verifiedpermissions_identity_source_example.override_logical_id("example")
```

### OpenID Connect Configuration Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.verifiedpermissions_identity_source import VerifiedpermissionsIdentitySource
from imports.aws.verifiedpermissions_policy_store import VerifiedpermissionsPolicyStore
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = VerifiedpermissionsPolicyStore(self, "example",
            validation_settings=[VerifiedpermissionsPolicyStoreValidationSettings(
                mode="STRICT"
            )
            ]
        )
        aws_verifiedpermissions_identity_source_example =
        VerifiedpermissionsIdentitySource(self, "example_1",
            configuration=[VerifiedpermissionsIdentitySourceConfiguration(
                open_id_connect_configuration=[VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfiguration(
                    entity_id_prefix="MyOIDCProvider",
                    group_configuration=[VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationGroupConfiguration(
                        group_claim="groups",
                        group_entity_type="MyCorp::UserGroup"
                    )
                    ],
                    issuer="https://auth.example.com",
                    token_selection=[VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationTokenSelection(
                        access_token_only=[VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfigurationTokenSelectionAccessTokenOnly(
                            audiences=["https://myapp.example.com"],
                            principal_id_claim="sub"
                        )
                        ]
                    )
                    ]
                )
                ]
            )
            ],
            policy_store_id=example.id,
            principal_entity_type="MyCorp::User"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_verifiedpermissions_identity_source_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `policy_store_id` - (Required) Specifies the ID of the policy store in which you want to store this identity source.
* `configuration`- (Required) Specifies the details required to communicate with the identity provider (IdP) associated with this identity source. See [Configuration](#configuration) below.
* `principal_entity_type`- (Optional) Specifies the namespace and data type of the principals generated for identities authenticated by the new identity source.

### Configuration

* `cognito_user_pool_configuration` - (Required) Specifies the configuration details of an Amazon Cognito user pool that Verified Permissions can use as a source of authenticated identities as entities. See [Cognito User Pool Configuration](#cognito-user-pool-configuration) below.
* `open_id_connect_configuration` - (Required) Specifies the configuration details of an OpenID Connect (OIDC) identity provider, or identity source, that Verified Permissions can use to generate entities from authenticated identities. See [Open ID Connect Configuration](#open-id-connect-configuration) below.

#### Cognito User Pool Configuration

* `user_pool_arn` - (Required) The Amazon Resource Name (ARN) of the Amazon Cognito user pool that contains the identities to be authorized.
* `client_ids` - (Optional) The unique application client IDs that are associated with the specified Amazon Cognito user pool.
* `group_configuration` - (Optional) The type of entity that a policy store maps to groups from an Amazon Cognito user pool identity source. See [Group Configuration](#group-configuration) below.

#### Group Configuration

* `group_entity_type` - (Required) The name of the schema entity type that's mapped to the user pool group. Defaults to `AWS::CognitoGroup`.

#### Open ID Connect Configuration

* `issuer` - (Required) The issuer URL of an OIDC identity provider. This URL must have an OIDC discovery endpoint at the path `.well-known/openid-configuration`.
* `token_selection` - (Required) The token type that you want to process from your OIDC identity provider. Your policy store can process either identity (ID) or access tokens from a given OIDC identity source. See [Token Selection](#token-selection) below.
* `entity_id_prefix` - (Optional) A descriptive string that you want to prefix to user entities from your OIDC identity provider.
* `group_configuration` - (Optional) The type of entity that a policy store maps to groups from an Amazon Cognito user pool identity source. See [Group Configuration](#open-id-group-configuration) below.

#### Token Selection

* `access_token_only` - (Optional) The OIDC configuration for processing access tokens. See [Access Token Only](#access-token-only) below.
* `identity_token_only` - (Optional) The OIDC configuration for processing identity (ID) tokens. See [Identity Token Only](#identity-token-only) below.

#### Access Token Only

* `audiences` - (Optional) The access token aud claim values that you want to accept in your policy store.
* `principal_id_claim` - (Optional) The claim that determines the principal in OIDC access tokens.

#### Identity Token Only

* `client_ids` - (Optional) The ID token audience, or client ID, claim values that you want to accept in your policy store from an OIDC identity provider.
* `group_entity_type` - (Optional) The claim that determines the principal in OIDC access tokens.

#### Open ID Group Configuration

* `group_claim` - (Required) The token claim that you want Verified Permissions to interpret as group membership. For example, `groups`.
* `group_entity_type` - (Required) The policy store entity type that you want to map your users' group claim to. For example, `MyCorp::UserGroup`. A group entity type is an entity that can have a user entity type as a member.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `policy_id` - The Policy ID of the policy.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Verified Permissions Identity Source using the `policy_store_id:identity_source_id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.verifiedpermissions_identity_source import VerifiedpermissionsIdentitySource
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VerifiedpermissionsIdentitySource.generate_config_for_import(self, "example", "policy-store-id-12345678:identity-source-id-12345678")
```

Using `terraform import`, import Verified Permissions Identity Source using the `policy_store_id:identity_source_id`. For example:

```console
% terraform import aws_verifiedpermissions_identity_source.example policy-store-id-12345678:identity-source-id-12345678
```

<!-- cache-key: cdktf-0.20.8 input-40fc3d53236ce6505da2ee2eb9ac375dab9980f28583f61953bb22abc101d905 -->