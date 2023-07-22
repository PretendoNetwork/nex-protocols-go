// Package storage_manager implements the StorageManager NEX protocol
package storage_manager

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

// StorageManagerProtocol handles the StorageManager NEX protocol
type StorageManagerProtocol struct {
	Server                    *nex.Server
	acquireCardIDHandler      func(err error, client *nex.Client, callID uint32)
	activateWithCardIDHandler func(err error, client *nex.Client, callID uint32, unknown uint8, cardID uint64)
}

// Setup initializes the protocol
func (protocol *StorageManagerProtocol) Setup() {
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

// NewStorageManagerProtocol returns a new StorageManagerProtocol
func NewStorageManagerProtocol(server *nex.Server) *StorageManagerProtocol {
	protocol := &StorageManagerProtocol{Server: server}

	protocol.Setup()

	return protocol
}
