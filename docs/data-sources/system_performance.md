---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_system_performance Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Health and Performance.
  This data source gives the aggregated performance indicators. The data can be retrieved for the last 3 months.
---

# dnacenter_system_performance (Data Source)

It performs read operation on Health and Performance.

- This data source gives the aggregated performance indicators. The data can be retrieved for the last 3 months.

## Example Usage

```terraform
data "dnacenter_system_performance" "example" {
  provider   = dnacenter
  end_time   = 1609459200
  function   = "string"
  kpi        = "string"
  start_time = 1609459200
}

output "dnacenter_system_performance_example" {
  value = data.dnacenter_system_performance.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **end_time** (Number) endTime query parameter. This is the epoch end time in milliseconds upto which performance indicator need to be fetched
- **function** (String) function query parameter. Valid values: sum,average,max
- **id** (String) The ID of this resource.
- **kpi** (String) kpi query parameter. Valid values: cpu,memory,network
- **start_time** (Number) startTime query parameter. This is the epoch start time in milliseconds from which performance indicator need to be fetched

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **host_name** (String)
- **kpis** (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis))
- **version** (String)

<a id="nestedobjatt--item--kpis"></a>
### Nested Schema for `item.kpis`

Read-Only:

- **cpu** (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--cpu))
- **memory** (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--memory))
- **network_rx_rate** (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--network_rx_rate))
- **network_tx_rate** (List of Object) (see [below for nested schema](#nestedobjatt--item--kpis--network_tx_rate))

<a id="nestedobjatt--item--kpis--cpu"></a>
### Nested Schema for `item.kpis.cpu`

Read-Only:

- **units** (String)
- **utilization** (String)


<a id="nestedobjatt--item--kpis--memory"></a>
### Nested Schema for `item.kpis.memory`

Read-Only:

- **units** (String)
- **utilization** (String)


<a id="nestedobjatt--item--kpis--network_rx_rate"></a>
### Nested Schema for `item.kpis.network_rx_rate`

Read-Only:

- **units** (String)
- **utilization** (String)


<a id="nestedobjatt--item--kpis--network_tx_rate"></a>
### Nested Schema for `item.kpis.network_tx_rate`

Read-Only:

- **units** (String)
- **utilization** (String)

