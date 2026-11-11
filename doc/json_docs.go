package doc

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type jsonFlag struct {
	Name      string `json:"name"`
	Shorthand string `json:"shorthand,omitempty"`
	Usage     string `json:"usage"`
	Default   string `json:"default"`
}

type jsonCommand struct {
	Name     string        `json:"name"`
	Path     string        `json:"path"`
	Short    string        `json:"short"`
	Long     string        `json:"long,omitempty"`
	Example  string        `json:"example,omitempty"`
	Flags    []jsonFlag    `json:"flags,omitempty"`
	Children []jsonCommand `json:"children,omitempty"`
}

func buildJSON(cmd *cobra.Command) jsonCommand {
	cmd.SetOutput(io.Discard)
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	doc := jsonCommand{
		Name:    cmd.Name(),
		Path:    cmd.CommandPath(),
		Short:   cmd.Short,
		Long:    cmd.Long,
		Example: cmd.Example,
	}
	cmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		doc.Flags = append(doc.Flags, jsonFlag{Name: f.Name, Shorthand: f.Shorthand, Usage: f.Usage, Default: f.DefValue})
	})
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		doc.Children = append(doc.Children, buildJSON(c))
	}
	return doc
}

// GenJSON creates the JSON documentation of cmd and its subcommands.
func GenJSON(cmd *cobra.Command, w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(buildJSON(cmd))
}

// GenJSONTree writes one JSON file per command into dir.
func GenJSONTree(cmd *cobra.Command, dir string) error {
	os.MkdirAll(dir, 0777)
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := GenJSONTree(c, dir); err != nil {
			return err
		}
	}

	f, err := os.Create(filepath.Join(dir, commandSlug(cmd)+".json"))
	if err != nil {
		return err
	}
	return GenJSON(cmd, f)
}
