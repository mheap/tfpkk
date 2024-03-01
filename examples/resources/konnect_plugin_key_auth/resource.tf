resource "konnect_plugin_key_auth" "my_pluginkeyauth" {
  config = {
    anonymous        = "...my_anonymous..."
    hide_credentials = false
    key_in_body      = true
    key_in_header    = true
    key_in_query     = false
    key_names = [
      "...",
    ]
    run_on_preflight = false
  }
  control_plane_id = "3147a05e-0578-4dd2-90d5-6ee0701be332"
  enabled          = true
}