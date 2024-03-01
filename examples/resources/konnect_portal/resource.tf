resource "konnect_portal" "my_portal" {
  auto_approve_applications            = false
  auto_approve_developers              = false
  custom_client_domain                 = "grumpy-water.org"
  custom_domain                        = "lumbering-circuit.biz"
  default_application_auth_strategy_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  is_public                            = false
  portal_id                            = "c6162d8c-c023-4f8b-9dc2-003cf7925657"
  rbac_enabled                         = true
}