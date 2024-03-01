resource "konnect_service" "my_service" {
  ca_certificates = [
    "...",
  ]
  client_certificate = {
    id = "4b3676d9-dc7a-453e-875a-7c2c465dec6b"
  }
  connect_timeout  = 5
  control_plane_id = "cafb4e58-66ce-489d-b047-31b77b31cb14"
  enabled          = false
  host             = "...my_host..."
  name             = "Estelle Padberg"
  path             = "...my_path..."
  port             = 1
  protocol         = "https"
  read_timeout     = 1
  retries          = 1
  tags = [
    "...",
  ]
  tls_verify       = false
  tls_verify_depth = 8
  url              = "...my_url..."
  write_timeout    = 8
}