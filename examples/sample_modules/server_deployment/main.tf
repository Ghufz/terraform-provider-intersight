terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "1.0.46"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkeyfile = var.secretkey
  endpoint = var.endpoint
}

//terraform import intersight_server_profile.server1 <moid>

resource "intersight_server_profile" "server1" {
  name="server1"
  action="Deploy"
}
