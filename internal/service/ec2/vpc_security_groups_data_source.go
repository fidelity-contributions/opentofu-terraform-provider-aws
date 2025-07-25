// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_security_groups", name="Security Groups")
func dataSourceSecurityGroups() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceSecurityGroupsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrARNs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrFilter: customFiltersSchema(),
			names.AttrIDs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrTags: tftags.TagsSchemaComputed(),
			"vpc_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceSecurityGroupsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := meta.(*conns.AWSClient)
	conn := c.EC2Client(ctx)

	input := &ec2.DescribeSecurityGroupsInput{}

	input.Filters = append(input.Filters, newTagFilterList(
		svcTags(tftags.New(ctx, d.Get(names.AttrTags).(map[string]any))),
	)...)

	input.Filters = append(input.Filters, newCustomFilterList(
		d.Get(names.AttrFilter).(*schema.Set),
	)...)

	if len(input.Filters) == 0 {
		input.Filters = nil
	}

	output, err := findSecurityGroups(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 Security Groups: %s", err)
	}

	var arns, securityGroupIDs, vpcIDs []string

	for _, v := range output {
		ownerID, sgID := aws.ToString(v.OwnerId), aws.ToString(v.GroupId)
		arns = append(arns, securityGroupARN(ctx, c, ownerID, sgID))
		securityGroupIDs = append(securityGroupIDs, sgID)
		vpcIDs = append(vpcIDs, aws.ToString(v.VpcId))
	}

	d.SetId(c.Region(ctx))
	d.Set(names.AttrARNs, arns)
	d.Set(names.AttrIDs, securityGroupIDs)
	d.Set("vpc_ids", vpcIDs)

	return diags
}
