package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/project"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/proto"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/render"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/version"
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
	rootCmd.AddCommand(version.VersionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
