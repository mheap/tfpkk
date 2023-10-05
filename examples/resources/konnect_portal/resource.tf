resource "konnect_portal" "my_portal" {
  auto_approve_applications = true
  auto_approve_developers   = false
  custom_client_domain      = "weird-depressive.org"
  custom_domain             = "fantastic-babe.org"
  is_public                 = false
  portal_id                 = "8c3b3ed8-234b-4415-9cde-1e1cd2dd1e30"
  rbac_enabled              = false
}