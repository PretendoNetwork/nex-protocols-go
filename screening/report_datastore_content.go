// Package protocol implements the Screening protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportDataStoreContent(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ReportDataStoreContent == nil {
		globals.Logger.Warning("Screening::ReportDataStoreContent not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Screening::ReportDataStoreContent STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	rmcMessage, errorCode := protocol.ReportDataStoreContent(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
