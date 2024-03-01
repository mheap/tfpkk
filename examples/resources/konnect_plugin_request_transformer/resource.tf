resource "konnect_plugin_request_transformer" "my_pluginrequesttransformer" {
  config = {
    add = {
      body = [
        "...",
      ]
      headers = [
        "...",
      ]
      querystring = [
        "...",
      ]
    }
    append = {
      body = [
        "...",
      ]
      headers = [
        "...",
      ]
      querystring = [
        "...",
      ]
    }
    http_method = "...my_http_method..."
    remove = {
      body = [
        "...",
      ]
      headers = [
        "...",
      ]
      querystring = [
        "...",
      ]
    }
    rename = {
      body = [
        "...",
      ]
      headers = [
        "...",
      ]
      querystring = [
        "...",
      ]
    }
    replace = {
      body = [
        "...",
      ]
      headers = [
        "...",
      ]
      querystring = [
        "...",
      ]
      uri = "http://parallel-mobility.name"
    }
  }
  control_plane_id = "7799dc9b-5606-450c-a4b3-676d9dc7a53e"
  enabled          = false
}