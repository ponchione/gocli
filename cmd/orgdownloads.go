package cmd

import "gocli/core"

type OrganizeCommand struct {
	Name        string
	Description string
}

func init() {
	core.Commands["org"] = &OrganizeCommand{
		Name:        "org",
		Description: "org will organize the contents of your Downloads folder.",
	}
}

func (o *OrganizeCommand) Execute(args []string) error {
	return nil
}

func (o *OrganizeCommand) Help() string {
	return o.Description
}
