package utils

import (
	"github.com/fatih/structs"
	"github.com/hyperdidact/gotty/backend/localcommand"
	"github.com/hyperdidact/gotty/cmd"
	"github.com/hyperdidact/gotty/server"
	"github.com/spf13/cobra"
	"testing"
)

type Test struct {
	name string
	cmd  *cobra.Command
}

type ExpectedFlagInfo struct {
	Name         string
	Shorthand    string
	Description  string
	DefaultValue string
}

func TestGenerateCobraFlags(t *testing.T) {
	tests := []struct {
		Name                 string
		VariadicStructsInput []interface{}
		ExpectedFlags         map[string]ExpectedFlagInfo
	}{
		{
			Name: "string (anonymous)",
			VariadicStructsInput: []interface{}{
				struct {
					Address string `hcl:"address" flagName:"address" flagSName:"a" flagDescribe:"IP address to listen" default:"0.0.0.0"`
				}{},
			},
			ExpectedFlags: map[string]ExpectedFlagInfo{
				"address": {
					Name:         "address",
					Shorthand:    "a",
					Description:  "IP address to listen",
					DefaultValue: "0.0.0.0",
				},
			},
		},
		{
			Name: "bool (anonymous)",
			VariadicStructsInput: []interface{}{
				struct {
					PermitWrite bool `hcl:"permit_write" flagName:"permit-write" flagSName:"w" flagDescribe:"Permit clients to write to the TTY (BE CAREFUL)" default:"false"`
				}{},
			},
			ExpectedFlags: map[string]ExpectedFlagInfo{
				"permit-write": {
				Name:         "permit-write",
				Shorthand:    "w",
				Description:  "Permit clients to write to the TTY (BE CAREFUL)",
				DefaultValue: "false",
			},},

		},
		{
			Name: "int (anonymous)",
			VariadicStructsInput: []interface{}{
				struct {
					RandomUrlLength int `hcl:"random_url_length" flagName:"random-url-length" flagDescribe:"Random URL length" default:"8"`
				}{},
			},
			ExpectedFlags: map[string]ExpectedFlagInfo{
				"random-url-length": {
					Name:         "random-url-length",
					Description:  "Random URL length",
					DefaultValue: "8",
				},},

		},
		{
			Name: "server options",
			VariadicStructsInput: []interface{}{
				server.Options{},
			},
			ExpectedFlags: generateExpectedFlags(server.Options{}),
		},
		{
			Name: "localcommand option",
			VariadicStructsInput: []interface{}{
				localcommand.Options{},
			},
			ExpectedFlags: generateExpectedFlags(localcommand.Options{}),
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			testOptionSlice := test.VariadicStructsInput
			for idx := range testOptionSlice {
				e := &(testOptionSlice[idx])
				if err := ApplyDefaultValues(*e); err != nil {
					t.Errorf("failed to apply default values: %v", err)
					return
				}
			}

			root := cmd.NewRootCmd("")
			gotCmd, _, err := GenerateCobraFlags(root, testOptionSlice...)
			if err != nil {
				t.Errorf("failed to generate cobra flags: %v", err)
				return
			}
			for name, wantFlag := range test.ExpectedFlags {
				gotFlag := gotCmd.Flag(name)

				if gotFlag == nil {
					t.Errorf("Wanted a flag named '%v', but didn't get it.", name)
					//we're done, can't test further
					return
				}

				if gotFlag.Name != wantFlag.Name {
					t.Errorf("Wanted a flag named '%v', but got '%v'", wantFlag.Name, gotFlag.Name )
					t.Fail()
				}

				if gotFlag.Shorthand != wantFlag.Shorthand {
					t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", name, wantFlag.Shorthand, gotFlag.Shorthand)
					t.Fail()
				}

				if gotFlag.DefValue != wantFlag.DefaultValue {
					t.Errorf("Wanted flag '%v' to have a default value of '%v', but got '%v'", name, wantFlag.DefaultValue , gotFlag.DefValue)
					t.Fail()
				}

				if gotFlag.Usage != wantFlag.Description {
					t.Errorf("Wanted flag '%v' to have a description of '%v', but got '%v'", name, wantFlag.Description , gotFlag.Usage)
					t.Fail()
				}
				t.Logf("Flag '%v', Shorthand '%v', Default Value '%v', Description '%v'", name, gotFlag.Shorthand, gotFlag.DefValue, gotFlag.Usage)
			}



		})
	}
}

func generateExpectedFlags(options ...interface{}) map[string]ExpectedFlagInfo {
	 var result map[string]ExpectedFlagInfo = map[string]ExpectedFlagInfo{}

	for _, struct_ := range options {
		o := structs.New(struct_)
		for _, field := range o.Fields() {
			flagName := field.Tag("flagName")
			if flagName == "" {
				continue
			}
			flag := ExpectedFlagInfo{
				Name:         flagName,
				Shorthand:    field.Tag("flagSName"),
				Description:  field.Tag("flagDescribe"),
				DefaultValue: field.Tag("default"),
			}
			result[flagName] = flag

		}}
	return result
}


