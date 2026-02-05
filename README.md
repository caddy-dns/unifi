UniFi module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with UniFi Network.

## Caddy module name

```
dns.providers.unifi
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "unifi",
				"api_key": "YOUR_PROVIDER_API_KEY",
				"base_url": "YOUR_UNIFI_NETWORK_API_URL",
				"site_id": "YOUR_UNIFI_SITE_ID"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns unifi {
		api_key {env.YOUR_PROVIDER_API_KEY}
		base_url {env.YOUR_UNIFI_NETWORK_API_URL}
		site_id {env.YOUR_UNIFI_SITE_ID}
	}
}
```

```
# one site
tls {
	dns unifi {
		api_key {env.YOUR_PROVIDER_API_KEY}
		base_url {env.YOUR_UNIFI_NETWORK_API_URL}
		site_id {env.YOUR_UNIFI_SITE_ID}
	}
}
```

## More information
For information on what values to use for the configuration options, check out the [libdns/unifi](https://github.com/libdns/unifi) package.
