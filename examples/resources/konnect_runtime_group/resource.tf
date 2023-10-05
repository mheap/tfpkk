resource "konnect_runtime_group" "my_runtimegroup" {
  auth_type    = "pinned_client_certs"
  cluster_type = "CLUSTER_TYPE_HYBRID"
  description  = "A test runtime group for exploration."
  labels       = "{ \"see\": \"documentation\" }"
  name         = "Test Runtime Group"
}