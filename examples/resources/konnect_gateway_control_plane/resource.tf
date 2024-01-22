resource "konnect_gateway_control_plane" "my_gatewaycontrolplane" {
  auth_type     = "pinned_client_certs"
  cloud_gateway = false
  cluster_type  = "CLUSTER_TYPE_HYBRID"
  description   = "A test control plane for exploration."
  labels        = "{ \"see\": \"documentation\" }"
  name          = "Test Control Plane"
}