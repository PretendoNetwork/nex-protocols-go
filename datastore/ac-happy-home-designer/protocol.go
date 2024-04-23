// Package protocol implements the Animal Crossing: Happy Home Designer protocol
package protocol

import (
	"github.com/PretendoNetwork/nex-go/v2"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
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

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Animal Crossing: Happy Home Designer) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
}

// NewProtocol returns a new DataStore (Animal Crossing: Happy Home Designer) protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.dataStoreProtocol.SetEndpoint(endpoint)

	return protocol
}
