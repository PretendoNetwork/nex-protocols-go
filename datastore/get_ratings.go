// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetRatings(packet nex.PacketInterface) {
	if protocol.GetRatings == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::GetRatings not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var dataIDs types.List[types.UInt64]
	var accessPassword types.UInt64

	var err error

	err = dataIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRatings(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, dataIDs, accessPassword)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = accessPassword.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRatings(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), packet, callID, dataIDs, accessPassword)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRatings(nil, packet, callID, dataIDs, accessPassword)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
