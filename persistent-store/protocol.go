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
	Server                     *nex.Server
	findByGroupHandler         func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32) uint32
	insertItemHandler          func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, bufData []byte, bReplace bool) uint32
	removeItemHandler          func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string) uint32
	getItemHandler             func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string) uint32
	insertCustomItemHandler    func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, hData *nex.DataHolder, bReplace bool) uint32
	getCustomItemHandler       func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string) uint32
	findItemsBySQLQueryHandler func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, strQuery string) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodFindByGroup:
				go protocol.handleFindByGroup(packet)
			case MethodInsertItem:
				go protocol.handleInsertItem(packet)
			case MethodRemoveItem:
				go protocol.handleRemoveItem(packet)
			case MethodGetItem:
				go protocol.handleGetItem(packet)
			case MethodInsertCustomItem:
				go protocol.handleInsertCustomItem(packet)
			case MethodGetCustomItem:
				go protocol.handleGetCustomItem(packet)
			case MethodFindItemsBySQLQuery:
				go protocol.handleFindItemsBySQLQuery(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Persistent Store method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Persistent Store protocol
func NewProtocol(server *nex.Server) *Protocol {
	persistentStoreProtocol := &Protocol{Server: server}

	persistentStoreProtocol.Setup()

	return persistentStoreProtocol
}
