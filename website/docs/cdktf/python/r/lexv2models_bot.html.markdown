---
subcategory: "Lex V2 Models"
layout: "aws"
page_title: "AWS: aws_lexv2models_bot"
description: |-
  Terraform resource for managing an AWS Lex V2 Models Bot.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lexv2models_bot

Terraform resource for managing an AWS Lex V2 Models Bot.

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
from imports.aws.iam_role import IamRole
from imports.aws.lexv2_models_bot import Lexv2ModelsBot
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = IamRole(self, "example",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "lexv2.amazonaws.com"
                        },
                        "Sid": ""
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            name="example",
            tags={
                "created_by": "aws"
            }
        )
        aws_lexv2_models_bot_example = Lexv2ModelsBot(self, "example_1",
            data_privacy=[{
                "child_directed": False
            }
            ],
            description="Example description",
            idle_session_ttl_in_seconds=60,
            name="example",
            role_arn=example.arn,
            tags={
                "foo": "bar"
            },
            type="Bot"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_lexv2_models_bot_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `name` - Name of the bot. The bot name must be unique in the account that creates the bot. Type String. Length Constraints: Minimum length of 1. Maximum length of 100.
* `data_privacy` - Provides information on additional privacy protections Amazon Lex should use with the bot's data. See [`data_privacy`](#data-privacy)
* `idle_session_ttl_in_seconds` - Time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot. You can specify between 60 (1 minute) and 86,400 (24 hours) seconds.
* `role_arn` - ARN of an IAM role that has permission to access the bot.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `members` - List of bot members in a network to be created. See [`bot_members`](#bot-members).
* `tags` - List of tags to add to the bot. You can only add tags when you create a bot.
* `type` - Type of a bot to create. Possible values are `"Bot"` and `"BotNetwork"`.
* `description` - Description of the bot. It appears in lists to help you identify a particular bot.
* `test_bot_alias_tags` - List of tags to add to the test alias for a bot. You can only add tags when you create a bot.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Unique identifier for a particular bot.

### Data Privacy

* `child_directed` (Required) -  For each Amazon Lex bot created with the Amazon Lex Model Building Service, you must specify whether your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to the Children's Online Privacy Protection Act (COPPA) by specifying true or false in the childDirected field.

### Bot Members

* `alias_id` (Required) - Alias ID of a bot that is a member of this network of bots.
* `alias_name` (Required) - Alias name of a bot that is a member of this network of bots.
* `id` (Required) - Unique ID of a bot that is a member of this network of bots.
* `name` (Required) - Unique name of a bot that is a member of this network of bots.
* `version` (Required) - Version of a bot that is a member of this network of bots.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Lex V2 Models Bot using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lexv2_models_bot import Lexv2ModelsBot
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Lexv2ModelsBot.generate_config_for_import(self, "example", "bot-id-12345678")
```

Using `terraform import`, import Lex V2 Models Bot using the `id`. For example:

```console
% terraform import aws_lexv2models_bot.example bot-id-12345678
```

<!-- cache-key: cdktf-0.20.8 input-c5bdaf469481c5ff54e192168c1f568be10936293114614e9b100dc991489bdc -->