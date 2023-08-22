resource "konnect_runtime_group" "my_runtimegroup" {
    cluster_type = "CLUSTER_TYPE_HYBRID"
            description = "A test runtime group for exploration."
            labels = "{ \"see\": \"documentation\" }"
            name = "Test Runtime Group"
        }