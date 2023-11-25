//go:build !windows
// +build !windows

package irsdk

import (
	"errors"

	"github.com/mpapenbr/goirsdk/yaml"
)

var ErrNotSupported = errors.New("not supported on this platform")

type Irsdk struct {
	irsdkYaml yaml.IrsdkYaml // parsed yaml data
}

type ApiConfig struct {
	WaitForSimTimeout int
}

func NewIrsdk() *Irsdk {
	return &Irsdk{}
}

func NewIrsdkWithConfig(config ApiConfig) *Irsdk {
	return &Irsdk{}
}

func CheckIfSimIsRunning() bool {
	return false
}

func (irsdk *Irsdk) Close() {
}

func (irsdk *Irsdk) WaitForValidData() bool {
	return false
}

// returns true if new valid data is copied from iRacing telemetry to this Irdsk struct
//
//nolint:lll  //by design
func (irsdk *Irsdk) GetData() bool {
	return false
}

func (irsdk *Irsdk) GetYaml() (*yaml.IrsdkYaml, error) {
	return &irsdk.irsdkYaml, nil
}

func (irsdk *Irsdk) GetLatestYaml() *yaml.IrsdkYaml {
	return &irsdk.irsdkYaml
}

func (irsdk *Irsdk) GetYamlString() string {
	return ""
}

func (irsdk *Irsdk) RepairedYaml(s string) string {
	return s
}

func (irsdk *Irsdk) GetValue(name string) (any, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetIntValue(name string) (int32, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetIntValues(name string) ([]int32, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetFloatValue(name string) (float32, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetDoubleValue(name string) (float64, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetDoubleValues(name string) ([]float64, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetFloatValues(name string) ([]float32, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetBoolValue(name string) (bool, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetBoolValues(name string) ([]bool, error) {
	return nil, ErrNotSupported
}

func (irsdk *Irsdk) GetValueKeys() []string {
	return nil, ErrNotSupported
}
