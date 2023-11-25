// build +windows
package irsdk

import (
	"github.com/mpapenbr/goirsdk/yaml"
)

type Irsdk interface {
	Close()

	WaitForValidData()

	GetData() bool

	GetYaml() (*yaml.IrsdkYaml, error)

	GetLatestYaml() *yaml.IrsdkYaml

	GetYamlString() string

	RepairedYaml(s string) string

	GetValue(name string) (any, error)
	GetIntValue(name string) (int32, error)
	GetIntValues(name string) ([]int32, error)
	GetFloatValue(name string) (float32, error)
	GetDoubleValue(name string) (float64, error)
	GetDoubleValues(name string) ([]float64, error)
	GetFloatValues(name string) ([]float32, error)
	GetBoolValue(name string) (bool, error)
	GetBoolValues(name string) ([]bool, error)
	GetValueKeys() []string
}
