resource "konnect_route" "my_route" {
  control_plane_id           = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  https_redirect_status_code = "426"
  name                       = "Nicholas Rice"
  path_handling              = "v0"
  preserve_host              = false
  regex_priority             = 0
  request_buffering          = true
  response_buffering         = false
  strip_path                 = true
}