// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetList == nil {
		globals.Logger.Warning("Friends::GetList not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	byRelationship, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.GetList(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReversed, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.GetList(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetList(nil, packet, callID, byRelationship, bReversed)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
