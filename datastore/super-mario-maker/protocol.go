// Package datastore_super_mario_maker implements the Super Mario Maker DataStore NEX protocol
package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the DataStore (SMM) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodGetObjectInfos is the method ID for the method GetObjectInfos
	MethodGetObjectInfos = 0x2D

	// MethodRateCustomRanking is the method ID for the method RateCustomRanking
	MethodRateCustomRanking = 0x30

	// MethodGetCustomRankingByDataID is the method ID for the method GetCustomRankingByDataID
	MethodGetCustomRankingByDataID = 0x32

	// MethodAddToBufferQueues is the method ID for the method AddToBufferQueues
	MethodAddToBufferQueues = 0x35

	// MethodGetBufferQueue is the method ID for the method GetBufferQueue
	MethodGetBufferQueue = 0x36

	// MethodCompleteAttachFile is the method ID for the method CompleteAttachFile
	MethodCompleteAttachFile = 0x39

	// MethodPrepareAttachFile is the method ID for the method PrepareAttachFile
	MethodPrepareAttachFile = 0x3B

	// MethodGetApplicationConfig is the method ID for the method GetApplicationConfig
	MethodGetApplicationConfig = 0x3D

	// MethodFollowingsLatestCourseSearchObject is the method ID for the method FollowingsLatestCourseSearchObject
	MethodFollowingsLatestCourseSearchObject = 0x41

	// MethodRecommendedCourseSearchObject is the method ID for the method RecommendedCourseSearchObject
	MethodRecommendedCourseSearchObject = 0x42

	// MethodSuggestedCourseSearchObject is the method ID for the method SuggestedCourseSearchObject
	MethodSuggestedCourseSearchObject = 0x44

	// MethodUploadCourseRecord is the method ID for the method UploadCourseRecord
	MethodUploadCourseRecord = 0x47

	// MethodGetCourseRecord is the method ID for the method GetCourseRecord
	MethodGetCourseRecord = 0x48

	// MethodGetApplicationConfigString is the method ID for the method GetApplicationConfigString
	MethodGetApplicationConfigString = 0x4A

	// MethodGetDeletionReason is the method ID for the method GetDeletionReason
	MethodGetDeletionReason = 0x4C

	// MethodGetMetasWithCourseRecord is the method ID for the method GetMetasWithCourseRecord
	MethodGetMetasWithCourseRecord = 0x4E

	// MethodCheckRateCustomRankingCounter is the method ID for the method CheckRateCustomRankingCounter
	MethodCheckRateCustomRankingCounter = 0x4F

	// MethodCTRPickUpCourseSearchObject is the method ID for the method CTRPickUpCourseSearchObject
	MethodCTRPickUpCourseSearchObject = 0x52
)

var patchedMethods = []uint32{
	MethodGetObjectInfos,
	MethodRateCustomRanking,
	MethodGetCustomRankingByDataID,
	MethodAddToBufferQueues,
	MethodGetBufferQueue,
	MethodCompleteAttachFile,
	MethodPrepareAttachFile,
	MethodGetApplicationConfig,
	MethodFollowingsLatestCourseSearchObject,
	MethodRecommendedCourseSearchObject,
	MethodSuggestedCourseSearchObject,
	MethodUploadCourseRecord,
	MethodGetCourseRecord,
	MethodGetApplicationConfigString,
	MethodGetDeletionReason,
	MethodGetMetasWithCourseRecord,
	MethodCheckRateCustomRankingCounter,
	MethodCTRPickUpCourseSearchObject,
}

