---
subcategory: "WorkSpaces"
layout: "aws"
page_title: "AWS: aws_workspaces_image"
description: |-
  Get information about Workspaces image.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_workspaces_image

Use this data source to get information about a Workspaces image.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsWorkspacesImage } from "./.gen/providers/aws/data-aws-workspaces-image";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsWorkspacesImage(this, "example", {
      imageId: "wsi-ten5h0y19",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `imageId` - (Required) ID of the image.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `name` - The name of the image.
* `description` - The description of the image.
* `os` - The operating system that the image is running.
* `requiredTenancy` - Specifies whether the image is running on dedicated hardware. When Bring Your Own License (BYOL) is enabled, this value is set to DEDICATED. For more information, see [Bring Your Own Windows Desktop Images](https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html).
* `state` - The status of the image.

<!-- cache-key: cdktf-0.20.8 input-4cd58afbd9c55bd1e014ef19622a2d45727d3d26a3ab4c1a116e5b541d802978 -->