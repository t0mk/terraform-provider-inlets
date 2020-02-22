package inlets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/inlets/inletsctl/pkg/provision"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"packet_auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PACKET_AUTH_TOKEN", nil),
			},
			"digitalocean_auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DO_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"inlets_packet": resourcePacket(),
		},
		ConfigureFunc: providerConfigure,
	}
}

type Config struct {
	PacketAuthToken string
	DOAuthToken     string
	//...
}

func (c Config) getPacketProvisioner() *provision.PacketProvisioner {
	pp, _ := provision.NewPacketProvisioner(c.PacketAuthToken)
	return pp
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		PacketAuthToken: d.Get("packet_auth_token").(string),
		DOAuthToken:     d.Get("digitalocean_auth_token").(string),
	}
	return config, nil
}
