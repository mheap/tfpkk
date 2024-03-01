resource "konnect_plugin_basic_auth" "my_pluginbasicauth" {
  config = {
    anonymous        = "...my_anonymous..."
    hide_credentials = false
  }
  control_plane_id = "4b414b98-adfc-4ac1-b20c-f8fe297d892f"
  enabled          = false
}