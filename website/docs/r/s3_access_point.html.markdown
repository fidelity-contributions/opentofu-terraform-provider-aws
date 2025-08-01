---
subcategory: "S3 Control"
layout: "aws"
page_title: "AWS: aws_s3_access_point"
description: |-
  Manages an S3 Access Point.
---

# Resource: aws_s3_access_point

Provides a resource to manage an S3 Access Point.

~> **NOTE on Access Points and Access Point Policies:** Terraform provides both a standalone [Access Point Policy](s3control_access_point_policy.html) resource and an Access Point resource with a resource policy defined in-line. You cannot use an Access Point with in-line resource policy in conjunction with an Access Point Policy resource. Doing so will cause a conflict of policies and will overwrite the access point's resource policy.

-> Advanced usage: To use a custom API endpoint for this Terraform resource, use the [`s3control` endpoint provider configuration](/docs/providers/aws/index.html#s3control), not the `s3` endpoint provider configuration.

-> This resource can be used with s3 directory buckets. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets.html) for more information.

## Example Usage

### AWS Partition General Purpose Bucket

```terraform
resource "aws_s3_bucket" "example" {
  bucket = "example"
}

resource "aws_s3_access_point" "example" {
  bucket = aws_s3_bucket.example.id
  name   = "example"
}
```

### S3 on Outposts Bucket

```terraform
resource "aws_s3control_bucket" "example" {
  bucket = "example"
}

resource "aws_s3_access_point" "example" {
  bucket = aws_s3control_bucket.example.arn
  name   = "example"

  # VPC must be specified for S3 on Outposts
  vpc_configuration {
    vpc_id = aws_vpc.example.id
  }
}

resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}
```

### AWS Partition Directory Bucket

```terraform
data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_s3_directory_bucket" "example" {
  bucket = "example--zoneId--x-s3"
  location {
    name = data.aws_availability_zones.available.zone_ids[0]
  }
}

resource "aws_s3_access_point" "example" {
  bucket = aws_s3_directory_bucket.test.bucket
  name   = "example--zoneId--xa-s3"
}
```

## Argument Reference

The following arguments are required:

* `bucket` - (Required) Name of an AWS Partition S3 General Purpose Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
* `name` - (Required) Name you want to assign to this access point. See the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-access-points.html?icmpid=docs_amazons3_console#access-points-names) for naming conditions.

The following arguments are optional:

* `account_id` - (Optional) AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the Terraform AWS provider.
* `bucket_account_id` - (Optional) AWS account ID associated with the S3 bucket associated with this access point.
* `policy` - (Optional) Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `aws_s3control_access_point_policy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
* `public_access_block_configuration` - (Optional) Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `tags` - (Optional) Map of tags to assign to the bucket. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `vpc_configuration` - (Optional) Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.

### public_access_block_configuration Configuration Block

The following arguments are optional:

* `block_public_acls` - (Optional) Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `true`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
    * PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.
    * PUT Object calls fail if the request includes a public ACL.
    * PUT Bucket calls fail if the request includes a public ACL.
* `block_public_policy` - (Optional) Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `true`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
    * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
* `ignore_public_acls` - (Optional) Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `true`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
    * Ignore all public ACLs on buckets in this account and any objects that they contain.
* `restrict_public_buckets` - (Optional) Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `true`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
    * Only the bucket owner and AWS Services can access buckets with public policies.

### vpc_configuration Configuration Block

The following arguments are required:

* `vpc_id` - (Required)  This access point will only allow connections from the specified VPC ID.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `alias` - Alias of the S3 Access Point.
* `arn` - ARN of the S3 Access Point.
* `domain_name` - DNS domain name of the S3 Access Point in the format _`name`_-_`account_id`_.s3-accesspoint._region_.amazonaws.com.
Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
* `endpoints` - VPC endpoints for the S3 Access Point.
* `has_public_access_policy` - Indicates whether this access point currently has a policy that allows public access.
* `id` - For Access Point of an AWS Partition S3 Bucket, the AWS account ID and access point name separated by a colon (`:`). For S3 on Outposts Bucket, the ARN of the Access Point.
* `network_origin` - Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import this resource using the `account_id` and `name` separated by a colon (`:`) for Access Points associated with an AWS Partition S3 Bucket or the ARN for Access Points associated with an S3 on Outposts Bucket. For example:

Import using the `account_id` and `name` separated by a colon (`:`) for Access Points associated with an AWS Partition S3 Bucket:

```terraform
import {
  to = aws_s3_access_point.example
  id = "123456789012:example"
}
```

Import using the ARN for Access Points associated with an S3 on Outposts Bucket:

```terraform
import {
  to = aws_s3_access_point.example
  id = "arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-1234567890123456/accesspoint/example"
}
```

**Using `terraform import` to import.** For example:

Import using the `account_id` and `name` separated by a colon (`:`) for Access Points associated with an AWS Partition S3 Bucket:

```console
% terraform import aws_s3_access_point.example 123456789012:example
```

Import using the ARN for Access Points associated with an S3 on Outposts Bucket:

```console
% terraform import aws_s3_access_point.example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-1234567890123456/accesspoint/example
```
