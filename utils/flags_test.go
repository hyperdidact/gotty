package utils

import (
	"github.com/hyperdidact/gotty/backend/localcommand"
	"github.com/hyperdidact/gotty/cmd"
	"github.com/hyperdidact/gotty/server"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"

	"github.com/codegangsta/cli"
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

type OptionsString struct {
	Address string `hcl:"address" flagName:"address" flagSName:"a" flagDescribe:"IP address to listen" default:"0.0.0.0"`
}

type OptionsBool struct {
	PermitWrite bool `hcl:"permit_write" flagName:"permit-write" flagSName:"w" flagDescribe:"Permit clients to write to the TTY (BE CAREFUL)" default:"false"`
}

type OptionsInt struct {
	RandomUrlLength int `hcl:"random_url_length" flagName:"random-url-length" flagDescribe:"Random URL length" default:"8"`
}

// TestGenerateCobraFlags, tests that the GenerateCobraFlags function
// add flags to a cobra command that are equivalent to codegangsta/cli
func TestGenerateCobraStringFlags(t *testing.T) {
	testOptions := &OptionsString{}
	if err := ApplyDefaultValues(testOptions); err != nil {
		t.Errorf("failed to apply default values: %v", err)
		return
	}
	root := cmd.NewRootCmd("")
	got, _, err := GenerateCobraFlags(root, testOptions)
	if err != nil {
		t.Errorf("failed to generate cobra flags: %v", err)
		return
	}
	name := "address"
	shorthand := "a"
	defaultValue := "0.0.0.0"
	flag := got.Flag(name)
	if flag == nil {
		t.Errorf("Wanted a flag named '%v', but didn't get it.", name)
		//we're done, can't test further
		return
	}
	if flag.Name != name {
		t.Errorf("Wanted a flag named '%v', but got '%v'", name, flag.Name)
		t.Fail()
	}
	if flag.Shorthand != shorthand {
		t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.Shorthand, shorthand)
		t.Fail()
	}
	if flag.DefValue != defaultValue {
		t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.DefValue, defaultValue)
		t.Fail()
	}
	if t.Failed() {
		return
	}
	t.Logf("Flag '%v', Shorthand '%v', Default Value '%v'", flag.Name, flag.Shorthand, flag.DefValue)
}

func TestGenerateCobraBooleanFlags(t *testing.T) {
	testOptions := &OptionsBool{}
	if err := ApplyDefaultValues(testOptions); err != nil {
		t.Errorf("failed to apply default values: %v", err)
		return
	}
	root := cmd.NewRootCmd("")
	got, _, err := GenerateCobraFlags(root, testOptions)
	if err != nil {
		t.Errorf("failed to generate cobra flags: %v", err)
		return
	}
	name := "permit-write"
	shorthand := "w"
	defaultValue := "false"
	flag := got.Flag(name)
	if flag == nil {
		t.Errorf("Wanted a flag named '%v', but didn't get it.", name)
		//we're done, can't test further
		return
	}
	if flag.Name != name {
		t.Errorf("Wanted a flag named '%v', but got '%v'", name, flag.Name)
		t.Fail()
	}
	if flag.Shorthand != shorthand {
		t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.Shorthand, shorthand)
		t.Fail()
	}
	if flag.DefValue != defaultValue {
		t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.DefValue, defaultValue)
		t.Fail()
	}
	if t.Failed() {
		return
	}
	t.Logf("Flag '%v', Shorthand '%v', Default Value '%v'", flag.Name, flag.Shorthand, flag.DefValue)
}

