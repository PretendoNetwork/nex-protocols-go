package secure_connection

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection protocol
	ProtocolID = 0xB

	// MethodRegister is the method ID for the method Register
	MethodRegister = 0x1

	// MethodRequestConnectionData is the method ID for the method RequestConnectionData
	MethodRequestConnectionData = 0x2

	// MethodRequestURLs is the method ID for the method RequestURLs
	MethodRequestURLs = 0x3

	// MethodRegisterEx is the method ID for the method RegisterEx
	MethodRegisterEx = 0x4

	// MethodTestConnectivity is the method ID for the method TestConnectivity
	MethodTestConnectivity = 0x5

	// MethodUpdateURLs is the method ID for the method UpdateURLs
	MethodUpdateURLs = 0x6

	// MethodReplaceURL is the method ID for the method ReplaceURL
	MethodReplaceURL = 0x7

	// MethodSendReport is the method ID for the method SendReport
	MethodSendReport = 0x8
)

// SecureConnectionProtocol handles the Secure Connection nex protocol
type SecureConnectionProtocol struct {
	Server                       *nex.Server
	RegisterHandler              func(err error, client *nex.Client, callID uint32, stationUrls []*nex.StationURL)
	RequestConnectionDataHandler func(err error, client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)
	RequestURLsHandler           func(err error, client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)
	RegisterExHandler            func(err error, client *nex.Client, callID uint32, stationUrls []*nex.StationURL, loginData *nex.DataHolder)
	TestConnectivityHandler      func(err error, client *nex.Client, callID uint32)
	UpdateURLsHandler            func(err error, client *nex.Client, callID uint32, stationUrls []*nex.StationURL)
	ReplaceURLHandler            func(err error, client *nex.Client, callID uint32, oldStation *nex.StationURL, newStation *nex.StationURL)
	SendReportHandler            func(err error, client *nex.Client, callID uint32, reportID uint32, report []byte)
}

// Setup initializes the protocol
func (protocol *SecureConnectionProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *SecureConnectionProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodRegister:
		go protocol.HandleRegister(packet)
	case MethodRequestConnectionData:
		go protocol.HandleRequestConnectionData(packet)
	case MethodRequestURLs:
		go protocol.HandleRequestURLs(packet)
	case MethodRegisterEx:
		go protocol.HandleRegisterEx(packet)
	case MethodTestConnectivity:
		go protocol.HandleTestConnectivity(packet)
	case MethodUpdateURLs:
		go protocol.HandleUpdateURLs(packet)
	case MethodReplaceURL:
		go protocol.HandleReplaceURL(packet)
	case MethodSendReport:
		go protocol.HandleSendReport(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Secure Connection method ID: %#v\n", request.MethodID())
	}
}

// NewSecureConnectionProtocol returns a new SecureConnectionProtocol
func NewSecureConnectionProtocol(server *nex.Server) *SecureConnectionProtocol {
	protocol := &SecureConnectionProtocol{Server: server}

	protocol.Setup()

	return protocol
}
