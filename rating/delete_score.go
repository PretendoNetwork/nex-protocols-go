// Package protocol implements the Rating protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeleteScore(packet nex.PacketInterface) {
	if protocol.DeleteScore == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Rating::DeleteScore not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var category types.UInt32
	var uniqueID types.UInt64

	var err error

	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteScore(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, category, uniqueID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteScore(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, category, uniqueID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteScore(nil, packet, callID, category, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
