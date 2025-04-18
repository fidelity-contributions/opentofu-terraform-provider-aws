// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfsagemaker "github.com/hashicorp/terraform-provider-aws/internal/service/sagemaker"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccSageMakerImageVersion_basic(t *testing.T) {
	ctx := acctest.Context(t)
	if os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE") == "" {
		t.Skip("Environment variable SAGEMAKER_IMAGE_VERSION_BASE_IMAGE is not set")
	}

	var image sagemaker.DescribeImageVersionOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_image_version.test"
	baseImage := os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SageMakerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckImageVersionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccImageVersionConfig_basic(rName, baseImage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageVersionExists(ctx, resourceName, &image),
					resource.TestCheckResourceAttr(resourceName, "image_name", rName),
					resource.TestCheckResourceAttr(resourceName, "base_image", baseImage),
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "1"),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, "image_arn", "sagemaker", fmt.Sprintf("image/%s", rName)),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "sagemaker", fmt.Sprintf("image-version/%s/1", rName)),
					resource.TestCheckResourceAttrSet(resourceName, "container_image"),
					resource.TestCheckResourceAttr(resourceName, "horovod", acctest.CtFalse),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSageMakerImageVersion_full(t *testing.T) {
	ctx := acctest.Context(t)
	if os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE") == "" {
		t.Skip("Environment variable SAGEMAKER_IMAGE_VERSION_BASE_IMAGE is not set")
	}

	var image sagemaker.DescribeImageVersionOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNameUpdate := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_image_version.test"
	baseImage := os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SageMakerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckImageVersionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccImageVersionConfig_full(rName, baseImage, rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageVersionExists(ctx, resourceName, &image),
					resource.TestCheckResourceAttr(resourceName, "image_name", rName),
					resource.TestCheckResourceAttr(resourceName, "base_image", baseImage),
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "1"),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, "image_arn", "sagemaker", fmt.Sprintf("image/%s", rName)),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "sagemaker", fmt.Sprintf("image-version/%s/1", rName)),
					resource.TestCheckResourceAttrSet(resourceName, "container_image"),
					resource.TestCheckResourceAttr(resourceName, "horovod", acctest.CtFalse),
					resource.TestCheckResourceAttr(resourceName, "processor", "CPU"),
					resource.TestCheckResourceAttr(resourceName, "vendor_guidance", "STABLE"),
					resource.TestCheckResourceAttr(resourceName, "release_notes", rName),
					resource.TestCheckResourceAttr(resourceName, "job_type", "TRAINING"),
					resource.TestCheckResourceAttr(resourceName, "ml_framework", "TensorFlow 1.1"),
					resource.TestCheckResourceAttr(resourceName, "programming_lang", "Python 3.8"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccImageVersionConfig_full(rName, baseImage, rNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageVersionExists(ctx, resourceName, &image),
					resource.TestCheckResourceAttr(resourceName, "image_name", rName),
					resource.TestCheckResourceAttr(resourceName, "base_image", baseImage),
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "1"),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, "image_arn", "sagemaker", fmt.Sprintf("image/%s", rName)),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "sagemaker", fmt.Sprintf("image-version/%s/1", rName)),
					resource.TestCheckResourceAttrSet(resourceName, "container_image"),
					resource.TestCheckResourceAttr(resourceName, "horovod", acctest.CtFalse),
					resource.TestCheckResourceAttr(resourceName, "processor", "CPU"),
					resource.TestCheckResourceAttr(resourceName, "vendor_guidance", "STABLE"),
					resource.TestCheckResourceAttr(resourceName, "release_notes", rNameUpdate),
					resource.TestCheckResourceAttr(resourceName, "job_type", "TRAINING"),
					resource.TestCheckResourceAttr(resourceName, "ml_framework", "TensorFlow 1.1"),
					resource.TestCheckResourceAttr(resourceName, "programming_lang", "Python 3.8"),
				),
			},
		},
	})
}

func TestAccSageMakerImageVersion_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	if os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE") == "" {
		t.Skip("Environment variable SAGEMAKER_IMAGE_VERSION_BASE_IMAGE is not set")
	}

	var image sagemaker.DescribeImageVersionOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_image_version.test"
	baseImage := os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SageMakerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckImageVersionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccImageVersionConfig_basic(rName, baseImage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageVersionExists(ctx, resourceName, &image),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceImageVersion(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSageMakerImageVersion_Disappears_image(t *testing.T) {
	ctx := acctest.Context(t)
	if os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE") == "" {
		t.Skip("Environment variable SAGEMAKER_IMAGE_VERSION_BASE_IMAGE is not set")
	}

	var image sagemaker.DescribeImageVersionOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_image_version.test"
	baseImage := os.Getenv("SAGEMAKER_IMAGE_VERSION_BASE_IMAGE")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SageMakerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckImageVersionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccImageVersionConfig_basic(rName, baseImage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageVersionExists(ctx, resourceName, &image),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceImage(), "aws_sagemaker_image.test"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckImageVersionDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_sagemaker_image_version" {
				continue
			}

			_, err := tfsagemaker.FindImageVersionByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return fmt.Errorf("reading SageMaker AI Image Version (%s): %w", rs.Primary.ID, err)
			}

			return fmt.Errorf("SageMaker AI Image Version %q still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckImageVersionExists(ctx context.Context, n string, image *sagemaker.DescribeImageVersionOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sagmaker Image ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerClient(ctx)
		resp, err := tfsagemaker.FindImageVersionByName(ctx, conn, rs.Primary.ID)
		if err != nil {
			return err
		}

		*image = *resp

		return nil
	}
}

func testAccImageVersionConfigBase(rName string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
  name               = %[1]q
  assume_role_policy = data.aws_iam_policy_document.test.json
}

data "aws_iam_policy_document" "test" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["sagemaker.${data.aws_partition.current.dns_suffix}"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "test" {
  role       = aws_iam_role.test.name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/AmazonSageMakerFullAccess"
}

resource "aws_sagemaker_image" "test" {
  image_name = %[1]q
  role_arn   = aws_iam_role.test.arn

  depends_on = [aws_iam_role_policy_attachment.test]
}
`, rName)
}

func testAccImageVersionConfig_basic(rName, baseImage string) string {
	return testAccImageVersionConfigBase(rName) + fmt.Sprintf(`
resource "aws_sagemaker_image_version" "test" {
  image_name = aws_sagemaker_image.test.id
  base_image = %[1]q
}
`, baseImage)
}

func testAccImageVersionConfig_full(rName, baseImage, notes string) string {
	return testAccImageVersionConfigBase(rName) + fmt.Sprintf(`
resource "aws_sagemaker_image_version" "test" {
  image_name       = aws_sagemaker_image.test.id
  base_image       = %[1]q
  job_type         = "TRAINING"
  processor        = "CPU"
  release_notes    = %[2]q
  vendor_guidance  = "STABLE"
  ml_framework     = "TensorFlow 1.1"
  programming_lang = "Python 3.8"
}
`, baseImage, notes)
}
