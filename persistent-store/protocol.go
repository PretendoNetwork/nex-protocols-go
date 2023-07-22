// Package persistent_store implements the Persistent Store protocol
package persistent_store

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Nintendo Persistent Store protocol
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

// PersistentStoreProtocol handles the Persistent Store protocol
type PersistentStoreProtocol struct {
	Server                     *nex.Server
	findByGroupHandler         func(err error, client *nex.Client, callID uint32, uiGroup uint32)
	insertItemHandler          func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string, bufData []byte, bReplace bool)
	removeItemHandler          func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string)
	getItemHandler             func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string)
	insertCustomItemHandler    func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string, hData *nex.DataHolder, bReplace bool)
	getCustomItemHandler       func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string)
	findItemsBySQLQueryHandler func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string, strQuery string)
}

// Setup initializes the protocol
func (protocol *PersistentStoreProtocol) Setup() {
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
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported Persistent Store method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewPersistentStoreProtocol returns a new PersistentStoreProtocol
func NewPersistentStoreProtocol(server *nex.Server) *PersistentStoreProtocol {
	persistentStoreProtocol := &PersistentStoreProtocol{Server: server}

	persistentStoreProtocol.Setup()

	return persistentStoreProtocol
}
