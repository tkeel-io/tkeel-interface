package main

import (
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/version"
	"log"

	"github.com/spf13/cobra"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/project"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/proto"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/render"
<<<<<<< feat/add-markdown-link
=======
	"github.com/tkeel-io/tkeel-interface/tool/pkg/version"
>>>>>>> main
)

var rootCmd = &cobra.Command{
	Use:     "tkeel-tool",
	Short:   "An elegant toolkit for Go microservices.",
	Long:    `An elegant toolkit for Go microservices.`,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(render.CmdMarkdown)
<<<<<<< feat/add-markdown-link
	rootCmd.AddCommand(version.VersionCmd)
=======
>>>>>>> main
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
