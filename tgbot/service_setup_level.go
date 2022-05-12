package tgbot

type SetupLevel uint8

const (
	Empty SetupLevel = 1 << iota
	Instance
	Configuration
	Init
	Workers
	Repository
	Handlers
)

func Set(b, flag SetupLevel) SetupLevel    { return b | flag }
func Clear(b, flag SetupLevel) SetupLevel  { return b &^ flag }
func Toggle(b, flag SetupLevel) SetupLevel { return b ^ flag }
func Has(b, flag SetupLevel) bool          { return b&flag != 0 }
