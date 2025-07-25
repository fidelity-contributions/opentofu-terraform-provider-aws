---
subcategory: "Elasticsearch"
layout: "aws"
page_title: "AWS: aws_elasticsearch_domain_saml_options"
description: |-
  Terraform resource for managing SAML authentication options for an AWS Elasticsearch Domain.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_elasticsearch_domain_saml_options

Manages SAML authentication options for an AWS Elasticsearch Domain.

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
from imports.aws.elasticsearch_domain import ElasticsearchDomain
from imports.aws.elasticsearch_domain_saml_options import ElasticsearchDomainSamlOptions
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = ElasticsearchDomain(self, "example",
            cluster_config=ElasticsearchDomainClusterConfig(
                instance_type="r4.large.elasticsearch"
            ),
            domain_name="example",
            elasticsearch_version="1.5",
            snapshot_options=ElasticsearchDomainSnapshotOptions(
                automated_snapshot_start_hour=23
            ),
            tags={
                "Domain": "TestDomain"
            }
        )
        aws_elasticsearch_domain_saml_options_example =
        ElasticsearchDomainSamlOptions(self, "example_1",
            domain_name=example.domain_name,
            saml_options=ElasticsearchDomainSamlOptionsSamlOptions(
                enabled=True,
                idp=ElasticsearchDomainSamlOptionsSamlOptionsIdp(
                    entity_id="https://example.com",
                    metadata_content=Token.as_string(Fn.file("./saml-metadata.xml"))
                )
            )
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_elasticsearch_domain_saml_options_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `domain_name` - (Required) Name of the domain.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `saml_options` - (Optional) The SAML authentication options for an AWS Elasticsearch Domain.

### saml_options

* `enabled` - (Required) Whether SAML authentication is enabled.
* `idp` - (Optional) Information from your identity provider.
* `master_backend_role` - (Optional) This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
* `master_user_name` - (Optional) This username from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
* `roles_key` - (Optional) Element of the SAML assertion to use for backend roles. Default is roles.
* `session_timeout_minutes` - (Optional) Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
* `subject_key` - (Optional) Custom SAML attribute to use for user names. Default is an empty string - `""`. This will cause Elasticsearch to use the `NameID` element of the `Subject`, which is the default location for name identifiers in the SAML specification.

#### idp

* `entity_id` - (Required) The unique Entity ID of the application in SAML Identity Provider.
* `metadata_content` - (Required) The Metadata of the SAML application in xml format.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The name of the domain the SAML options are associated with.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Elasticsearch domains using the `domain_name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.elasticsearch_domain_saml_options import ElasticsearchDomainSamlOptions
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ElasticsearchDomainSamlOptions.generate_config_for_import(self, "example", "domain_name")
```

Using `terraform import`, import Elasticsearch domains using the `domain_name`. For example:

```console
% terraform import aws_elasticsearch_domain_saml_options.example domain_name
```

<!-- cache-key: cdktf-0.20.8 input-cec75046e058ee7a4e585f5f83ccab6fc30b664e1b2dc01fdbe4cef49dca6dcf -->