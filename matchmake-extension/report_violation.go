// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportViolation sets the ReportViolation handler function
func (protocol *Protocol) ReportViolation(handler func(err error, client *nex.Client, callID uint32, pid uint32, userName string, violationCode uint32) uint32) {
	protocol.reportViolationHandler = handler
}

func (protocol *Protocol) handleReportViolation(packet nex.PacketInterface) {
	if protocol.reportViolationHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ReportViolation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.reportViolationHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0, "", 0)
		return
	}

	userName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.reportViolationHandler(fmt.Errorf("Failed to read userName from parameters. %s", err.Error()), client, callID, 0, "", 0)
		return
	}

	violationCode, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.reportViolationHandler(fmt.Errorf("Failed to read violationCode from parameters. %s", err.Error()), client, callID, 0, "", 0)
		return
	}

	go protocol.reportViolationHandler(nil, client, callID, pid, userName, violationCode)
}
