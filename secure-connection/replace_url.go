// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReplaceURL(packet nex.PacketInterface) {
	var err error

	if protocol.ReplaceURL == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SecureConnection::ReplaceURL not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	target := types.NewStationURL("")
	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReplaceURL(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	url := types.NewStationURL("")
	err = url.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReplaceURL(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReplaceURL(nil, packet, callID, target, url)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
