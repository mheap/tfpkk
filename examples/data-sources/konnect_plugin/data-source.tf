data "konnect_plugin" "my_plugin" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  filter_tags      = "...my_filter_tags..."
  id               = "response-ratelimiting"
  offset           = "...my_offset..."
  size             = 2
}