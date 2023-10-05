data "konnect_service" "my_service" {
  ca_certificates = [
    "...",
  ]
  client_certificate = {
    id = "68047b4a-ee85-40e6-b47c-d76d7e58a875"
  }
  id               = "22220ac4-5556-4c1f-9c73-a004321d7b9c"
  runtime_group_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  tags = [
    "...",
  ]
  tls_verify       = false
  tls_verify_depth = 6
  url              = "...my_url..."
}