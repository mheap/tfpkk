resource "konnect_plugin" "my_plugin" {
  config = {
    hour   = 500
    minute = 500
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
  instance_name    = "rate-limiting-foo"
  name             = "rate-limiting"
}