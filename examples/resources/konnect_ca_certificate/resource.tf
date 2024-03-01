resource "konnect_ca_certificate" "my_cacertificate" {
  cert             = "...my_cert..."
  cert_digest      = "...my_cert_digest..."
  control_plane_id = "5784651a-3c5b-44a8-bd71-693bd366a47d"
}