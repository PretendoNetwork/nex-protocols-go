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

	if protocol.SendReport == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SecureConnection::SendReport not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	reportID := types.NewPrimitiveU32(0)
	err = reportID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SendReport(fmt.Errorf("Failed to read reportID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	reportData := types.NewQBuffer(nil)
	err = reportData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SendReport(fmt.Errorf("Failed to read reportData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendReport(nil, packet, callID, reportID, reportData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
