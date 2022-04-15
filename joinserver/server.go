package main

import (
	"context"
	"strings"

	cli "github.com/laurentiuNiculae/go-rpc-plugins/plugins/clicommand"
)

type CLIServer struct {
	cli.UnimplementedCLICommandServer
}

func (s CLIServer) Command(ctx context.Context, args *cli.CLIArgs) (*cli.CLIResponse, error) {
	var resp *cli.CLIResponse = &cli.CLIResponse{}

	resp.Message = "I am combining the args: " + strings.Join(args.GetArgs(), "|")

	return resp, nil
}
