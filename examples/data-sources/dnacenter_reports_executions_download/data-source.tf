
data "dnacenter__reports_executions_download" "example" {
  provider     = dnacenter
  dirpath      = "string"
  execution_id = "string"
  report_id    = "string"
}