// Package protocol implements the Screening protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportDataStoreContent sets the ReportDataStoreContent handler function
func (protocol *Protocol) ReportDataStoreContent(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.reportDataStoreContentHandler = handler
}

func (protocol *Protocol) handleReportDataStoreContent(packet nex.PacketInterface) {
	globals.Logger.Warning("Screening::ReportDataStoreContent STUBBED")

	if protocol.reportDataStoreContentHandler == nil {
		globals.Logger.Warning("Screening::ReportDataStoreContent not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
}
