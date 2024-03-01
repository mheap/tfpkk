resource "konnect_plugin_key_auth" "my_pluginkeyauth" {
  config = {
    anonymous        = "...my_anonymous..."
    hide_credentials = false
    key_in_body      = true
    key_in_header    = false
    key_in_query     = true
    key_names = [
      "...",
    ]
    run_on_preflight = false
  }
  control_plane_id = "ebf7e058-032e-4a6f-a6c6-162d8cc023f8"
  enabled          = false
}