resource "konnect_service" "my_service" {
  ca_certificates = [
    "4e3ad2e4-0bc4-4638-8e34-c84a417ba39b",
  ]
  client_certificate = {
    id = "4e3ad2e4-0bc4-4638-8e34-c84a417ba39b"
  }
  connect_timeout  = 6000
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
  host             = "example.com"
  name             = "my-service"
  path             = "/some_api"
  port             = 80
  protocol         = "http"
  read_timeout     = 6000
  retries          = 5
  tags = [
    "user-level",
  ]
  tls_verify       = true
  tls_verify_depth = "respected"
  write_timeout    = 6000
}