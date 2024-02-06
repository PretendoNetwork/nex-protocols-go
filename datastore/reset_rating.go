// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleResetRating(packet nex.PacketInterface) {
	var err error

	if protocol.ResetRating == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::ResetRating not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	target := datastore_types.NewDataStoreRatingTarget()
	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ResetRating(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	accessPassword := types.NewPrimitiveU64(0)
	err = accessPassword.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ResetRating(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ResetRating(nil, packet, callID, target, accessPassword)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
