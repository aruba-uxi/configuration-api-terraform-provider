---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "uxi_sensor Data Source - uxi"
subcategory: ""
description: |-
  
---

# uxi_sensor (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))

### Read-Only

- `address_note` (String)
- `ethernet_mac_address` (String)
- `id` (String) The ID of this resource.
- `latitude` (Number)
- `longitude` (Number)
- `model_number` (String)
- `name` (String)
- `notes` (String)
- `pcap_mode` (String)
- `serial` (String)
- `wifi_mac_address` (String)

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `sensor_id` (String)