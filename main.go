package main

import (
	"flag"
	"log"
	"os"

	"github.com/syucream/avro-util/pkg/commands"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("specify subcommand")
	}
	subCmd := os.Args[1]

	subFlag := flag.NewFlagSet(subCmd, flag.ExitOnError)
	schema := subFlag.String("schema", "", "Specify path/to/schema.json")
	format := subFlag.String("format", "ocf", "Specify format(raw, soe, ocf)")
	file := subFlag.String("file", "", "Specify path/to/avro binary file")
	err := subFlag.Parse(os.Args[2:])

	flags := commands.Flags{
		PathToSchema: *schema,
		Format:       *format,
		PathToFile:   *file,
	}

	err = commands.ExecuteCommand(subCmd, flags)
	if err != nil {
		log.Fatal("error occured during executing the command: ", err)
	}
}
