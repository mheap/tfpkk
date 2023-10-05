data "konnect_portal_auth" "my_portalauth" {
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
  portal_id = "99386df2-a43e-42b9-b37e-1745c41205d9"
}