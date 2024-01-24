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

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Nintendo Badge Arcade) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	server nex.ServerInterface
	dataStoreProtocol
	GetMetaByOwnerID func(err error, packet nex.PacketInterface, callID uint32, param *datastore_nintendo_badge_arcade_types.DataStoreGetMetaByOwnerIDParam) (*nex.RMCMessage, uint32)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.dataStoreProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetMetaByOwnerID:
		protocol.handleGetMetaByOwnerID(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported DataStoreBadgeArcade method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new DataStore (Nintendo Badge Arcade) protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}
	protocol.dataStoreProtocol.SetServer(server)

	return protocol
}
