// Package protocol implements the IPToLocationAdmin Protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetLocaleCode(packet nex.PacketInterface) {
	if protocol.SetLocaleCode == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Localization::SetLocaleCode not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var localCode types.String

	err := localCode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetLocaleCode(fmt.Errorf("failed to read localCode from parameters. %s", err.Error()), packet, callID, localCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetLocaleCode(nil, packet, callID, localCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
