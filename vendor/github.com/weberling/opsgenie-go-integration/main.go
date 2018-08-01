package main

import (
	"os"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	commands "github.com/weberling/opsgenie-go-integration/commands"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("Command required")
	}

	command, err := commands.Factory(args)
	if err != nil {
		panic(err)
	}
	command.Call(getOpsGenieCli())
}

func getOpsGenieCli() *ogcli.OpsGenieClient {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(os.Getenv("OPSGENIE_API_KEY"))
	return cli
}
