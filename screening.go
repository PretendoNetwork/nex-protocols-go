package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ScreeningProtocolID is the protocol ID for the Screening protocol
	ScreeningProtocolID = 0x7C

	// ScreeningMethodReportDataStoreContent is the method ID for the method ReportDataStoreContent
	ScreeningMethodReportDataStoreContent = 0x1

	// ScreeningMethodReportUser is the method ID for the method ReportUser
	ScreeningMethodReportUser = 0x2
)

// ScreeningProtocol handles the Screening protocol
type ScreeningProtocol struct {
	server                        *nex.Server
	ReportDataStoreContentHandler func(err error, client *nex.Client, callID uint32)
	ReportUserHandler             func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (screeningProtocol *ScreeningProtocol) Setup() {
	nexServer := screeningProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if ScreeningProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case ScreeningMethodReportDataStoreContent:
				go screeningProtocol.handleReportDataStoreContent(packet)
				break
			case ScreeningMethodReportUser:
				go screeningProtocol.handleReportUser(packet)
				break
			default:
				fmt.Printf("Unsupported Screening method ID: %#v\n", request.MethodID())
				break
			}
		}
	})
}

func (screeningProtocol *ScreeningProtocol) handleReportDataStoreContent(packet nex.PacketInterface) {
	logger.Warning("ScreeningProtocol::ReportDataStoreContent STUBBED")

	if screeningProtocol.ReportDataStoreContentHandler == nil {
		logger.Warning("ScreeningProtocol::ReportDataStoreContent not implemented")
		go respondNotImplemented(packet, ScreeningProtocolID)
		return
	}

}

func (screeningProtocol *ScreeningProtocol) handleReportUser(packet nex.PacketInterface) {
	logger.Warning("ScreeningProtocol::ReportUser STUBBED")

	if screeningProtocol.ReportUserHandler == nil {
		logger.Warning("ScreeningProtocol::ReportUser not implemented")
		go respondNotImplemented(packet, ScreeningProtocolID)
		return
	}

}

// ReportDataStoreContent sets the ReportDataStoreContent handler function
func (screeningProtocol *ScreeningProtocol) ReportDataStoreContent(handler func(err error, client *nex.Client, callID uint32)) {
	screeningProtocol.ReportDataStoreContentHandler = handler
}

// ReportUser sets the ReportUser handler function
func (screeningProtocol *ScreeningProtocol) ReportUser(handler func(err error, client *nex.Client, callID uint32)) {
	screeningProtocol.ReportUserHandler = handler
}

// NewScreeningProtocol returns a new ScreeningProtocol
func NewScreeningProtocol(server *nex.Server) *ScreeningProtocol {
	screeningProtocol := &ScreeningProtocol{server: server}

	screeningProtocol.Setup()

	return screeningProtocol
}
