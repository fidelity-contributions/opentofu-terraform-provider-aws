---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "AWS: aws_s3_bucket_intelligent_tiering_configuration"
description: |-
  Provides an S3 Intelligent-Tiering configuration resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_s3_bucket_intelligent_tiering_configuration

Provides an [S3 Intelligent-Tiering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html) configuration resource.

-> This resource cannot be used with S3 directory buckets.

## Example Usage

### Add intelligent tiering configuration for entire S3 bucket

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_bucket_intelligent_tiering_configuration import S3BucketIntelligentTieringConfiguration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = S3Bucket(self, "example",
            bucket="example"
        )
        S3BucketIntelligentTieringConfiguration(self, "example-entire-bucket",
            bucket=example.id,
            name="EntireBucket",
            tiering=[S3BucketIntelligentTieringConfigurationTiering(
                access_tier="DEEP_ARCHIVE_ACCESS",
                days=180
            ), S3BucketIntelligentTieringConfigurationTiering(
                access_tier="ARCHIVE_ACCESS",
                days=125
            )
            ]
        )
```

### Add intelligent tiering configuration with S3 object filter

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_bucket_intelligent_tiering_configuration import S3BucketIntelligentTieringConfiguration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = S3Bucket(self, "example",
            bucket="example"
        )
        S3BucketIntelligentTieringConfiguration(self, "example-filtered",
            bucket=example.id,
            filter=S3BucketIntelligentTieringConfigurationFilter(
                prefix="documents/",
                tags={
                    "class": "blue",
                    "priority": "high"
                }
            ),
            name="ImportantBlueDocuments",
            status="Disabled",
            tiering=[S3BucketIntelligentTieringConfigurationTiering(
                access_tier="ARCHIVE_ACCESS",
                days=125
            )
            ]
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `bucket` - (Required) Name of the bucket this intelligent tiering configuration is associated with.
* `name` - (Required) Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
* `status` - (Optional) Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
* `filter` - (Optional) Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
* `tiering` - (Required) S3 Intelligent-Tiering storage class tiers of the configuration (documented below).

The `filter` configuration supports the following:

* `prefix` - (Optional) Object key name prefix that identifies the subset of objects to which the configuration applies.
* `tags` - (Optional) All of these tags must exist in the object's tag set in order for the configuration to apply.

The `tiering` configuration supports the following:

* `access_tier` - (Required) S3 Intelligent-Tiering access tier. Valid values: `ARCHIVE_ACCESS`, `DEEP_ARCHIVE_ACCESS`.
* `days` - (Required) Number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import S3 bucket intelligent tiering configurations using `bucket:name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket_intelligent_tiering_configuration import S3BucketIntelligentTieringConfiguration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        S3BucketIntelligentTieringConfiguration.generate_config_for_import(self, "myBucketEntireBucket", "my-bucket:EntireBucket")
```

Using `terraform import`, import S3 bucket intelligent tiering configurations using `bucket:name`. For example:

```console
% terraform import aws_s3_bucket_intelligent_tiering_configuration.my-bucket-entire-bucket my-bucket:EntireBucket
```

<!-- cache-key: cdktf-0.20.8 input-080fd992ca26c97c18fa263c805f1286bc3acb805fee64480665442bdf06189d -->