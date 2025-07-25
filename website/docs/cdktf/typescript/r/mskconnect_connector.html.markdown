---
subcategory: "Managed Streaming for Kafka Connect"
layout: "aws"
page_title: "AWS: aws_mskconnect_connector"
description: |-
  Provides an Amazon MSK Connect Connector resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_mskconnect_connector

Provides an Amazon MSK Connect Connector resource.

## Example Usage

### Basic configuration

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { MskconnectConnector } from "./.gen/providers/aws/mskconnect-connector";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new MskconnectConnector(this, "example", {
      capacity: {
        autoscaling: {
          maxWorkerCount: 2,
          mcuCount: 1,
          minWorkerCount: 1,
          scaleInPolicy: {
            cpuUtilizationPercentage: 20,
          },
          scaleOutPolicy: {
            cpuUtilizationPercentage: 80,
          },
        },
      },
      connectorConfiguration: {
        "connector.class":
          "com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector",
        "tasks.max": "1",
        topics: "example",
      },
      kafkaCluster: {
        apacheKafkaCluster: {
          bootstrapServers: Token.asString(
            awsMskClusterExample.bootstrapBrokersTls
          ),
          vpc: {
            securityGroups: [Token.asString(awsSecurityGroupExample.id)],
            subnets: [example1.id, example2.id, example3.id],
          },
        },
      },
      kafkaClusterClientAuthentication: {
        authenticationType: "NONE",
      },
      kafkaClusterEncryptionInTransit: {
        encryptionType: "TLS",
      },
      kafkaconnectVersion: "2.7.1",
      name: "example",
      plugin: [
        {
          customPlugin: {
            arn: Token.asString(awsMskconnectCustomPluginExample.arn),
            revision: Token.asNumber(
              awsMskconnectCustomPluginExample.latestRevision
            ),
          },
        },
      ],
      serviceExecutionRoleArn: Token.asString(awsIamRoleExample.arn),
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `capacity` - (Required) Information about the capacity allocated to the connector. See [`capacity` Block](#capacity-block) for details.
* `connectorConfiguration` - (Required) A map of keys to values that represent the configuration for the connector.
* `kafkaCluster` - (Required) Specifies which Apache Kafka cluster to connect to. See [`kafkaCluster` Block](#kafka_cluster-block) for details.
* `kafkaClusterClientAuthentication` - (Required) Details of the client authentication used by the Apache Kafka cluster. See [`kafkaClusterClientAuthentication` Block](#kafka_cluster_client_authentication-block) for details.
* `kafkaClusterEncryptionInTransit` - (Required) Details of encryption in transit to the Apache Kafka cluster. See [`kafkaClusterEncryptionInTransit` Block](#kafka_cluster_encryption_in_transit-block) for details.
* `kafkaconnectVersion` - (Required) The version of Kafka Connect. It has to be compatible with both the Apache Kafka cluster's version and the plugins.
* `name` - (Required) The name of the connector.
* `plugin` - (Required) Specifies which plugins to use for the connector. See [`plugin` Block](#plugin-block) for details.
* `serviceExecutionRoleArn` - (Required) The Amazon Resource Name (ARN) of the IAM role used by the connector to access the Amazon Web Services resources that it needs. The types of resources depends on the logic of the connector. For example, a connector that has Amazon S3 as a destination must have permissions that allow it to write to the S3 destination bucket.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) A summary description of the connector.
* `logDelivery` - (Optional) Details about log delivery. See [`logDelivery` Block](#log_delivery-block) for details.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `workerConfiguration` - (Optional) Specifies which worker configuration to use with the connector. See [`workerConfiguration` Block](#worker_configuration-block) for details.

### `capacity` Block

The `capacity` configuration block supports the following arguments:

* `autoscaling` - (Optional) Information about the auto scaling parameters for the connector. See [`autoscaling` Block](#autoscaling-block) for details.
* `provisionedCapacity` - (Optional) Details about a fixed capacity allocated to a connector. See [`provisionedCapacity` Block](#provisioned_capacity-block) for details.

### `autoscaling` Block

The `autoscaling` configuration block supports the following arguments:

* `maxWorkerCount` - (Required) The maximum number of workers allocated to the connector.
* `mcuCount` - (Optional) The number of microcontroller units (MCUs) allocated to each connector worker. Valid values: `1`, `2`, `4`, `8`. The default value is `1`.
* `minWorkerCount` - (Required) The minimum number of workers allocated to the connector.
* `scaleInPolicy` - (Optional) The scale-in policy for the connector. See [`scaleInPolicy` Block](#scale_in_policy-block) for details.
* `scaleOutPolicy` - (Optional) The scale-out policy for the connector. See [`scaleOutPolicy` Block](#scale_out_policy-block) for details.

### `scaleInPolicy` Block

The `scaleInPolicy` configuration block supports the following arguments:

* `cpuUtilizationPercentage` - (Required) Specifies the CPU utilization percentage threshold at which you want connector scale in to be triggered.

### `scaleOutPolicy` Block

The `scaleOutPolicy` configuration block supports the following arguments:

* `cpuUtilizationPercentage` - (Required) The CPU utilization percentage threshold at which you want connector scale out to be triggered.

### `provisionedCapacity` Block

The `provisionedCapacity` configuration block supports the following arguments:

* `mcuCount` - (Optional) The number of microcontroller units (MCUs) allocated to each connector worker. Valid values: `1`, `2`, `4`, `8`. The default value is `1`.
* `workerCount` - (Required) The number of workers that are allocated to the connector.

### `kafkaCluster` Block

The `kafkaCluster` configuration block supports the following arguments:

* `apacheKafkaCluster` - (Required) The Apache Kafka cluster to which the connector is connected. See [`apacheKafkaCluster` Block](#apache_kafka_cluster-block) for details.

### `apacheKafkaCluster` Block

The `apacheKafkaCluster` configuration block supports the following arguments:

* `bootstrapServers` - (Required) The bootstrap servers of the cluster.
* `vpc` - (Required) Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster. See [`vpc` Block](#vpc-block) for details.

### `vpc` Block

The `vpc` configuration block supports the following arguments:

* `securityGroups` - (Required) The security groups for the connector.
* `subnets` - (Required) The subnets for the connector.

### `kafkaClusterClientAuthentication` Block

The `kafkaClusterClientAuthentication` configuration block supports the following arguments:

* `authenticationType` - (Optional) The type of client authentication used to connect to the Apache Kafka cluster. Valid values: `IAM`, `NONE`. A value of `NONE` means that no client authentication is used. The default value is `NONE`.

### `kafkaClusterEncryptionInTransit` Block

The `kafkaClusterEncryptionInTransit` configuration block supports the following arguments:

* `encryptionType` - (Optional) The type of encryption in transit to the Apache Kafka cluster. Valid values: `PLAINTEXT`, `TLS`. The default values is `PLAINTEXT`.

### `logDelivery` Block

The `logDelivery` configuration block supports the following arguments:

* `workerLogDelivery` - (Required) The workers can send worker logs to different destination types. This configuration specifies the details of these destinations. See [`workerLogDelivery` Block](#worker_log_delivery-block) for details.

### `workerLogDelivery` Block

The `workerLogDelivery` configuration block supports the following arguments:

* `cloudwatchLogs` - (Optional) Details about delivering logs to Amazon CloudWatch Logs. See [`cloudwatchLogs` Block](#cloudwatch_logs-block) for details.
* `firehose` - (Optional) Details about delivering logs to Amazon Kinesis Data Firehose. See [`firehose` Block](#firehose-block) for details.
* `s3` - (Optional) Details about delivering logs to Amazon S3. See [`s3` Block](#s3-block) for deetails.

### `cloudwatchLogs` Block

The `cloudwatchLogs` configuration block supports the following arguments:

* `enabled` - (Optional) Whether log delivery to Amazon CloudWatch Logs is enabled.
* `logGroup` - (Required) The name of the CloudWatch log group that is the destination for log delivery.

### `firehose` Block

The `firehose` configuration block supports the following arguments:

* `deliveryStream` - (Optional) The name of the Kinesis Data Firehose delivery stream that is the destination for log delivery.
* `enabled` - (Required) Specifies whether connector logs get delivered to Amazon Kinesis Data Firehose.

### `s3` Block

The `s3` configuration block supports the following arguments:

* `bucket` - (Optional) The name of the S3 bucket that is the destination for log delivery.
* `enabled` - (Required) Specifies whether connector logs get sent to the specified Amazon S3 destination.
* `prefix` - (Optional) The S3 prefix that is the destination for log delivery.

### `plugin` Block

The `plugin` configuration block supports the following argumens:

* `customPlugin` - (Required) Details about a custom plugin. See [`customPlugin` Block](#custom_plugin-block) for details.

### `customPlugin` Block

The `customPlugin` configuration block supports the following arguments:

* `arn` - (Required) The Amazon Resource Name (ARN) of the custom plugin.
* `revision` - (Required) The revision of the custom plugin.

### `workerConfiguration` Block

The `workerConfiguration` configuration block supports the following arguments:

* `arn` - (Required) The Amazon Resource Name (ARN) of the worker configuration.
* `revision` - (Required) The revision of the worker configuration.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the connector.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `version` - The current version of the connector.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `20m`)
* `update` - (Default `20m`)
* `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import MSK Connect Connector using the connector's `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { MskconnectConnector } from "./.gen/providers/aws/mskconnect-connector";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    MskconnectConnector.generateConfigForImport(
      this,
      "example",
      "arn:aws:kafkaconnect:eu-central-1:123456789012:connector/example/264edee4-17a3-412e-bd76-6681cfc93805-3"
    );
  }
}

```

Using `terraform import`, import MSK Connect Connector using the connector's `arn`. For example:

```console
% terraform import aws_mskconnect_connector.example 'arn:aws:kafkaconnect:eu-central-1:123456789012:connector/example/264edee4-17a3-412e-bd76-6681cfc93805-3'
```

<!-- cache-key: cdktf-0.20.8 input-4ee0d5808327d2a40eab5d24e8c22103cc72f071acfec6cc118484dbe97264a5 -->