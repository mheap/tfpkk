data "konnect_consumer" "my_consumer" {
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  filter_tags      = "...my_filter_tags..."
  id               = "e94215b0-9198-42ce-bf0b-98afff31b2a7"
  offset           = "...my_offset..."
  size             = 5
}