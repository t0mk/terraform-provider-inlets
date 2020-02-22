package inlets

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/inlets/inletsctl/pkg/provision"
)

func resourcePacket() *schema.Resource {
	sch := baseHostSchema()
	return &schema.Resource{
		Create: resourcePacketCreate,
		Read:   resourcePacketRead,
		Delete: resourcePacketDelete,
		Schema: sch,
	}
}

func resourcePacketCreate(d *schema.ResourceData, meta interface{}) error {
	provisioner := meta.(Config).getPacketProvisioner()

	hostConf, err := resourceDataToBasicHost(d, "packet")
	if err != nil {
		return err
	}
	if hostConf == nil {
		return fmt.Errorf("Couldn't process resource params to host configuration")
	}

	res, err := provisioner.Provision(*hostConf)
	if err != nil {
		return err
	}
	d.SetId(res.ID)

	stateAfterWait, err := waitForHostState(d, []string{"active", "failed"}, []string{"queued", "provisioning"},
		provisioner, res.ID)
	if err != nil {
		d.SetId("")
		return err
	}
	if stateAfterWait != "active" {
		d.SetId("")
		return fmt.Errorf("Provisioning of host %s failed", res.ID)
	}
	return resourcePacketRead(d, meta)
}

func resourcePacketRead(d *schema.ResourceData, meta interface{}) error {
	provisioner := meta.(Config).getPacketProvisioner()
	ph, err := provisioner.Status(d.Id())

	if err != nil {
		return err
	}

	d.Set("ip", ph.IP)
	d.Set("status", ph.Status)

	return nil
}

func resourcePacketDelete(d *schema.ResourceData, meta interface{}) error {
	adds := d.Get("additional").(map[string]interface{})

	hdr := provision.HostDeleteRequest{
		ID:        d.Id(),
		IP:        d.Get("ip").(string),
		ProjectID: adds["project_id"].(string),
	}

	provisioner := meta.(Config).getPacketProvisioner()

	err := provisioner.Delete(hdr)
	if err != nil {
		return err
	}
	return nil
}
