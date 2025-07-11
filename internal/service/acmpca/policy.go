// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acmpca

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/sdkv2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_acmpca_policy", name="Policy")
// @ArnIdentity("resource_arn")
// @V60SDKv2Fix
// @Testing(generator=false)
func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourcePolicyPut,
		ReadWithoutTimeout:   resourcePolicyRead,
		UpdateWithoutTimeout: resourcePolicyPut,
		DeleteWithoutTimeout: resourcePolicyDelete,

		Schema: map[string]*schema.Schema{
			names.AttrPolicy: sdkv2.IAMPolicyDocumentSchemaRequired(),
			names.AttrResourceARN: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePolicyPut(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ACMPCAClient(ctx)

	policy, err := structure.NormalizeJsonString(d.Get(names.AttrPolicy).(string))
	if err != nil {
		return sdkdiag.AppendFromErr(diags, err)
	}

	resourceARN := d.Get(names.AttrResourceARN).(string)
	input := acmpca.PutPolicyInput{
		Policy:      aws.String(policy),
		ResourceArn: aws.String(resourceARN),
	}

	_, err = conn.PutPolicy(ctx, &input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "putting ACM PCA Policy (%s): %s", resourceARN, err)
	}

	if d.IsNewResource() {
		d.SetId(resourceARN)
	}

	return append(diags, resourcePolicyRead(ctx, d, meta)...)
}

func resourcePolicyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ACMPCAClient(ctx)

	policy, err := findPolicyByARN(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] ACM PCA Policy (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading ACM PCA Policy (%s): %s", d.Id(), err)
	}

	d.Set(names.AttrPolicy, policy)
	d.Set(names.AttrResourceARN, d.Id())

	return diags
}

func resourcePolicyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ACMPCAClient(ctx)

	log.Printf("[DEBUG] Deleting ACM PCA Policy: %s", d.Id())
	input := acmpca.DeletePolicyInput{
		ResourceArn: aws.String(d.Id()),
	}
	_, err := conn.DeletePolicy(ctx, &input)

	if errs.IsA[*types.ResourceNotFoundException](err) ||
		errs.IsA[*types.RequestAlreadyProcessedException](err) ||
		errs.IsA[*types.RequestInProgressException](err) ||
		errs.IsAErrorMessageContains[*types.InvalidRequestException](err, "Self-signed policy can not be revoked") {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting ACM PCA Policy (%s): %s", d.Id(), err)
	}

	return diags
}

func findPolicyByARN(ctx context.Context, conn *acmpca.Client, arn string) (*string, error) {
	input := acmpca.GetPolicyInput{
		ResourceArn: aws.String(arn),
	}

	output, err := conn.GetPolicy(ctx, &input)

	if errs.IsA[*types.ResourceNotFoundException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.Policy == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.Policy, nil
}
