// Package screening implements the Screening NEX protocol
package screening

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the Screening protocol
	ProtocolID = 0x7C

	// MethodReportDataStoreContent is the method ID for the method ReportDataStoreContent
	MethodReportDataStoreContent = 0x1

	// MethodReportUser is the method ID for the method ReportUser
	MethodReportUser = 0x2
)

// ScreeningProtocol handles the Screening protocol
type ScreeningProtocol struct {
	Server                        *nex.Server
	ReportDataStoreContentHandler func(err error, client *nex.Client, callID uint32)
	ReportUserHandler             func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *ScreeningProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *ScreeningProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodReportDataStoreContent:
		go protocol.handleReportDataStoreContent(packet)
	case MethodReportUser:
		go protocol.handleReportUser(packet)
	default:
		fmt.Printf("Unsupported Screening method ID: %#v\n", request.MethodID())
	}
}

// NewScreeningProtocol returns a new ScreeningProtocol
func NewScreeningProtocol(server *nex.Server) *ScreeningProtocol {
	protocol := &ScreeningProtocol{Server: server}

	protocol.Setup()

	return protocol
}
