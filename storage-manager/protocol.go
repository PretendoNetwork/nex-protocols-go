// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the StorageManager protocol
	ProtocolID = 0x6E

	// MethodAcquireCardID is the method ID for the method AcquireCardID
	MethodAcquireCardID = 0x4

	// MethodActivateWithCardID is the method ID for the method ActivateWithCardID
	MethodActivateWithCardID = 0x5
)

// Protocol stores all the RMC method handlers for the StorageManager protocol and listens for requests
type Protocol struct {
	Server                    *nex.Server
	acquireCardIDHandler      func(err error, client *nex.Client, callID uint32)
	activateWithCardIDHandler func(err error, client *nex.Client, callID uint32, unknown uint8, cardID uint64)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodAcquireCardID:
				go protocol.handleAcquireCardID(packet)
			case MethodActivateWithCardID:
				go protocol.handleActivateWithCardID(packet)
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported StorageManager method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new StorageManager protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
