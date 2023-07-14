// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testExampleResource,
			},
			// Attempting an update with the same config, expect noop
			{
				Config: testExampleResource,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("scaffolding_example.test", plancheck.ResourceActionNoop),
					},
				},
			},
			// Attempting an update with timeouts added, would expect in-place update but replace occur instead
			{
				Config: testExampleResourceWithTimeouts,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("scaffolding_example.test", plancheck.ResourceActionUpdate),
					},
				},
			},
		},
	})
}

const testExampleResource = `resource "scaffolding_example" "test" {}`
const testExampleResourceWithTimeouts = `resource "scaffolding_example" "test" {
  timeouts = {
    create = "60s"
  }
}`
