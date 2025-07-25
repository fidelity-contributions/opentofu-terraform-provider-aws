---
subcategory: "Lightsail"
layout: "aws"
page_title: "AWS: aws_lightsail_lb_https_redirection_policy"
description: |-
  Manages HTTPS redirection for a Lightsail Load Balancer.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lightsail_lb_https_redirection_policy

Manages HTTPS redirection for a Lightsail Load Balancer.

Use this resource to configure automatic redirection of HTTP traffic to HTTPS on a Lightsail Load Balancer. A valid certificate must be attached to the load balancer before enabling HTTPS redirection.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LightsailLb } from "./.gen/providers/aws/lightsail-lb";
import { LightsailLbCertificate } from "./.gen/providers/aws/lightsail-lb-certificate";
import { LightsailLbCertificateAttachment } from "./.gen/providers/aws/lightsail-lb-certificate-attachment";
import { LightsailLbHttpsRedirectionPolicy } from "./.gen/providers/aws/lightsail-lb-https-redirection-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new LightsailLb(this, "example", {
      healthCheckPath: "/",
      instancePort: Token.asNumber("80"),
      name: "example-load-balancer",
      tags: {
        foo: "bar",
      },
    });
    const awsLightsailLbCertificateExample = new LightsailLbCertificate(
      this,
      "example_1",
      {
        domainName: "example.com",
        lbName: example.id,
        name: "example-load-balancer-certificate",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLightsailLbCertificateExample.overrideLogicalId("example");
    const awsLightsailLbCertificateAttachmentExample =
      new LightsailLbCertificateAttachment(this, "example_2", {
        certificateName: Token.asString(awsLightsailLbCertificateExample.name),
        lbName: example.name,
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLightsailLbCertificateAttachmentExample.overrideLogicalId("example");
    const awsLightsailLbHttpsRedirectionPolicyExample =
      new LightsailLbHttpsRedirectionPolicy(this, "example_3", {
        enabled: true,
        lbName: example.name,
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLightsailLbHttpsRedirectionPolicyExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `enabled` - (Required) Whether to enable HTTP to HTTPS redirection. `true` to activate HTTP to HTTPS redirection or `false` to deactivate HTTP to HTTPS redirection.
* `lbName` - (Required) Name of the load balancer to which you want to enable HTTP to HTTPS redirection.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Name used for this load balancer (matches `lbName`).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_lightsail_lb_https_redirection_policy` using the `lbName` attribute. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LightsailLbHttpsRedirectionPolicy } from "./.gen/providers/aws/lightsail-lb-https-redirection-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    LightsailLbHttpsRedirectionPolicy.generateConfigForImport(
      this,
      "example",
      "example-load-balancer"
    );
  }
}

```

Using `terraform import`, import `aws_lightsail_lb_https_redirection_policy` using the `lbName` attribute. For example:

```console
% terraform import aws_lightsail_lb_https_redirection_policy.example example-load-balancer
```

<!-- cache-key: cdktf-0.20.8 input-3bccf4d90ec56a5692ae9dbadac4bd524f62d9188018be78eb8ed39d2ee8f3a4 -->