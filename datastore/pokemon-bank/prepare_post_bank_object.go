// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePreparePostBankObject(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.PreparePostBankObject == nil {
		globals.Logger.Warning("DataStorePokemonBank::PreparePostBankObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	slotID := types.NewPrimitiveU16(0)
	err = slotID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PreparePostBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	size := types.NewPrimitiveU32(0)
	err = size.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PreparePostBankObject(fmt.Errorf("Failed to read size from parameters. %s", err.Error()), packet, callID, nil, nil)
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
