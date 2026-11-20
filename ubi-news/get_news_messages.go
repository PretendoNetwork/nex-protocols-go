// Package protocol implements the UbiNews protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetNewsMessages(packet nex.PacketInterface) {
	if protocol.GetNewsMessages == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiNews::GetNewsMessages not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var resultRange types.ResultRange

	err := resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetNewsMessages(fmt.Errorf("failed to read resultRange from parameters. %s", err.Error()), packet, callID, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetNewsMessages(nil, packet, callID, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
