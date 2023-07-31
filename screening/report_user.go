// Package protocol implements the Screening protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportUser sets the ReportUser handler function
func (protocol *Protocol) ReportUser(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.reportUserHandler = handler
}

func (protocol *Protocol) handleReportUser(packet nex.PacketInterface) {
	if protocol.reportUserHandler == nil {
		globals.Logger.Warning("Screening::ReportUser not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("Screening::ReportUser STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	go protocol.reportUserHandler(nil, client, callID, packet.Payload())
}
