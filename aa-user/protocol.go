// Package aauser implements the AAUser NEX protocol
package aauser

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	aauser_types "github.com/PretendoNetwork/nex-protocols-go/aa-user/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the AAUser protocol
	ProtocolID = 0x7B

	// MethodRegisterApplication is the method ID for the method RegisterApplication
	MethodRegisterApplication = 0x1

	// MethodUnregisterApplication is the method ID for the method UnregisterApplication
	MethodUnregisterApplication = 0x2

	// MethodSetApplicationInfo is the method ID for the method RegisterApplication
	MethodSetApplicationInfo = 0x3

	// MethodGetApplicationInfo is the method ID for the method GetApplicationInfo
	MethodGetApplicationInfo = 0x4
)

// AAUserProtocol handles the AAUser NEX protocol
type AAUserProtocol struct {
	Server                       *nex.Server
	registerApplicationHandler   func(err error, client *nex.Client, callID uint32, titleID uint64)
	unregisterApplicationHandler func(err error, client *nex.Client, callID uint32, titleID uint64)
	setApplicationInfoHandler    func(err error, client *nex.Client, callID uint32, applicationInfo []*aauser_types.ApplicationInfo)
	getApplicationInfoHandler    func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *AAUserProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodRegisterApplication:
				go protocol.handleRegisterApplication(packet)
			case MethodUnregisterApplication:
				go protocol.handleUnregisterApplication(packet)
			case MethodSetApplicationInfo:
				go protocol.handleSetApplicationInfo(packet)
			case MethodGetApplicationInfo:
				go protocol.handleGetApplicationInfo(packet)
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported AAUser method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewAAUserProtocol returns a new AAUserProtocol
func NewAAUserProtocol(server *nex.Server) *AAUserProtocol {
	protocol := &AAUserProtocol{Server: server}

	protocol.Setup()

	return protocol
}
