terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "0.19.5"
    }
  }
}

provider "konnect" {
  # Configuration options
}