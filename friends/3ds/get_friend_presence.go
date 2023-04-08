package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPresence sets the GetFriendPresence handler function
func (protocol *Friends3DSProtocol) GetFriendPresence(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	protocol.GetFriendPresenceHandler = handler
}

func (protocol *Friends3DSProtocol) HandleGetFriendPresence(packet nex.PacketInterface) {
	if protocol.GetFriendPresenceHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPresence not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	PidList := parametersStream.ReadListUInt32LE()

	go protocol.GetFriendPresenceHandler(nil, client, callID, PidList)
}
