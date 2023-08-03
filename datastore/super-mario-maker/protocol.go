// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Super Mario Maker) protocol
	ProtocolID = 0x73

	// MethodGetObjectInfos is the method ID for the GetObjectInfos method
	MethodGetObjectInfos = 0x2D

	// MethodGetMetaByOwnerID is the method ID for the GetMetaByOwnerID method
	MethodGetMetaByOwnerID = 0x2E

	// MethodCustomSearchObject is the method ID for the CustomSearchObject method
	MethodCustomSearchObject = 0x2F

	// MethodRateCustomRanking is the method ID for the RateCustomRanking method
	MethodRateCustomRanking = 0x30

	// MethodGetCustomRanking is the method ID for the GetCustomRanking method
	MethodGetCustomRanking = 0x31

	// MethodGetCustomRankingByDataID is the method ID for the GetCustomRankingByDataID method
	MethodGetCustomRankingByDataID = 0x32

	// MethodDeleteCustomRanking is the method ID for the DeleteCustomRanking method
	MethodDeleteCustomRanking = 0x33

	// MethodAddToBufferQueue is the method ID for the AddToBufferQueue method
	MethodAddToBufferQueue = 0x34

	// MethodAddToBufferQueues is the method ID for the AddToBufferQueues method
	MethodAddToBufferQueues = 0x35

	// MethodGetBufferQueue is the method ID for the GetBufferQueue method
	MethodGetBufferQueue = 0x36

	// MethodGetBufferQueues is the method ID for the GetBufferQueues method
	MethodGetBufferQueues = 0x37

	// MethodClearBufferQueues is the method ID for the ClearBufferQueues method
	MethodClearBufferQueues = 0x38

	// MethodCompleteAttachFile is the method ID for the CompleteAttachFile method
	MethodCompleteAttachFile = 0x39

	// MethodCompleteAttachFileV1 is the method ID for the CompleteAttachFileV1 method
	MethodCompleteAttachFileV1 = 0x3A

	// MethodPrepareAttachFile is the method ID for the PrepareAttachFile method
	MethodPrepareAttachFile = 0x3B

	// MethodConditionalSearchObject is the method ID for the ConditionalSearchObject method
	MethodConditionalSearchObject = 0x3C

	// MethodGetApplicationConfig is the method ID for the GetApplicationConfig method
	MethodGetApplicationConfig = 0x3D

	// MethodSetApplicationConfig is the method ID for the SetApplicationConfig method
	MethodSetApplicationConfig = 0x3E

	// MethodDeleteApplicationConfig is the method ID for the DeleteApplicationConfig method
	MethodDeleteApplicationConfig = 0x3F

	// MethodLatestCourseSearchObject is the method ID for the LatestCourseSearchObject method
	MethodLatestCourseSearchObject = 0x40

	// MethodFollowingsLatestCourseSearchObject is the method ID for the FollowingsLatestCourseSearchObject method
	MethodFollowingsLatestCourseSearchObject = 0x41

	// MethodRecommendedCourseSearchObject is the method ID for the RecommendedCourseSearchObject method
	MethodRecommendedCourseSearchObject = 0x42

	// MethodScoreRangeCascadedSearchObject is the method ID for the ScoreRangeCascadedSearchObject method
	MethodScoreRangeCascadedSearchObject = 0x43

	// MethodSuggestedCourseSearchObject is the method ID for the SuggestedCourseSearchObject method
	MethodSuggestedCourseSearchObject = 0x44

	// MethodPreparePostObjectWithOwnerIDAndDataID is the method ID for the PreparePostObjectWithOwnerIDAndDataID method
	MethodPreparePostObjectWithOwnerIDAndDataID = 0x45

	// MethodCompletePostObjectWithOwnerID is the method ID for the CompletePostObjectWithOwnerID method
	MethodCompletePostObjectWithOwnerID = 0x46

	// MethodUploadCourseRecord is the method ID for the UploadCourseRecord method
	MethodUploadCourseRecord = 0x47

	// MethodGetCourseRecord is the method ID for the GetCourseRecord method
	MethodGetCourseRecord = 0x48

	// MethodDeleteCourseRecord is the method ID for the DeleteCourseRecord method
	MethodDeleteCourseRecord = 0x49

	// MethodGetApplicationConfigString is the method ID for the GetApplicationConfigString method
	MethodGetApplicationConfigString = 0x4A

	// MethodSetApplicationConfigString is the method ID for the SetApplicationConfigString method
	MethodSetApplicationConfigString = 0x4B

	// MethodGetDeletionReason is the method ID for the GetDeletionReason method
	MethodGetDeletionReason = 0x4C

	// MethodSetDeletionReason is the method ID for the SetDeletionReason method
	MethodSetDeletionReason = 0x4D

	// MethodGetMetasWithCourseRecord is the method ID for the GetMetasWithCourseRecord method
	MethodGetMetasWithCourseRecord = 0x4E

	// MethodCheckRateCustomRankingCounter is the method ID for the CheckRateCustomRankingCounter method
	MethodCheckRateCustomRankingCounter = 0x4F

	// MethodResetRateCustomRankingCounter is the method ID for the ResetRateCustomRankingCounter method
	MethodResetRateCustomRankingCounter = 0x50

	// MethodBestScoreRateCourseSearchObject is the method ID for the BestScoreRateCourseSearchObject method
	MethodBestScoreRateCourseSearchObject = 0x51

	// MethodCTRPickUpCourseSearchObject is the method ID for the CTRPickUpCourseSearchObject method
	MethodCTRPickUpCourseSearchObject = 0x52

	// MethodSetCachedRanking is the method ID for the SetCachedRanking method
	MethodSetCachedRanking = 0x53

	// MethodDeleteCachedRanking is the method ID for the DeleteCachedRanking method
	MethodDeleteCachedRanking = 0x54

	// MethodChangePlayablePlatform is the method ID for the ChangePlayablePlatform method
	MethodChangePlayablePlatform = 0x55

	// MethodSearchUnknownPlatformObjects is the method ID for the SearchUnknownPlatformObjects method
	MethodSearchUnknownPlatformObjects = 0x56

	// MethodReportCourse is the method ID for the ReportCourse method
	MethodReportCourse = 0x57
)

