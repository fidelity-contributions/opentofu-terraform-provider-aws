---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_gateway_response"
description: |-
  Provides an API Gateway Gateway Response for a REST API Gateway.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_api_gateway_gateway_response

Provides an API Gateway Gateway Response for a REST API Gateway.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayGatewayResponse } from "./.gen/providers/aws/api-gateway-gateway-response";
import { ApiGatewayRestApi } from "./.gen/providers/aws/api-gateway-rest-api";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const main = new ApiGatewayRestApi(this, "main", {
      name: "MyDemoAPI",
    });
    new ApiGatewayGatewayResponse(this, "test", {
      responseParameters: {
        "gatewayresponse.header.Authorization": "'Basic'",
      },
      responseTemplates: {
        "application/json": '{\\"message\\":$context.error.messageString}',
      },
      responseType: "UNAUTHORIZED",
      restApiId: main.id,
      statusCode: "401",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `restApiId` - (Required) String identifier of the associated REST API.
* `responseType` - (Required) Response type of the associated GatewayResponse.
* `statusCode` - (Optional) HTTP status code of the Gateway Response.
* `responseTemplates` - (Optional) Map of templates used to transform the response body.
* `responseParameters` - (Optional) Map of parameters (paths, query strings and headers) of the Gateway Response.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_api_gateway_gateway_response` using `REST-API-ID/RESPONSE-TYPE`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayGatewayResponse } from "./.gen/providers/aws/api-gateway-gateway-response";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ApiGatewayGatewayResponse.generateConfigForImport(
      this,
      "example",
      "12345abcde/UNAUTHORIZED"
    );
  }
}

```

Using `terraform import`, import `aws_api_gateway_gateway_response` using `REST-API-ID/RESPONSE-TYPE`. For example:

```console
% terraform import aws_api_gateway_gateway_response.example 12345abcde/UNAUTHORIZED
```

<!-- cache-key: cdktf-0.20.8 input-a14260c5eb0c0d600fe4e89fb6f7dce663f8f56bb8fa71be9037b76a5f392f2e -->