resource "konnect_service" "my_service" {
  ca_certificates = [
    "...",
  ]
  client_certificate = {
    id = "d83db6eb-aaee-41a3-86b1-d4eec2f5947f"
  }
  connect_timeout  = 4
  created_at       = 1422386534
  enabled          = false
  host             = "...my_host..."
  id               = "dd6fee3f-9487-4643-914a-c664acba9111"
  name             = "Robert Denesik"
  path             = "...my_path..."
  port             = 1
  protocol         = "...my_protocol..."
  read_timeout     = 9
  retries          = 6
  runtime_group_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  tags = [
    "...",
  ]
  tls_verify       = true
  tls_verify_depth = 3
  updated_at       = 5
  url              = "...my_url..."
  write_timeout    = 2
}