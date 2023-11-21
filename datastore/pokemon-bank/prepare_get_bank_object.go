// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePrepareGetBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PrepareGetBankObject == nil {
		globals.Logger.Warning("DataStorePokemonBank::PrepareGetBankObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.PrepareGetBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	applicationID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.PrepareGetBankObject(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PrepareGetBankObject(nil, packet, callID, slotID, applicationID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
