---
subcategory: "EC2 Image Builder"
layout: "aws"
page_title: "AWS: aws_imagebuilder_image"
description: |-
    Manages an Image Builder Image
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_imagebuilder_image

Manages an Image Builder Image.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ImagebuilderImage } from "./.gen/providers/aws/imagebuilder-image";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new ImagebuilderImage(this, "example", {
      distributionConfigurationArn: Token.asString(
        awsImagebuilderDistributionConfigurationExample.arn
      ),
      imageRecipeArn: Token.asString(awsImagebuilderImageRecipeExample.arn),
      infrastructureConfigurationArn: Token.asString(
        awsImagebuilderInfrastructureConfigurationExample.arn
      ),
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `infrastructureConfigurationArn` - (Required) Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `containerRecipeArn` - (Optional) - Amazon Resource Name (ARN) of the container recipe.
* `distributionConfigurationArn` - (Optional) Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
* `enhancedImageMetadataEnabled` - (Optional) Whether additional information about the image being created is collected. Defaults to `true`.
* `executionRole` - (Optional) Amazon Resource Name (ARN) of the service-linked role to be used by Image Builder to [execute workflows](https://docs.aws.amazon.com/imagebuilder/latest/userguide/manage-image-workflows.html).
* `imageRecipeArn` - (Optional) Amazon Resource Name (ARN) of the image recipe.
* `imageTestsConfiguration` - (Optional) Configuration block with image tests configuration. Detailed below.
* `imageScanningConfiguration` - (Optional) Configuration block with image scanning configuration. Detailed below.
* `workflow` - (Optional) Configuration block with the workflow configuration. Detailed below.
* `tags` - (Optional) Key-value map of resource tags for the Image Builder Image. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### image_tests_configuration

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `imageTestsEnabled` - (Optional) Whether image tests are enabled. Defaults to `true`.
* `timeoutMinutes` - (Optional) Number of minutes before image tests time out. Valid values are between `60` and `1440`. Defaults to `720`.

### image_scanning_configuration

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `imageScanningEnabled` - (Optional) Indicates whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image. Defaults to `false`.
* `ecrConfiguration` - (Optional) Configuration block with ECR configuration. Detailed below.

### ecr_configuration

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `repositoryName` - (Optional) The name of the container repository that Amazon Inspector scans to identify findings for your container images.
* `containerTags` - (Optional) Set of tags for Image Builder to apply to the output container image that that Amazon Inspector scans.

### workflow

The following arguments are required:

* `workflowArn` - (Required) Amazon Resource Name (ARN) of the Image Builder Workflow.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `onFailure` - (Optional) The action to take if the workflow fails. Must be one of `CONTINUE` or `ABORT`.
* `parallelGroup` - (Optional) The parallel group in which to run a test Workflow.
* `parameter` - (Optional) Configuration block for the workflow parameters. Detailed below.

### parameter

The following arguments are required:

* `name` - (Required) The name of the Workflow parameter.
* `value` - (Required) The value of the Workflow parameter.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the image.
* `dateCreated` - Date the image was created.
* `platform` - Platform of the image.
* `osVersion` - Operating System version of the image.
* `outputResources` - List of objects with resources created by the image.
    * `amis` - Set of objects with each Amazon Machine Image (AMI) created.
        * `accountId` - Account identifier of the AMI.
        * `description` - Description of the AMI.
        * `image` - Identifier of the AMI.
        * `name` - Name of the AMI.
        * `region` - Region of the AMI.
    * `containers` - Set of objects with each container image created and stored in the output repository.
        * `image_uris` - Set of URIs for created containers.
        * `region` - Region of the container image.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `version` - Version of the image.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `60m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_imagebuilder_image` resources using the Amazon Resource Name (ARN). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ImagebuilderImage } from "./.gen/providers/aws/imagebuilder-image";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ImagebuilderImage.generateConfigForImport(
      this,
      "example",
      "arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1"
    );
  }
}

```

Using `terraform import`, import `aws_imagebuilder_image` resources using the Amazon Resource Name (ARN). For example:

```console
% terraform import aws_imagebuilder_image.example arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1
```

<!-- cache-key: cdktf-0.20.8 input-84945ccad1d221b0febf006ba9383b2ed6762cea24116932ed658f9d86911932 -->