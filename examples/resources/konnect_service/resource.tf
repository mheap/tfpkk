resource "konnect_service" "my_service" {
  ca_certificates = [
    "...",
  ]
  client_certificate = {
    id = "ae90248b-41b6-43f9-80d7-66a42b1f9779"
  }
  connect_timeout  = 6
  control_plane_id = "e7f64e8e-9cbc-48cd-95e3-42c7dd4186c1"
  enabled          = false
  host             = "...my_host..."
  name             = "April Leffler"
  path             = "...my_path..."
  port             = 6
  protocol         = "http"
  read_timeout     = 10
  retries          = 2
  tags = [
    "...",
  ]
  tls_verify       = false
  tls_verify_depth = 6
  url              = "...my_url..."
  write_timeout    = 0
}