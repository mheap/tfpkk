terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "0.2.2"
    }
  }
}

provider "konnect" {
  # Configuration options
}