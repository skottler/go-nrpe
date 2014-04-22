package main

import (
        "github.com/mitchellh/cli"
        "fmt"
        "io/ioutil"
        "log"
        "os"
)

func main() {
        os.Exit(realMain())
}

func realMain() int {
        log.SetOutput(ioutil.Discard)

        args := os.Args[1:]

        for _, arg := range args {
                if arg == "-v" || arg == "--version" {
                        newArgs := make([]string, len(args)+1)
                        newArgs[0] = "version"
                        copy(newArgs[1:], args)
                        args = newArgs
                        break
                }
        }

        cli := &cli.CLI{
                Args:     args,
                HelpFunc: cli.BasicHelpFunc("nrpe"),
       }

       exitCode, err := cli.Run()

       if err != nil {
               fmt.Fprintf(os.Stderr, "Error during execution: %s\n", err.Error())
       }

        return exitCode
}
