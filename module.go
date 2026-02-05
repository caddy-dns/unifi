package unifi

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/unifi"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *unifi.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.template",
		New: func() caddy.Module { return &Provider{new(unifi.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.ApiKey = repl.ReplaceAll(p.ApiKey, "")
	p.BaseUrl = repl.ReplaceAll(p.BaseUrl, "")
	p.SiteId = repl.ReplaceAll(p.SiteId, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	unifi {
//	    api_key string
//	    base_url string
//	    site_id string
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if d.NextArg() {
					p.ApiKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "base_url":
				if d.NextArg() {
					p.BaseUrl = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "site_id":
				if d.NextArg() {
					p.SiteId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.ApiKey == "" {
		return d.Err("missing API Key")
	}
	if p.BaseUrl == "" {
		return d.Err("missing base URL")
	}
	if p.SiteId == "" {
		return d.Err("missing site ID")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
