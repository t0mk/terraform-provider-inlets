package inlets

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/inlets/inletsctl/pkg/provision"
)

// this could be sourced from imported inlets repo
var defaultParams = map[string]map[string]interface{}{
	"packet": {
		"plan":   "t1.small.x86",
		"region": "ewr1",
		"os":     "ubuntu_16_04",
	},
}

func baseHostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"plan": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"os": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"userdata": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"additional": {
			Type: schema.TypeMap,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			ForceNew: true,
		},

		"ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"zone": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func getParam(d *schema.ResourceData, typ, paramName string) (interface{}, error) {
	if val, ok := d.GetOk(paramName); ok {
		return val, nil
	} else {
		defVal, ok := defaultParams[typ][paramName]
		if !ok {
			return nil, fmt.Errorf("No default for [%s][%s]", typ, paramName)
		}
		return defVal, nil
	}
}

func resourceDataToBasicHost(d *schema.ResourceData, typ string) (*provision.BasicHost, error) {
	bh := provision.BasicHost{}
	region, err := getParam(d, typ, "region")
	if err != nil {
		return nil, err
	}
	bh.Region = region.(string)

	plan, err := getParam(d, typ, "plan")
	if err != nil {
		return nil, err
	}
	bh.Plan = plan.(string)

	os, err := getParam(d, typ, "os")
	if err != nil {
		return nil, err
	}
	bh.OS = os.(string)

	name, err := getParam(d, typ, "name")
	if err != nil {
		return nil, err
	}
	bh.Name = name.(string)

	userData, err := getParam(d, typ, "userdata")
	if err != nil {
		return nil, err
	}
	bh.UserData = userData.(string)

	if additional, ok := d.GetOk("additional"); ok {
		ifMap := additional.(map[string]interface{})
		addMap := map[string]string{}
		for k, v := range ifMap {
			addMap[k] = v.(string)
		}
		bh.Additional = addMap
	} else {
		bh.Additional = nil
	}

	return &bh, nil
}

func waitForHostState(d *schema.ResourceData, targets []string, pending []string,
	p interface{}, id string) (string, error) {
	provisioner := p.(provision.Provisioner)

	stateConf := &resource.StateChangeConf{
		Pending: pending,
		Target:  targets,
		Refresh: func() (interface{}, string, error) {
			h, err := provisioner.Status(id)
			if err == nil {
				return h.Status, h.Status, nil
			}
			return "error", "error", err
		},
		Timeout:    60 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}
	attrval, err := stateConf.WaitForState()

	return attrval.(string), err
}
