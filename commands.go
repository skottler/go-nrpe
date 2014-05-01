package main

import (
        "github.com/mitchellh/cli"
        "github.com/skottler/go-nrpe/command"
        "os"
)

var Commands map[string]cli.CommandFactory

func init() {
        ui := &cli.BasicUi{Writer: os.Stdout}

        Commands = map[string]cli.CommandFactory {
                "version": func() (cli.Command, error) {
                        return &command.VersionCommand{
                                Ui:                ui,
                        }, nil
                },
        }
}
