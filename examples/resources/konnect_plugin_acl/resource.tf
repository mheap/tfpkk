resource "konnect_plugin_acl" "my_pluginacl" {
  config = {
    allow = [
      "...",
    ]
    deny = [
      "...",
    ]
    hide_groups_header      = false
    include_consumer_groups = false
  }
  control_plane_id = "47f68f22-2d7a-4337-adaa-8f99e761e948"
  enabled          = false
}