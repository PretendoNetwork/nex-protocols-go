// Package screening implements the Screening NEX protocol
package screening

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportDataStoreContent sets the ReportDataStoreContent handler function
func (protocol *ScreeningProtocol) ReportDataStoreContent(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.ReportDataStoreContentHandler = handler
}

func (protocol *ScreeningProtocol) handleReportDataStoreContent(packet nex.PacketInterface) {
	globals.Logger.Warning("Screening::ReportDataStoreContent STUBBED")

	if protocol.ReportDataStoreContentHandler == nil {
		globals.Logger.Warning("Screening::ReportDataStoreContent not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
}
