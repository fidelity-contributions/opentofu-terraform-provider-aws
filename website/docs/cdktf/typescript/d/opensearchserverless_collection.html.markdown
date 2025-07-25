---
subcategory: "OpenSearch Serverless"
layout: "aws"
page_title: "AWS: aws_opensearchserverless_collection"
description: |-
  Terraform data source for managing an AWS OpenSearch Serverless Collection.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_opensearchserverless_collection

Terraform data source for managing an AWS OpenSearch Serverless Collection.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsOpensearchserverlessCollection } from "./.gen/providers/aws/data-aws-opensearchserverless-collection";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsOpensearchserverlessCollection(this, "example", {
      name: "example",
    });
  }
}

```

## Argument Reference

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `id` - (Optional) ID of the collection.
* `name` - (Optional) Name of the collection.

~> Exactly one of `id` or `name` is required.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the collection.
* `collectionEndpoint` - Collection-specific endpoint used to submit index, search, and data upload requests to an OpenSearch Serverless collection.
* `createdDate` - Date the Collection was created.
* `dashboardEndpoint` - Collection-specific endpoint used to access OpenSearch Dashboards.
* `description` - Description of the collection.
* `failureCode` - A failure code associated with the collection.
* `failureReason` - A failure reason associated with the collection.
* `kmsKeyArn` - The ARN of the Amazon Web Services KMS key used to encrypt the collection.
* `lastModifiedDate` - Date the Collection was last modified.
* `standbyReplicas` - Indicates whether standby replicas should be used for a collection.
* `tags` - A map of tags to assign to the collection.
* `type` - Type of collection.

<!-- cache-key: cdktf-0.20.8 input-6073cb0e683097ba3f5c043592179682637605534db023aa3bf046c05569f14e -->