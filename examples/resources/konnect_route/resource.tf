resource "konnect_route" "my_route" {
  control_plane_id           = "0ce8052a-e316-4190-a7be-94aa09997799"
  https_redirect_status_code = "308"
  name                       = "Marshall Price"
  path_handling              = "v0"
  preserve_host              = false
  regex_priority             = 4
  request_buffering          = false
  response_buffering         = true
  strip_path                 = true
}