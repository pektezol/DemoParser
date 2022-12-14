package packets

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/packets/classes"
	"github.com/pektezol/demoparser/packets/messages"
)

const MSSC = 2

func ParsePacket(reader *bitreader.ReaderType) (status int) {
	messageType := reader.TryReadInt8()
	messageTick := reader.TryReadInt32()
	messageSlot := reader.TryReadInt8()
	_ = messageSlot
	switch messageType {
	case 1:
		signOn := SignOn{
			PacketInfo:  classes.ParseCmdInfo(reader, MSSC),
			InSequence:  int32(reader.TryReadInt32()),
			OutSequence: int32(reader.TryReadInt32()),
		}
		size := int(reader.TryReadInt32())
		signOn.Data = messages.ParseMessage(reader.TryReadBytesToSlice(size))
		// fmt.Printf("[%d] (%d) {%d} SignOn: %v\n", messageTick, messageType, messageSlot, signOn)
		return 1
	case 2:
		packet := Packet{
			PacketInfo:  classes.ParseCmdInfo(reader, MSSC),
			InSequence:  int32(reader.TryReadInt32()),
			OutSequence: int32(reader.TryReadInt32()),
		}
		size := int(reader.TryReadInt32())
		packet.Data = messages.ParseMessage(reader.TryReadBytesToSlice(size))
		// fmt.Printf("[%d] (%d) Packet: %v\n", messageTick, messageType, packet)
		return 2
	case 3:
		syncTick := SyncTick{}
		fmt.Printf("[%d] (%d) SyncTick: %v\n", messageTick, messageType, syncTick)
		return 3
	case 4:
		size := int(reader.TryReadInt32())
		var consoleCmd ConsoleCmd
		consoleCmd.Data = reader.TryReadStringLen(size)
		// fmt.Printf("[%d] (%d) ConsoleCmd: %s\n", messageTick, messageType, consoleCmd.Data)
		return 4
	case 5: // TODO: UserCmd - Buttons
		userCmd := UserCmd{
			Cmd: int32(reader.TryReadInt32()),
		}
		size := int(reader.TryReadInt32())
		userCmd.Data = classes.ParseUserCmdInfo(reader.TryReadBytesToSlice(size))
		// fmt.Printf("[%d] (%d) UserCmd: %v\n", messageTick, messageType, userCmd)
		return 5
	case 6: // TODO: DataTables
		// datatables := DataTables{
		//  	Size: int32(reader.TryReadInt32()),
		// }
		size := int(reader.TryReadInt32())
		reader.SkipBytes(size)
		// datatables.Data = classes.ParseDataTable(reader.TryReadBytesToSlice(int(datatables.Size)))
		// fmt.Printf("[%d] (%d) DataTables: %v\n", messageTick, messageType, datatables)
		return 6
	case 7:
		stop := Stop{
			RemainingData: nil,
		}
		fmt.Printf("[%d] (%d) Stop: %v\n", messageTick, messageType, stop)
		return 7
	case 8: // TODO: CustomData
		reader.SkipBytes(4)
		size := int(reader.TryReadInt32())
		reader.SkipBytes(size)
		// fmt.Printf("[%d] (%d) CustomData: \n", messageTick, messageType)
		return 8
	case 9: // TODO: StringTables - Data
		var stringTables StringTables
		size := int(reader.TryReadInt32())
		stringTables.Data = classes.ParseStringTable(reader.TryReadBytesToSlice(size))
		// fmt.Printf("[%d] (%d) StringTables: %v\n", messageTick, messageType, stringTables)
		return 9
	default:
		return 0
	}
}
