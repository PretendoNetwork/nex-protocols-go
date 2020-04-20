package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreSMMProtocolID is the protocol ID for the DataStore (SMM) protocol. ID is the same as the DataStore protocol
	DataStoreSMMProtocolID = 0x73
)

// DataStoreSMMProtocol handles the DataStore (SMM) nex protocol. Embeds DataStoreProtocol
type DataStoreSMMProtocol struct {
	server *nex.Server
	DataStoreProtocol
}

// Setup initializes the protocol
func (dataStoreSMMProtocol *DataStoreSMMProtocol) Setup() {
	nexServer := dataStoreSMMProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if DataStoreSMMProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case DataStoreMethodGetMeta:
				go dataStoreSMMProtocol.handleGetMeta(packet)
			default:
				fmt.Printf("Unsupported DataStoreSMM method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewDataStoreSMMProtocol returns a new DataStoreSMMProtocol
func NewDataStoreSMMProtocol(server *nex.Server) *DataStoreSMMProtocol {
	dataStoreSMMProtocol := &DataStoreSMMProtocol{server: server}
	dataStoreSMMProtocol.DataStoreProtocol.server = server

	dataStoreSMMProtocol.Setup()

	return dataStoreSMMProtocol
}
