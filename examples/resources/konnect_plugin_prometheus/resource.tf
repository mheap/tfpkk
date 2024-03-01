resource "konnect_plugin_prometheus" "my_pluginprometheus" {
  config = {
    bandwidth_metrics       = false
    latency_metrics         = true
    per_consumer            = true
    status_code_metrics     = false
    upstream_health_metrics = true
  }
  control_plane_id = "03cf7925-6576-4d53-8d35-5e9e09c1380f"
  enabled          = false
}