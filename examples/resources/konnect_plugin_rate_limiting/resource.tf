resource "konnect_plugin_rate_limiting" "my_pluginratelimiting" {
  config = {
    day                 = 28.92
    error_code          = 78.53
    error_message       = "...my_error_message..."
    fault_tolerant      = true
    header_name         = "...my_header_name..."
    hide_client_headers = false
    hour                = 62.77
    limit_by            = "consumer-group"
    minute              = 78.36
    month               = 33.09
    path                = "...my_path..."
    policy              = "local"
    redis_database      = 4
    redis_host          = "...my_redis_host..."
    redis_password      = "...my_redis_password..."
    redis_port          = 9
    redis_server_name   = "...my_redis_server_name..."
    redis_ssl           = true
    redis_ssl_verify    = false
    redis_timeout       = 47.1
    redis_username      = "...my_redis_username..."
    second              = 23.44
    sync_rate           = 99.14
    year                = 96.16
  }
  control_plane_id = "023cc3c0-e0ce-4805-aae3-16190a7be94a"
  enabled          = true
}