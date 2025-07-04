// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transfer

import (
	"context"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkDataSource("aws_transfer_connector", name="Connector")
func newConnectorDataSource(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &connectorDataSource{}, nil
}

const (
	DSNameConnector = "Connector Data Source"
)

type connectorDataSource struct {
	framework.DataSourceWithModel[connectorDataSourceModel]
}

func (d *connectorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Connector object was expanded
			"access_role": schema.StringAttribute{
				Computed: true,
			},
			names.AttrARN: schema.StringAttribute{
				Computed: true,
			},
			"as2_config": framework.DataSourceComputedListOfObjectAttribute[dsAs2Config](ctx),
			names.AttrID: schema.StringAttribute{
				CustomType: fwtypes.RegexpType,
				Required:   true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexache.MustCompile(`c-([0-9a-f]{17})`),
						""),
					stringvalidator.LengthAtMost(19),
					stringvalidator.LengthAtLeast(19),
				},
			},
			"logging_role": schema.StringAttribute{
				Computed: true,
			},
			"security_policy_name": schema.StringAttribute{
				Computed: true,
			},
			"service_managed_egress_ip_addresses": schema.ListAttribute{
				CustomType:  fwtypes.ListOfStringType,
				Computed:    true,
				ElementType: types.StringType,
			},
			"sftp_config":  framework.DataSourceComputedListOfObjectAttribute[dsSftpConfig](ctx),
			names.AttrTags: tftags.TagsAttributeComputedOnly(),
			names.AttrURL: schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *connectorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	conn := d.Meta().TransferClient(ctx)

	var data connectorDataSourceModel
	var describeConnectorInput transfer.DescribeConnectorInput
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !data.ConnectorId.IsNull() || !data.ConnectorId.IsUnknown() {
		describeConnectorInput.ConnectorId = data.ConnectorId.ValueStringPointer()
	}

	description, err := conn.DescribeConnector(ctx, &describeConnectorInput)

	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.Transfer, create.ErrActionReading, DSNameConnector, data.SecurityPolicyName.String(), err),
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(flex.Flatten(ctx, description.Connector, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tags := keyValueTags(ctx, description.Connector.Tags).IgnoreAWS().IgnoreConfig(d.Meta().IgnoreTagsConfig(ctx))
	data.Tags = tftags.FlattenStringValueMap(ctx, tags.Map())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type connectorDataSourceModel struct {
	framework.WithRegionModel
	ARN                             types.String                                  `tfsdk:"arn"`
	AccessRole                      types.String                                  `tfsdk:"access_role"`
	As2Config                       fwtypes.ListNestedObjectValueOf[dsAs2Config]  `tfsdk:"as2_config"`
	ConnectorId                     fwtypes.Regexp                                `tfsdk:"id"`
	LoggingRole                     types.String                                  `tfsdk:"logging_role"`
	SecurityPolicyName              types.String                                  `tfsdk:"security_policy_name"`
	ServiceManagedEgressIpAddresses fwtypes.ListOfString                          `tfsdk:"service_managed_egress_ip_addresses"`
	SftpConfig                      fwtypes.ListNestedObjectValueOf[dsSftpConfig] `tfsdk:"sftp_config"`
	Tags                            tftags.Map                                    `tfsdk:"tags"`
	Url                             types.String                                  `tfsdk:"url"`
}

type dsAs2Config struct {
	BasicAuthSecretId   types.String `tfsdk:"basic_auth_secret_id"`
	Compression         types.String `tfsdk:"compression"`
	EncryptionAlgorithm types.String `tfsdk:"encryption_algorithm"`
	LocalProfileId      types.String `tfsdk:"local_profile_id"`
	MdnResponse         types.String `tfsdk:"mdn_response"`
	MdnSigningAlgorithm types.String `tfsdk:"mdn_signing_algorithm"`
	MessageSubject      types.String `tfsdk:"message_subject"`
	PartnerProfileId    types.String `tfsdk:"partner_profile_id"`
	SigningAlgorithm    types.String `tfsdk:"singing_algorithm"`
}

type dsSftpConfig struct {
	TrustedHostKeys fwtypes.ListValueOf[types.String] `tfsdk:"trusted_host_keys"`
	UserSecretId    types.String                      `tfsdk:"user_secret_id"`
}
