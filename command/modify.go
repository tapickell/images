package command

import (
	"fmt"
	"os"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

type Modify struct {
	provider string
}

func NewModify(config *Config) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Modify{
			provider: config.Provider,
		}, nil
	}
}

func (m *Modify) Help() string {
	if m.provider == "" {
		defaultHelp := `Usage: images modify [options]

  Modifies images properties. Each providers sub options are different.

Options:

  -provider                  Provider to be used to modify images
`
		return defaultHelp
	}

	return Help("modify", m.provider)
}

func (m *Modify) Run(args []string) int {
	if m.provider == "" {
		fmt.Print(m.Help())
		return 1
	}

	if flags.Has("help", args) {
		fmt.Print(m.Help())
		return 1
	}

	p, err := Provider(m.provider, args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	mr, ok := p.(Modifier)
	if !ok {
		err := fmt.Errorf("'%s' doesn't support listing images", m.provider)
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	if err := mr.Modify(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	return 0
}

func (m *Modify) Synopsis() string {
	return "Modify image properties"
}
