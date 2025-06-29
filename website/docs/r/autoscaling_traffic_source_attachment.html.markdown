---
subcategory: "Auto Scaling"
layout: "aws"
page_title: "AWS: aws_autoscaling_traffic_source_attachment"
description: |-
  Terraform resource for managing an AWS Auto Scaling Traffic Source Attachment.
---

# Resource: aws_autoscaling_traffic_source_attachment

Attaches a traffic source to an Auto Scaling group.

~> **NOTE on Auto Scaling Groups, Attachments and Traffic Source Attachments:** Terraform provides standalone [Attachment](autoscaling_attachment.html) (for attaching Classic Load Balancers and Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target groups) and Traffic Source Attachment (for attaching Load Balancers and VPC Lattice target groups) resources and an [Auto Scaling Group](autoscaling_group.html) resource with `load_balancers`, `target_group_arns` and `traffic_source` attributes. Do not use the same traffic source in more than one of these resources. Doing so will cause a conflict of attachments. A [`lifecycle` configuration block](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html) can be used to suppress differences if necessary.

## Example Usage

### Basic Usage

```terraform
resource "aws_autoscaling_traffic_source_attachment" "example" {
  autoscaling_group_name = aws_autoscaling_group.example.id

  traffic_source {
    identifier = aws_lb_target_group.example.arn
    type       = "elbv2"
  }
}
```

## Argument Reference

This resource supports the following arguments:

- `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
- `autoscaling_group_name` - (Required) The name of the Auto Scaling group.
- `traffic_source` - (Required) The unique identifiers of a traffic sources.

`traffic_source` supports the following:

- `identifier` - (Required) Identifies the traffic source. For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.
- `type` - (Required) Provides additional context for the value of `identifier`.
  The following lists the valid values:
  `elb` if `identifier` is the name of a Classic Load Balancer.
  `elbv2` if `identifier` is the ARN of an Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target group.
  `vpc-lattice` if `identifier` is the ARN of a VPC Lattice target group.

## Attribute Reference

This resource exports no additional attributes.
