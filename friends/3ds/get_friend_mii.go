package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendMii sets the GetFriendMii handler function
func (protocol *Friends3DSProtocol) GetFriendMii(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	protocol.GetFriendMiiHandler = handler
}

func (protocol *Friends3DSProtocol) HandleGetFriendMii(packet nex.PacketInterface) {
	if protocol.GetFriendMiiHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendMiiHandler not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	PidList := parametersStream.ReadListUInt32LE()

	go protocol.GetFriendMiiHandler(nil, client, callID, PidList)
}
