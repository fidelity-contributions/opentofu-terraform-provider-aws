---
subcategory: "DynamoDB"
layout: "aws"
page_title: "AWS: aws_dynamodb_table_replica"
description: |-
  Provides a DynamoDB table replica resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dynamodb_table_replica

Provides a DynamoDB table replica resource for [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html).

~> **Note:** Use `lifecycle` [`ignore_changes`](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#ignore_changes) for `replica` in the associated [aws_dynamodb_table](/docs/providers/aws/r/dynamodb_table.html) configuration.

~> **Note:** Do not use the `replica` configuration block of [aws_dynamodb_table](/docs/providers/aws/r/dynamodb_table.html) together with this resource as the two configuration options are mutually exclusive.

## Example Usage

### Basic Example

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DynamodbTable } from "./.gen/providers/aws/dynamodb-table";
import { DynamodbTableReplicaA } from "./.gen/providers/aws/dynamodb-table-replica";
import { AwsProvider } from "./.gen/providers/aws/provider";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const main = new AwsProvider(this, "aws", {
      alias: "main",
      region: "us-west-2",
    });
    const alt = new AwsProvider(this, "aws_1", {
      alias: "alt",
      region: "us-east-2",
    });
    const example = new DynamodbTable(this, "example", {
      attribute: [
        {
          name: "BrodoBaggins",
          type: "S",
        },
      ],
      billingMode: "PAY_PER_REQUEST",
      hashKey: "BrodoBaggins",
      lifecycle: {
        ignoreChanges: [replica],
      },
      name: "TestTable",
      provider: main,
      streamEnabled: true,
      streamViewType: "NEW_AND_OLD_IMAGES",
    });
    const awsDynamodbTableReplicaExample = new DynamodbTableReplicaA(
      this,
      "example_3",
      {
        globalTableArn: example.arn,
        provider: alt,
        tags: {
          Name: "IZPAWS",
          Pozo: "Amargo",
        },
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDynamodbTableReplicaExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `globalTableArn` - (Required) ARN of the _main_ or global table which this resource will replicate.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `kmsKeyArn` - (Optional, Forces new resource) ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
* `deletionProtectionEnabled` - (Optional) Whether deletion protection is enabled (true) or disabled (false) on the table replica.
* `pointInTimeRecovery` - (Optional) Whether to enable Point In Time Recovery for the table replica. Default is `false`.
* `tableClassOverride` - (Optional, Forces new resource) Storage class of the table replica. Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS`. If not used, the table replica will use the same class as the global table.
* `tags` - (Optional) Map of tags to populate on the created table. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the table replica.
* `id` - Name of the table and region of the main global table joined with a semicolon (_e.g._, `TableName:us-east-1`).
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `20m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DynamoDB table replicas using the `table-name:main-region`. For example:

~> **Note:** When importing, use the region where the initial or _main_ global table resides, _not_ the region of the replica.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DynamodbTableReplicaA } from "./.gen/providers/aws/dynamodb-table-replica";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DynamodbTableReplicaA.generateConfigForImport(
      this,
      "example",
      "TestTable:us-west-2"
    );
  }
}

```

Using `terraform import`, import DynamoDB table replicas using the `table-name:main-region`. For example:

~> **Note:** When importing, use the region where the initial or _main_ global table resides, _not_ the region of the replica.

```console
% terraform import aws_dynamodb_table_replica.example TestTable:us-west-2
```

<!-- cache-key: cdktf-0.20.8 input-2bd54a095e0c9e3e17d78ac35ddaab469bc8af424f0966b983b6eb918f983a80 -->