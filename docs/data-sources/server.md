---
subcategory: "Server"
---


# Data Source: ncloud_server

This module can be useful for getting detail of Server instance created before.

## Example Usage

#### Basic usage

```terraform
variable "instance_no" {}

data "ncloud_server" "server" {
  id = var.instance_no
}
```

#### Usage of using filter

```terraform
variable "subnet_no" {}
variable "name" {}

data "ncloud_server" "server" {
  filter {
    name   = "subnet_no"
    values = [var.subnet_no]
  }

  filter {
    name   = "name"
    values = [var.name]
  }
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The ID of the specific Server instance to retrieve.
* `filter` - (Optional) Custom filter block as described below.
  * `name` - (Required) The name of the field to filter by
  * `values` - (Required) Set of values that are accepted for the given field.
  * `regex` - (Optional) is `values` treated as a regular expression. 

## Attributes Reference

* `instance_no` - The ID of Server instance. (It is the same result as `id`)
* `name` - The name of Server instance.
* `description` - Description of the server.
* `server_image_product_code` - Server image product code.
* `server_product_code` - Server product code.
* `cpu_count` - number of CPUs
* `memory_size` - The size of the memory in bytes.
* `platform_type` - Platform type code
* `public_ip` - Public IP
* `base_block_storage_disk_type` - Base block storage disk type code
* `base_block_storage_disk_detail_type` - Base block storage disk detail type code
* `member_server_image_no` - The ID of Member server image.
* `login_key_name` - The login key name to encrypt with the public key.
* `is_protect_server_termination` - Whether is protect return when creating.
* `zone` - Available zone where the Server instance placed.
* `vpc_no` - The ID of the associated VPC. 
* `subnet_no` - The ID of the associated Subnet.
* `network_interface` - List of Network Interface.
  * `network_interface_no` - The ID of Network interface.
  * `order` - Order of network interfaces to be assigned to the server to create.
  * `subnet_no` - Subnet ID of the network interface.
  * `private_ip` - IP address of the network interface.
* `init_script_no` - The ID of Init script.
* `placement_group_no` - The ID of Physical placement group.
* `is_encrypted_base_block_storage_volume` - Whether to encrypt basic block storage if server image is RHV.
* `hypervisor_type` - Hypervisor type. (`XEN` or `KVM`)
* `server_image_number` - Server image number.
* `server_spec_code` - Server spec code.
