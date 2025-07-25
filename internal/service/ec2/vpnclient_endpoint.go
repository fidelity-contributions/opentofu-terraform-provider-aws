// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_ec2_client_vpn_endpoint", name="Client VPN Endpoint")
// @Tags(identifierAttribute="id")
// @Testing(tagsTest=false)
func resourceClientVPNEndpoint() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceClientVPNEndpointCreate,
		ReadWithoutTimeout:   resourceClientVPNEndpointRead,
		DeleteWithoutTimeout: resourceClientVPNEndpointDelete,
		UpdateWithoutTimeout: resourceClientVPNEndpointUpdate,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_options": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MaxItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_directory_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"root_certificate_chain_arn": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidARN,
						},
						"saml_provider_arn": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidARN,
						},
						"self_service_saml_provider_arn": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidARN,
						},
						names.AttrType: {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							ValidateDiagFunc: enum.Validate[awstypes.ClientVpnAuthenticationType](),
						},
					},
				},
			},
			"client_cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"client_connect_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						names.AttrEnabled: {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"lambda_function_arn": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: verify.ValidARN,
						},
					},
				},
			},
			"client_login_banner_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"banner_text": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringLenBetween(0, 1400),
						},
						names.AttrEnabled: {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"client_route_enforcement_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enforced": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"connection_log_options": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudwatch_log_group": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cloudwatch_log_stream": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						names.AttrEnabled: {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			names.AttrDescription: {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disconnect_on_session_timeout": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			names.AttrDNSName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrSecurityGroupIDs: {
				Type:     schema.TypeSet,
				MinItems: 1,
				MaxItems: 5,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"self_service_portal": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          awstypes.SelfServicePortalDisabled,
				ValidateDiagFunc: enum.Validate[awstypes.SelfServicePortal](),
			},
			"self_service_portal_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_certificate_arn": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidARN,
			},
			"session_timeout_hours": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      24,
				ValidateFunc: validation.IntInSlice([]int{8, 10, 12, 24}),
			},
			"split_tunnel": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"transport_protocol": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Default:          awstypes.TransportProtocolUdp,
				ValidateDiagFunc: enum.Validate[awstypes.TransportProtocol](),
			},
			names.AttrVPCID: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
				ValidateFunc: validation.IntInSlice([]int{
					443,
					1194,
				}),
			},
		},
	}
}

func resourceClientVPNEndpointCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	input := &ec2.CreateClientVpnEndpointInput{
		ClientCidrBlock:      aws.String(d.Get("client_cidr_block").(string)),
		ClientToken:          aws.String(id.UniqueId()),
		ServerCertificateArn: aws.String(d.Get("server_certificate_arn").(string)),
		SplitTunnel:          aws.Bool(d.Get("split_tunnel").(bool)),
		TagSpecifications:    getTagSpecificationsIn(ctx, awstypes.ResourceTypeClientVpnEndpoint),
		TransportProtocol:    awstypes.TransportProtocol(d.Get("transport_protocol").(string)),
		VpnPort:              aws.Int32(int32(d.Get("vpn_port").(int))),
	}

	if v, ok := d.GetOk("authentication_options"); ok && v.(*schema.Set).Len() > 0 {
		input.AuthenticationOptions = expandClientVPNAuthenticationRequests(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("client_connect_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.ClientConnectOptions = expandClientConnectOptions(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk("client_login_banner_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.ClientLoginBannerOptions = expandClientLoginBannerOptions(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk("client_route_enforcement_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.ClientRouteEnforcementOptions = expandClientRouteEnforcementOptions(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk("connection_log_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
		input.ConnectionLogOptions = expandConnectionLogOptions(v.([]any)[0].(map[string]any))
	}

	if v, ok := d.GetOk(names.AttrDescription); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("disconnect_on_session_timeout"); ok {
		input.DisconnectOnSessionTimeout = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("dns_servers"); ok && len(v.([]any)) > 0 {
		input.DnsServers = flex.ExpandStringValueList(v.([]any))
	}

	if v, ok := d.GetOk(names.AttrSecurityGroupIDs); ok {
		input.SecurityGroupIds = flex.ExpandStringValueSet(v.(*schema.Set))
	}

	if v, ok := d.GetOk("self_service_portal"); ok {
		input.SelfServicePortal = awstypes.SelfServicePortal(v.(string))
	}

	if v, ok := d.GetOk("session_timeout_hours"); ok {
		input.SessionTimeoutHours = aws.Int32(int32(v.(int)))
	}

	if v, ok := d.GetOk(names.AttrVPCID); ok {
		input.VpcId = aws.String(v.(string))
	}

	output, err := conn.CreateClientVpnEndpoint(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating EC2 Client VPN Endpoint: %s", err)
	}

	d.SetId(aws.ToString(output.ClientVpnEndpointId))

	return append(diags, resourceClientVPNEndpointRead(ctx, d, meta)...)
}

func resourceClientVPNEndpointRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := meta.(*conns.AWSClient)
	conn := c.EC2Client(ctx)

	ep, err := findClientVPNEndpointByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EC2 Client VPN Endpoint (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 Client VPN Endpoint (%s): %s", d.Id(), err)
	}

	d.Set(names.AttrARN, clientVPNEndpointARN(ctx, c, d.Id()))
	if err := d.Set("authentication_options", flattenClientVPNAuthentications(ep.AuthenticationOptions)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting authentication_options: %s", err)
	}
	d.Set("client_cidr_block", ep.ClientCidrBlock)
	if ep.ClientConnectOptions != nil {
		if err := d.Set("client_connect_options", []any{flattenClientConnectResponseOptions(ep.ClientConnectOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting client_connect_options: %s", err)
		}
	} else {
		d.Set("client_connect_options", nil)
	}
	if ep.ClientLoginBannerOptions != nil {
		if err := d.Set("client_login_banner_options", []any{flattenClientLoginBannerResponseOptions(ep.ClientLoginBannerOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting client_login_banner_options: %s", err)
		}
	} else {
		d.Set("client_login_banner_options", nil)
	}
	if ep.ClientRouteEnforcementOptions != nil {
		if err := d.Set("client_route_enforcement_options", []any{flattenClientRouteEnforcementOptions(ep.ClientRouteEnforcementOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting client_route_enforcement_options: %s", err)
		}
	} else {
		d.Set("client_route_enforcement_options", nil)
	}
	if ep.ConnectionLogOptions != nil {
		if err := d.Set("connection_log_options", []any{flattenConnectionLogResponseOptions(ep.ConnectionLogOptions)}); err != nil {
			return sdkdiag.AppendErrorf(diags, "setting connection_log_options: %s", err)
		}
	} else {
		d.Set("connection_log_options", nil)
	}
	d.Set(names.AttrDescription, ep.Description)
	d.Set("disconnect_on_session_timeout", ep.DisconnectOnSessionTimeout)
	d.Set(names.AttrDNSName, ep.DnsName)
	d.Set("dns_servers", aws.StringSlice(ep.DnsServers))
	d.Set(names.AttrSecurityGroupIDs, aws.StringSlice(ep.SecurityGroupIds))
	if aws.ToString(ep.SelfServicePortalUrl) != "" {
		d.Set("self_service_portal", awstypes.SelfServicePortalEnabled)
	} else {
		d.Set("self_service_portal", awstypes.SelfServicePortalDisabled)
	}
	d.Set("self_service_portal_url", ep.SelfServicePortalUrl)
	d.Set("server_certificate_arn", ep.ServerCertificateArn)
	d.Set("session_timeout_hours", ep.SessionTimeoutHours)
	d.Set("split_tunnel", ep.SplitTunnel)
	d.Set("transport_protocol", ep.TransportProtocol)
	d.Set(names.AttrVPCID, ep.VpcId)
	d.Set("vpn_port", ep.VpnPort)

	setTagsOut(ctx, ep.Tags)

	return diags
}

func resourceClientVPNEndpointUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	if d.HasChangesExcept(names.AttrTags, names.AttrTagsAll) {
		var waitForClientConnectResponseOptionsUpdate bool
		input := &ec2.ModifyClientVpnEndpointInput{
			ClientVpnEndpointId: aws.String(d.Id()),
		}

		if d.HasChange("client_connect_options") {
			waitForClientConnectResponseOptionsUpdate = true

			if v, ok := d.GetOk("client_connect_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
				input.ClientConnectOptions = expandClientConnectOptions(v.([]any)[0].(map[string]any))
			}
		}

		if d.HasChange("client_login_banner_options") {
			if v, ok := d.GetOk("client_login_banner_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
				input.ClientLoginBannerOptions = expandClientLoginBannerOptions(v.([]any)[0].(map[string]any))
			}
		}

		if d.HasChange("client_route_enforcement_options") {
			if v, ok := d.GetOk("client_route_enforcement_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
				input.ClientRouteEnforcementOptions = expandClientRouteEnforcementOptions(v.([]any)[0].(map[string]any))
			}
		}

		if d.HasChange("connection_log_options") {
			if v, ok := d.GetOk("connection_log_options"); ok && len(v.([]any)) > 0 && v.([]any)[0] != nil {
				input.ConnectionLogOptions = expandConnectionLogOptions(v.([]any)[0].(map[string]any))
			}
		}

		if d.HasChange(names.AttrDescription) {
			input.Description = aws.String(d.Get(names.AttrDescription).(string))
		}

		if d.HasChange("disconnect_on_session_timeout") {
			input.DisconnectOnSessionTimeout = aws.Bool(d.Get("disconnect_on_session_timeout").(bool))
		}

		if d.HasChange("dns_servers") {
			dnsServers := d.Get("dns_servers").([]any)
			enabled := len(dnsServers) > 0

			input.DnsServers = &awstypes.DnsServersOptionsModifyStructure{
				Enabled: aws.Bool(enabled),
			}
			if enabled {
				input.DnsServers.CustomDnsServers = flex.ExpandStringValueList(dnsServers)
			}
		}

		if d.HasChange(names.AttrSecurityGroupIDs) {
			input.SecurityGroupIds = flex.ExpandStringValueSet(d.Get(names.AttrSecurityGroupIDs).(*schema.Set))
			// "InvalidParameterValue: Security Groups cannot be modified without specifying Vpc Id"
			input.VpcId = aws.String(d.Get(names.AttrVPCID).(string))
		}

		if d.HasChange("self_service_portal") {
			input.SelfServicePortal = awstypes.SelfServicePortal(d.Get("self_service_portal").(string))
		}

		if d.HasChange("session_timeout_hours") {
			input.SessionTimeoutHours = aws.Int32(int32(d.Get("session_timeout_hours").(int)))
		}

		if d.HasChange("server_certificate_arn") {
			input.ServerCertificateArn = aws.String(d.Get("server_certificate_arn").(string))
		}

		if d.HasChange("split_tunnel") {
			input.SplitTunnel = aws.Bool(d.Get("split_tunnel").(bool))
		}

		if d.HasChange("vpn_port") {
			input.VpnPort = aws.Int32(int32(d.Get("vpn_port").(int)))
		}

		if d.HasChange(names.AttrVPCID) {
			input.VpcId = aws.String(d.Get(names.AttrVPCID).(string))
		}

		if _, err := conn.ModifyClientVpnEndpoint(ctx, input); err != nil {
			return sdkdiag.AppendErrorf(diags, "modifying EC2 Client VPN Endpoint (%s): %s", d.Id(), err)
		}

		if waitForClientConnectResponseOptionsUpdate {
			if _, err := waitClientVPNEndpointClientConnectResponseOptionsUpdated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
				return sdkdiag.AppendErrorf(diags, "waiting for EC2 Client VPN Endpoint (%s) ClientConnectResponseOptions update: %s", d.Id(), err)
			}
		}
	}

	return append(diags, resourceClientVPNEndpointRead(ctx, d, meta)...)
}

func resourceClientVPNEndpointDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	log.Printf("[DEBUG] Deleting EC2 Client VPN Endpoint: %s", d.Id())
	input := ec2.DeleteClientVpnEndpointInput{
		ClientVpnEndpointId: aws.String(d.Id()),
	}
	_, err := conn.DeleteClientVpnEndpoint(ctx, &input)

	if tfawserr.ErrCodeEquals(err, errCodeInvalidClientVPNEndpointIdNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting EC2 Client VPN Endpoint (%s): %s", d.Id(), err)
	}

	if _, err := waitClientVPNEndpointDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for EC2 Client VPN Endpoint (%s) delete: %s", d.Id(), err)
	}

	return diags
}

func expandClientVPNAuthenticationRequest(tfMap map[string]any) *awstypes.ClientVpnAuthenticationRequest {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.ClientVpnAuthenticationRequest{}

	var authnType string
	if v, ok := tfMap[names.AttrType].(string); ok && v != "" {
		authnType = v
		apiObject.Type = awstypes.ClientVpnAuthenticationType(v)
	}

	switch authnType {
	case string(awstypes.ClientVpnAuthenticationTypeCertificateAuthentication):
		if v, ok := tfMap["root_certificate_chain_arn"].(string); ok && v != "" {
			apiObject.MutualAuthentication = &awstypes.CertificateAuthenticationRequest{
				ClientRootCertificateChainArn: aws.String(v),
			}
		}

	case string(awstypes.ClientVpnAuthenticationTypeDirectoryServiceAuthentication):
		if v, ok := tfMap["active_directory_id"].(string); ok && v != "" {
			apiObject.ActiveDirectory = &awstypes.DirectoryServiceAuthenticationRequest{
				DirectoryId: aws.String(v),
			}
		}

	case string(awstypes.ClientVpnAuthenticationTypeFederatedAuthentication):
		if v, ok := tfMap["saml_provider_arn"].(string); ok && v != "" {
			apiObject.FederatedAuthentication = &awstypes.FederatedAuthenticationRequest{
				SAMLProviderArn: aws.String(v),
			}

			if v, ok := tfMap["self_service_saml_provider_arn"].(string); ok && v != "" {
				apiObject.FederatedAuthentication.SelfServiceSAMLProviderArn = aws.String(v)
			}
		}
	}

	return apiObject
}

func expandClientVPNAuthenticationRequests(tfList []any) []awstypes.ClientVpnAuthenticationRequest {
	if len(tfList) == 0 {
		return nil
	}

	var apiObjects []awstypes.ClientVpnAuthenticationRequest

	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]any)

		if !ok {
			continue
		}

		apiObject := expandClientVPNAuthenticationRequest(tfMap)

		if apiObject == nil {
			continue
		}

		apiObjects = append(apiObjects, *apiObject)
	}

	return apiObjects
}

func flattenClientVPNAuthentication(apiObject awstypes.ClientVpnAuthentication) map[string]any {
	tfMap := map[string]any{}
	tfMap[names.AttrType] = apiObject.Type

	if apiObject.MutualAuthentication != nil {
		if v := apiObject.MutualAuthentication.ClientRootCertificateChain; v != nil {
			tfMap["root_certificate_chain_arn"] = aws.ToString(v)
		}
	} else if apiObject.ActiveDirectory != nil {
		if v := apiObject.ActiveDirectory.DirectoryId; v != nil {
			tfMap["active_directory_id"] = aws.ToString(v)
		}
	} else if apiObject.FederatedAuthentication != nil {
		if v := apiObject.FederatedAuthentication.SamlProviderArn; v != nil {
			tfMap["saml_provider_arn"] = aws.ToString(v)
		}

		if v := apiObject.FederatedAuthentication.SelfServiceSamlProviderArn; v != nil {
			tfMap["self_service_saml_provider_arn"] = aws.ToString(v)
		}
	}

	return tfMap
}

func flattenClientVPNAuthentications(apiObjects []awstypes.ClientVpnAuthentication) []any {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []any

	for _, apiObject := range apiObjects {
		tfList = append(tfList, flattenClientVPNAuthentication(apiObject))
	}

	return tfList
}

func expandClientConnectOptions(tfMap map[string]any) *awstypes.ClientConnectOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.ClientConnectOptions{}

	var enabled bool
	if v, ok := tfMap[names.AttrEnabled].(bool); ok {
		enabled = v
	}

	if enabled {
		if v, ok := tfMap["lambda_function_arn"].(string); ok && v != "" {
			apiObject.LambdaFunctionArn = aws.String(v)
		}
	}

	apiObject.Enabled = aws.Bool(enabled)

	return apiObject
}

func flattenClientConnectResponseOptions(apiObject *awstypes.ClientConnectResponseOptions) map[string]any {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]any{}

	if v := apiObject.Enabled; v != nil {
		tfMap[names.AttrEnabled] = v
	}

	if v := apiObject.LambdaFunctionArn; v != nil {
		tfMap["lambda_function_arn"] = aws.ToString(v)
	}

	return tfMap
}

