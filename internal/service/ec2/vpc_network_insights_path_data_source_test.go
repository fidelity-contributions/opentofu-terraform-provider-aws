// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccVPCNetworkInsightsPathDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	datasourceName := "data.aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCNetworkInsightsPathDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrDestination, resourceName, names.AttrDestination),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrDestinationARN, resourceName, names.AttrDestinationARN),
					resource.TestCheckResourceAttrPair(datasourceName, "destination_ip", resourceName, "destination_ip"),
					resource.TestCheckResourceAttrPair(datasourceName, "destination_port", resourceName, "destination_port"),
					resource.TestCheckResourceAttr(datasourceName, "filter_at_destination.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "filter_at_source.#", "1"),
					resource.TestCheckResourceAttrPair(datasourceName, "filter_at_source.0.destination_port_range.0.from_port", resourceName, "destination_port"),
					resource.TestCheckResourceAttrPair(datasourceName, "filter_at_source.0.destination_port_range.0.to_port", resourceName, "destination_port"),
					resource.TestCheckResourceAttrPair(datasourceName, "network_insights_path_id", resourceName, names.AttrID),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrProtocol, resourceName, names.AttrProtocol),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrSource, resourceName, names.AttrSource),
					resource.TestCheckResourceAttrPair(datasourceName, "source_arn", resourceName, "source_arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "source_ip", resourceName, "source_ip"),
					resource.TestCheckResourceAttrPair(datasourceName, acctest.CtTagsPercent, resourceName, acctest.CtTagsPercent),
				),
			},
		},
	})
}

func TestAccVPCNetworkInsightsPathDataSource_filterAtSource(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	datasourceName := "data.aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCNetworkInsightsPathDataSourceConfig_filterAtDestination(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(datasourceName, "filter_at_source", resourceName, "filter_at_source"),
					resource.TestCheckResourceAttrPair(datasourceName, "network_insights_path_id", resourceName, names.AttrID),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrProtocol, resourceName, names.AttrProtocol),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrSource, resourceName, names.AttrSource),
					resource.TestCheckResourceAttrPair(datasourceName, "source_arn", resourceName, "source_arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "source_ip", resourceName, "source_ip"),
					resource.TestCheckResourceAttrPair(datasourceName, acctest.CtTagsPercent, resourceName, acctest.CtTagsPercent),
				),
			},
		},
	})
}

func TestAccVPCNetworkInsightsPathDataSource_filterAtDestination(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	datasourceName := "data.aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCNetworkInsightsPathDataSourceConfig_filterAtSource(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrDestination, resourceName, names.AttrDestination),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrDestinationARN, resourceName, names.AttrDestinationARN),
					resource.TestCheckResourceAttrPair(datasourceName, "filter_at_destination", resourceName, "filter_at_destination"),
					resource.TestCheckResourceAttrPair(datasourceName, "network_insights_path_id", resourceName, names.AttrID),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrProtocol, resourceName, names.AttrProtocol),
					resource.TestCheckResourceAttrPair(datasourceName, names.AttrSource, resourceName, names.AttrSource),
					resource.TestCheckResourceAttrPair(datasourceName, acctest.CtTagsPercent, resourceName, acctest.CtTagsPercent),
				),
			},
		},
	})
}

func testAccVPCNetworkInsightsPathDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

  subnet_id = aws_subnet.test[0].id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_network_insights_path" "test" {
  source           = aws_network_interface.test[0].id
  destination      = aws_network_interface.test[1].id
  destination_port = 443
  protocol         = "tcp"

  tags = {
    Name = %[1]q
  }
}

data "aws_ec2_network_insights_path" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id
}
`, rName))
}

func testAccVPCNetworkInsightsPathDataSourceConfig_filterAtSource(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

  subnet_id = aws_subnet.test[0].id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_network_insights_path" "test" {
  source = aws_network_interface.test[0].id
  filter_at_source {
    destination_address = aws_network_interface.test[1].private_ip
    destination_port_range {
      from_port = 80
      to_port   = 80
    }
    source_address = aws_network_interface.test[0].private_ip
    source_port_range {
      from_port = 0
      to_port   = 65535
    }
  }
  protocol = "tcp"

  tags = {
    Name = %[1]q
  }
}

data "aws_ec2_network_insights_path" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id
}
`, rName))
}

func testAccVPCNetworkInsightsPathDataSourceConfig_filterAtDestination(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

  subnet_id = aws_subnet.test[0].id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_network_insights_path" "test" {
  source      = aws_network_interface.test[0].id
  destination = aws_network_interface.test[1].id
  filter_at_destination {
    destination_address = aws_network_interface.test[1].private_ip
    destination_port_range {
      from_port = 80
      to_port   = 80
    }
    source_address = aws_network_interface.test[0].private_ip
    source_port_range {
      from_port = 0
      to_port   = 65535
    }
  }
  protocol = "tcp"

  tags = {
    Name = %[1]q
  }
}

data "aws_ec2_network_insights_path" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id
}
`, rName))
}
