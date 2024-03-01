resource "konnect_plugin_basic_auth" "my_pluginbasicauth" {
  config = {
    anonymous        = "...my_anonymous..."
    hide_credentials = true
  }
  control_plane_id = "2ffa6239-c32b-4362-8f48-82477011de20"
  enabled          = true
}