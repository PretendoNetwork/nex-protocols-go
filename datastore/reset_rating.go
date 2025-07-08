// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleResetRating(packet nex.PacketInterface) {
	if protocol.ResetRating == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::ResetRating not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	target := datastore_types.NewDataStoreRatingTarget()
	var updatePassword types.UInt64

	var err error

	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ResetRating(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, target, updatePassword)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = updatePassword.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ResetRating(fmt.Errorf("Failed to read updatePassword from parameters. %s", err.Error()), packet, callID, target, updatePassword)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ResetRating(nil, packet, callID, target, updatePassword)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
