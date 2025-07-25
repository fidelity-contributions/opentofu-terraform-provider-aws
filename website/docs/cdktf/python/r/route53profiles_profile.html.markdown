---
subcategory: "Route 53 Profiles"
layout: "aws"
page_title: "AWS: aws_route53profiles_profile"
description: |-
  Terraform resource for managing an AWS Route 53 Profile.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_route53profiles_profile

Terraform resource for managing an AWS Route 53 Profile.

## Example Usage

### Empty Profile

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_profiles_profile import Route53ProfilesProfile
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Route53ProfilesProfile(self, "example",
            name="example",
            tags={
                "Environment": "dev"
            }
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the Profile.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Profile.
* `id` - ID of the Profile.
* `name` - Name of the Profile.
* `share_status` - Share status of the Profile.
* `status` - Status of the Profile.
* `status_message` - Status message of the Profile.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `read` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Route 53 Profiles Profile using the `example_id_arg`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.route53_profiles_profile import Route53ProfilesProfile
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Route53ProfilesProfile.generate_config_for_import(self, "example", "rp-12345678")
```

Using `terraform import`, import Route 53 Profiles Profile using the `example`. For example:

```console
% terraform import aws_route53profiles_profile.example rp-12345678
```

<!-- cache-key: cdktf-0.20.8 input-2a86de21f45a30228220544390d7382f7c3f500eb5806ac8e9fe1ba9a8ef322b -->