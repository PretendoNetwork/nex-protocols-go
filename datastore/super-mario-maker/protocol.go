package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
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

	// MethodGetCustomRankingByDataId is the method ID for the method GetCustomRankingByDataId
	MethodGetCustomRankingByDataId = 0x32

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

	// MethodSuggestedCourseSearchObjectis the method ID for the method SuggestedCourseSearchObject
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
	MethodGetCustomRankingByDataId,
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

// DataStoreSuperMarioMakerProtocol handles the DataStore (SMM) nex protocol. Embeds DataStoreProtocol
type DataStoreSuperMarioMakerProtocol struct {
	Server *nex.Server
	datastore.DataStoreProtocol
	GetObjectInfosHandler                     func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	RateCustomRankingHandler                  func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*DataStoreRateCustomRankingParam)
	GetCustomRankingByDataIdHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam)
	AddToBufferQueuesHandler                  func(err error, client *nex.Client, callID uint32, params []*BufferQueueParam, buffers [][]byte)
	GetBufferQueueHandler                     func(err error, client *nex.Client, callID uint32, bufferQueueParam *BufferQueueParam)
	CompleteAttachFileHandler                 func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *datastore.DataStoreCompletePostParam)
	PrepareAttachFileHandler                  func(err error, client *nex.Client, callID uint32, dataStoreAttachFileParam *DataStoreAttachFileParam)
	GetApplicationConfigHandler               func(err error, client *nex.Client, callID uint32, applicationID uint32)
	FollowingsLatestCourseSearchObjectHandler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore.DataStoreSearchParam, extraData []string)
	RecommendedCourseSearchObjectHandler      func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore.DataStoreSearchParam, extraData []string)
	SuggestedCourseSearchObjectHandler        func(err error, client *nex.Client, callID uint32, param *datastore.DataStoreSearchParam, extraData []string)
	UploadCourseRecordHandler                 func(err error, client *nex.Client, callID uint32, param *DataStoreUploadCourseRecordParam)
	GetCourseRecordHandler                    func(err error, client *nex.Client, callID uint32, param *DataStoreGetCourseRecordParam)
	GetApplicationConfigStringHandler         func(err error, client *nex.Client, callID uint32, applicationID uint32)
	GetDeletionReasonHandler                  func(err error, client *nex.Client, callID uint32, dataIdLst []uint64)
	GetMetasWithCourseRecordHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*DataStoreGetCourseRecordParam, dataStoreGetMetaParam *datastore.DataStoreGetMetaParam)
	CheckRateCustomRankingCounterHandler      func(err error, client *nex.Client, callID uint32, applicationID uint32)
	CTRPickUpCourseSearchObjectHandler        func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore.DataStoreSearchParam, extraData []string)
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

func (protocol *DataStoreSuperMarioMakerProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetObjectInfos:
		go protocol.HandleGetObjectInfos(packet)
	case MethodRateCustomRanking:
		go protocol.HandleRateCustomRanking(packet)
	case MethodGetCustomRankingByDataId:
		go protocol.HandleGetCustomRankingByDataId(packet)
	case MethodAddToBufferQueues:
		go protocol.HandleAddToBufferQueues(packet)
	case MethodGetBufferQueue:
		go protocol.HandleGetBufferQueue(packet)
	case MethodCompleteAttachFile:
		go protocol.HandleCompleteAttachFile(packet)
	case MethodPrepareAttachFile:
		go protocol.HandlePrepareAttachFile(packet)
	case MethodGetApplicationConfig:
		go protocol.HandleGetApplicationConfig(packet)
	case MethodFollowingsLatestCourseSearchObject:
		go protocol.HandleFollowingsLatestCourseSearchObject(packet)
	case MethodRecommendedCourseSearchObject:
		go protocol.HandleRecommendedCourseSearchObject(packet)
	case MethodSuggestedCourseSearchObject:
		go protocol.HandleSuggestedCourseSearchObject(packet)
	case MethodUploadCourseRecord:
		go protocol.HandleUploadCourseRecord(packet)
	case MethodGetCourseRecord:
		go protocol.HandleGetCourseRecord(packet)
	case MethodGetApplicationConfigString:
		go protocol.HandleGetApplicationConfigString(packet)
	case MethodGetDeletionReason:
		go protocol.HandleGetDeletionReason(packet)
	case MethodGetMetasWithCourseRecord:
		go protocol.HandleGetMetasWithCourseRecord(packet)
	case MethodCheckRateCustomRankingCounter:
		go protocol.HandleCheckRateCustomRankingCounter(packet)
	case MethodCTRPickUpCourseSearchObject:
		go protocol.HandleCTRPickUpCourseSearchObject(packet)
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
