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
      uri = "http://vacant-emanate.name"
    }
  }
  control_plane_id = "b663982d-5de8-40c6-a3db-61e38c8a40f3"
  enabled          = true
}