package datastore_nintendo_badge_arcade

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Badge Arcade) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodGetMetaByOwnerId is the method ID for GetMetaByOwnerId
	MethodGetMetaByOwnerId = 0x2D
)

var patchedMethods = []uint32{
	MethodGetMetaByOwnerId,
}

// DataStoreNintendoBadgeArcadeProtocol handles the DataStore (Badge Arcade) nex protocol. Embeds DataStoreProtocol
type DataStoreNintendoBadgeArcadeProtocol struct {
	Server *nex.Server
	datastore.DataStoreProtocol
	GetMetaByOwnerIdHandler func(err error, client *nex.Client, callID uint32, param *DataStoreGetMetaByOwnerIdParam)
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

func (protocol *DataStoreNintendoBadgeArcadeProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetMetaByOwnerId:
		go protocol.HandleGetMetaByOwnerId(packet)
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
