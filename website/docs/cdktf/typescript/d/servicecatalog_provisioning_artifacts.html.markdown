---
subcategory: "Service Catalog"
layout: "aws"
page_title: "AWS: aws_servicecatalog_provisioning_artifacts"
description: |-
  Provides information on Service Catalog Provisioning Artifacts
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_servicecatalog_provisioning_artifacts

Lists the provisioning artifacts for the specified product.

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
import { DataAwsServicecatalogProvisioningArtifacts } from "./.gen/providers/aws/data-aws-servicecatalog-provisioning-artifacts";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsServicecatalogProvisioningArtifacts(this, "example", {
      productId: "prod-yakog5pdriver",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `productId` - (Required) Product identifier.

The following arguments are optional:

* `acceptLanguage` - (Optional) Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `provisioningArtifactDetails` - List with information about the provisioning artifacts. See details below.

### provisioning_artifact_details

* `active` - Indicates whether the product version is active.
* `createdTime` - The UTC time stamp of the creation time.
* `description` - The description of the provisioning artifact.
* `guidance` - Information set by the administrator to provide guidance to end users about which provisioning artifacts to use.
* `id` - The identifier of the provisioning artifact.
* `name` - The name of the provisioning artifact.
* `type` - The type of provisioning artifact.

<!-- cache-key: cdktf-0.20.8 input-70d3d0ef3ce0371f0258d9d9db5e9876444b44930ec3530badd4b8fe021fee18 -->