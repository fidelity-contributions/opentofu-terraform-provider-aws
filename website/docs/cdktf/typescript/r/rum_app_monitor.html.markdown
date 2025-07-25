---
subcategory: "CloudWatch RUM"
layout: "aws"
page_title: "AWS: aws_rum_app_monitor"
description: |-
  Provides a CloudWatch RUM App Monitor resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_rum_app_monitor

Provides a CloudWatch RUM App Monitor resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RumAppMonitor } from "./.gen/providers/aws/rum-app-monitor";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new RumAppMonitor(this, "example", {
      domain: "localhost",
      name: "example",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the log stream.
* `appMonitorConfiguration` - (Optional) configuration data for the app monitor. See [app_monitor_configuration](#app_monitor_configuration) below.
* `cwLogEnabled` - (Optional) Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges. Default value is `false`.
* `customEvents` - (Optional) Specifies whether this app monitor allows the web client to define and send custom events. If you omit this parameter, custom events are `DISABLED`. See [custom_events](#custom_events) below.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### app_monitor_configuration

* `allowCookies` - (Optional) If you set this to `true`, RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
* `domain` - (Optional) The top-level internet domain name for which your application has administrative authority. Exactly one of `domain` or `domainList` must be specified.
* `domainList` - (Optional) A list of internet domain names for which your application has administrative authority. Exactly one of `domain` or `domainList` must be specified.
* `enableXray` - (Optional) If you set this to `true`, RUM enables X-Ray tracing for the user sessions that RUM samples. RUM adds an X-Ray trace header to allowed HTTP requests. It also records an X-Ray segment for allowed HTTP requests.
* `excludedPages` - (Optional) A list of URLs in your website or application to exclude from RUM data collection.
* `favoritePages` - (Optional) A list of pages in the CloudWatch RUM console that are to be displayed with a "favorite" icon.
* `guestRoleArn` - (Optional) The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
* `identityPoolId` - (Optional) The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
* `includedPages` - (Optional)  If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
* `sessionSampleRate` - (Optional) Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. Default value is `0.1`.
* `telemetries` - (Optional) An array that lists the types of telemetry data that this app monitor is to collect. Valid values are `errors`, `performance`, and `http`.

### custom_events

* `status` - (Optional) Specifies whether this app monitor allows the web client to define and send custom events. The default is for custom events to be `DISABLED`. Valid values are `DISABLED` and `ENABLED`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) specifying the app monitor.
* `id` - The CloudWatch RUM name as it is the identifier of a RUM.
* `appMonitorId` - The unique ID of the app monitor. Useful for JS templates.
* `cwLogGroup` - The name of the log group where the copies are stored.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cloudwatch RUM App Monitor using the `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RumAppMonitor } from "./.gen/providers/aws/rum-app-monitor";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    RumAppMonitor.generateConfigForImport(this, "example", "example");
  }
}

```

Using `terraform import`, import Cloudwatch RUM App Monitor using the `name`. For example:

```console
% terraform import aws_rum_app_monitor.example example
```

<!-- cache-key: cdktf-0.20.8 input-b79cc3763c3e90bb4872796d248f7220a9871eaa19f8fdfb9dd75fca3b2e1748 -->