module github.com/CiscoDevNet/terraform-provider-intersight

go 1.15

require (
	github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.2
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
)

replace github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0 => ./intersight_gosdk
