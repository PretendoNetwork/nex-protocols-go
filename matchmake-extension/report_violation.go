// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportViolation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ReportViolation == nil {
		globals.Logger.Warning("MatchmakeExtension::ReportViolation not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadPID()
	if err != nil {
		_, errorCode = protocol.ReportViolation(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	userName, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.ReportViolation(fmt.Errorf("Failed to read userName from parameters. %s", err.Error()), packet, callID, nil, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	violationCode, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.ReportViolation(fmt.Errorf("Failed to read violationCode from parameters. %s", err.Error()), packet, callID, nil, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ReportViolation(nil, packet, callID, pid, userName, violationCode)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
