resource "konnect_vault" "my_vault" {
  config           = "{ \"see\": \"documentation\" }"
  control_plane_id = "3e91c3d5-d6c8-40f5-ae90-43d79c5ff0ad"
  description      = "...my_description..."
  name             = "Forrest Russel"
  prefix           = "...my_prefix..."
}