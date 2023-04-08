package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByPrincipalID sets the AddFriendByPrincipalID handler function
func (protocol *Friends3DSProtocol) AddFriendByPrincipalID(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pid uint32)) {
	protocol.AddFriendByPrincipalIDHandler = handler
}

func (protocol *Friends3DSProtocol) HandleAddFriendByPrincipalID(packet nex.PacketInterface) {
	if protocol.AddFriendByPrincipalIDHandler == nil {
		globals.Logger.Warning("Friends3DS::AddFriendByPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc := parametersStream.ReadUInt64LE()
	pid := parametersStream.ReadUInt32LE()

	go protocol.AddFriendByPrincipalIDHandler(nil, client, callID, lfc, pid)
}
