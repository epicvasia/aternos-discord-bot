package message

import aternos "github.com/sleeyax/aternos-api"

const (
	Red    = 0xED4245
	Green  = 0x57F287
	Blue   = 0x3498db
	Yellow = 0xFEE75C
	Orange = 0xF39C12
	Aqua   = 0x1abc9c
	Gray   = 0x979C9F
	Blurple = 0x5865F2
)

var colorMap = map[aternos.ServerStatus]int{
	aternos.Online:    Green,
	aternos.Offline:   Red,
	aternos.Starting:  Yellow,
	aternos.Stopping:  Yellow,
	aternos.Loading:   Blurple,
	aternos.Preparing: Blurple,
	aternos.Saving:    Gray,
}