func expandClientLoginBannerOptions(tfMap map[string]any) *awstypes.ClientLoginBannerOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.ClientLoginBannerOptions{}

	var enabled bool
	if v, ok := tfMap[names.AttrEnabled].(bool); ok {
		enabled = v
	}

	if enabled {
		if v, ok := tfMap["banner_text"].(string); ok && v != "" {
			apiObject.BannerText = aws.String(v)
		}
	}

	apiObject.Enabled = aws.Bool(enabled)

	return apiObject
}

func flattenClientLoginBannerResponseOptions(apiObject *awstypes.ClientLoginBannerResponseOptions) map[string]any {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]any{}

	if v := apiObject.BannerText; v != nil {
		tfMap["banner_text"] = aws.ToString(v)
	}

	if v := apiObject.Enabled; v != nil {
		tfMap[names.AttrEnabled] = v
	}

	return tfMap
}

func expandConnectionLogOptions(tfMap map[string]any) *awstypes.ConnectionLogOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.ConnectionLogOptions{}

	var enabled bool
	if v, ok := tfMap[names.AttrEnabled].(bool); ok {
		enabled = v
	}

	if enabled {
		if v, ok := tfMap["cloudwatch_log_group"].(string); ok && v != "" {
			apiObject.CloudwatchLogGroup = aws.String(v)
		}

		if v, ok := tfMap["cloudwatch_log_stream"].(string); ok && v != "" {
			apiObject.CloudwatchLogStream = aws.String(v)
		}
	}

	apiObject.Enabled = aws.Bool(enabled)

	return apiObject
}

