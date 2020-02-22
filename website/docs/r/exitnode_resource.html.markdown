---
layout: "inlets"
page_title: "Inlets: exitnode resource"
sidebar_current: "docs-inlets-resource"
description: |-
  Sample resource in the Terraform provider Inlets.
---

# exitnode

Exitnode resource in the Terraform provider inlets.

## Example Usage

```hcl
resource "inlets_exitnode" "example" {
  sample_attribute = "foo"
}
```

## Argument Reference

The following arguments are supported:

* `sample_attribute` - Sample attribute.

