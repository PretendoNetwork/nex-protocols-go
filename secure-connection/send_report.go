// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSendReport(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SendReport == nil {
		globals.Logger.Warning("SecureConnection::SendReport not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	reportID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.SendReport(fmt.Errorf("Failed to read reportID from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	reportData, err := parametersStream.ReadQBuffer()
	if err != nil {
		_, errorCode = protocol.SendReport(fmt.Errorf("Failed to read reportData from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SendReport(nil, packet, callID, reportID, reportData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
