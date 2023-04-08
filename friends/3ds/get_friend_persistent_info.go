package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPersistentInfo sets the GetFriendPersistentInfo handler function
func (protocol *Friends3DSProtocol) GetFriendPersistentInfo(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	protocol.GetFriendPersistentInfoHandler = handler
}

func (protocol *Friends3DSProtocol) HandleGetFriendPersistentInfo(packet nex.PacketInterface) {
	if protocol.GetFriendPersistentInfoHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPersistentInfo not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	PidList := parametersStream.ReadListUInt32LE()

	go protocol.GetFriendPersistentInfoHandler(nil, client, callID, PidList)
}
