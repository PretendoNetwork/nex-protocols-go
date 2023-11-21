// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReplaceURL(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ReplaceURL == nil {
		globals.Logger.Warning("SecureConnection::ReplaceURL not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStationURL()
	if err != nil {
		_, errorCode = protocol.ReplaceURL(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	url, err := parametersStream.ReadStationURL()
	if err != nil {
		_, errorCode = protocol.ReplaceURL(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ReplaceURL(nil, packet, callID, target, url)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
