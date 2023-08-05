// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePresence sets the UpdatePresence handler function
func (protocol *Protocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame bool) uint32) {
	protocol.updatePresenceHandler = handler
}

func (protocol *Protocol) handleUpdatePresence(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updatePresenceHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePresence not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nintendoPresence, err := parametersStream.ReadStructure(friends_3ds_types.NewNintendoPresence())
	if err != nil {
		errorCode = protocol.updatePresenceHandler(fmt.Errorf("Failed to read nintendoPresence from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	showGame, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.updatePresenceHandler(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), client, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updatePresenceHandler(nil, client, callID, nintendoPresence.(*friends_3ds_types.NintendoPresence), showGame)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
