resource "konnect_plugin_cors" "my_plugincors" {
  config = {
    credentials = true
    exposed_headers = [
      "...",
    ]
    headers = [
      "...",
    ]
    max_age = 48.63
    methods = [
      "POST",
    ]
    origins = [
      "...",
    ]
    preflight_continue = false
    private_network    = true
  }
  control_plane_id = "8940adb9-b2c8-45e4-8b66-3982d5de80c6"
  enabled          = true
}