// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePrepareGetBankObject(packet nex.PacketInterface) {
	if protocol.PrepareGetBankObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::PrepareGetBankObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	slotID := types.NewPrimitiveU16(0)
	applicationID := types.NewPrimitiveU16(0)

	var err error

	err = slotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PrepareGetBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = applicationID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PrepareGetBankObject(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PrepareGetBankObject(nil, packet, callID, slotID, applicationID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
