Terraform Provider Inlets
==================

Building The Provider
---------------------

Clone repository and run `go build` in root.

Run `make` to install.

Using the provider
----------------------

After you installed (with `make` in root of this repo), go to your Terraform directory and run `tf init -plugin-dir $GOPATH/bin`.

Example
---------------------------

Only `inlets_packet` is working right now. To run a host in Packet, you should 
- export your Packet API token to PACKET_AUTH_TOKEN envvar
- choose a project where you want to run the host and find out its UUID
- put userdata to `./userdata.yml`
- use following ttemplate:

```hcl
resource "inlets_packet" test {
    name = "inletstest"
    userdata = file("./userdata.yml") 
    additional = {"project_id": "52123fb2-ee46-4673-93a8-de2c2bdba33b"}
}
```

Notes
--------------
The provider assumes a lot of defaults, see the top of [host_common.go](inlets/host_common.go)

