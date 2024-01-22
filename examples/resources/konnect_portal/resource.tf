resource "konnect_portal" "my_portal" {
  auto_approve_applications = false
  auto_approve_developers   = true
  custom_client_domain      = "submissive-neuropsychiatry.biz"
  custom_domain             = "evergreen-bladder.org"
  is_public                 = true
  portal_id                 = "c2314257-8465-41a3-85b4-a87d71693bd3"
  rbac_enabled              = false
}