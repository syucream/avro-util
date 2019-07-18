package commands

import "fmt"

type formatType int

const (
	formatTypeUnknown formatType = iota
	formatTypeRaw
	formatTypeSoe
	formatTypeOcf
)

var (
	formatTypeMap = map[string]formatType{
		"raw": formatTypeRaw,
		"soe": formatTypeSoe,
		"ocf": formatTypeOcf,
	}
	commandMap = map[string]func(Flags) error{
		"random": Random,
		"tojson": ToJson,
	}
)

var (
	ErrUnknownCommand    = fmt.Errorf("unknown command")
	ErrUnknownFormatType = fmt.Errorf("unknown format type")
)
