---
subcategory: "ElastiCache"
layout: "aws"
page_title: "AWS: aws_elasticache_serverless_cache"
description: |-
  Provides an ElastiCache Serverless Cache resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_elasticache_serverless_cache

Provides an ElastiCache Serverless Cache resource which manages memcached, redis or valkey.

## Example Usage

### Memcached Serverless

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.elasticache_serverless_cache import ElasticacheServerlessCache
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ElasticacheServerlessCache(self, "example",
            cache_usage_limits=[ElasticacheServerlessCacheCacheUsageLimits(
                data_storage=[ElasticacheServerlessCacheCacheUsageLimitsDataStorage(
                    maximum=10,
                    unit="GB"
                )
                ],
                ecpu_per_second=[ElasticacheServerlessCacheCacheUsageLimitsEcpuPerSecond(
                    maximum=5000
                )
                ]
            )
            ],
            description="Test Server",
            engine="memcached",
            kms_key_id=test.arn,
            major_engine_version="1.6",
            name="example",
            security_group_ids=[Token.as_string(aws_security_group_test.id)],
            subnet_ids=Token.as_list(Fn.lookup_nested(aws_subnet_test, ["*", "id"]))
        )
```

### Redis OSS Serverless

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.elasticache_serverless_cache import ElasticacheServerlessCache
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ElasticacheServerlessCache(self, "example",
            cache_usage_limits=[ElasticacheServerlessCacheCacheUsageLimits(
                data_storage=[ElasticacheServerlessCacheCacheUsageLimitsDataStorage(
                    maximum=10,
                    unit="GB"
                )
                ],
                ecpu_per_second=[ElasticacheServerlessCacheCacheUsageLimitsEcpuPerSecond(
                    maximum=5000
                )
                ]
            )
            ],
            daily_snapshot_time="09:00",
            description="Test Server",
            engine="redis",
            kms_key_id=test.arn,
            major_engine_version="7",
            name="example",
            security_group_ids=[Token.as_string(aws_security_group_test.id)],
            snapshot_retention_limit=1,
            subnet_ids=Token.as_list(Fn.lookup_nested(aws_subnet_test, ["*", "id"]))
        )
```

### Valkey Serverless

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.elasticache_serverless_cache import ElasticacheServerlessCache
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ElasticacheServerlessCache(self, "example",
            cache_usage_limits=[ElasticacheServerlessCacheCacheUsageLimits(
                data_storage=[ElasticacheServerlessCacheCacheUsageLimitsDataStorage(
                    maximum=10,
                    unit="GB"
                )
                ],
                ecpu_per_second=[ElasticacheServerlessCacheCacheUsageLimitsEcpuPerSecond(
                    maximum=5000
                )
                ]
            )
            ],
            daily_snapshot_time="09:00",
            description="Test Server",
            engine="valkey",
            kms_key_id=test.arn,
            major_engine_version="7",
            name="example",
            security_group_ids=[Token.as_string(aws_security_group_test.id)],
            snapshot_retention_limit=1,
            subnet_ids=Token.as_list(Fn.lookup_nested(aws_subnet_test, ["*", "id"]))
        )
```

## Argument Reference

The following arguments are required:

* `engine` - (Required) Name of the cache engine to be used for this cache cluster. Valid values are `memcached`, `redis` or `valkey`.
* `name` - (Required) The Cluster name which serves as a unique identifier to the serverless cache

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `cache_usage_limits` - (Optional) Sets the cache usage limits for storage and ElastiCache Processing Units for the cache. See [`cache_usage_limits` Block](#cache_usage_limits-block) for details.
* `daily_snapshot_time` - (Optional) The daily time that snapshots will be created from the new serverless cache. Only supported for engine types `"redis"` or `"valkey"`. Defaults to `0`.
* `description` - (Optional) User-provided description for the serverless cache. The default is NULL.
* `kms_key_id` - (Optional) ARN of the customer managed key for encrypting the data at rest. If no KMS key is provided, a default service key is used.
* `major_engine_version` - (Optional) The version of the cache engine that will be used to create the serverless cache.
  See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html) in the AWS Documentation for supported versions.
* `security_group_ids` - (Optional) A list of the one or more VPC security groups to be associated with the serverless cache. The security group will authorize traffic access for the VPC end-point (private-link). If no other information is given this will be the VPC’s Default Security Group that is associated with the cluster VPC end-point.
* `snapshot_arns_to_restore` - (Optional, Redis only) The list of ARN(s) of the snapshot that the new serverless cache will be created from. Available for Redis only.
* `snapshot_retention_limit` - (Optional, Redis only) The number of snapshots that will be retained for the serverless cache that is being created. As new snapshots beyond this limit are added, the oldest snapshots will be deleted on a rolling basis. Available for Redis only.
* `subnet_ids` - (Optional) A list of the identifiers of the subnets where the VPC endpoint for the serverless cache will be deployed. All the subnetIds must belong to the same VPC.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `user_group_id` - (Optional) The identifier of the UserGroup to be associated with the serverless cache. Available for Redis only. Default is NULL.

### `cache_usage_limits` Block

The `cache_usage_limits` configuration block supports the following arguments:

* `data_storage` - The maximum data storage limit in the cache, expressed in Gigabytes. See [`data_storage` Block](#data_storage-block) for details.
* `ecpu_per_second` - The configuration for the number of ElastiCache Processing Units (ECPU) the cache can consume per second. See [`ecpu_per_second` Block](#ecpu_per_second-block) for details.

### `data_storage` Block

The `data_storage` configuration block supports the following arguments:

* `minimum` - The lower limit for data storage the cache is set to use. Must be between 1 and 5,000.
* `maximum` - The upper limit for data storage the cache is set to use. Must be between 1 and 5,000.
* `unit` - The unit that the storage is measured in, in GB.

### `ecpu_per_second` Block

The `ecpu_per_second` configuration block supports the following arguments:

* `minimum` - The minimum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.
* `maximum` - The maximum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the serverless cache.
* `create_time` - Timestamp of when the serverless cache was created.
* `endpoint` - Represents the information required for client programs to connect to a cache node. See [`endpoint` Block](#endpoint-block) for details.
* `full_engine_version` - The name and version number of the engine the serverless cache is compatible with.
* `major_engine_version` - The version number of the engine the serverless cache is compatible with.
* `reader_endpoint` - Represents the information required for client programs to connect to a cache node. See [`reader_endpoint` Block](#reader_endpoint-block) for details.
* `status` - The current status of the serverless cache. The allowed values are CREATING, AVAILABLE, DELETING, CREATE-FAILED and MODIFYING.

### `endpoint` Block

The `endpoint` configuration block exports the following attributes:

* `address` - The DNS hostname of the cache node.
* `port` - The port number that the cache engine is listening on. Set as integer.

### `reader_endpoint` Block

The `reader_endpoint` configuration block exports the following attributes:

* `address` - The DNS hostname of the cache node.
* `port` - The port number that the cache engine is listening on. Set as integer.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `40m`)
- `update` - (Default `80m`)
- `delete` - (Default `40m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ElastiCache Serverless Cache using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.elasticache_serverless_cache import ElasticacheServerlessCache
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ElasticacheServerlessCache.generate_config_for_import(self, "myCluster", "my_cluster")
```

Using `terraform import`, import ElastiCache Serverless Cache using the `name`. For example:

```console
% terraform import aws_elasticache_serverless_cache.my_cluster my_cluster
```

<!-- cache-key: cdktf-0.20.8 input-fd7e6a4928845cb5b7620232e3794fd67db6eb15b312647121bf7594dd418875 -->