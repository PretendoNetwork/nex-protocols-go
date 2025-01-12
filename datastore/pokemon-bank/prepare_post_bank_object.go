// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handlePreparePostBankObject(packet nex.PacketInterface) {
	if protocol.PreparePostBankObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::PreparePostBankObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var slotID types.UInt16
	var size types.UInt32

	var err error

	err = slotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, slotID, size)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = size.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostBankObject(fmt.Errorf("Failed to read size from parameters. %s", err.Error()), packet, callID, slotID, size)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PreparePostBankObject(nil, packet, callID, slotID, size)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
