// Package protocol implements the Nintendo Badge Arcade DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/datastore/nintendo-badge-arcade/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Nintendo Badge Arcade) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodGetMetaByOwnerID is the method ID for GetMetaByOwnerID
	MethodGetMetaByOwnerID = 0x2D
)

var patchedMethods = []uint32{
	MethodGetMetaByOwnerID,
}

type datastoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Nintendo Badge Arcade) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	Server *nex.Server
	datastoreProtocol
	GetMetaByOwnerIDHandler func(err error, client *nex.Client, callID uint32, param *datastore_nintendo_badge_arcade_types.DataStoreGetMetaByOwnerIDParam)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {

	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.datastoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetMetaByOwnerID:
		go protocol.handleGetMetaByOwnerID(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStoreBadgeArcade method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new DataStore (Nintendo Badge Arcade) protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.datastoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
