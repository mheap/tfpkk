resource "konnect_upstream" "my_upstream" {
  algorithm                 = "least-connections"
  control_plane_id          = "72613fae-60e7-4e8c-bdae-5ccdf3b1488d"
  hash_fallback             = "cookie"
  hash_fallback_header      = "...my_hash_fallback_header..."
  hash_fallback_query_arg   = "...my_hash_fallback_query_arg..."
  hash_fallback_uri_capture = "...my_hash_fallback_uri_capture..."
  hash_on                   = "cookie"
  hash_on_cookie            = "...my_hash_on_cookie..."
  hash_on_cookie_path       = "...my_hash_on_cookie_path..."
  hash_on_header            = "...my_hash_on_header..."
  hash_on_query_arg         = "...my_hash_on_query_arg..."
  hash_on_uri_capture       = "...my_hash_on_uri_capture..."
  host_header               = "...my_host_header..."
  name                      = "Ms. Isabel Dooley"
  slots                     = 2
  use_srv_name              = true
}