var patchedMethods = []uint32{
	MethodGetObjectInfos,
	MethodGetMetaByOwnerID,
	MethodCustomSearchObject,
	MethodRateCustomRanking,
	MethodGetCustomRanking,
	MethodGetCustomRankingByDataID,
	MethodDeleteCustomRanking,
	MethodAddToBufferQueue,
	MethodAddToBufferQueues,
	MethodGetBufferQueue,
	MethodGetBufferQueues,
	MethodClearBufferQueues,
	MethodCompleteAttachFile,
	MethodCompleteAttachFileV1,
	MethodPrepareAttachFile,
	MethodConditionalSearchObject,
	MethodGetApplicationConfig,
	MethodSetApplicationConfig,
	MethodDeleteApplicationConfig,
	MethodLatestCourseSearchObject,
	MethodFollowingsLatestCourseSearchObject,
	MethodRecommendedCourseSearchObject,
	MethodScoreRangeCascadedSearchObject,
	MethodSuggestedCourseSearchObject,
	MethodPreparePostObjectWithOwnerIDAndDataID,
	MethodCompletePostObjectWithOwnerID,
	MethodUploadCourseRecord,
	MethodGetCourseRecord,
	MethodDeleteCourseRecord,
	MethodGetApplicationConfigString,
	MethodSetApplicationConfigString,
	MethodGetDeletionReason,
	MethodSetDeletionReason,
	MethodGetMetasWithCourseRecord,
	MethodCheckRateCustomRankingCounter,
	MethodResetRateCustomRankingCounter,
	MethodBestScoreRateCourseSearchObject,
	MethodCTRPickUpCourseSearchObject,
	MethodSetCachedRanking,
	MethodDeleteCachedRanking,
	MethodChangePlayablePlatform,
	MethodSearchUnknownPlatformObjects,
	MethodReportCourse,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Super Mario Maker) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	Server *nex.Server
	dataStoreProtocol
	getObjectInfosHandler                        func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	getMetaByOwnerIDHandler                      func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetMetaByOwnerIDParam)
	customSearchObjectHandler                    func(err error, client *nex.Client, callID uint32, condition uint32, param *datastore_types.DataStoreSearchParam)
	rateCustomRankingHandler                     func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam)
	getCustomRankingHandler                      func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingParam)
	getCustomRankingByDataIDHandler              func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIDParam)
	deleteCustomRankingHandler                   func(err error, client *nex.Client, callID uint32, dataIDList []uint64)
	addToBufferQueueHandler                      func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam, buffer []byte)
	addToBufferQueuesHandler                     func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam, buffers [][]byte)
	getBufferQueueHandler                        func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam)
	getBufferQueuesHandler                       func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam)
	clearBufferQueuesHandler                     func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam)
	completeAttachFileHandler                    func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParam)
	completeAttachFileV1Handler                  func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParamV1)
	prepareAttachFileHandler                     func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreAttachFileParam)
	conditionalSearchObjectHandler               func(err error, client *nex.Client, callID uint32, condition uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	getApplicationConfigHandler                  func(err error, client *nex.Client, callID uint32, applicationID uint32)
	setApplicationConfigHandler                  func(err error, client *nex.Client, callID uint32, applicationID uint32, key uint32, value int32)
	deleteApplicationConfigHandler               func(err error, client *nex.Client, callID uint32, applicationID uint32, key uint32)
	latestCourseSearchObjectHandler              func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	followingsLatestCourseSearchObjectHandler    func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	recommendedCourseSearchObjectHandler         func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	scoreRangeCascadedSearchObjectHandler        func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	suggestedCourseSearchObjectHandler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	preparePostObjectWithOwnerIDAndDataIDHandler func(err error, client *nex.Client, callID uint32, ownerID uint32, dataID uint64, param *datastore_types.DataStorePreparePostParam)
	completePostObjectWithOwnerIDHandler         func(err error, client *nex.Client, callID uint32, ownerID uint32, param *datastore_types.DataStoreCompletePostParam)
	uploadCourseRecordHandler                    func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam)
	getCourseRecordHandler                       func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam)
	deleteCourseRecordHandler                    func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam)
	getApplicationConfigStringHandler            func(err error, client *nex.Client, callID uint32, applicationID uint32)
	setApplicationConfigStringHandler            func(err error, client *nex.Client, callID uint32, applicationID uint32, key uint32, value string)
	getDeletionReasonHandler                     func(err error, client *nex.Client, callID uint32, dataIDLst []uint64)
	setDeletionReasonHandler                     func(err error, client *nex.Client, callID uint32, dataIDLst []uint64, deletionReason uint32)
	getMetasWithCourseRecordHandler              func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam, metaParam *datastore_types.DataStoreGetMetaParam)
	checkRateCustomRankingCounterHandler         func(err error, client *nex.Client, callID uint32, applicationID uint32)
	resetRateCustomRankingCounterHandler         func(err error, client *nex.Client, callID uint32, applicationID uint32)
	bestScoreRateCourseSearchObjectHandler       func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	cTRPickUpCourseSearchObjectHandler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)
	setCachedRankingHandler                      func(err error, client *nex.Client, callID uint32, rankingType string, rankingArgs []string, dataIDLst []uint64)
	deleteCachedRankingHandler                   func(err error, client *nex.Client, callID uint32, rankingType string, rankingArgs []string)
	changePlayablePlatformHandler                func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam)
	searchUnknownPlatformObjectsHandler          func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	reportCourseHandler                          func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreReportCourseParam)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.dataStoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetObjectInfos:
		go protocol.handleGetObjectInfos(packet)
	case MethodGetMetaByOwnerID:
		go protocol.handleGetMetaByOwnerID(packet)
	case MethodCustomSearchObject:
		go protocol.handleCustomSearchObject(packet)
	case MethodRateCustomRanking:
		go protocol.handleRateCustomRanking(packet)
	case MethodGetCustomRanking:
		go protocol.handleGetCustomRanking(packet)
	case MethodGetCustomRankingByDataID:
		go protocol.handleGetCustomRankingByDataID(packet)
	case MethodDeleteCustomRanking:
		go protocol.handleDeleteCustomRanking(packet)
	case MethodAddToBufferQueue:
		go protocol.handleAddToBufferQueue(packet)
	case MethodAddToBufferQueues:
		go protocol.handleAddToBufferQueues(packet)
	case MethodGetBufferQueue:
		go protocol.handleGetBufferQueue(packet)
	case MethodGetBufferQueues:
		go protocol.handleGetBufferQueues(packet)
	case MethodClearBufferQueues:
		go protocol.handleClearBufferQueues(packet)
	case MethodCompleteAttachFile:
		go protocol.handleCompleteAttachFile(packet)
	case MethodCompleteAttachFileV1:
		go protocol.handleCompleteAttachFileV1(packet)
	case MethodPrepareAttachFile:
		go protocol.handlePrepareAttachFile(packet)
	case MethodConditionalSearchObject:
		go protocol.handleConditionalSearchObject(packet)
	case MethodGetApplicationConfig:
		go protocol.handleGetApplicationConfig(packet)
	case MethodSetApplicationConfig:
		go protocol.handleSetApplicationConfig(packet)
	case MethodDeleteApplicationConfig:
		go protocol.handleDeleteApplicationConfig(packet)
	case MethodLatestCourseSearchObject:
		go protocol.handleLatestCourseSearchObject(packet)
	case MethodFollowingsLatestCourseSearchObject:
		go protocol.handleFollowingsLatestCourseSearchObject(packet)
	case MethodRecommendedCourseSearchObject:
		go protocol.handleRecommendedCourseSearchObject(packet)
	case MethodScoreRangeCascadedSearchObject:
		go protocol.handleScoreRangeCascadedSearchObject(packet)
	case MethodSuggestedCourseSearchObject:
		go protocol.handleSuggestedCourseSearchObject(packet)
	case MethodPreparePostObjectWithOwnerIDAndDataID:
		go protocol.handlePreparePostObjectWithOwnerIDAndDataID(packet)
	case MethodCompletePostObjectWithOwnerID:
		go protocol.handleCompletePostObjectWithOwnerID(packet)
	case MethodUploadCourseRecord:
		go protocol.handleUploadCourseRecord(packet)
	case MethodGetCourseRecord:
		go protocol.handleGetCourseRecord(packet)
	case MethodDeleteCourseRecord:
		go protocol.handleDeleteCourseRecord(packet)
	case MethodGetApplicationConfigString:
		go protocol.handleGetApplicationConfigString(packet)
	case MethodSetApplicationConfigString:
		go protocol.handleSetApplicationConfigString(packet)
	case MethodGetDeletionReason:
		go protocol.handleGetDeletionReason(packet)
	case MethodSetDeletionReason:
		go protocol.handleSetDeletionReason(packet)
	case MethodGetMetasWithCourseRecord:
		go protocol.handleGetMetasWithCourseRecord(packet)
	case MethodCheckRateCustomRankingCounter:
		go protocol.handleCheckRateCustomRankingCounter(packet)
	case MethodResetRateCustomRankingCounter:
		go protocol.handleResetRateCustomRankingCounter(packet)
	case MethodBestScoreRateCourseSearchObject:
		go protocol.handleBestScoreRateCourseSearchObject(packet)
	case MethodCTRPickUpCourseSearchObject:
		go protocol.handleCTRPickUpCourseSearchObject(packet)
	case MethodSetCachedRanking:
		go protocol.handleSetCachedRanking(packet)
	case MethodDeleteCachedRanking:
		go protocol.handleDeleteCachedRanking(packet)
	case MethodChangePlayablePlatform:
		go protocol.handleChangePlayablePlatform(packet)
	case MethodSearchUnknownPlatformObjects:
		go protocol.handleSearchUnknownPlatformObjects(packet)
	case MethodReportCourse:
		go protocol.handleReportCourse(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStore (Super Mario Maker) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new DataStoreSuperMarioMaker protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.dataStoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
