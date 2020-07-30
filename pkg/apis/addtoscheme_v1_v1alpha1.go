package apis

import (
	"github.com/epmd-edp/go-go-operator-sdk-postgresql/pkg/apis/v1/v1alpha1"
)

func init() {
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
