// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleResetRatings(packet nex.PacketInterface) {
	if protocol.ResetRatings == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::ResetRatings not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var transactional types.Bool
	var dataIDs types.List[types.UInt64]

	var err error

	err = dataIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPasswordInfos(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, dataIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = transactional.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ResetRatings(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, dataIDs, transactional)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ResetRatings(nil, packet, callID, dataIDs, transactional)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
