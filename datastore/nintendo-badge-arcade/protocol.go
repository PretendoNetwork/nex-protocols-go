package datastore_nintendo_badge_arcade

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/datastore/nintendo-badge-arcade/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Badge Arcade) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodGetMetaByOwnerID is the method ID for GetMetaByOwnerID
	MethodGetMetaByOwnerID = 0x2D
)

var patchedMethods = []uint32{
	MethodGetMetaByOwnerID,
}

// DataStoreNintendoBadgeArcadeProtocol handles the DataStore (Badge Arcade) nex protocol. Embeds DataStoreProtocol
type DataStoreNintendoBadgeArcadeProtocol struct {
	Server *nex.Server
	datastore.DataStoreProtocol
	GetMetaByOwnerIDHandler func(err error, client *nex.Client, callID uint32, param *datastore_nintendo_badge_arcade_types.DataStoreGetMetaByOwnerIDParam)
}

// Setup initializes the protocol
func (protocol *DataStoreNintendoBadgeArcadeProtocol) Setup() {

	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.DataStoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *DataStoreNintendoBadgeArcadeProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetMetaByOwnerID:
		go protocol.handleGetMetaByOwnerID(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStoreBadgeArcade method ID: %#v\n", request.MethodID())
	}
}

// NewDataStoreNintendoBadgeArcadeProtocol returns a new DataStoreNintendoBadgeArcadeProtocol
func NewDataStoreNintendoBadgeArcadeProtocol(server *nex.Server) *DataStoreNintendoBadgeArcadeProtocol {
	protocol := &DataStoreNintendoBadgeArcadeProtocol{Server: server}
	protocol.DataStoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
