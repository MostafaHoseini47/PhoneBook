package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

type Server struct{}

func (cmd Server) command(trap chan os.Signal) *cobra.Command {
}

func (cmd *Server) run(conf config.Config, trap chan os.Signal) {
}
