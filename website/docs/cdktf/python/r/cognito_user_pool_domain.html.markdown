---
subcategory: "Cognito IDP (Identity Provider)"
layout: "aws"
page_title: "AWS: aws_cognito_user_pool_domain"
description: |-
  Provides a Cognito User Pool Domain resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cognito_user_pool_domain

Provides a Cognito User Pool Domain resource.

## Example Usage

### Amazon Cognito domain

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cognito_user_pool import CognitoUserPool
from imports.aws.cognito_user_pool_domain import CognitoUserPoolDomain
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = CognitoUserPool(self, "example",
            name="example-pool"
        )
        CognitoUserPoolDomain(self, "main",
            domain="example-domain",
            user_pool_id=example.id
        )
```

### Custom Cognito domain

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cognito_user_pool import CognitoUserPool
from imports.aws.cognito_user_pool_domain import CognitoUserPoolDomain
from imports.aws.data_aws_route53_zone import DataAwsRoute53Zone
from imports.aws.route53_record import Route53Record
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = CognitoUserPool(self, "example",
            name="example-pool"
        )
        main = CognitoUserPoolDomain(self, "main",
            certificate_arn=cert.arn,
            domain="auth.example.com",
            user_pool_id=example.id
        )
        data_aws_route53_zone_example = DataAwsRoute53Zone(self, "example_2",
            name="example.com"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_route53_zone_example.override_logical_id("example")
        Route53Record(self, "auth-cognito-A",
            alias=Route53RecordAlias(
                evaluate_target_health=False,
                name=main.cloudfront_distribution,
                zone_id=main.cloudfront_distribution_zone_id
            ),
            name=main.domain,
            type="A",
            zone_id=Token.as_string(data_aws_route53_zone_example.zone_id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `domain` - (Required) For custom domains, this is the fully-qualified domain name, such as auth.example.com. For Amazon Cognito prefix domains, this is the prefix alone, such as auth.
* `user_pool_id` - (Required) The user pool ID.
* `certificate_arn` - (Optional) The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
* `managed_login_version` - (Optional) A version number that indicates the state of managed login for your domain. Valid values: `1` for hosted UI (classic), `2` for the newer managed login with the branding designer.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `aws_account_id` - The AWS account ID for the user pool owner.
* `cloudfront_distribution` - The Amazon CloudFront endpoint (e.g. `dpp0gtxikpq3y.cloudfront.net`) that you use as the target of the alias that you set up with your Domain Name Service (DNS) provider.
* `cloudfront_distribution_arn` - The URL of the CloudFront distribution. This is required to generate the ALIAS `aws_route53_record`
* `cloudfront_distribution_zone_id` - The Route 53 hosted zone ID of the CloudFront distribution.
* `s3_bucket` - The S3 bucket where the static files for this domain are stored.
* `version` - The app version.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cognito User Pool Domains using the `domain`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cognito_user_pool_domain import CognitoUserPoolDomain
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CognitoUserPoolDomain.generate_config_for_import(self, "main", "auth.example.org")
```

Using `terraform import`, import Cognito User Pool Domains using the `domain`. For example:

```console
% terraform import aws_cognito_user_pool_domain.main auth.example.org
```

<!-- cache-key: cdktf-0.20.8 input-1e2f05961159c89a97a1134a2ca27e5d3a840b2e2d98f56ba6ba61db0d56ee4d -->