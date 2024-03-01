terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "0.19.4"
    }
  }
}

provider "konnect" {
  # Configuration options
}