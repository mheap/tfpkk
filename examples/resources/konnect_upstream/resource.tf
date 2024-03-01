resource "konnect_upstream" "my_upstream" {
  algorithm                 = "least-connections"
  control_plane_id          = "e7f64e8e-9cbc-48cd-95e3-42c7dd4186c1"
  hash_fallback             = "header"
  hash_fallback_header      = "...my_hash_fallback_header..."
  hash_fallback_query_arg   = "...my_hash_fallback_query_arg..."
  hash_fallback_uri_capture = "...my_hash_fallback_uri_capture..."
  hash_on                   = "ip"
  hash_on_cookie            = "...my_hash_on_cookie..."
  hash_on_cookie_path       = "...my_hash_on_cookie_path..."
  hash_on_header            = "...my_hash_on_header..."
  hash_on_query_arg         = "...my_hash_on_query_arg..."
  hash_on_uri_capture       = "...my_hash_on_uri_capture..."
  host_header               = "...my_host_header..."
  name                      = "Hilda Flatley"
  slots                     = 2
  use_srv_name              = false
}