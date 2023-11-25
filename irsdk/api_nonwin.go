//go:build !windows
// +build !windows

package irsdk

import (
	"errors"

	"github.com/mpapenbr/goIrsdk/yaml"
)

var ErrNotSupported = errors.New("not supported on this platform")

type IrsdkDummy struct {
	IrsdkDummyYaml yaml.IrsdkYaml // parsed yaml data
}

func (IrsdkDummy *IrsdkDummy) Close() {
}

func (IrsdkDummy *IrsdkDummy) WaitForValidData() bool {
	return false
}

// returns true if new valid data is copied from iRacing telemetry to this Irdsk struct
//
//nolint:lll  //by design
func (IrsdkDummy *IrsdkDummy) GetData() bool {
	return false
}

func (IrsdkDummy *IrsdkDummy) GetYaml() (*yaml.IrsdkDummyYaml, error) {
	return &IrsdkDummy.IrsdkDummyYaml, nil
}

func (IrsdkDummy *IrsdkDummy) GetLatestYaml() *yaml.IrsdkDummyYaml {
	return &IrsdkDummy.IrsdkDummyYaml
}

func (IrsdkDummy *IrsdkDummy) GetYamlString() string {
	return ""
}

func (IrsdkDummy *IrsdkDummy) RepairedYaml(s string) string {
	return s
}

func (IrsdkDummy *IrsdkDummy) GetValue(name string) (any, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetIntValue(name string) (int32, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetIntValues(name string) ([]int32, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetFloatValue(name string) (float32, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetDoubleValue(name string) (float64, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetDoubleValues(name string) ([]float64, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetFloatValues(name string) ([]float32, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetBoolValue(name string) (bool, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetBoolValues(name string) ([]bool, error) {
	return nil, ErrNotSupported
}

func (IrsdkDummy *IrsdkDummy) GetValueKeys() []string {
	return nil, ErrNotSupported
}
