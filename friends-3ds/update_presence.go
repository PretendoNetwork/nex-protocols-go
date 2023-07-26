// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePresence sets the UpdatePresence handler function
func (protocol *Protocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame bool)) {
	protocol.updatePresenceHandler = handler
}

func (protocol *Protocol) handleUpdatePresence(packet nex.PacketInterface) {
	if protocol.updatePresenceHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePresence not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nintendoPresence, err := parametersStream.ReadStructure(friends_3ds_types.NewNintendoPresence())
	if err != nil {
		go protocol.updatePresenceHandler(fmt.Errorf("Failed to read nintendoPresence from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	showGame, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.updatePresenceHandler(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	go protocol.updatePresenceHandler(nil, client, callID, nintendoPresence.(*friends_3ds_types.NintendoPresence), showGame)
}
