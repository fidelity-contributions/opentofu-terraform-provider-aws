// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logs

import (
	"context"
	"log"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	awstypes "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/retry"
	tfslices "github.com/hashicorp/terraform-provider-aws/internal/slices"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_cloudwatch_log_destination", name="Destination")
// @Tags(identifierAttribute="arn")
// @Testing(existsType="github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types;awstypes;awstypes.Destination")
// @Testing(existsTakesT=true)
// @Testing(destroyTakesT=true)
func resourceDestination() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceDestinationCreate,
		ReadWithoutTimeout:   resourceDestinationRead,
		UpdateWithoutTimeout: resourceDestinationUpdate,
		DeleteWithoutTimeout: resourceDestinationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.Any(
					validation.StringLenBetween(1, 512),
					validation.StringMatch(regexache.MustCompile(`[^:*]*`), ""),
				),
			},
			names.AttrRoleARN: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidARN,
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			names.AttrTargetARN: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidARN,
			},
		},
	}
}

const (
	propagationTimeout = 2 * time.Minute
)

func resourceDestinationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).LogsClient(ctx)

	name := d.Get(names.AttrName).(string)
	input := &cloudwatchlogs.PutDestinationInput{
		DestinationName: aws.String(name),
		RoleArn:         aws.String(d.Get(names.AttrRoleARN).(string)),
		TargetArn:       aws.String(d.Get(names.AttrTargetARN).(string)),
	}

	outputRaw, err := tfresource.RetryWhenIsA[*awstypes.InvalidParameterException](ctx, propagationTimeout, func() (any, error) {
		return conn.PutDestination(ctx, input)
	})

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating CloudWatch Logs Destination (%s): %s", name, err)
	}

	destination := outputRaw.(*cloudwatchlogs.PutDestinationOutput).Destination
	d.SetId(aws.ToString(destination.DestinationName))

	// Although PutDestinationInput has a Tags field, specifying tags there results in
	// "InvalidParameterException: Could not deliver test message to specified destination. Check if the destination is valid."
	if err := createTags(ctx, conn, aws.ToString(destination.Arn), getTagsIn(ctx)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting CloudWatch Logs Destination (%s) tags: %s", d.Id(), err)
	}

	return append(diags, resourceDestinationRead(ctx, d, meta)...)
}

func resourceDestinationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).LogsClient(ctx)

	destination, err := findDestinationByName(ctx, conn, d.Id())

	if !d.IsNewResource() && retry.NotFound(err) {
		log.Printf("[WARN] CloudWatch Logs Destination (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading CloudWatch Logs Destination (%s): %s", d.Id(), err)
	}

	d.Set(names.AttrARN, destination.Arn)
	d.Set(names.AttrName, destination.DestinationName)
	d.Set(names.AttrRoleARN, destination.RoleArn)
	d.Set(names.AttrTargetARN, destination.TargetArn)

	return diags
}

func resourceDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).LogsClient(ctx)

	if d.HasChangesExcept(names.AttrTags, names.AttrTagsAll) {
		input := &cloudwatchlogs.PutDestinationInput{
			DestinationName: aws.String(d.Id()),
			RoleArn:         aws.String(d.Get(names.AttrRoleARN).(string)),
			TargetArn:       aws.String(d.Get(names.AttrTargetARN).(string)),
		}

		_, err := tfresource.RetryWhenIsA[*awstypes.InvalidParameterException](ctx, propagationTimeout, func() (any, error) {
			return conn.PutDestination(ctx, input)
		})

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating CloudWatch Logs Destination (%s): %s", d.Id(), err)
		}
	}

	return append(diags, resourceDestinationRead(ctx, d, meta)...)
}

func resourceDestinationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).LogsClient(ctx)

	log.Printf("[INFO] Deleting CloudWatch Logs Destination: %s", d.Id())
	_, err := conn.DeleteDestination(ctx, &cloudwatchlogs.DeleteDestinationInput{
		DestinationName: aws.String(d.Id()),
	})

	if errs.IsA[*awstypes.ResourceNotFoundException](err) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting CloudWatch Logs Destination (%s): %s", d.Id(), err)
	}

	return diags
}

func findDestinationByName(ctx context.Context, conn *cloudwatchlogs.Client, name string) (*awstypes.Destination, error) {
	input := cloudwatchlogs.DescribeDestinationsInput{
		DestinationNamePrefix: aws.String(name),
	}

	return findDestination(ctx, conn, &input, func(v *awstypes.Destination) bool {
		return aws.ToString(v.DestinationName) == name
	})
}

func findDestination(ctx context.Context, conn *cloudwatchlogs.Client, input *cloudwatchlogs.DescribeDestinationsInput, filter tfslices.Predicate[*awstypes.Destination]) (*awstypes.Destination, error) {
	output, err := findDestinations(ctx, conn, input, filter, tfslices.WithReturnFirstMatch)

	if err != nil {
		return nil, err
	}

	return tfresource.AssertSingleValueResult(output)
}

func findDestinations(ctx context.Context, conn *cloudwatchlogs.Client, input *cloudwatchlogs.DescribeDestinationsInput, filter tfslices.Predicate[*awstypes.Destination], optFns ...tfslices.FinderOptionsFunc) ([]awstypes.Destination, error) {
	var output []awstypes.Destination
	opts := tfslices.NewFinderOptions(optFns)

	pages := cloudwatchlogs.NewDescribeDestinationsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)

		if err != nil {
			return nil, err
		}

		for _, v := range page.Destinations {
			if filter(&v) {
				output = append(output, v)
				if opts.ReturnFirstMatch() {
					return output, nil
				}
			}
		}
	}

	return output, nil
}
