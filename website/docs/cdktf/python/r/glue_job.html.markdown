---
subcategory: "Glue"
layout: "aws"
page_title: "AWS: aws_glue_job"
description: |-
  Provides an Glue Job resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_glue_job

Provides a Glue Job resource.

-> Glue functionality, such as monitoring and logging of jobs, is typically managed with the `default_arguments` argument. See the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the Glue developer guide for additional information.

## Example Usage

### Python Glue Job

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_job import GlueJob
from imports.aws.iam_role import IamRole
from imports.aws.s3_object import S3Object
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        glue_job_role = IamRole(self, "glue_job_role",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "glue.amazonaws.com"
                        }
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            name="glue-job-role"
        )
        S3Object(self, "glue_etl_script",
            bucket=glue_scripts.id,
            key="jobs/etl_job.py",
            source="jobs/etl_job.py"
        )
        GlueJob(self, "etl_job",
            command=GlueJobCommand(
                name="glueetl",
                python_version="3",
                script_location="s3://${" + glue_scripts.bucket + "}/jobs/etl_job.py"
            ),
            connections=[example.name],
            default_arguments={
                "--continuous-log-logGroup": "/aws-glue/jobs",
                "--enable-auto-scaling": "true",
                "--enable-continuous-cloudwatch-log": "true",
                "--enable-continuous-log-filter": "true",
                "--enable-metrics": "",
                "--job-language": "python"
            },
            description="An example Glue ETL job",
            execution_class="STANDARD",
            execution_property=GlueJobExecutionProperty(
                max_concurrent_runs=1
            ),
            glue_version="5.0",
            max_retries=0,
            name="example-etl-job",
            notification_property=GlueJobNotificationProperty(
                notify_delay_after=3
            ),
            number_of_workers=2,
            role_arn=glue_job_role.arn,
            tags={
                "ManagedBy": "AWS"
            },
            timeout=2880,
            worker_type="G.1X"
        )
```

### Pythonshell Job

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_job import GlueJob
from imports.aws.iam_role import IamRole
from imports.aws.s3_object import S3Object
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        glue_job_role = IamRole(self, "glue_job_role",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "glue.amazonaws.com"
                        }
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            name="glue-job-role"
        )
        S3Object(self, "python_shell_script",
            bucket=glue_scripts.id,
            key="jobs/shell_job.py",
            source="jobs/shell_job.py"
        )
        GlueJob(self, "python_shell_job",
            command=GlueJobCommand(
                name="pythonshell",
                python_version="3.9",
                script_location="s3://${" + glue_scripts.bucket + "}/jobs/shell_job.py"
            ),
            connections=[example.name],
            default_arguments={
                "--continuous-log-logGroup": "/aws-glue/jobs",
                "--enable-continuous-cloudwatch-log": "true",
                "--job-language": "python",
                "library-set": "analytics"
            },
            description="An example Python shell job",
            execution_property=GlueJobExecutionProperty(
                max_concurrent_runs=1
            ),
            max_capacity=Token.as_number("0.0625"),
            max_retries=0,
            name="example-python-shell-job",
            role_arn=glue_job_role.arn,
            tags={
                "ManagedBy": "AWS"
            },
            timeout=2880
        )
```

### Ray Job

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_job import GlueJob
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlueJob(self, "example",
            command=GlueJobCommand(
                name="glueray",
                python_version="3.9",
                runtime="Ray2.4",
                script_location="s3://${" + aws_s3_bucket_example.bucket + "}/example.py"
            ),
            glue_version="4.0",
            name="example",
            role_arn=Token.as_string(aws_iam_role_example.arn),
            worker_type="Z.2X"
        )
```

### Scala Job

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_job import GlueJob
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlueJob(self, "example",
            command=GlueJobCommand(
                script_location="s3://${" + aws_s3_bucket_example.bucket + "}/example.scala"
            ),
            default_arguments={
                "--job-language": "scala"
            },
            name="example",
            role_arn=Token.as_string(aws_iam_role_example.arn)
        )
```

### Streaming Job

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_job import GlueJob
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlueJob(self, "example",
            command=GlueJobCommand(
                name="gluestreaming",
                script_location="s3://${" + aws_s3_bucket_example.bucket + "}/example.script"
            ),
            name="example streaming job",
            role_arn=Token.as_string(aws_iam_role_example.arn)
        )
