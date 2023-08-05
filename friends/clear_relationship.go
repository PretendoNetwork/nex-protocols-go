// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearRelationship sets the ClearRelationship handler function
func (protocol *Protocol) ClearRelationship(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32) uint32) {
	protocol.clearRelationshipHandler = handler
}

func (protocol *Protocol) handleClearRelationship(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.clearRelationshipHandler == nil {
		globals.Logger.Warning("Friends::ClearRelationship not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.clearRelationshipHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.clearRelationshipHandler(nil, client, callID, uiPlayer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
