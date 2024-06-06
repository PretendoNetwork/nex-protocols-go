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

	// MethodGetMetaByOwnerID is the method ID for GetMetaByOwnerID
	MethodGetMetaByOwnerID = 0x2E

	// MethodGetMetaByUniqueID is the method ID for GetMetaByUniqueID
	MethodGetMetaByUniqueID = 0x2F

	// MethodSearchHouseNew is the method ID for SearchHouseNew
	MethodSearchHouseNew = 0x30

	// MethodSearchHousePopular is the method ID for SearchHousePopular
	MethodSearchHousePopular = 0x31

	// MethodSearchHouseResident is the method ID for SearchHouseResident
	MethodSearchHouseResident = 0x32

	// MethodSearchHouseContest is the method ID for SearchHouseContest
	MethodSearchHouseContest = 0x33

	// MethodSearchHouseContestRandom is the method ID for SearchHouseContestRandom
	MethodSearchHouseContestRandom = 0x34

	// MethodAddToBufferQueue is the method ID for AddToBufferQueue
	MethodAddToBufferQueue = 0x35

	// MethodGetBufferQueue is the method ID for GetBufferQueue
	MethodGetBufferQueue = 0x36

	// MethodGetBufferQueues is the method ID for GetBufferQueues
	MethodGetBufferQueues = 0x37

	// MethodClearBufferQueues is the method ID for ClearBufferQueues
	MethodClearBufferQueues = 0x38

	// MethodGetContestEntryCount is the method ID for GetContestEntryCount
	MethodGetContestEntryCount = 0x39
)

var patchedMethods = []uint32{
	MethodGetObjectInfos,
	MethodGetMetaByOwnerId,
	MethodGetMetaByUniqueId,
	MethodSearchHouseNew,
	MethodSearchHousePopular,
	MethodSearchHouseResident,
	MethodSearchHouseContest,
	MethodSearchHouseContestRandom,
	MethodAddToBufferQueue,
	MethodGetBufferQueue,
	MethodGetBufferQueues,
	MethodClearBufferQueues,
	MethodGetContestEntryCount,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Animal Crossing: Happy Home Designer) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	GetObjectInfos           func(err error, packet nex.PacketInterface, callId uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	GetMetaByOwnerId         func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreGetMetaByOwnerIdParam) (*nex.RMCMessage, *nex.Error)
	GetMetaByUniqueId        func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreGetMetaByUniqueIdParam) (*nex.RMCMessage, *nex.Error)
	SearchHouseNew           func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreSearchHouseParam) (*nex.RMCMessage, *nex.Error)
	SearchHousePopular       func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreSearchHouseParam) (*nex.RMCMessage, *nex.Error)
	SearchHouseResident      func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreSearchHouseParam) (*nex.RMCMessage, *nex.Error)
	SearchHouseContest       func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreSearchHouseParam) (*nex.RMCMessage, *nex.Error)
	SearchHouseContestRandom func(err error, packet nex.PacketInterface, callId uint32, param *datastore_ac_happy_home_designer_types.DataStoreSearchHouseParam) (*nex.RMCMessage, *nex.Error)
	AddToBufferQueue         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_ac_happy_home_designer_types.BufferQueueParam, buffer *types.QBuffer) (*nex.RMCMessage, *nex.Error)
	GetBufferQueue           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_ac_happy_home_designer_types.BufferQueueParam) (*nex.RMCMessage, *nex.Error)
	GetBufferQueues          func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_ac_happy_home_designer_types.BufferQueueParam]) (*nex.RMCMessage, *nex.Error)
	ClearBufferQueues        func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_ac_happy_home_designer_types.BufferQueueParam]) (*nex.RMCMessage, *nex.Error)
	GetContestEntryCount     func(err error, packet nex.PacketInterface, callID uint32, pEntries *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
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
	case MethodGetMetaByUniqueId:
		protocol.handleGetMetaByUniqueId(packet)
	case MethodSearchHouseNew:
		protocol.handleSearchHouseNew(packet)
	case MethodSearchHousePopular:
		protocol.handleSearchHousePopular(packet)
	case MethodSearchHouseResident:
		protocol.handleSearchHouseResident(packet)
	case MethodSearchHouseContest:
		protocol.handleSearchHouseContest(packet)
	case MethodSearchHouseContestRandom:
		protocol.handleSearchHouseContestRandom(packet)
	case MethodAddToBufferQueue:
		protocol.handleAddToBufferQueue(packet)
	case MethodGetBufferQueue:
		protocol.handleGetBufferQueue(packet)
	case MethodGetBufferQueues:
		protocol.handleGetBufferQueues(packet)
	case MethodClearBufferQueues:
		protocol.handleClearBufferQueues(packet)
	case MethodGetContestEntryCount:
		protocol.handleGetContestEntryCount(packet)
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
