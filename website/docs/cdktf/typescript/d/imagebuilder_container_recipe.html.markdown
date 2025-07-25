---
subcategory: "EC2 Image Builder"
layout: "aws"
page_title: "AWS: aws_imagebuilder_container_recipe"
description: |-
    Provides details about an Image Builder Container Recipe
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_imagebuilder_container_recipe

Provides details about an Image builder Container Recipe.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsImagebuilderContainerRecipe } from "./.gen/providers/aws/data-aws-imagebuilder-container-recipe";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsImagebuilderContainerRecipe(this, "example", {
      arn: "arn:aws:imagebuilder:us-east-1:aws:container-recipe/example/1.0.0",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `arn` - (Required) ARN of the container recipe.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `component` - List of objects with components for the container recipe.
    * `componentArn` - ARN of the Image Builder Component.
    * `parameter` - Set of parameters that are used to configure the component.
        * `name` - Name of the component parameter.
        * `value` - Value of the component parameter.
* `containerType` - Type of the container.
* `dateCreated` - Date the container recipe was created.
* `description` - Description of the container recipe.
* `dockerfileTemplateData` - Dockerfile template used to build the image.
* `encrypted` - Flag that indicates if the target container is encrypted.
* `instanceConfiguration` - List of objects with instance configurations for building and testing container images.
    * `blockDeviceMapping` - Set of objects with block device mappings for the instance configuration.
        * `deviceName` - Name of the device. For example, `/dev/sda` or `/dev/xvdb`.
        * `ebs` - Single list of object with Elastic Block Storage (EBS) block device mapping settings.
            * `deleteOnTermination` - Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
            * `encrypted` - Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
            * `iops` - Number of Input/Output (I/O) operations per second to provision for an `io1` or `io2` volume.
            * `kmsKeyId` - ARN of the Key Management Service (KMS) Key for encryption.
            * `snapshotId` - Identifier of the EC2 Volume Snapshot.
            * `throughput` - For GP3 volumes only. The throughput in MiB/s that the volume supports.
            * `volumeSize` - Size of the volume, in GiB.
            * `volumeType` - Type of the volume. For example, `gp2` or `io2`.
        * `noDevice` - Whether to remove a mapping from the parent image.
        * `virtualName` - Virtual device name. For example, `ephemeral0`. Instance store volumes are numbered starting from 0.
    * `image` - AMI ID of the base image for container build and test instance.
* `kmsKeyId` - KMS key used to encrypt the container image.
* `name` - Name of the container recipe.
* `owner` - Owner of the container recipe.
* `parentImage` - Base image for the container recipe.
* `platform` - Platform of the container recipe.
* `tags` - Key-value map of resource tags for the container recipe.
* `targetRepository` - Destination repository for the container image.
    * `repositoryName` - Name of the container repository where the output container image is stored. The name is prefixed by the repository location.
    * `service` - Service in which this image is registered.
* `version` - Version of the container recipe.
* `workingDirectory` - Working directory used during build and test workflows.

<!-- cache-key: cdktf-0.20.8 input-3908c3e50c991cc678c2d5ffb916c74c8b4fd8325d45d59b8e993e519c7fb856 -->