---
subcategory: "EFS (Elastic File System)"
layout: "aws"
page_title: "AWS: aws_efs_access_point"
description: |-
  Provides an Elastic File System (EFS) access point.
---

# Resource: aws_efs_access_point

Provides an Elastic File System (EFS) access point.

## Example Usage

```terraform
resource "aws_efs_access_point" "test" {
  file_system_id = aws_efs_file_system.foo.id
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `file_system_id` - (Required) ID of the file system for which the access point is intended.
* `posix_user` - (Optional) Operating system user and group applied to all file system requests made using the access point. [Detailed](#posix_user) below.
* `root_directory`- (Optional) Directory on the Amazon EFS file system that the access point provides access to. [Detailed](#root_directory) below.
* `tags` - (Optional) Key-value mapping of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### posix_user

* `gid` - (Required) POSIX group ID used for all file system operations using this access point.
* `secondary_gids` - (Optional) Secondary POSIX group IDs used for all file system operations using this access point.
* `uid` - (Required) POSIX user ID used for all file system operations using this access point.

### root_directory

The access point exposes the specified file system path as the root directory of your file system to applications using the access point. NFS clients using the access point can only access data in the access point's RootDirectory and it's subdirectories.

* `creation_info` - (Optional) POSIX IDs and permissions to apply to the access point's Root Directory. See [Creation Info](#creation_info) below.
* `path` - (Optional) Path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide `creation_info`.

### creation_info

If the `path` specified does not exist, EFS creates the root directory using the `creation_info` settings when a client connects to an access point.

* `owner_gid` - (Required) POSIX group ID to apply to the `root_directory`.
* `owner_uid` - (Required) POSIX user ID to apply to the `root_directory`.
* `permissions` - (Required) POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the access point.
* `file_system_arn` - ARN of the file system.
* `id` - ID of the access point.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the EFS access points using the `id`. For example:

```terraform
import {
  to = aws_efs_access_point.test
  id = "fsap-52a643fb"
}
```

Using `terraform import`, import the EFS access points using the `id`. For example:

```console
% terraform import aws_efs_access_point.test fsap-52a643fb
```
