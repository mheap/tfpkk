resource "konnect_plugin_exit_transformer" "my_pluginexittransformer" {
  config = {
    functions = [
      "...",
    ]
    handle_unexpected = true
    handle_unknown    = true
  }
  control_plane_id = "b61e38c8-a40f-43d0-95fe-8244b6b57de0"
  enabled          = false
}