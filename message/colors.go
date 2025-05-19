package message

import aternos "github.com/sleeyax/aternos-api"

var colorMap = map[aternos.ServerStatus]int{
	aternos.Online:    0x57F287, // green
	aternos.Offline:   0xED4245, // red
	aternos.Starting:  0xFEE75C, // yellow
	aternos.Stopping:  0xFEE75C, // yellow
	aternos.Loading:   0x5865F2, // blurple
	aternos.Preparing: 0x5865F2, // blurple
	aternos.Saving:    0x979C9F, // gray
}