resource "konnect_route" "my_route" {
  created_at = 8
  headers    = {}
  hosts = [
    "foo.example.com",
  ]
  https_redirect_status_code = 10
  id                         = "56c4566c-14cc-4132-9011-4139fcbbe50a"
  methods = [
    "...",
  ]
  name          = "Ada Jast"
  path_handling = "...my_path_handling..."
  paths = [
    "...",
  ]
  preserve_host = false
  protocols = [
    "...",
  ]
  regex_priority     = 5
  request_buffering  = false
  response_buffering = false
  runtime_group_id   = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  service = {
    id = "eab13803-3779-407e-8d63-0db43275a870"
  }
  snis = [
    "...",
  ]
  strip_path = false
  tags = [
    "...",
  ]
  updated_at = 7
}