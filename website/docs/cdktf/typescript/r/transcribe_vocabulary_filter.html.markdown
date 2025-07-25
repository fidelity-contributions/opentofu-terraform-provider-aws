---
subcategory: "Transcribe"
layout: "aws"
page_title: "AWS: aws_transcribe_vocabulary_filter"
description: |-
  Terraform resource for managing an AWS Transcribe VocabularyFilter.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_transcribe_vocabulary_filter

Terraform resource for managing an AWS Transcribe VocabularyFilter.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { TranscribeVocabularyFilter } from "./.gen/providers/aws/transcribe-vocabulary-filter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new TranscribeVocabularyFilter(this, "example", {
      languageCode: "en-US",
      tags: {
        tag1: "value1",
        tag2: "value3",
      },
      vocabularyFilterName: "example",
      words: ["cars", "bucket"],
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `languageCode` - (Required) The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
* `vocabularyFilterName` - (Required) The name of the VocabularyFilter.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `vocabularyFilterFileUri` - (Optional) The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
* `tags` - (Optional) A map of tags to assign to the VocabularyFilter. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `words` - (Optional) - A list of terms to include in the vocabulary. Conflicts with `vocabularyFilterFileUri` argument.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - VocabularyFilter name.
* `arn` - ARN of the VocabularyFilter.
* `downloadUri` - Generated download URI.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Transcribe VocabularyFilter using the `vocabularyFilterName`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { TranscribeVocabularyFilter } from "./.gen/providers/aws/transcribe-vocabulary-filter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    TranscribeVocabularyFilter.generateConfigForImport(
      this,
      "example",
      "example-name"
    );
  }
}

```

Using `terraform import`, import Transcribe VocabularyFilter using the `vocabularyFilterName`. For example:

```console
% terraform import aws_transcribe_vocabulary_filter.example example-name
```

<!-- cache-key: cdktf-0.20.8 input-776ce9ebfe48d55f7e97f237fc10b3df02aeea4d4c2b1ff0907348615fb09b5a -->