// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRegisterEx(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RegisterEx == nil {
		globals.Logger.Warning("SecureConnection::RegisterEx not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	vecMyURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		_, errorCode = protocol.RegisterEx(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	hCustomData, err := parametersStream.ReadDataHolder()
	if err != nil {
		_, errorCode = protocol.RegisterEx(fmt.Errorf("Failed to read hCustomData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RegisterEx(nil, packet, callID, vecMyURLs, hCustomData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
