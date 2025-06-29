// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glue_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccGlue_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"CatalogTableOptimizer": {
			acctest.CtBasic:                 testAccCatalogTableOptimizer_basic,
			"deleteOrphanFileConfiguration": testAccCatalogTableOptimizer_DeleteOrphanFileConfiguration,
			acctest.CtDisappears:            testAccCatalogTableOptimizer_disappears,
			"retentionConfiguration":        testAccCatalogTableOptimizer_RetentionConfiguration,
			"update":                        testAccCatalogTableOptimizer_update,
		},
		"DataCatalogEncryptionSettings": {
			acctest.CtBasic: testAccDataCatalogEncryptionSettings_basic,
			"dataSource":    testAccDataCatalogEncryptionSettingsDataSource_basic,
		},
		"ResourcePolicy": {
			acctest.CtBasic:      testAccResourcePolicy_basic,
			"update":             testAccResourcePolicy_update,
			"hybrid":             testAccResourcePolicy_hybrid,
			acctest.CtDisappears: testAccResourcePolicy_disappears,
			"equivalent":         testAccResourcePolicy_ignoreEquivalent,
			"Identity":           testAccGlueResourcePolicy_IdentitySerial,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
