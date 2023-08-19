terraform {
  required_providers {
    Konnect = {
      source  = "Kong/Konnect"
      version = "0.2.2"
    }
  }
}

provider "Konnect" {
  # Configuration options
}