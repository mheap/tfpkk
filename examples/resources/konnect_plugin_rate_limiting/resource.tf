resource "konnect_plugin_rate_limiting" "my_pluginratelimiting" {
  config = {
    day                 = 12.89
    error_code          = 96.67
    error_message       = "...my_error_message..."
    fault_tolerant      = true
    header_name         = "...my_header_name..."
    hide_client_headers = false
    hour                = 39.53
    limit_by            = "credential"
    minute              = 20.49
    month               = 60.23
    path                = "...my_path..."
    policy              = "redis"
    redis_database      = 2
    redis_host          = "...my_redis_host..."
    redis_password      = "...my_redis_password..."
    redis_port          = 1
    redis_server_name   = "...my_redis_server_name..."
    redis_ssl           = false
    redis_ssl_verify    = true
    redis_timeout       = 41.49
    redis_username      = "...my_redis_username..."
    second              = 18.67
    sync_rate           = 1.27
    year                = 97.97
  }
  control_plane_id = "48824770-11de-4205-b77e-a8940adb9b2c"
  enabled          = false
}