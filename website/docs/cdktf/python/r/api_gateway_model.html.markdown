---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_model"
description: |-
  Provides a Model for a REST API Gateway.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_api_gateway_model

Provides a Model for a REST API Gateway.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.api_gateway_model import ApiGatewayModel
from imports.aws.api_gateway_rest_api import ApiGatewayRestApi
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        my_demo_api = ApiGatewayRestApi(self, "MyDemoAPI",
            description="This is my API for demonstration purposes",
            name="MyDemoAPI"
        )
        ApiGatewayModel(self, "MyDemoModel",
            content_type="application/json",
            description="a JSON schema",
            name="user",
            rest_api_id=my_demo_api.id,
            schema=Token.as_string(
                Fn.jsonencode({
                    "type": "object"
                }))
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `rest_api_id` - (Required) ID of the associated REST API
* `name` - (Required) Name of the model
* `description` - (Optional) Description of the model
* `content_type` - (Required) Content type of the model
* `schema` - (Required) Schema of the model in a JSON form

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the model

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_api_gateway_model` using `REST-API-ID/NAME`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.api_gateway_model import ApiGatewayModel
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ApiGatewayModel.generate_config_for_import(self, "example", "12345abcde/example")
```

Using `terraform import`, import `aws_api_gateway_model` using `REST-API-ID/NAME`. For example:

```console
% terraform import aws_api_gateway_model.example 12345abcde/example
```

<!-- cache-key: cdktf-0.20.8 input-4e52221db4a67e14520da48490df7cffd8ac899f9dcbe5133045bd95a2f5338b -->