```

### Enabling CloudWatch Logs and Metrics

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_log_group import CloudwatchLogGroup
from imports.aws.glue_job import GlueJob
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, command, name, roleArn):
        super().__init__(scope, name)
        example = CloudwatchLogGroup(self, "example",
            name="example",
            retention_in_days=14
        )
        aws_glue_job_example = GlueJob(self, "example_1",
            default_arguments={
                "--continuous-log-logGroup": example.name,
                "--enable-continuous-cloudwatch-log": "true",
                "--enable-continuous-log-filter": "true",
                "--enable-metrics": ""
            },
            command=command,
            name=name,
            role_arn=role_arn
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_glue_job_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `command` - (Required) The command of the job. Defined below.
* `connections` - (Optional) The list of connections used for this job.
* `default_arguments` - (Optional) The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
* `non_overridable_arguments` - (Optional) Non-overridable arguments for this job, specified as name-value pairs.
* `description` - (Optional) Description of the job.
* `execution_property` - (Optional) Execution property of the job. Defined below.
* `glue_version` - (Optional) The version of glue to use, for example "1.0". Ray jobs should set this to 4.0 or greater. For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
* `job_mode` - (Optional) Describes how a job was created. Valid values are `SCRIPT`, `NOTEBOOK` and `VISUAL`.
* `job_run_queuing_enabled` - (Optional) Specifies whether job run queuing is enabled for the job runs for this job. A value of true means job run queuing is enabled for the job runs. If false or not populated, the job runs will not be considered for queueing.
* `execution_class` - (Optional) Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
* `maintenance_window` - (Optional) Specifies the day of the week and hour for the maintenance window for streaming jobs.
* `max_capacity` - (Optional) The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
* `max_retries` - (Optional) The maximum number of times to retry this job if it fails.
* `name` - (Required) The name you assign to this job. It must be unique in your account.
* `notification_property` - (Optional) Notification property of the job. Defined below.
* `role_arn` - (Required) The ARN of the IAM role associated with this job.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `timeout` - (Optional) The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
* `security_configuration` - (Optional) The name of the Security Configuration to be associated with the job.
* `source_control_details` - (Optional) The details for a source control configuration for a job, allowing synchronization of job artifacts to or from a remote repository. Defined below.
* `worker_type` - (Optional) The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, G.2X, or G.025X for Spark jobs. Accepts the value Z.2X for Ray jobs.
    * For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
    * For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
    * For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
    * For the G.4X worker type, each worker maps to 4 DPU (16 vCPUs, 64 GB of memory) with 256GB disk (approximately 235GB free), and provides 1 executor per worker. Recommended for memory-intensive jobs. Only available for Glue version 3.0. Available AWS Regions: US East (Ohio), US East (N. Virginia), US West (Oregon), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo), Canada (Central), Europe (Frankfurt), Europe (Ireland), and Europe (Stockholm).
    * For the G.8X worker type, each worker maps to 8 DPU (32 vCPUs, 128 GB of memory) with 512GB disk (approximately 487GB free), and provides 1 executor per worker. Recommended for memory-intensive jobs. Only available for Glue version 3.0. Available AWS Regions: US East (Ohio), US East (N. Virginia), US West (Oregon), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo), Canada (Central), Europe (Frankfurt), Europe (Ireland), and Europe (Stockholm).
    * For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for low volume streaming jobs. Only available for Glue version 3.0.
    * For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPU, 64 GB of m emory, 128 GB disk), and provides up to 8 Ray workers based on the autoscaler.
* `number_of_workers` - (Optional) The number of workers of a defined workerType that are allocated when a job runs.

### command Argument Reference

* `name` - (Optional) The name of the job command. Defaults to `glueetl`. Use `pythonshell` for Python Shell Job Type, `glueray` for Ray Job Type, or `gluestreaming` for Streaming Job Type. `max_capacity` needs to be set if `pythonshell` is chosen.
* `script_location` - (Required) Specifies the S3 path to a script that executes a job.
* `python_version` - (Optional) The Python version being used to execute a Python shell job. Allowed values are 2, 3 or 3.9. Version 3 refers to Python 3.11 when `glue_version` is set to 5.0.
* `runtime` - (Optional) In Ray jobs, runtime is used to specify the versions of Ray, Python and additional libraries available in your environment. This field is not used in other job types. For supported runtime environment values, see [Working with Ray jobs](https://docs.aws.amazon.com/glue/latest/dg/ray-jobs-section.html#author-job-ray-runtimes) in the Glue Developer Guide.

### execution_property Argument Reference

* `max_concurrent_runs` - (Optional) The maximum number of concurrent runs allowed for a job. The default is 1.

### notification_property Argument Reference

* `notify_delay_after` - (Optional) After a job run starts, the number of minutes to wait before sending a job run delay notification.

### source_control_details Argument Reference

* `auth_strategy` - (Optional) The type of authentication, which can be an authentication token stored in Amazon Web Services Secrets Manager, or a personal access token. Valid values are: `PERSONAL_ACCESS_TOKEN` and `AWS_SECRETS_MANAGER`.
* `auth_token` - (Optional) The value of an authorization token.
* `branch` - (Optional) A branch in the remote repository.
* `folder` - (Optional) A folder in the remote repository.
* `last_commit_id` - (Optional) The last commit ID for a commit in the remote repository.
* `owner` - (Optional) The owner of the remote repository that contains the job artifacts.
* `provider` - (Optional) The provider for the remote repository. Valid values are: `GITHUB`, `GITLAB`, `BITBUCKET`, and `AWS_CODE_COMMIT`.
* `repository` - (Optional) The name of the remote repository that contains the job artifacts.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of Glue Job
* `id` - Job name
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Glue Jobs using `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_job import GlueJob
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlueJob.generate_config_for_import(self, "myJob", "MyJob")
```

Using `terraform import`, import Glue Jobs using `name`. For example:

```console
% terraform import aws_glue_job.MyJob MyJob
```

<!-- cache-key: cdktf-0.20.8 input-37cb6d025e4d5572f3e6049132ab8f1159c5e0df753835dcc2fa2c3a575d5a37 -->