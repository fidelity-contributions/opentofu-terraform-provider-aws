---
subcategory: "DataSync"
layout: "aws"
page_title: "AWS: aws_datasync_location_nfs"
description: |-
  Manages an AWS DataSync NFS Location
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_datasync_location_nfs

Manages an NFS Location within AWS DataSync.

~> **NOTE:** The DataSync Agents must be available before creating this resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DatasyncLocationNfs } from "./.gen/providers/aws/datasync-location-nfs";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DatasyncLocationNfs(this, "example", {
      onPremConfig: {
        agentArns: [Token.asString(awsDatasyncAgentExample.arn)],
      },
      serverHostname: "nfs.example.com",
      subdirectory: "/exported/path",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `mountOptions` - (Optional) Configuration block containing mount options used by DataSync to access the NFS Server.
* `onPremConfig` - (Required) Configuration block containing information for connecting to the NFS File System.
* `serverHostname` - (Required) Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
* `subdirectory` - (Required) Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
* `tags` - (Optional) Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### mount_options Argument Reference

The `mountOptions` configuration block supports the following arguments:

* `version` - (Optional) The specific NFS version that you want DataSync to use for mounting your NFS share. Valid values: `AUTOMATIC`, `NFS3`, `NFS4_0` and `NFS4_1`. Default: `AUTOMATIC`

### on_prem_config Argument Reference

The `onPremConfig` configuration block supports the following arguments:

* `agentArns` - (Required) List of Amazon Resource Names (ARNs) of the DataSync Agents used to connect to the NFS server.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Amazon Resource Name (ARN) of the DataSync Location.
* `arn` - Amazon Resource Name (ARN) of the DataSync Location.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_datasync_location_nfs` using the DataSync Task Amazon Resource Name (ARN). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DatasyncLocationNfs } from "./.gen/providers/aws/datasync-location-nfs";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DatasyncLocationNfs.generateConfigForImport(
      this,
      "example",
      "arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567"
    );
  }
}

```

Using `terraform import`, import `aws_datasync_location_nfs` using the DataSync Task Amazon Resource Name (ARN). For example:

```console
% terraform import aws_datasync_location_nfs.example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```

<!-- cache-key: cdktf-0.20.8 input-f240fbe4a449dfba36c7ea5e857026af75a0722e9f076d806f78dde5203da54d -->