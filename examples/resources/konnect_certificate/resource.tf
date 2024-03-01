resource "konnect_certificate" "my_certificate" {
  cert             = "...my_cert..."
  cert_alt         = "...my_cert_alt..."
  control_plane_id = "0a225645-09c2-47f8-8438-1f8d1881e2d9"
  key              = "...my_key..."
  key_alt          = "...my_key_alt..."
}