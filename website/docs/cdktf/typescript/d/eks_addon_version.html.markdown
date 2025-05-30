---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "AWS: aws_eks_addon_version"
description: |-
  Retrieve information about versions of an EKS add-on
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_eks_addon_version

Retrieve information about a specific EKS add-on version compatible with an EKS cluster version.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformOutput, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEksAddonVersion } from "./.gen/providers/aws/data-aws-eks-addon-version";
import { EksAddon } from "./.gen/providers/aws/eks-addon";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const defaultVar = new DataAwsEksAddonVersion(this, "default", {
      addonName: "vpc-cni",
      kubernetesVersion: example.version,
    });
    const latest = new DataAwsEksAddonVersion(this, "latest", {
      addonName: "vpc-cni",
      kubernetesVersion: example.version,
      mostRecent: true,
    });
    const cdktfTerraformOutputDefault = new TerraformOutput(this, "default_2", {
      value: defaultVar.version,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    cdktfTerraformOutputDefault.overrideLogicalId("default");
    const cdktfTerraformOutputLatest = new TerraformOutput(this, "latest_3", {
      value: latest.version,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    cdktfTerraformOutputLatest.overrideLogicalId("latest");
    new EksAddon(this, "vpc_cni", {
      addonName: "vpc-cni",
      addonVersion: Token.asString(latest.version),
      clusterName: example.name,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `addonName` – (Required) Name of the EKS add-on. The name must match one of
  the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
* `kubernetesVersion` – (Required) Version of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
* `mostRecent` - (Optional) Determines if the most recent or default version of the addon should be returned.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - Name of the add-on
* `version` - Version of the EKS add-on.

<!-- cache-key: cdktf-0.20.8 input-7c6f3385681241f89ada51ee4c5d85e4e13250c68f32fe30e511ca8a29a53eed -->