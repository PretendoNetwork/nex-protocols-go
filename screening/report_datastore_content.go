// Package protocol implements the Screening protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportDataStoreContent sets the ReportDataStoreContent handler function
func (protocol *Protocol) ReportDataStoreContent(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.reportDataStoreContentHandler = handler
}

func (protocol *Protocol) handleReportDataStoreContent(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportDataStoreContentHandler == nil {
		globals.Logger.Warning("Screening::ReportDataStoreContent not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Screening::ReportDataStoreContent STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	errorCode = protocol.reportDataStoreContentHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
