resource "konnect_portal" "my_portal" {
  auto_approve_applications            = true
  auto_approve_developers              = false
  custom_client_domain                 = "private-keep.net"
  custom_domain                        = "cute-scallion.biz"
  default_application_auth_strategy_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  is_public                            = false
  portal_id                            = "b31cb148-4dae-4625-911b-c6bb4fac99f0"
  rbac_enabled                         = false
}