---
subcategory: "Serverless Application Repository"
layout: "aws"
page_title: "AWS: aws_serverlessapplicationrepository_application"
description: |-
  Get information on a AWS Serverless Application Repository application
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_serverlessapplicationrepository_application

Use this data source to get information about an AWS Serverless Application Repository application. For example, this can be used to determine the required `capabilities` for an application.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_serverlessapplicationrepository_application import DataAwsServerlessapplicationrepositoryApplication
from imports.aws.serverlessapplicationrepository_cloudformation_stack import ServerlessapplicationrepositoryCloudformationStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsServerlessapplicationrepositoryApplication(self, "example",
            application_id="arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication"
        )
        aws_serverlessapplicationrepository_cloudformation_stack_example =
        ServerlessapplicationrepositoryCloudformationStack(self, "example_1",
            application_id=Token.as_string(example.application_id),
            capabilities=Token.as_list(example.required_capabilities),
            name="Example",
            semantic_version=Token.as_string(example.semantic_version)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_serverlessapplicationrepository_cloudformation_stack_example.override_logical_id("example")
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `application_id` - (Required) ARN of the application.
* `semantic_version` - (Optional) Requested version of the application. By default, retrieves the latest version.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `application_id` - ARN of the application.
* `name` - Name of the application.
* `required_capabilities` - A list of capabilities describing the permissions needed to deploy the application.
* `source_code_url` - URL pointing to the source code of the application version.
* `template_url` - URL pointing to the Cloud Formation template for the application version.

<!-- cache-key: cdktf-0.20.8 input-944b82ac5dd958af9069727776a1201fed695ed3d6cfd0589c43e546316c2d39 -->