resource "konnect_portal_auth" "my_portalauth" {
  basic_auth_enabled      = true
  konnect_mapping_enabled = false
  oidc_auth_enabled       = false
  oidc_claim_mappings = {
    email  = "email"
    groups = "custom-group-claim"
    name   = "name"
  }
  oidc_client_id     = "...my_oidc_client_id..."
  oidc_client_secret = "...my_oidc_client_secret..."
  oidc_issuer        = "...my_oidc_issuer..."
  oidc_scopes = [
    "...",
  ]
  oidc_team_mapping_enabled = true
  portal_id                 = "170048b6-07f4-48ea-80d1-8e1ac97517b6"
}