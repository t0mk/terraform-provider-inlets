package inlets

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var packetTestProjectEnvVar = "PACKET_TEST_PROJECT"

func TestAccPacketHostBasic(t *testing.T) {
	testUD := `#cloud-config
runcmd:
 - [ ls, -l, / ]
`
	resourceName := "inlets_packet.test"
	proj := os.Getenv(packetTestProjectEnvVar)
	if proj == "" {
		t.Fatalf("Set %s to test in Packet", packetTestProjectEnvVar)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckPacket(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPacketHostBasic_Config(testUD, proj),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ip"),
				),
			},
		},
	})
}

func testAccPacketHostBasic_Config(ud, proj string) string {
	return fmt.Sprintf(`
locals {
	ud = <<EOS
%sEOS
}

resource "inlets_packet" "test" {
  name             = "tfacc-inlets-packet-device"
  userdata         = local.ud
  additional       = {"project_id": "%s"}
}`, ud, proj)
}
