// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearRelationship sets the ClearRelationship handler function
func (protocol *FriendsProtocol) ClearRelationship(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32)) {
	protocol.clearRelationshipHandler = handler
}

func (protocol *FriendsProtocol) handleClearRelationship(packet nex.PacketInterface) {
	if protocol.clearRelationshipHandler == nil {
		globals.Logger.Warning("Friends::ClearRelationship not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.clearRelationshipHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.clearRelationshipHandler(nil, client, callID, uiPlayer)
}
