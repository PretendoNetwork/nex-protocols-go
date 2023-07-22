// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPresence sets the GetFriendPresence handler function
func (protocol *Friends3DSProtocol) GetFriendPresence(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	protocol.getFriendPresenceHandler = handler
}

func (protocol *Friends3DSProtocol) handleGetFriendPresence(packet nex.PacketInterface) {
	if protocol.getFriendPresenceHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPresence not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pidList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getFriendPresenceHandler(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getFriendPresenceHandler(nil, client, callID, pidList)
}