// DataStoreSuperMarioMakerProtocol handles the DataStore (SMM) NEX protocol. Embeds DataStoreProtocol
type DataStoreSuperMarioMakerProtocol struct {
	Server *nex.Server
	datastore.DataStoreProtocol
	GetObjectInfosHandler                     func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	RateCustomRankingHandler                  func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam)
	GetCustomRankingByDataIDHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIDParam *datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIDParam)
	AddToBufferQueuesHandler                  func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam, buffers [][]byte)
	GetBufferQueueHandler                     func(err error, client *nex.Client, callID uint32, bufferQueueParam *datastore_super_mario_maker_types.BufferQueueParam)
	CompleteAttachFileHandler                 func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *datastore_types.DataStoreCompletePostParam)
	PrepareAttachFileHandler                  func(err error, client *nex.Client, callID uint32, dataStoreAttachFileParam *datastore_super_mario_maker_types.DataStoreAttachFileParam)
	GetApplicationConfigHandler               func(err error, client *nex.Client, callID uint32, applicationID uint32)
	FollowingsLatestCourseSearchObjectHandler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore_types.DataStoreSearchParam, extraData []string)
	RecommendedCourseSearchObjectHandler      func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore_types.DataStoreSearchParam, extraData []string)
	SuggestedCourseSearchObjectHandler        func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	UploadCourseRecordHandler                 func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam)
	GetCourseRecordHandler                    func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam)
	GetApplicationConfigStringHandler         func(err error, client *nex.Client, callID uint32, applicationID uint32)
	GetDeletionReasonHandler                  func(err error, client *nex.Client, callID uint32, dataIDLst []uint64)
	GetMetasWithCourseRecordHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam, dataStoreGetMetaParam *datastore_types.DataStoreGetMetaParam)
	CheckRateCustomRankingCounterHandler      func(err error, client *nex.Client, callID uint32, applicationID uint32)
	CTRPickUpCourseSearchObjectHandler        func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore_types.DataStoreSearchParam, extraData []string)
}

// Setup initializes the protocol
func (protocol *DataStoreSuperMarioMakerProtocol) Setup() {
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
func (protocol *DataStoreSuperMarioMakerProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetObjectInfos:
		go protocol.handleGetObjectInfos(packet)
	case MethodRateCustomRanking:
		go protocol.handleRateCustomRanking(packet)
	case MethodGetCustomRankingByDataID:
		go protocol.handleGetCustomRankingByDataID(packet)
	case MethodAddToBufferQueues:
		go protocol.handleAddToBufferQueues(packet)
	case MethodGetBufferQueue:
		go protocol.handleGetBufferQueue(packet)
	case MethodCompleteAttachFile:
		go protocol.handleCompleteAttachFile(packet)
	case MethodPrepareAttachFile:
		go protocol.handlePrepareAttachFile(packet)
	case MethodGetApplicationConfig:
		go protocol.handleGetApplicationConfig(packet)
	case MethodFollowingsLatestCourseSearchObject:
		go protocol.handleFollowingsLatestCourseSearchObject(packet)
	case MethodRecommendedCourseSearchObject:
		go protocol.handleRecommendedCourseSearchObject(packet)
	case MethodSuggestedCourseSearchObject:
		go protocol.handleSuggestedCourseSearchObject(packet)
	case MethodUploadCourseRecord:
		go protocol.handleUploadCourseRecord(packet)
	case MethodGetCourseRecord:
		go protocol.handleGetCourseRecord(packet)
	case MethodGetApplicationConfigString:
		go protocol.handleGetApplicationConfigString(packet)
	case MethodGetDeletionReason:
		go protocol.handleGetDeletionReason(packet)
	case MethodGetMetasWithCourseRecord:
		go protocol.handleGetMetasWithCourseRecord(packet)
	case MethodCheckRateCustomRankingCounter:
		go protocol.handleCheckRateCustomRankingCounter(packet)
	case MethodCTRPickUpCourseSearchObject:
		go protocol.handleCTRPickUpCourseSearchObject(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStoreSMM method ID: %#v\n", request.MethodID())
	}
}

// NewDataStoreSuperMarioMakerProtocol returns a new DataStoreSuperMarioMakerProtocol
func NewDataStoreSuperMarioMakerProtocol(server *nex.Server) *DataStoreSuperMarioMakerProtocol {
	protocol := &DataStoreSuperMarioMakerProtocol{Server: server}
	protocol.DataStoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
