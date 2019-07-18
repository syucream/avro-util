package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/linkedin/goavro"
)

var (
	toJsonHandlerMap = map[formatType]func(string, string) (string, error){
		formatTypeRaw: toJsonRaw,
		// TODO soe
		formatTypeOcf: toJsonOcf,
	}
)

func toJsonRaw(pathToSchema string, pathToFile string) (string, error) {
	codec, err := generateCodec(pathToSchema)
	if err != nil {
		return "", err
	}

	fileData, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return "", err
	}

	datum, _, err := codec.NativeFromBinary(fileData)
	if err != nil {
		return "", err
	}

	jsonBin, err := codec.TextualFromNative(nil, datum)
	if err != nil {
		return "", err
	}

	return string(jsonBin), nil
}

func toJsonOcf(pathToSchema string, pathToFile string) (string, error) {
	fileData, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return "", err
	}

	rbuf := bytes.NewBuffer(fileData)
	ocfr, err := goavro.NewOCFReader(rbuf)

	// TODO it can catch some cases?
	ready := ocfr.Scan()
	if !ready {
		return "", ocfr.Err()
	}

	datum, err := ocfr.Read()
	if err != nil {
		return "", err
	}

	jsonBin, err := ocfr.Codec().TextualFromNative(nil, datum)
	if err != nil {
		return "", err
	}

	return string(jsonBin), nil
}

func ToJson(flags Flags) error {
	ft, err := getFormatType(flags.Format)
	if err != nil {
		return err
	}

	handler, ok := toJsonHandlerMap[ft]
	if !ok {
		return ErrUnknownFormatType
	}

	out, err := handler(flags.PathToSchema, flags.PathToFile)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}
