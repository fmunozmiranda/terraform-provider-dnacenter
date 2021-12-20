
data "dnacenter_event_series" "example" {
  provider   = dnacenter
  category   = "string"
  domain     = "string"
  end_time   = 1609459200
  event_ids  = "string"
  limit      = 1
  offset     = 1
  order      = "string"
  severity   = "string"
  sort_by    = "string"
  source     = "string"
  start_time = 1609459200
  sub_domain = "string"
  type       = "string"
}

output "dnacenter_event_series_example" {
  value = data.dnacenter_event_series.example.items
}
