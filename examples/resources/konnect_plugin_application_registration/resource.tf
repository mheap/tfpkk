resource "konnect_plugin_application_registration" "my_pluginapplicationregistration" {
  config = {
    auto_approve = true
    description  = "...my_description..."
    display_name = "...my_display_name..."
    show_issuer  = false
  }
  control_plane_id = "73147a05-e057-48dd-a90d-56ee0701be33"
  enabled          = true
}