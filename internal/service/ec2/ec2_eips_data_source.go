// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_eips", name="EIPs")
func dataSourceEIPs() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceEIPsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"allocation_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrFilter: customFiltersSchema(),
			"public_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrTags: tftags.TagsSchema(),
		},
	}
}

func dataSourceEIPsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	input := ec2.DescribeAddressesInput{}

	if tags, tagsOk := d.GetOk(names.AttrTags); tagsOk {
		input.Filters = append(input.Filters, newTagFilterList(
			svcTags(tftags.New(ctx, tags.(map[string]any))),
		)...)
	}

	if filters, filtersOk := d.GetOk(names.AttrFilter); filtersOk {
		input.Filters = append(input.Filters,
			newCustomFilterList(filters.(*schema.Set))...)
	}

	if len(input.Filters) == 0 {
		input.Filters = nil
	}

	output, err := findEIPs(ctx, conn, &input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 EIPs: %s", err)
	}

	var allocationIDs []string
	var publicIPs []string

	for _, v := range output {
		publicIPs = append(publicIPs, aws.ToString(v.PublicIp))

		if v.Domain == awstypes.DomainTypeVpc {
			allocationIDs = append(allocationIDs, aws.ToString(v.AllocationId))
		}
	}

	d.SetId(meta.(*conns.AWSClient).Region(ctx))
	d.Set("allocation_ids", allocationIDs)
	d.Set("public_ips", publicIPs)

	return diags
}
