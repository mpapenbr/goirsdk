package irsdk

import (
	"testing"
)

func Test_fixYaml(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"standard",
			args{`
		DriverInfo:
		  DriverCarIdx: 0
		  Drivers: 
		  	- CarIdx: 0
			  UserName: John Doe
			  TeamName: @Invalid Yaml team name
		  	- CarIdx: 1
			  UserName: Jane Doe
			  TeamName: Valid Yaml team name
		`},
			`
		DriverInfo:
		  DriverCarIdx: 0
		  Drivers: 
		  	- CarIdx: 0
			  UserName: "John Doe"
			  TeamName: "@Invalid Yaml team name"
		  	- CarIdx: 1
			  UserName: "Jane Doe"
			  TeamName: "Valid Yaml team name"
		`,
		},
	}
	irsdk := &Irsdk{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := irsdk.RepairedYaml(tt.args.s); got != tt.want {
				t.Errorf("fixYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}
