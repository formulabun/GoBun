package subcommand

import (
	"GoBun/functional/array"
	"GoBun/functional/strings"
	"fmt"
)

type Registry struct {
	options []Subcommand
}

type Subcommand struct {
	Name    string
	Options []string
	Run     func([]string) (help fmt.Stringer, err error)
}

func (r *Registry) Register(command Subcommand) *Registry {
	r.options = append(r.options, command)
	return r
}

func (r Registry) String() string {
	result := "{"
	result += strings.Join(array.Map(r.options, func(s Subcommand) string { return s.Name }), ",")
	return result + "}"
}

func (r *Registry) Run(command string, options []string) (help fmt.Stringer, err error) {
	var find = func(comm Subcommand) bool {
		for _, c := range comm.Options {
			if c == command {
				return true
			}
		}
		return false
	}

	var option = array.FindFirst(r.options, find)
	if option == nil {
		return nil, fmt.Errorf("cannot find subcommand %s", command)
	}

	help, err = option.Run(options)
	if err != nil {
		return nil, err
	}

	if help != nil {
		return strings.Stringer{fmt.Sprintf("%s %s", command, help)}, nil
	}

	return nil, nil
}
