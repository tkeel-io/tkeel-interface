package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/markdown"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/project"
	"github.com/tkeel-io/tkeel-interface/tool/cmd/artisan/proto"
	"github.com/tkeel-io/tkeel-interface/tool/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:     "tkeel-tool",
	Short:   "An elegant toolkit for Go microservices.",
	Long:    `An elegant toolkit for Go microservices.`,
	Version: fmt.Sprintf("%s [%s %s]", version.Version, version.GitCommit, version.BuildDate),
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(markdown.CmdMarkdown)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
