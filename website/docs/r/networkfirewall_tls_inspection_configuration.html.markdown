---
subcategory: "Network Firewall"
layout: "aws"
page_title: "AWS: aws_networkfirewall_tls_inspection_configuration"
description: |-
  Terraform resource for managing an AWS Network Firewall TLS Inspection Configuration.
---

# Resource: aws_networkfirewall_tls_inspection_configuration

Terraform resource for managing an AWS Network Firewall TLS Inspection Configuration.

## Example Usage

~> **NOTE:** You must configure either inbound inspection, outbound inspection, or both.

### Basic inbound/ingress inspection

```
resource "aws_networkfirewall_tls_inspection_configuration" "example" {
  name        = "example"
  description = "example"
  encryption_configuration {
    key_id = "AWS_OWNED_KMS_KEY"
    type   = "AWS_OWNED_KMS_KEY"
  }
  tls_inspection_configuration {
    server_certificate_configuration {
      server_certificate {
        resource_arn = aws_acm_certificate.example_1.arn
      }
      scope {
        protocols = [6]
        destination_ports {
          from_port = 443
          to_port   = 443
        }
        destination {
          address_definition = "0.0.0.0/0"
        }
        source_ports {
          from_port = 0
          to_port   = 65535
        }
        source {
          address_definition = "0.0.0.0/0"
        }
      }
    }
  }
}
```

### Basic outbound/engress inspection

```
resource "aws_networkfirewall_tls_inspection_configuration" "example" {
  name        = "example"
  description = "example"
  encryption_configuration {
    key_id = "AWS_OWNED_KMS_KEY"
    type   = "AWS_OWNED_KMS_KEY"
  }
  tls_inspection_configuration {
    server_certificate_configuration {
      certificate_authority_arn = aws_acm_certificate.example_1.arn
      check_certificate_revocation_status {
        revoked_status_action = "REJECT"
        unknown_status_action = "PASS"
      }
      scope {
        protocols = [6]
        destination_ports {
          from_port = 443
          to_port   = 443
        }
        destination {
          address_definition = "0.0.0.0/0"
        }
        source_ports {
          from_port = 0
          to_port   = 65535
        }
        source {
          address_definition = "0.0.0.0/0"
        }
      }
    }
  }
}
```

### Inbound with encryption configuration

```
resource "aws_kms_key" "example" {
  description             = "example"
  deletion_window_in_days = 7
}

resource "aws_networkfirewall_tls_inspection_configuration" "example" {
  name        = "example"
  description = "example"
  encryption_configuration {
    key_id = aws_kms_key.example.arn
    type   = "CUSTOMER_KMS"
  }
  tls_inspection_configuration {
    server_certificate_configuration {
      server_certificate {
        resource_arn = aws_acm_certificate.example_1.arn
      }
      scopes {
        protocols = [6]
        destination_ports {
          from_port = 443
          to_port   = 443
        }
        destinations {
          address_definition = "0.0.0.0/0"
        }
        source_ports {
          from_port = 0
          to_port   = 65535
        }
        sources {
          address_definition = "0.0.0.0/0"
        }
      }
    }
  }
}
```

### Outbound with encryption configuration

```
resource "aws_kms_key" "example" {
  description             = "example"
  deletion_window_in_days = 7
}

resource "aws_networkfirewall_tls_inspection_configuration" "example" {
  name        = "example"
  description = "example"
  encryption_configuration {
    key_id = aws_kms_key.example.arn
    type   = "CUSTOMER_KMS"
  }
  tls_inspection_configuration {
    server_certificate_configurations {
      certificate_authority_arn = aws_acm_certificate.example_1.arn
      check_certificate_revocation_status {
        revoked_status_action = "REJECT"
        unknown_status_action = "PASS"
      }
      scope {
        protocols = [6]
        destination_ports {
          from_port = 443
          to_port   = 443
        }
        destination {
          address_definition = "0.0.0.0/0"
        }
        source_ports {
          from_port = 0
          to_port   = 65535
        }
        source {
          address_definition = "0.0.0.0/0"
        }
      }
    }
  }
}
```

### Combined inbound and outbound

