// Package protocol implements the legacy Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	if protocol.GetCommonData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetCommonData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	endpoint := packet.Sender().Endpoint()

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var uniqueID types.UInt32

	var err error

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCommonData(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, uniqueID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCommonData(nil, packet, callID, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
