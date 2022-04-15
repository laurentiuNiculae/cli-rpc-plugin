package main

import (
	"context"
	"fmt"

	cli "github.com/laurentiuNiculae/go-rpc-plugins/plugins/clicommand"
)

type CLIServer struct {
	cli.UnimplementedCLICommandServer
}

func (s CLIServer) Command(ctx context.Context, cliArgs *cli.CLIArgs) (*cli.CLIResponse, error) {
	var resp *cli.CLIResponse = &cli.CLIResponse{}

	args := cliArgs.GetArgs()

	if len(args) == 0 {
		resp.Message = "Hello! You didn't send any arguments. Try some!"
	} else {
		resp.Message = fmt.Sprintf("Hello! You sent %d arguments!", len(args))
	}

	return resp, nil
}
