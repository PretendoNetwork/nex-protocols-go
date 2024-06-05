// Package protocol implements the Animal Crossing: Happy Home Designer protocol
package protocol

import (
	"fmt"
	"slices"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_ac_happy_home_designer_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/ac-happy-home-designer/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Animal Crossing: Happy Home Designer) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodGetObjectInfos is the method ID for GetObjectInfos
	MethodGetObjectInfos = 0x2D

	// MethodGetMetaByOwnerId is the method ID for GetMetaByOwnerId
	MethodGetMetaByOwnerId = 0x2E

	// MethodGetMetaByUniqueId is the method ID for GetMetaByUniqueId
	MethodGetMetaByUniqueId = 0x2F

	// MethodSearchHouseNew is the method ID for SearchHouseNew
	MethodSearchHouseNew = 0x30

	// MethodSearchHousePopular is the method ID for SearchHousePopular
	MethodSearchHousePopular = 0x31

	// MethodSearchHouseResident is the method ID for SearchHouseResident
	MethodSearchHouseResident = 0x32

	// MethodSearchHouseContent is the method ID for SearchHouseContent
	MethodSearchHouseContent = 0x33

	// MethodSearchHouseContestRandom is the method ID for SearchHouseContentRandom
	MethodSearchHouseContentRandom = 0x34

	// MethodAddToBufferQueue is the method ID for AddToBufferQueue
	MethodAddToBufferQueue = 0x35

	// MethodGetBufferQueue is the method ID for GetBufferQueue
	MethodGetBufferQueue = 0x36

	// MethodGetBufferQueues is the method ID for GetBufferQueues
	MethodGetBuffersQueues = 0x37

	// MethodClearBufferQueues is the method ID for ClearBufferQueues
	MethodClearBuffersQueues = 0x38

	// MethodGetContestEntryCount is the method ID for GetContestEntryCount
	MethodGetContestEntryCount = 0x39
)

var patchedMethods = []uint32{
	MethodGetObjectInfos,
	MethodGetMetaByOwnerId,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Animal Crossing: Happy Home Designer) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	GetObjectInfos   func(err error, packet nex.PacketInterface, callId uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	GetMetaByOwnerId func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreGetMetaByOwnerIdParam) (*nex.RMCMessage, *nex.Error)
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
	case MethodGetObjectInfos:
		protocol.handleGetObjectInfos(packet)
	case MethodGetMetaByOwnerId:
		protocol.handleGetMetaByOwnerId(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported DataStoreHappyHomeDesigner method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new DataStore (Animal Crossing: Happy Home Designer) protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.dataStoreProtocol.SetEndpoint(endpoint)

	return protocol
}
