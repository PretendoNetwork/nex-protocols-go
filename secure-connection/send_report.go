// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSendReport(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SendReport == nil {
		globals.Logger.Warning("SecureConnection::SendReport not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	reportID := types.NewPrimitiveU32(0)
	err = reportID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SendReport(fmt.Errorf("Failed to read reportID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	reportData := types.NewQBuffer(nil)
	err = reportData.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SendReport(fmt.Errorf("Failed to read reportData from parameters. %s", err.Error()), packet, callID, nil, nil)
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
