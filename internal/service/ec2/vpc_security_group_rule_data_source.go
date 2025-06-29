// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwflex "github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkDataSource("aws_vpc_security_group_rule", name="Security Group Rule")
func newSecurityGroupRuleDataSource(context.Context) (datasource.DataSourceWithConfigure, error) {
	d := &securityGroupRuleDataSource{}

	return d, nil
}

type securityGroupRuleDataSource struct {
	framework.DataSourceWithModel[securityGroupRuleDataSourceModel]
}

func (d *securityGroupRuleDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrARN: schema.StringAttribute{
				Computed: true,
			},
			"cidr_ipv4": schema.StringAttribute{
				Computed: true,
			},
			"cidr_ipv6": schema.StringAttribute{
				Computed: true,
			},
			names.AttrDescription: schema.StringAttribute{
				Computed: true,
			},
			"from_port": schema.Int64Attribute{
				Computed: true,
			},
			names.AttrID: framework.IDAttribute(),
			"ip_protocol": schema.StringAttribute{
				Computed: true,
			},
			"is_egress": schema.BoolAttribute{
				Computed: true,
			},
			"prefix_list_id": schema.StringAttribute{
				Computed: true,
			},
			"referenced_security_group_id": schema.StringAttribute{
				Computed: true,
			},
			"security_group_id": schema.StringAttribute{
				Computed: true,
			},
			"security_group_rule_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			names.AttrTags: tftags.TagsAttributeComputedOnly(),
			"to_port": schema.Int64Attribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			names.AttrFilter: customFiltersBlock(ctx),
		},
	}
}

func (d *securityGroupRuleDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data securityGroupRuleDataSourceModel
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	conn := d.Meta().EC2Client(ctx)
	ignoreTagsConfig := d.Meta().IgnoreTagsConfig(ctx)

	input := &ec2.DescribeSecurityGroupRulesInput{
		Filters: newCustomFilterListFramework(ctx, data.Filters),
	}

	if !data.SecurityGroupRuleID.IsNull() {
		input.SecurityGroupRuleIds = []string{fwflex.StringValueFromFramework(ctx, data.SecurityGroupRuleID)}
	}

	if len(input.Filters) == 0 {
		// Don't send an empty filters list; the EC2 API won't accept it.
		input.Filters = nil
	}

	output, err := findSecurityGroupRule(ctx, conn, input)

	if err != nil {
		response.Diagnostics.AddError("reading Security Group Rules", tfresource.SingularDataSourceFindError("Security Group Rule", err).Error())

		return
	}

	data.ID = fwflex.StringToFramework(ctx, output.SecurityGroupRuleId)
	data.ARN = d.securityGroupRuleARN(ctx, data.ID.ValueString())
	data.CIDRIPv4 = fwflex.StringToFramework(ctx, output.CidrIpv4)
	data.CIDRIPv6 = fwflex.StringToFramework(ctx, output.CidrIpv6)
	data.Description = fwflex.StringToFramework(ctx, output.Description)
	data.FromPort = fwflex.Int32ToFrameworkInt64(ctx, output.FromPort)
	data.IPProtocol = fwflex.StringToFramework(ctx, output.IpProtocol)
	data.IsEgress = fwflex.BoolToFramework(ctx, output.IsEgress)
	data.PrefixListID = fwflex.StringToFramework(ctx, output.PrefixListId)
	data.ReferencedSecurityGroupID = flattenReferencedSecurityGroup(ctx, output.ReferencedGroupInfo, d.Meta().AccountID(ctx))
	data.SecurityGroupID = fwflex.StringToFramework(ctx, output.GroupId)
	data.SecurityGroupRuleID = fwflex.StringToFramework(ctx, output.SecurityGroupRuleId)
	data.Tags = tftags.FlattenStringValueMap(ctx, keyValueTags(ctx, output.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map())
	data.ToPort = fwflex.Int32ToFrameworkInt64(ctx, output.ToPort)

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (d *securityGroupRuleDataSource) securityGroupRuleARN(ctx context.Context, id string) types.String {
	return types.StringValue(d.Meta().RegionalARN(ctx, names.EC2, fmt.Sprintf("security-group-rule/%s", id)))
}

type securityGroupRuleDataSourceModel struct {
	framework.WithRegionModel
	ARN                       types.String  `tfsdk:"arn"`
	CIDRIPv4                  types.String  `tfsdk:"cidr_ipv4"`
	CIDRIPv6                  types.String  `tfsdk:"cidr_ipv6"`
	Description               types.String  `tfsdk:"description"`
	Filters                   customFilters `tfsdk:"filter"`
	FromPort                  types.Int64   `tfsdk:"from_port"`
	ID                        types.String  `tfsdk:"id"`
	IPProtocol                types.String  `tfsdk:"ip_protocol"`
	IsEgress                  types.Bool    `tfsdk:"is_egress"`
	PrefixListID              types.String  `tfsdk:"prefix_list_id"`
	ReferencedSecurityGroupID types.String  `tfsdk:"referenced_security_group_id"`
	SecurityGroupID           types.String  `tfsdk:"security_group_id"`
	SecurityGroupRuleID       types.String  `tfsdk:"security_group_rule_id"`
	Tags                      tftags.Map    `tfsdk:"tags"`
	ToPort                    types.Int64   `tfsdk:"to_port"`
}
