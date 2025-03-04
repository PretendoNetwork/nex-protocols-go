// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleFindByID(packet nex.PacketInterface) {
	if protocol.FindByID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::FindByID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var lstID types.List[types.UInt32]

	err := lstID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByID(fmt.Errorf("Failed to read lstID from parameters. %s", err.Error()), packet, callID, lstID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByID(nil, packet, callID, lstID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
