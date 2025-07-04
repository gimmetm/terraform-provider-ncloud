---
subcategory: "Server"
---


# Resource: ncloud_public_ip

Provides a Public IP instance resource.

## Example Usage

```hcl
resource "ncloud_public_ip" "public_ip" {
  server_instance_no = "812345"
}
```

## Argument Reference

The following arguments are supported:

* `server_instance_no` - (Optional) Server instance number to assign after creating a public IP. You can get one by calling getPublicIpTargetServerInstanceList.
* `description` - (Optional) Public IP description.


## Attributes Reference

* `id` - The ID of Public IP.
* `public_ip_no` - The ID of Public IP. (It is the same result as `id`)
* `public_ip` - Public IP Address.
* `kind_type` - Public IP kind type

## Import

### `terraform import` command

* Public IP can be imported using the `id`. For example:

```console
$ terraform import ncloud_public_ip.rsc_name 12345
```

### `import` block

* In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Public IP using the `id`. For example:

```terraform
import {
  to = ncloud_public_ip.rsc_name
  id = "12345"
}
```
