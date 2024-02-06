// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportViolation(packet nex.PacketInterface) {
	var err error

	if protocol.ReportViolation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::ReportViolation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pid := types.NewPID(0)
	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportViolation(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	userName := types.NewString("")
	err = userName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportViolation(fmt.Errorf("Failed to read userName from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	violationCode := types.NewPrimitiveU32(0)
	err = violationCode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportViolation(fmt.Errorf("Failed to read violationCode from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportViolation(nil, packet, callID, pid, userName, violationCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
