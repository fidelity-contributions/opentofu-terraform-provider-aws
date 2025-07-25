---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_connect_peer"
description: |-
  Manages an AWS Network Manager Connect Peer.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_networkmanager_connect_peer

Manages an AWS Network Manager Connect Peer.

Use this resource to create a Connect peer in AWS Network Manager. Connect peers establish BGP sessions with your on-premises networks through Connect attachments, enabling dynamic routing between your core network and external networks.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_connect_attachment import NetworkmanagerConnectAttachment
from imports.aws.networkmanager_connect_peer import NetworkmanagerConnectPeer
from imports.aws.networkmanager_vpc_attachment import NetworkmanagerVpcAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = NetworkmanagerVpcAttachment(self, "example",
            core_network_id=Token.as_string(awscc_networkmanager_core_network_example.id),
            subnet_arns=Token.as_list(Fn.lookup_nested(aws_subnet_example, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example.arn)
        )
        aws_networkmanager_connect_attachment_example =
        NetworkmanagerConnectAttachment(self, "example_1",
            core_network_id=Token.as_string(awscc_networkmanager_core_network_example.id),
            edge_location=example.edge_location,
            options=NetworkmanagerConnectAttachmentOptions(
                protocol="GRE"
            ),
            transport_attachment_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_connect_attachment_example.override_logical_id("example")
        aws_networkmanager_connect_peer_example = NetworkmanagerConnectPeer(self, "example_2",
            bgp_options=NetworkmanagerConnectPeerBgpOptions(
                peer_asn=65000
            ),
            connect_attachment_id=Token.as_string(aws_networkmanager_connect_attachment_example.id),
            inside_cidr_blocks=["172.16.0.0/16"],
            peer_address="127.0.0.1"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_connect_peer_example.override_logical_id("example")
```

### Usage with attachment accepter

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_attachment_accepter import NetworkmanagerAttachmentAccepter
from imports.aws.networkmanager_connect_attachment import NetworkmanagerConnectAttachment
from imports.aws.networkmanager_connect_peer import NetworkmanagerConnectPeer
from imports.aws.networkmanager_vpc_attachment import NetworkmanagerVpcAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = NetworkmanagerVpcAttachment(self, "example",
            core_network_id=Token.as_string(awscc_networkmanager_core_network_example.id),
            subnet_arns=Token.as_list(Fn.lookup_nested(aws_subnet_example, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example.arn)
        )
        aws_networkmanager_attachment_accepter_example =
        NetworkmanagerAttachmentAccepter(self, "example_1",
            attachment_id=example.id,
            attachment_type=example.attachment_type
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_attachment_accepter_example.override_logical_id("example")
        aws_networkmanager_connect_attachment_example =
        NetworkmanagerConnectAttachment(self, "example_2",
            core_network_id=Token.as_string(awscc_networkmanager_core_network_example.id),
            depends_on=[aws_networkmanager_attachment_accepter_example],
            edge_location=example.edge_location,
            options=NetworkmanagerConnectAttachmentOptions(
                protocol="GRE"
            ),
            transport_attachment_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_connect_attachment_example.override_logical_id("example")
        example2 = NetworkmanagerAttachmentAccepter(self, "example2",
            attachment_id=Token.as_string(aws_networkmanager_connect_attachment_example.id),
            attachment_type=Token.as_string(aws_networkmanager_connect_attachment_example.attachment_type)
        )
        aws_networkmanager_connect_peer_example = NetworkmanagerConnectPeer(self, "example_4",
            bgp_options=NetworkmanagerConnectPeerBgpOptions(
                peer_asn=65500
            ),
            connect_attachment_id=Token.as_string(aws_networkmanager_connect_attachment_example.id),
            depends_on=[example2],
            inside_cidr_blocks=["172.16.0.0/16"],
            peer_address="127.0.0.1"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_connect_peer_example.override_logical_id("example")
```

### Usage with a Tunnel-less Connect attachment

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_connect_attachment import NetworkmanagerConnectAttachment
from imports.aws.networkmanager_connect_peer import NetworkmanagerConnectPeer
from imports.aws.networkmanager_vpc_attachment import NetworkmanagerVpcAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = NetworkmanagerVpcAttachment(self, "example",
            core_network_id=Token.as_string(awscc_networkmanager_core_network_example.id),
            subnet_arns=Token.as_list(Fn.lookup_nested(aws_subnet_example, ["*", "arn"])),
            vpc_arn=Token.as_string(aws_vpc_example.arn)
        )
        aws_networkmanager_connect_attachment_example =
        NetworkmanagerConnectAttachment(self, "example_1",
            core_network_id=Token.as_string(awscc_networkmanager_core_network_example.id),
            edge_location=example.edge_location,
            options=NetworkmanagerConnectAttachmentOptions(
                protocol="NO_ENCAP"
            ),
            transport_attachment_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_connect_attachment_example.override_logical_id("example")
        aws_networkmanager_connect_peer_example = NetworkmanagerConnectPeer(self, "example_2",
            bgp_options=NetworkmanagerConnectPeerBgpOptions(
                peer_asn=65000
            ),
            connect_attachment_id=Token.as_string(aws_networkmanager_connect_attachment_example.id),
            peer_address="127.0.0.1",
            subnet_arn=example2.arn
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_networkmanager_connect_peer_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `connect_attachment_id` - (Required) ID of the connection attachment.
* `peer_address` - (Required) Connect peer address.

The following arguments are optional:

* `bgp_options` - (Optional) Connect peer BGP options. See [bgp_options](#bgp_options) for more information.
* `core_network_address` - (Optional) Connect peer core network address.
* `inside_cidr_blocks` - (Optional) Inside IP addresses used for BGP peering. Required when the Connect attachment protocol is `GRE`. See [`aws_networkmanager_connect_attachment`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/networkmanager_connect_attachment) for details.
* `subnet_arn` - (Optional) Subnet ARN for the Connect peer. Required when the Connect attachment protocol is `NO_ENCAP`. See [`aws_networkmanager_connect_attachment`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/networkmanager_connect_attachment) for details.
* `tags` - (Optional) Key-value tags for the attachment. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### bgp_options

* `peer_asn` - (Optional) Peer ASN.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Connect peer.
* `configuration` - Configuration of the Connect peer.
* `connect_peer_id` - ID of the Connect peer.
* `core_network_id` - ID of a core network.
* `created_at` - Timestamp when the Connect peer was created.
* `edge_location` - Region where the peer is located.
* `state` - State of the Connect peer.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `delete` - (Default `15m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_networkmanager_connect_peer` using the connect peer ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkmanager_connect_peer import NetworkmanagerConnectPeer
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkmanagerConnectPeer.generate_config_for_import(self, "example", "connect-peer-061f3e96275db1acc")
```

Using `terraform import`, import `aws_networkmanager_connect_peer` using the connect peer ID. For example:

```console
% terraform import aws_networkmanager_connect_peer.example connect-peer-061f3e96275db1acc
```

<!-- cache-key: cdktf-0.20.8 input-00f38b8ee9a4df94fe5653ee1ce5e977c919f030fcff28603d3064e8d8fa09d9 -->