```terraform
resource "aws_networkfirewall_tls_inspection_configuration" "example" {
  name        = "example"
  description = "example"
  encryption_configuration {
    key_id = "AWS_OWNED_KMS_KEY"
    type   = "AWS_OWNED_KMS_KEY"
  }
  tls_inspection_configuration {
    server_certificate_configuration {
      certificate_authority_arn = aws_acm_certificate.example_1.arn
      check_certificate_revocation_status {
        revoked_status_action = "REJECT"
        unknown_status_action = "PASS"
      }
      server_certificate {
        resource_arn = aws_acm_certificate.example_2.arn
      }
      scope {
        protocols = [6]
        destination_ports {
          from_port = 443
          to_port   = 443
        }
        destination {
          address_definition = "0.0.0.0/0"
        }
        source_ports {
          from_port = 0
          to_port   = 65535
        }
        source {
          address_definition = "0.0.0.0/0"
        }
      }
    }
  }
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required, Forces new resource) Descriptive name of the TLS inspection configuration.
* `tls_inspection_configuration` - (Required) TLS inspection configuration block. Detailed below.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `description` - (Optional) Description of the TLS inspection configuration.
* `encryption_configuration` - (Optional) Encryption configuration block. Detailed below.

### Encryption Configuration

* `key_id` - (Optional) ARN of the Amazon Web Services Key Management Service (KMS) customer managed key.
* `type` - (Optional) Type of KMS key to use for encryption of your Network Firewall resources. Valid values: `AWS_OWNED_KMS_KEY`, `CUSTOMER_KMS`.

### TLS Inspection Configuration

* `server_certificate_configuration` - (Required) Server certificate configurations that are associated with the TLS configuration. Detailed below.

### Server Certificate Configuration

The `server_certificate_configuration` block supports the following arguments:

* `certificate_authority_arn` - (Optional) ARN of the imported certificate authority (CA) certificate within Certificate Manager (ACM) to use for outbound SSL/TLS inspection. See [Using SSL/TLS certificates with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html) for limitations on CA certificates.
* `check_certificate_revocation_status` (Optional) - Check Certificate Revocation Status block. Detailed below.
* `scope` (Required) - Scope block. Detailed below.
* `server_certificate` - (Optional) Server certificates to use for inbound SSL/TLS inspection. See [Using SSL/TLS certificates with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html).

### Check Certificate Revocation Status

The `check_certificate_revocation_status` block supports the following arguments:

~> **NOTE  To check the certificate revocation status, you must also specify a `certificate_authority_arn` in `server_certificate_configuration`.

`revoked_status_action` - (Optional) how Network Firewall processes traffic when it determines that the certificate presented by the server in the SSL/TLS connection has a revoked status. See [Checking certificate revocation status](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html#tls-inspection-check-certificate-revocation-status) for details. Valid values: `PASS`, `DROP`, `REJECT`.
`unknown_status_action` - (Optional) How Network Firewall processes traffic when it determines that the certificate presented by the server in the SSL/TLS connection has an unknown status, or a status that cannot be determined for any other reason, including when the service is unable to connect to the OCSP and CRL endpoints for the certificate. See [Checking certificate revocation status](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html#tls-inspection-check-certificate-revocation-status) for details. Valid values: `PASS`, `DROP`, `REJECT`.

### Scopes

The `scope` block supports the following arguments:

* `destination` - (Required) Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See [Destination](#destination) below for details.
* `destination_ports` - (Optional) Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See [Destination Ports](#destination-ports) below for details.
* `protocols` - (Optional) Set of protocols to inspect for, specified using the protocol's assigned internet protocol number (IANA). Network Firewall currently supports TCP only. Valid values: `6`
* `source` - (Optional) Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See [Source](#source) below for details.
* `source_ports` - (Optional) Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See [Source Ports](#source-ports) below for details.

### Destination

The `destination` block supports the following argument:

* `address_definition` - (Required)  An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.

### Destination Ports

The `destination_ports` block supports the following arguments:

* `from_ports` - (Required) The lower limit of the port range. This must be less than or equal to the `to_port`.
* `to_ports` - (Optional) The upper limit of the port range. This must be greater than or equal to the `from_port`.

### Source

The `source` block supports the following argument:

* `address_definition` - (Required)  An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.

### Source Ports

The `source_ports` block supports the following arguments:

* `from_port` - (Required) The lower limit of the port range. This must be less than or equal to the `to_port`.
* `to_port` - (Optional) The upper limit of the port range. This must be greater than or equal to the `from_port`.

### Server Certificates

The `server_certificate` block supports the following arguments:

* `resource_arn` - (Optional) ARN of the Certificate Manager SSL/TLS server certificate that's used for inbound SSL/TLS inspection.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the TLS Inspection Configuration.
* `certificate_authority` - Certificate Manager certificate block. See [Certificate Authority](#certificate-authority) below for details.
* `certificates` - List of certificate blocks describing certificates associated with the TLS inspection configuration. See [Certificates](#certificates) below for details.
* `number_of_associations` - Number of firewall policies that use this TLS inspection configuration.
* `tls_inspection_configuration_id` - A unique identifier for the TLS inspection configuration.
* `update_token` - String token used when updating the rule group.

### Certificate Authority

The `certificate_authority` block exports the following attributes:

* `certificate_arn` - ARN of the certificate.
* `certificate_serial` -  Serial number of the certificate.
* `status` - Status of the certificate.
* `status_message` - Details about the certificate status, including information about certificate errors.

### Certificates

The `certificates` block exports the following attributes:

* `certificate_arn` - ARN of the certificate.
* `certificate_serial` -  Serial number of the certificate.
* `status` - Status of the certificate.
* `status_message` - Details about the certificate status, including information about certificate errors.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Network Firewall TLS Inspection Configuration using the `arn`. For example:

```terraform
import {
  to = aws_networkfirewall_tls_inspection_configuration.example
  id = "arn:aws:network-firewall::<region>:<account_id>:tls-configuration/example"
}
```

Using `terraform import`, import Network Firewall TLS Inspection Configuration using the `arn`. For example:

```console
% terraform import aws_networkfirewall_tls_inspection_configuration.example arn:aws:network-firewall::<region>:<account_id>:tls-configuration/example
```