func flattenConnectionLogResponseOptions(apiObject *awstypes.ConnectionLogResponseOptions) map[string]any {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]any{}

	if v := apiObject.CloudwatchLogGroup; v != nil {
		tfMap["cloudwatch_log_group"] = aws.ToString(v)
	}

	if v := apiObject.CloudwatchLogStream; v != nil {
		tfMap["cloudwatch_log_stream"] = aws.ToString(v)
	}

	if v := apiObject.Enabled; v != nil {
		tfMap[names.AttrEnabled] = v
	}

	return tfMap
}

func expandClientRouteEnforcementOptions(tfMap map[string]any) *awstypes.ClientRouteEnforcementOptions {
	if tfMap == nil {
		return nil
	}

	apiObject := &awstypes.ClientRouteEnforcementOptions{}

	if v, ok := tfMap["enforced"].(bool); ok {
		apiObject.Enforced = aws.Bool(v)
	}

	return apiObject
}

func flattenClientRouteEnforcementOptions(apiObject *awstypes.ClientRouteEnforcementResponseOptions) map[string]any {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]any{}

	if v := apiObject.Enforced; v != nil {
		tfMap["enforced"] = v
	}

	return tfMap
}

func clientVPNEndpointARN(ctx context.Context, c *conns.AWSClient, clientVPNEndpointID string) string {
	return c.RegionalARN(ctx, names.EC2, "client-vpn-endpoint/"+clientVPNEndpointID)
}
