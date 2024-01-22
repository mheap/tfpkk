terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "0.7.9"
    }
  }
}

provider "konnect" {
  # Configuration options
}