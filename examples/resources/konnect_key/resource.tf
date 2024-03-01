resource "konnect_key" "my_key" {
  control_plane_id = "dee1b0c6-7928-48ab-9027-b2087a38791d"
  jwk              = "...my_jwk..."
  kid              = "...my_kid..."
  name             = "Austin Wisozk"
}