// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package generate

import (
	"github.com/Pedro-lmso-erp/erp/src/tools/logging"
)

const (
	// erpPath is the go import path of the base erp package
	erpPath = "github.com/Pedro-lmso-erp/erp"
	// ModelsPath is the go import path of the erp/models package
	ModelsPath = erpPath + "/src/models"
	// DatesPath is the go import path of the erp/models/types/dates package
	DatesPath = erpPath + "/src/models/types/dates"
	// PoolPath is the go import path of the autogenerated pool package
	PoolPath = "github.com/Pedro-lmso-erp/pool"
	// PoolModelPackage is the name of the pool package with model data
	PoolModelPackage = "h"
	// PoolQueryPackage is the name of the pool package with query dat
	PoolQueryPackage = "q"
	// PoolInterfacesPackage is the name of the pool packages with all model interfaces
	PoolInterfacesPackage = "m"
)

var (
	log logging.Logger
	// ModelMixins are the names of the mixins declared in the models package
	ModelMixins = map[string]bool{
		"CommonMixin":    true,
		"BaseMixin":      true,
		"ModelMixin":     true,
		"TransientMixin": true,
	}
	// MethodsToAdd are methods that are declared directly in the generated code.
	// Usually this is because they can't be declared in base_model due to not convertible arg or return types.
	methodsToAdd = map[string]bool{
		"Aggregates": true,
	}
)

func init() {
	log = logging.GetLogger("tools/generate")
}
