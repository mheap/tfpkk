resource "konnect_jwt" "my_jwt" {
  algorithm        = "HS256"
  control_plane_id = "d9d514a9-5a5c-4a8f-b2eb-0c0f7ab0b23a"
  key              = "...my_key..."
  rsa_public_key   = "...my_rsa_public_key..."
  secret           = "...my_secret..."
}