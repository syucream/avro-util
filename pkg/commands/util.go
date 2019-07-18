package commands

import (
	"io/ioutil"

	"github.com/linkedin/goavro"
)

// ExecuteCommand executes command specified given name
func ExecuteCommand(cmdName string, flags Flags) error {
	cmd, ok := commandMap[cmdName]
	if !ok {
		return ErrUnknownCommand
	}

	return cmd(flags)
}

func getFormatType(ft string) (formatType, error) {
	ftval, ok := formatTypeMap[ft]
	if !ok {
		return formatTypeUnknown, ErrUnknownFormatType
	}

	return ftval, nil
}

func generateCodec(pathToSchema string) (*goavro.Codec, error) {
	s, err := ioutil.ReadFile(pathToSchema)
	if err != nil {
		return nil, err
	}

	return goavro.NewCodec(string(s))
}
