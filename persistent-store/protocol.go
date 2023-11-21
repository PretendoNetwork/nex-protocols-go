// Package protocol implements the Persistent Store protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Persistent Store protocol
	ProtocolID = 0x18

	// MethodFindByGroup is the method ID for method FindByGroup
	MethodFindByGroup = 0x1

	// MethodInsertItem is the method ID for method InsertItem
	MethodInsertItem = 0x2

	// MethodRemoveItem is the method ID for method RemoveItem
	MethodRemoveItem = 0x3

	// MethodGetItem is the method ID for method GetItem
	MethodGetItem = 0x4

	// MethodInsertCustomItem is the method ID for method InsertCustomItem
	MethodInsertCustomItem = 0x5

	// MethodGetCustomItem is the method ID for method GetCustomItem
	MethodGetCustomItem = 0x6

	// MethodFindItemsBySQLQuery is the method ID for method FindItemsBySQLQuery
	MethodFindItemsBySQLQuery = 0x7
)

// Protocol handles the Persistent Store protocol
type Protocol struct {
	Server              nex.ServerInterface
	FindByGroup         func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32) (*nex.RMCMessage, uint32)
	InsertItem          func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, bufData []byte, bReplace bool) (*nex.RMCMessage, uint32)
	RemoveItem          func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string) (*nex.RMCMessage, uint32)
	GetItem             func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string) (*nex.RMCMessage, uint32)
	InsertCustomItem    func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, hData *nex.DataHolder, bReplace bool) (*nex.RMCMessage, uint32)
	GetCustomItem       func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string) (*nex.RMCMessage, uint32)
	FindItemsBySQLQuery func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, strQuery string) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodFindByGroup:
				protocol.handleFindByGroup(packet)
			case MethodInsertItem:
				protocol.handleInsertItem(packet)
			case MethodRemoveItem:
				protocol.handleRemoveItem(packet)
			case MethodGetItem:
				protocol.handleGetItem(packet)
			case MethodInsertCustomItem:
				protocol.handleInsertCustomItem(packet)
			case MethodGetCustomItem:
				protocol.handleGetCustomItem(packet)
			case MethodFindItemsBySQLQuery:
				protocol.handleFindItemsBySQLQuery(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Persistent Store method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Persistent Store protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	persistentStoreProtocol := &Protocol{Server: server}

	persistentStoreProtocol.Setup()

	return persistentStoreProtocol
}
