package commands

import (
	"fmt"
)

var (
	randomHandlerMap = map[formatType]func(string) ([]byte, error){
		formatTypeRaw: generateRandomRaw,
	}
)

func generateRandomRaw(pathToSchema string) ([]byte, error) {
	_, err := generateCodec(pathToSchema)
	if err != nil {
		return []byte{}, err
	}

	// TODO randomize

	return []byte{}, nil
}

func Random(flags Flags) error {
	ft, err := getFormatType(flags.Format)
	if err != nil {
		return err
	}

	handler, ftIsOk := randomHandlerMap[ft]
	if !ftIsOk {
		return ErrUnknownFormatType
	}

	out, err := handler(flags.PathToSchema)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}
