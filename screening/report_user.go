package screening

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportUser sets the ReportUser handler function
func (protocol *ScreeningProtocol) ReportUser(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.ReportUserHandler = handler
}

func (protocol *ScreeningProtocol) HandleReportUser(packet nex.PacketInterface) {
	globals.Logger.Warning("Screening::ReportUser STUBBED")

	if protocol.ReportUserHandler == nil {
		globals.Logger.Warning("Screening::ReportUser not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

}
