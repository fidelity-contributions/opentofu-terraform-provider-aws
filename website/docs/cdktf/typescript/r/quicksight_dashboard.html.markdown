---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_dashboard"
description: |-
  Manages a QuickSight Dashboard.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_dashboard

Resource for managing a QuickSight Dashboard.

## Example Usage

### From Source Template

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightDashboard } from "./.gen/providers/aws/quicksight-dashboard";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new QuicksightDashboard(this, "example", {
      dashboardId: "example-id",
      name: "example-name",
      sourceEntity: {
        sourceTemplate: {
          arn: source.arn,
          dataSetReferences: [
            {
              dataSetArn: dataset.arn,
              dataSetPlaceholder: "1",
            },
          ],
        },
      },
      versionDescription: "version",
    });
  }
}

```

### With Definition

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightDashboard } from "./.gen/providers/aws/quicksight-dashboard";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new QuicksightDashboard(this, "example", {
      dashboardId: "example-id",
      definition: {
        dataSetIdentifiersDeclarations: [
          {
            dataSetArn: dataset.arn,
            identifier: "1",
          },
        ],
        sheets: [
          {
            sheetId: "Example1",
            title: "Example",
            visuals: [
              {
                lineChartVisual: {
                  chartConfiguration: {
                    fieldWells: {
                      lineChartAggregatedFieldWells: {
                        category: [
                          {
                            categoricalDimensionField: {
                              column: {
                                columnName: "Column1",
                                dataSetIdentifier: "1",
                              },
                              fieldId: "1",
                            },
                          },
                        ],
                        values: [
                          {
                            categoricalMeasureField: {
                              aggregationFunction: "COUNT",
                              column: {
                                columnName: "Column1",
                                dataSetIdentifier: "1",
                              },
                              fieldId: "2",
                            },
                          },
                        ],
                      },
                    },
                  },
                  title: {
                    formatText: {
                      plainText: "Line Chart Example",
                    },
                  },
                  visualId: "LineChart",
                },
              },
            ],
          },
        ],
      },
      name: "example-name",
      versionDescription: "version",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `dashboardId` - (Required, Forces new resource) Identifier for the dashboard.
* `name` - (Required) Display name for the dashboard.
* `versionDescription` - (Required) A description of the current dashboard version being created/updated.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `awsAccountId` - (Optional, Forces new resource) AWS account ID.
* `dashboardPublishOptions` - (Optional) Options for publishing the dashboard. See [dashboard_publish_options](#dashboard_publish_options).
* `definition` - (Optional) A detailed dashboard definition. Only one of `definition` or `sourceEntity` should be configured. See [definition](#definition).
* `parameters` - (Optional) The parameters for the creation of the dashboard, which you want to use to override the default settings. A dashboard can have any type of parameters, and some parameters might accept multiple values. See [parameters](#parameters).
* `permissions` - (Optional) A set of resource permissions on the dashboard. Maximum of 64 items. See [permissions](#permissions).
* `sourceEntity` - (Optional) The entity that you are using as a source when you create the dashboard (template). Only one of `definition` or `sourceEntity` should be configured. See [source_entity](#source_entity).
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `themeArn` - (Optional) The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. The theme ARN must exist in the same AWS account where you create the dashboard.

### permissions

* `actions` - (Required) List of IAM actions to grant or revoke permissions on.
* `principal` - (Required) ARN of the principal. See the [ResourcePermission documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ResourcePermission.html) for the applicable ARN values.

### source_entity

* `sourceTemplate` - (Optional) The source template. See [source_template](#source_template).

### source_template

* `arn` - (Required) The Amazon Resource Name (ARN) of the resource.
* `dataSetReferences` - (Required) List of dataset references. See [data_set_references](#data_set_references).

### data_set_references

* `dataSetArn` - (Required) Dataset Amazon Resource Name (ARN).
* `dataSetPlaceholder` - (Required) Dataset placeholder.

### dashboard_publish_options

* `adHocFilteringOption` - (Optional) Ad hoc (one-time) filtering option. See [ad_hoc_filtering_option](#ad_hoc_filtering_option).
* `dataPointDrillUpDownOption` - (Optional) The drill-down options of data points in a dashboard. See [data_point_drill_up_down_option](#data_point_drill_up_down_option).
* `dataPointMenuLabelOption` - (Optional) The data point menu label options of a dashboard. See [data_point_menu_label_option](#data_point_menu_label_option).
* `dataPointTooltipOption` - (Optional) The data point tool tip options of a dashboard. See [data_point_tooltip_option](#data_point_tooltip_option).
* `exportToCsvOption` - (Optional) Export to .csv option. See [export_to_csv_option](#export_to_csv_option).
* `exportWithHiddenFieldsOption` - (Optional) Determines if hidden fields are exported with a dashboard. See [export_with_hidden_fields_option](#export_with_hidden_fields_option).
* `sheetControlsOption` - (Optional) Sheet controls option. See [sheet_controls_option](#sheet_controls_option).
* `sheetLayoutElementMaximizationOption` - (Optional) The sheet layout maximization options of a dashboard. See [sheet_layout_element_maximization_option](#sheet_layout_element_maximization_option).
* `visualAxisSortOption` - (Optional) The axis sort options of a dashboard. See [visual_axis_sort_option](#visual_axis_sort_option).
* `visualMenuOption` - (Optional) The menu options of a visual in a dashboard. See [visual_menu_option](#visual_menu_option).

### ad_hoc_filtering_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### data_point_drill_up_down_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### data_point_menu_label_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### data_point_tooltip_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### export_to_csv_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### export_with_hidden_fields_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### sheet_controls_option

* `visibilityState` - (Optional) Visibility state. Possibles values: EXPANDED, COLLAPSED.

### sheet_layout_element_maximization_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### visual_axis_sort_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### visual_menu_option

* `availabilityStatus` - (Optional) Availability status. Possibles values: ENABLED, DISABLED.

### parameters

* `dateTimeParameters` - (Optional) A list of parameters that have a data type of date-time. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DateTimeParameter.html).
* `decimalParameters` - (Optional) A list of parameters that have a data type of decimal. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalParameter.html).
* `integerParameters` - (Optional) A list of parameters that have a data type of integer. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_IntegerParameter.html).
* `stringParameters` - (Optional) A list of parameters that have a data type of string. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StringParameter.html).

### definition

* `dataSetIdentifiersDeclarations` - (Required) A list dataset identifier declarations. With this mapping,you can use dataset identifiers instead of dataset Amazon Resource Names (ARNs) throughout the dashboard's sub-structures. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSetIdentifierDeclaration.html).
* `analysisDefaults` - (Optional) The configuration for default analysis settings. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_AnalysisDefaults.html).
* `calculatedFields` - (Optional) A list of calculated field definitions for the dashboard. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CalculatedField.html).
* `columnConfigurations` - (Optional) A list of dashboard-level column configurations. Column configurations are used to set default formatting for a column that's used throughout a dashboard. See [AWS API Documentation for complete description](ttps://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnConfiguration.html).
* `filterGroups` - (Optional) A list of filter definitions for a dashboard. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FilterGroup.html). For more information, see [Filtering Data](https://docs.aws.amazon.com/quicksight/latest/user/filtering-visual-data.html) in Amazon QuickSight User Guide.
* `parametersDeclarations` - (Optional) A list of parameter declarations for a dashboard. Parameters are named variables that can transfer a value for use by an action or an object. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ParameterDeclaration.html). For more information, see [Parameters in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the Amazon QuickSight User Guide.
* `sheets` - (Optional) A list of sheet definitions for a dashboard. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_SheetDefinition.html).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the dashboard.
* `createdTime` - The time that the dashboard was created.
* `id` - A comma-delimited string joining AWS account ID and dashboard ID.
* `lastUpdatedTime` - The time that the dashboard was last updated.
* `sourceEntityArn` - Amazon Resource Name (ARN) of a template that was used to create this dashboard.
* `status` - The dashboard creation status.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).
* `versionNumber` - The version number of the dashboard version.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `5m`)
* `update` - (Default `5m`)
* `delete` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a QuickSight Dashboard using the AWS account ID and dashboard ID separated by a comma (`,`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightDashboard } from "./.gen/providers/aws/quicksight-dashboard";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    QuicksightDashboard.generateConfigForImport(
      this,
      "example",
      "123456789012,example-id"
    );
  }
}

```

Using `terraform import`, import a QuickSight Dashboard using the AWS account ID and dashboard ID separated by a comma (`,`). For example:

```console
% terraform import aws_quicksight_dashboard.example 123456789012,example-id
```

<!-- cache-key: cdktf-0.20.8 input-0dde427a0ea9320283cfea5d88db74f856433081838c7dfd8d83488fb5d678d1 -->