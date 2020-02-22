---
layout: "inlets"
page_title: "Inlets: data source exitnode"
sidebar_current: "docs-inlets-datasource"
description: |-
  Exitnode data source in the Terraform provider Inlets.
---

# exitnode 

Exitnode data source in the Terraform provider scaffolding.

## Example Usage

```hcl
data "inlets_exitnode" "example" {
  sample_attribute = "foo"
}
```

## Attributes Reference

* `sample_attribute` - Sample attribute.
