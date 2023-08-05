// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPresence sets the GetFriendPresence handler function
func (protocol *Protocol) GetFriendPresence(handler func(err error, client *nex.Client, callID uint32, pidList []uint32) uint32) {
	protocol.getFriendPresenceHandler = handler
}

func (protocol *Protocol) handleGetFriendPresence(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFriendPresenceHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPresence not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pidList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getFriendPresenceHandler(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getFriendPresenceHandler(nil, client, callID, pidList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
