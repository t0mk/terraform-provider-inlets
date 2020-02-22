package inlets

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"inlets_data_source": dataSourceScaffolding(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"inlets_resource": resourceScaffolding(),
		},
	}
}
