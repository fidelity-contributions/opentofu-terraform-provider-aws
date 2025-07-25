---
subcategory: "FSx"
layout: "aws"
page_title: "AWS: aws_fsx_data_repository_association"
description: |-
  Manages a FSx for Lustre Data Repository Association.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_fsx_data_repository_association

Manages a FSx for Lustre Data Repository Association. See [Linking your file system to an S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html) for more information.

~> **NOTE:** Data Repository Associations are only compatible with AWS FSx for Lustre File Systems and `PERSISTENT_2` deployment type.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { FsxDataRepositoryAssociation } from "./.gen/providers/aws/fsx-data-repository-association";
import { FsxLustreFileSystem } from "./.gen/providers/aws/fsx-lustre-file-system";
import { S3Bucket } from "./.gen/providers/aws/s3-bucket";
import { S3BucketAcl } from "./.gen/providers/aws/s3-bucket-acl";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new FsxLustreFileSystem(this, "example", {
      deploymentType: "PERSISTENT_2",
      perUnitStorageThroughput: 125,
      storageCapacity: 1200,
      subnetIds: [Token.asString(awsSubnetExample.id)],
    });
    const awsS3BucketExample = new S3Bucket(this, "example_1", {
      bucket: "my-bucket",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3BucketExample.overrideLogicalId("example");
    const awsS3BucketAclExample = new S3BucketAcl(this, "example_2", {
      acl: "private",
      bucket: Token.asString(awsS3BucketExample.id),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsS3BucketAclExample.overrideLogicalId("example");
    const awsFsxDataRepositoryAssociationExample =
      new FsxDataRepositoryAssociation(this, "example_3", {
        dataRepositoryPath: "s3://${" + awsS3BucketExample.id + "}",
        fileSystemId: example.id,
        fileSystemPath: "/my-bucket",
        s3: {
          autoExportPolicy: {
            events: ["NEW", "CHANGED", "DELETED"],
          },
          autoImportPolicy: {
            events: ["NEW", "CHANGED", "DELETED"],
          },
        },
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsFsxDataRepositoryAssociationExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `batchImportMetaDataOnCreate` - (Optional) Set to true to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created. Defaults to `false`.
* `dataRepositoryPath` - (Required) The path to the Amazon S3 data repository that will be linked to the file system. The path must be an S3 bucket s3://myBucket/myPrefix/. This path specifies where in the S3 data repository files will be imported from or exported to. The same S3 bucket cannot be linked more than once to the same file system.
* `fileSystemId` - (Required) The ID of the Amazon FSx file system to on which to create a data repository association.
* `fileSystemPath` - (Required) A path on the file system that points to a high-level directory (such as `/ns1/`) or subdirectory (such as `/ns1/subdir/`) that will be mapped 1-1 with `dataRepositoryPath`. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path `/ns1/`, then you cannot link another data repository with file system path `/ns1/ns2`. This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
* `importedFileChunkSize` - (Optional) For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
* `s3` - (Optional) See the [`s3` configuration](#s3-arguments) block. Max of 1.
The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
* `deleteDataInFilesystem` - (Optional) Set to true to delete files from the file system upon deleting this data repository association. Defaults to `false`.
* `tags` - (Optional) A map of tags to assign to the data repository association. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

#### S3 arguments

* `autoExportPolicy` - (Optional) Specifies the type of updated objects that will be automatically exported from your file system to the linked S3 bucket. See the [`events` configuration](#events-arguments) block.
* `autoImportPolicy` - (Optional) Specifies the type of updated objects that will be automatically imported from the linked S3 bucket to your file system. See the [`events` configuration](#events-arguments) block.

#### Events arguments

* `events` - (Optional) A list of file event types to automatically export to your linked S3 bucket or import from the linked S3 bucket. Valid values are `NEW`, `CHANGED`, `DELETED`. Max of 3.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name of the file system.
* `id` - Identifier of the data repository association, e.g., `dra-12345678`
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `update` - (Default `10m`)
* `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import FSx Data Repository Associations using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { FsxDataRepositoryAssociation } from "./.gen/providers/aws/fsx-data-repository-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    FsxDataRepositoryAssociation.generateConfigForImport(
      this,
      "example",
      "dra-0b1cfaeca11088b10"
    );
  }
}

```

Using `terraform import`, import FSx Data Repository Associations using the `id`. For example:

```console
% terraform import aws_fsx_data_repository_association.example dra-0b1cfaeca11088b10
```

<!-- cache-key: cdktf-0.20.8 input-9ad72e70ffef7525be8136216b00f8f967766e2d6c5548dc98f2afdddf22c787 -->