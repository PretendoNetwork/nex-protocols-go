// Package protocol implements the Screening protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportUser sets the ReportUser handler function
func (protocol *Protocol) ReportUser(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.reportUserHandler = handler
}

func (protocol *Protocol) handleReportUser(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportUserHandler == nil {
		globals.Logger.Warning("Screening::ReportUser not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Screening::ReportUser STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	errorCode = protocol.reportUserHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
