---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_eip_association"
description: |-
  Provides an AWS EIP Association
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_eip_association

Provides an AWS EIP Association as a top level resource, to associate and disassociate Elastic IPs from AWS Instances and Network Interfaces.

~> **NOTE:** Do not use this resource to associate an EIP to `aws_lb` or `aws_nat_gateway` resources. Instead use the `allocationId` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.

~> **NOTE:** `aws_eip_association` is useful in scenarios where EIPs are either pre-existing or distributed to customers or users and therefore cannot be changed.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Eip } from "./.gen/providers/aws/eip";
import { EipAssociation } from "./.gen/providers/aws/eip-association";
import { Instance } from "./.gen/providers/aws/instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new Eip(this, "example", {
      domain: "vpc",
    });
    const web = new Instance(this, "web", {
      ami: "ami-21f78e11",
      availabilityZone: "us-west-2a",
      instanceType: "t2.micro",
      tags: {
        Name: "HelloWorld",
      },
    });
    new EipAssociation(this, "eip_assoc", {
      allocationId: example.id,
      instanceId: web.id,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `allocationId` - (Optional, Forces new resource) ID of the associated Elastic IP.
  This argument is required despite being optional at the resource level due to legacy support for EC2-Classic networking.
* `allowReassociation` - (Optional, Forces new resource) Whether to allow an Elastic IP address to be re-associated.
  Defaults to `true`.
* `instanceId` - (Optional, Forces new resource) ID of the instance.
  The instance must have exactly one attached network interface.
  You can specify either the instance ID or the network interface ID, but not both.
* `networkInterfaceId` - (Optional, Forces new resource) ID of the network interface.
  If the instance has more than one network interface, you must specify a network interface ID.
  You can specify either the instance ID or the network interface ID, but not both.
* `privateIpAddress` - (Optional, Forces new resource) Primary or secondary private IP address to associate with the Elastic IP address.
  If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
* `publicIp` - (Optional, Forces new resource, **Deprecated** since [EC2-Classic netwworking has retired](https://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/)) Address of the associated Elastic IP.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID that represents the association of the Elastic IP address with an instance.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EIP Assocations using their association IDs. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EipAssociation } from "./.gen/providers/aws/eip-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    EipAssociation.generateConfigForImport(this, "test", "eipassoc-ab12c345");
  }
}

```

Using `terraform import`, import EIP Assocations using their association IDs. For example:

```console
% terraform import aws_eip_association.test eipassoc-ab12c345
```

<!-- cache-key: cdktf-0.20.8 input-5271e2228d8190613bd6a1d712122d540cdb8448c79d8b1eccb670f9815733b7 -->