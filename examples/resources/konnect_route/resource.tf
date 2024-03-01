resource "konnect_route" "my_route" {
  control_plane_id           = "64f9a795-812c-4ac8-9640-320bb6b8f1cf"
  https_redirect_status_code = "307"
  name                       = "Andrea Swift"
  path_handling              = "v1"
  preserve_host              = false
  regex_priority             = 3
  request_buffering          = true
  response_buffering         = false
  strip_path                 = true
}