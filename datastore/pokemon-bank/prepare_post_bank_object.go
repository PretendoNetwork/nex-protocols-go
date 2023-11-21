// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePreparePostBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PreparePostBankObject == nil {
		globals.Logger.Warning("DataStorePokemonBank::PreparePostBankObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.PreparePostBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	size, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.PreparePostBankObject(fmt.Errorf("Failed to read size from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PreparePostBankObject(nil, packet, callID, slotID, size)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
