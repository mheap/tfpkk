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
      log_payloads   = false
      log_statistics = true
    }
    model = {
      name = "Shawna Howe"
      options = {
        anthropic_version   = "...my_anthropic_version..."
        azure_api_version   = "...my_azure_api_version..."
        azure_deployment_id = "...my_azure_deployment_id..."
        azure_instance      = "...my_azure_instance..."
        llama2_format       = "raw"
        max_tokens          = 1
        mistral_format      = "openai"
        temperature         = 87.03
        top_k               = 5
        top_p               = 66.26
        upstream_url        = "...my_upstream_url..."
      }
      provider = "azure"
    }
    route_type = "llm/v1/chat"
  }
  control_plane_id = "72daa8f9-9e76-41e9-88d9-18ebc0992a67"
  enabled          = true
}