func TestGenerateCobraIntFlags(t *testing.T) {
	testOptions := &OptionsInt{}
	if err := ApplyDefaultValues(testOptions); err != nil {
		t.Errorf("failed to apply default values: %v", err)
		return
	}
	root := cmd.NewRootCmd("")
	got, _, err := GenerateCobraFlags(root, testOptions)
	if err != nil {
		t.Errorf("failed to generate cobra flags: %v", err)
		return
	}
	name := "random-url-length"
	shorthand := ""
	defaultValue := "8"
	flag := got.Flag(name)
	if flag == nil {
		t.Errorf("Wanted a flag named '%v', but didn't get it.", name)
		//we're done, can't test further
		return
	}
	if flag.Name != name {
		t.Errorf("Wanted a flag named '%v', but got '%v'", name, flag.Name)
		t.Fail()
	}
	if flag.Shorthand != shorthand {
		t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.Shorthand, shorthand)
		t.Fail()
	}
	if flag.DefValue != defaultValue {
		t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.DefValue, defaultValue)
		t.Fail()
	}
	if t.Failed() {
		return
	}
	t.Logf("Flag '%v', Shorthand '%v', Default Value '%v'", flag.Name, flag.Shorthand, flag.DefValue)
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
				Description:  "IP address to listen",
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
				t.Logf("Flag '%v', Shorthand '%v', Default Value '%v', Description '%v'", name, gotFlag.Shorthand, gotFlag.DefValue, gotFlag.Usage)
			}

			//if flag.Name != test.FlagName {
			//	t.Errorf("Wanted a flag named '%v', but got '%v'", test.FlagName, flag.Name)
			//	t.Fail()
			//}
			//if flag.Shorthand != test.Shorthand {
			//	t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.Shorthand, test.Shorthand)
			//	t.Fail()
			//}
			//if flag.DefValue != defaultValue {
			//	t.Errorf("Wanted flag '%v' to have a shorthand of '%v', but got '%v'", flag.Name, flag.DefValue, defaultValue)
			//	t.Fail()
			//}
			//if t.Failed() {
			//	return
			//}

		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

// TestGenerateCobraFlags, tests that the GenerateCobraFlags function
// add flags to a cobra command that are equivalent to codegangsta/cli
func TestGenerateCobraFlagsOld(t *testing.T) {

	tests := []Test{}

	tests = funcName(t, tests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cmd == nil {
				t.Error("Invalid test command - it's nil")
				return
			}
			got, _, err := GenerateCobraFlags(tt.cmd)
			if got == nil {
				t.Error("Nil command returned")
				return
			}
			if err != nil {
				t.Errorf("Error %v", err)
				return
			}
			if !got.HasFlags() {
				t.Error("No flags found")
			}
		})
	}
}

func funcName(t *testing.T, tests []Test) []Test {
	originalCliFlags, flagMappings, err := originalCliFlags()
	if err != nil {
		t.Errorf("Failed to generate original flags: %v", err)
		t.Fail()
	}

	t.Logf("cliFlags: %# v", pretty.Formatter(originalCliFlags))
	t.Logf("flagMappings: %# v", pretty.Formatter(flagMappings))

	for _, oflg_ := range originalCliFlags {
		switch oflg := oflg_.(type) {
		case cli.StringFlag:
			t.Logf("cli.StringFlag: Flagname %v %v", oflg.GetName(), pretty.Formatter(oflg))
			tests = append(tests, Test{name: "string:" + oflg.GetName(), cmd: cmd.NewRootCmd("")})
		case cli.BoolFlag:
			t.Logf("cli.BoolFlag:   Flagname %v %v", oflg.GetName(), pretty.Formatter(oflg))
			tests = append(tests, Test{name: "bool:" + oflg.GetName(), cmd: cmd.NewRootCmd("")})
		case cli.IntFlag:
			t.Logf("cli.IntFlag:    Flagname %v %v", oflg.GetName(), pretty.Formatter(oflg))
			tests = append(tests, Test{name: "int:" + oflg.GetName(), cmd: cmd.NewRootCmd("")})
		default:
			t.Errorf("Unknown Flag Type: Flagname %v %v", oflg.GetName(), pretty.Formatter(oflg))
		}
	}
	return tests
}

func originalCliFlags() (flags []cli.Flag, mappings map[string]string, err error) {
	//original cliFlag generation code:
	appOptions := &server.Options{}
	if err := ApplyDefaultValues(appOptions); err != nil {
		return nil, nil, err
	}
	backendOptions := &localcommand.Options{}
	if err := ApplyDefaultValues(backendOptions); err != nil {
		return nil, nil, err
	}

	cliFlags, flagMappings, err := GenerateFlags(appOptions, backendOptions)
	if err != nil {
		return nil, nil, err
	}

	return cliFlags, flagMappings, nil
}
