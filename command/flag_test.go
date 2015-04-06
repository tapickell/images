package command

import "testing"

func TestIsFlag(t *testing.T) {
	var flags = []struct {
		flag   string
		isFlag bool
	}{
		{flag: "--foo", isFlag: true},
		{flag: "--foo=bar", isFlag: true},
		{flag: "-foo", isFlag: true},
		{flag: "-foo=bar", isFlag: true},
		{flag: "-f", isFlag: true},
		{flag: "-f=bar", isFlag: true},
		{flag: "f=bar", isFlag: false},
		{flag: "f=", isFlag: false},
		{flag: "f", isFlag: false},
		{flag: "", isFlag: false},
	}

	for _, f := range flags {
		is := isFlag(f.flag)
		if is != f.isFlag {
			t.Errorf("flag: %s\n\twant: %s\n\tgot : %s\n", f.flag, f.isFlag, is)
		}
	}
}

func TestParseFlag(t *testing.T) {
	var flags = []struct {
		flag string
		name string
	}{
		{name: "foo", flag: "--foo"},
		{name: "foo", flag: "-foo"},
		{name: "foo=bar", flag: "-foo=bar"},
		{name: "foo=", flag: "-foo="},
		{name: "foo=b", flag: "-foo=b"},
		{name: "f", flag: "-f"},
		{name: "f", flag: "--f"},
		{name: "", flag: "---f"},
		{name: "", flag: "f"},
		{name: "", flag: "--"},
		{name: "", flag: "-"},
	}

	for _, f := range flags {
		name, _ := parseFlag(f.flag)
		if name != f.name {
			t.Errorf("flag: %s\n\twant: %s\n\tgot : %s\n", f.flag, f.name, name)
		}
	}

}

func TestParseValue(t *testing.T) {
	var flags = []struct {
		flag  string
		name  string
		value string
	}{
		{flag: "foo=bar", name: "foo", value: "bar"},
		{flag: "foo=b", name: "foo", value: "b"},
		{flag: "f=", name: "f", value: ""},
		{flag: "f", name: "f", value: ""},
		{flag: "", name: "", value: ""},
	}

	for _, f := range flags {
		name, value := parseValue(f.flag)
		if value != f.value {
			t.Errorf("parsing value from flag: %s\n\twant: %s\n\tgot : %s\n",
				f.flag, f.value, value)
		}

		if name != f.name {
			t.Errorf("parsing name from flag: %s\n\twant: %s\n\tgot : %s\n",
				f.flag, f.name, name)
		}
	}

}

func TestParseProvider(t *testing.T) {
	var arguments = []struct {
		args  []string
		value string
	}{
		{args: []string{"--provider=aws"}, value: "aws"},
		{args: []string{"-provider=aws"}, value: "aws"},
		{args: []string{"-provider=aws,do"}, value: "aws,do"},
		{args: []string{"-p=aws", "aws"}, value: "aws"},
		{args: []string{"--provider", "aws"}, value: "aws"},
		{args: []string{"-provider", "aws"}, value: "aws"},
		{args: []string{"-p", "aws"}, value: "aws"},
		{args: []string{"-p"}, value: ""},
		{args: []string{"-p="}, value: ""},
		{args: []string{"-p=", "--foo"}, value: ""},
		{args: []string{"-p", "--foo"}, value: ""},
		{args: []string{"-provider", "--foo"}, value: ""},
		{args: []string{"--provider", "--foo"}, value: ""},
	}

	for _, args := range arguments {
		value, _ := parseFlagValue("provider", args.args)

		if value != args.value {
			t.Errorf("parsing args: %v\n\twant: %s\n\tgot : %s\n",
				args.args, args.value, value)
		}
	}
}