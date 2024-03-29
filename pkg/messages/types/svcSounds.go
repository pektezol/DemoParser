package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcSounds struct {
	ReliableSound bool
	SoundCount    uint8
	Length        uint16
	Data          []byte
}

func ParseSvcSounds(reader *bitreader.Reader) SvcSounds {
	svcSounds := SvcSounds{
		ReliableSound: reader.TryReadBool(),
	}
	if svcSounds.ReliableSound {
		svcSounds.SoundCount = 1
		svcSounds.Length = uint16(reader.TryReadUInt8())
	} else {
		svcSounds.SoundCount = reader.TryReadUInt8()
		svcSounds.Length = reader.TryReadUInt16()
	}
	svcSounds.Data = reader.TryReadBitsToSlice(uint64(svcSounds.Length))
	writer.TempAppendLine("\t\tReliable Sound: %t", svcSounds.ReliableSound)
	writer.TempAppendLine("\t\tSound Count: %d", svcSounds.SoundCount)
	writer.TempAppendLine("\t\tData: %v", svcSounds.Data)
	return svcSounds
}
