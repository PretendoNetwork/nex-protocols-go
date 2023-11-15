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
	Server nex.ServerInterface
	dataStoreProtocol
	GetObjectInfos                        func(err error, packet nex.PacketInterface, callID uint32, dataIDs []uint64) (*nex.RMCMessage, uint32)
	GetMetaByOwnerID                      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetMetaByOwnerIDParam) (*nex.RMCMessage, uint32)
	CustomSearchObject                    func(err error, packet nex.PacketInterface, callID uint32, condition uint32, param *datastore_types.DataStoreSearchParam) (*nex.RMCMessage, uint32)
	RateCustomRanking                     func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam) (*nex.RMCMessage, uint32)
	GetCustomRanking                      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingParam) (*nex.RMCMessage, uint32)
	GetCustomRankingByDataID              func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIDParam) (*nex.RMCMessage, uint32)
	DeleteCustomRanking                   func(err error, packet nex.PacketInterface, callID uint32, dataIDList []uint64) (*nex.RMCMessage, uint32)
	AddToBufferQueue                      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam, buffer []byte) (*nex.RMCMessage, uint32)
	AddToBufferQueues                     func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam, buffers [][]byte) (*nex.RMCMessage, uint32)
	GetBufferQueue                        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam) (*nex.RMCMessage, uint32)
	GetBufferQueues                       func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam) (*nex.RMCMessage, uint32)
	ClearBufferQueues                     func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam) (*nex.RMCMessage, uint32)
	CompleteAttachFile                    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, uint32)
	CompleteAttachFileV1                  func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) (*nex.RMCMessage, uint32)
	PrepareAttachFile                     func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreAttachFileParam) (*nex.RMCMessage, uint32)
	ConditionalSearchObject               func(err error, packet nex.PacketInterface, callID uint32, condition uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	GetApplicationConfig                  func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32) (*nex.RMCMessage, uint32)
	SetApplicationConfig                  func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32, key uint32, value int32) (*nex.RMCMessage, uint32)
	DeleteApplicationConfig               func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32, key uint32) (*nex.RMCMessage, uint32)
	LatestCourseSearchObject              func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	FollowingsLatestCourseSearchObject    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	RecommendedCourseSearchObject         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	ScoreRangeCascadedSearchObject        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	SuggestedCourseSearchObject           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	PreparePostObjectWithOwnerIDAndDataID func(err error, packet nex.PacketInterface, callID uint32, ownerID uint32, dataID uint64, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, uint32)
	CompletePostObjectWithOwnerID         func(err error, packet nex.PacketInterface, callID uint32, ownerID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, uint32)
	UploadCourseRecord                    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam) (*nex.RMCMessage, uint32)
	GetCourseRecord                       func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam) (*nex.RMCMessage, uint32)
	DeleteCourseRecord                    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam) (*nex.RMCMessage, uint32)
	GetApplicationConfigString            func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32) (*nex.RMCMessage, uint32)
	SetApplicationConfigString            func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32, key uint32, value string) (*nex.RMCMessage, uint32)
	GetDeletionReason                     func(err error, packet nex.PacketInterface, callID uint32, dataIDLst []uint64) (*nex.RMCMessage, uint32)
	SetDeletionReason                     func(err error, packet nex.PacketInterface, callID uint32, dataIDLst []uint64, deletionReason uint32) (*nex.RMCMessage, uint32)
	GetMetasWithCourseRecord              func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam, metaParam *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, uint32)
	CheckRateCustomRankingCounter         func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32) (*nex.RMCMessage, uint32)
	ResetRateCustomRankingCounter         func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32) (*nex.RMCMessage, uint32)
	BestScoreRateCourseSearchObject       func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	CTRPickUpCourseSearchObject           func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) (*nex.RMCMessage, uint32)
	SetCachedRanking                      func(err error, packet nex.PacketInterface, callID uint32, rankingType string, rankingArgs []string, dataIDLst []uint64) (*nex.RMCMessage, uint32)
	DeleteCachedRanking                   func(err error, packet nex.PacketInterface, callID uint32, rankingType string, rankingArgs []string) (*nex.RMCMessage, uint32)
	ChangePlayablePlatform                func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam) (*nex.RMCMessage, uint32)
	SearchUnknownPlatformObjects          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	ReportCourse                          func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreReportCourseParam) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID) {
				protocol.HandlePacket(packet)
			} else {
				protocol.dataStoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodGetObjectInfos:
		protocol.handleGetObjectInfos(packet)
	case MethodGetMetaByOwnerID:
		protocol.handleGetMetaByOwnerID(packet)
	case MethodCustomSearchObject:
		protocol.handleCustomSearchObject(packet)
	case MethodRateCustomRanking:
		protocol.handleRateCustomRanking(packet)
	case MethodGetCustomRanking:
		protocol.handleGetCustomRanking(packet)
	case MethodGetCustomRankingByDataID:
		protocol.handleGetCustomRankingByDataID(packet)
	case MethodDeleteCustomRanking:
		protocol.handleDeleteCustomRanking(packet)
	case MethodAddToBufferQueue:
		protocol.handleAddToBufferQueue(packet)
	case MethodAddToBufferQueues:
		protocol.handleAddToBufferQueues(packet)
	case MethodGetBufferQueue:
		protocol.handleGetBufferQueue(packet)
	case MethodGetBufferQueues:
		protocol.handleGetBufferQueues(packet)
	case MethodClearBufferQueues:
		protocol.handleClearBufferQueues(packet)
	case MethodCompleteAttachFile:
		protocol.handleCompleteAttachFile(packet)
	case MethodCompleteAttachFileV1:
		protocol.handleCompleteAttachFileV1(packet)
	case MethodPrepareAttachFile:
		protocol.handlePrepareAttachFile(packet)
	case MethodConditionalSearchObject:
		protocol.handleConditionalSearchObject(packet)
	case MethodGetApplicationConfig:
		protocol.handleGetApplicationConfig(packet)
	case MethodSetApplicationConfig:
		protocol.handleSetApplicationConfig(packet)
	case MethodDeleteApplicationConfig:
		protocol.handleDeleteApplicationConfig(packet)
	case MethodLatestCourseSearchObject:
		protocol.handleLatestCourseSearchObject(packet)
	case MethodFollowingsLatestCourseSearchObject:
		protocol.handleFollowingsLatestCourseSearchObject(packet)
	case MethodRecommendedCourseSearchObject:
		protocol.handleRecommendedCourseSearchObject(packet)
	case MethodScoreRangeCascadedSearchObject:
		protocol.handleScoreRangeCascadedSearchObject(packet)
	case MethodSuggestedCourseSearchObject:
		protocol.handleSuggestedCourseSearchObject(packet)
	case MethodPreparePostObjectWithOwnerIDAndDataID:
		protocol.handlePreparePostObjectWithOwnerIDAndDataID(packet)
	case MethodCompletePostObjectWithOwnerID:
		protocol.handleCompletePostObjectWithOwnerID(packet)
	case MethodUploadCourseRecord:
		protocol.handleUploadCourseRecord(packet)
	case MethodGetCourseRecord:
		protocol.handleGetCourseRecord(packet)
	case MethodDeleteCourseRecord:
		protocol.handleDeleteCourseRecord(packet)
	case MethodGetApplicationConfigString:
		protocol.handleGetApplicationConfigString(packet)
	case MethodSetApplicationConfigString:
		protocol.handleSetApplicationConfigString(packet)
	case MethodGetDeletionReason:
		protocol.handleGetDeletionReason(packet)
	case MethodSetDeletionReason:
		protocol.handleSetDeletionReason(packet)
	case MethodGetMetasWithCourseRecord:
		protocol.handleGetMetasWithCourseRecord(packet)
	case MethodCheckRateCustomRankingCounter:
		protocol.handleCheckRateCustomRankingCounter(packet)
	case MethodResetRateCustomRankingCounter:
		protocol.handleResetRateCustomRankingCounter(packet)
	case MethodBestScoreRateCourseSearchObject:
		protocol.handleBestScoreRateCourseSearchObject(packet)
	case MethodCTRPickUpCourseSearchObject:
		protocol.handleCTRPickUpCourseSearchObject(packet)
	case MethodSetCachedRanking:
		protocol.handleSetCachedRanking(packet)
	case MethodDeleteCachedRanking:
		protocol.handleDeleteCachedRanking(packet)
	case MethodChangePlayablePlatform:
		protocol.handleChangePlayablePlatform(packet)
	case MethodSearchUnknownPlatformObjects:
		protocol.handleSearchUnknownPlatformObjects(packet)
	case MethodReportCourse:
		protocol.handleReportCourse(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported DataStore (Super Mario Maker) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new DataStoreSuperMarioMaker protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.dataStoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
