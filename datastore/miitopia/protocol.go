// Package protocol implements the Miitopia DataStore protocol
package protocol

import (
	"github.com/PretendoNetwork/nex-go/v2"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_miitopia_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/miitopia/types"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Miitopia) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodSearchMii is the method ID for SearchMii
	MethodSearchMii = 0x2F
)

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Miitopia) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	SearchMii func(err error, packet nex.PacketInterface, callId uint32, param *datastore_miitopia_types.MiiTubeSearchParam) (*nex.RMCMessage, *nex.Error)
}
