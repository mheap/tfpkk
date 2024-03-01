resource "konnect_plugin_ai_proxy" "my_pluginaiproxy" {
  config = {
    auth = {
      header_name    = "...my_header_name..."
      header_value   = "...my_header_value..."
      param_location = "body"
      param_name     = "...my_param_name..."
      param_value    = "...my_param_value..."
    }
    logging = {
      log_payloads   = true
      log_statistics = false
    }
    model = {
      name = "Ms. Wm Rolfson"
      options = {
        anthropic_version   = "...my_anthropic_version..."
        azure_api_version   = "...my_azure_api_version..."
        azure_deployment_id = "...my_azure_deployment_id..."
        azure_instance      = "...my_azure_instance..."
        llama2_format       = "raw"
        max_tokens          = 6
        mistral_format      = "openai"
        temperature         = 48.68
        top_k               = 5
        top_p               = 45.65
        upstream_url        = "...my_upstream_url..."
      }
      provider = "azure"
    }
    route_type = "llm/v1/completions"
  }
  control_plane_id = "414b98ad-fcac-4172-8cf8-fe297d892f5c"
  enabled          = true
}