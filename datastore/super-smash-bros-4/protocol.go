package datastore_super_smash_bros_4

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the DataStore (Smash4) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodPostProfile is the method ID for the method PostProfile
	MethodPostProfile = 0x2D

	// MethodGetProfiles is the method ID for the method GetProfiles
	MethodGetProfiles = 0x2E

	// MethodSendPlayReport is the method ID for the method SendPlayReport
	MethodSendPlayReport = 0x2F

	// MethodGetWorldPlayReport is the method ID for the method GetWorldPlayReport
	MethodGetWorldPlayReport = 0x30

	// MethodGetReplayMeta is the method ID for the method GetReplayMeta
	MethodGetReplayMeta = 0x31

	// MethodPrepareGetReplay is the method ID for the method PrepareGetReplay
	MethodPrepareGetReplay = 0x32

	// MethodPreparePostReplay is the method ID for the method PreparePostReplay
	MethodPreparePostReplay = 0x33

	// MethodCompletePostReplay is the method ID for the method CompletePostReplay
	MethodCompletePostReplay = 0x34

	// MethodCheckPostReplay is the method ID for the method CheckPostReplay
	MethodCheckPostReplay = 0x35

	// MethodGetNextReplay is the method ID for the method GetNextReplay
	MethodGetNextReplay = 0x36

	// MethodPreparePostSharedData is the method ID for the method PreparePostSharedData
	MethodPreparePostSharedData = 0x37

	// MethodCompletePostSharedData is the method ID for the method CompletePostSharedData
	MethodCompletePostSharedData = 0x38

	// MethodSearchSharedData is the method ID for the method SearchSharedData
	MethodSearchSharedData = 0x39

	// MethodGetApplicationConfig is the method ID for the method GetApplicationConfig
	MethodGetApplicationConfig = 0x3A

	// MethodSearchReplay is the method ID for the method SearchReplay
	MethodSearchReplay = 0x3B

	// MethodPostFightingPowerScore is the method ID for the method PostFightingPowerScore
	MethodPostFightingPowerScore = 0x3C

	// MethodGetFightingPowerChart is the method ID for the method GetFightingPowerChart
	MethodGetFightingPowerChart = 0x3D

	// MethodGetFightingPowerChartAll is the method ID for the method GetFightingPowerChartAll
	MethodGetFightingPowerChartAll = 0x3E

	// MethodReportSharedData is the method ID for the method ReportSharedData
	MethodReportSharedData = 0x3F
)

var patchedMethods = []uint32{
	MethodPostProfile,
	MethodGetProfiles,
	MethodSendPlayReport,
	MethodGetWorldPlayReport,
	MethodGetReplayMeta,
	MethodPrepareGetReplay,
	MethodPreparePostReplay,
	MethodCompletePostReplay,
	MethodCheckPostReplay,
	MethodGetNextReplay,
	MethodPreparePostSharedData,
	MethodCompletePostSharedData,
	MethodSearchSharedData,
	MethodGetApplicationConfig,
	MethodSearchReplay,
	MethodPostFightingPowerScore,
	MethodGetFightingPowerChart,
	MethodGetFightingPowerChartAll,
	MethodReportSharedData,
}

// DataStoreSuperSmashBros4Protocol handles the DataStore (Smash4) nex protocol. Embeds DataStoreProtocol
type DataStoreSuperSmashBros4Protocol struct {
	Server *nex.Server
	datastore.DataStoreProtocol

	PostProfileHandler              func(err error, client *nex.Client, callID uint32, param *DataStorePostProfileParam)
	GetProfilesHandler              func(err error, client *nex.Client, callID uint32, pidList []uint32)
	SendPlayReportHandler           func(err error, client *nex.Client, callID uint32, playReport []int32)
	GetWorldPlayReportHandler       func(err error, client *nex.Client, callID uint32)
	GetReplayMetaHandler            func(err error, client *nex.Client, callID uint32, param *DataStoreGetReplayMetaParam)
	PrepareGetReplayHandler         func(err error, client *nex.Client, callID uint32, param *DataStorePrepareGetReplayParam)
	PreparePostReplayHandler        func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)
	CompletePostReplayHandler       func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostReplayParam)
	CheckPostReplayHandler          func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)
	GetNextReplayHandler            func(err error, client *nex.Client, callID uint32)
	PreparePostSharedDataHandler    func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostSharedDataParam)
	CompletePostSharedDataHandler   func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostSharedDataParam)
	SearchSharedDataHandler         func(err error, client *nex.Client, callID uint32, param *DataStoreSearchSharedDataParam)
	GetApplicationConfigHandler     func(err error, client *nex.Client, callID uint32, applicationID uint32)
	SearchReplayHandler             func(err error, client *nex.Client, callID uint32, param *DataStoreSearchReplayParam)
	PostFightingPowerScoreHandler   func(err error, client *nex.Client, callID uint32, params []*DataStorePostFightingPowerScoreParam)
	GetFightingPowerChartHandler    func(err error, client *nex.Client, callID uint32, mode uint8)
	GetFightingPowerChartAllHandler func(err error, client *nex.Client, callID uint32)
	ReportSharedDataHandler         func(err error, client *nex.Client, callID uint32, dataID uint64)
}

// Setup initializes the protocol
func (protocol *DataStoreSuperSmashBros4Protocol) Setup() {
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

func (protocol *DataStoreSuperSmashBros4Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodPostProfile:
		go protocol.HandlePostProfile(packet)
	case MethodGetProfiles:
		go protocol.HandleGetProfiles(packet)
	case MethodSendPlayReport:
		go protocol.HandleSendPlayReport(packet)
	case MethodGetWorldPlayReport:
		go protocol.HandleGetWorldPlayReport(packet)
	case MethodGetReplayMeta:
		go protocol.HandleGetReplayMeta(packet)
	case MethodPrepareGetReplay:
		go protocol.HandlePrepareGetReplay(packet)
	case MethodPreparePostReplay:
		go protocol.HandlePreparePostReplay(packet)
	case MethodCompletePostReplay:
		go protocol.HandleCompletePostReplay(packet)
	case MethodCheckPostReplay:
		go protocol.HandleCheckPostReplay(packet)
	case MethodGetNextReplay:
		go protocol.HandleGetNextReplay(packet)
	case MethodPreparePostSharedData:
		go protocol.HandlePreparePostSharedData(packet)
	case MethodCompletePostSharedData:
		go protocol.HandleCompletePostSharedData(packet)
	case MethodSearchSharedData:
		go protocol.HandleSearchSharedData(packet)
	case MethodGetApplicationConfig:
		go protocol.HandleGetApplicationConfig(packet)
	case MethodSearchReplay:
		go protocol.HandleSearchReplay(packet)
	case MethodPostFightingPowerScore:
		go protocol.HandlePostFightingPowerScore(packet)
	case MethodGetFightingPowerChart:
		go protocol.HandleGetFightingPowerChart(packet)
	case MethodGetFightingPowerChartAll:
		go protocol.HandleGetFightingPowerChartAll(packet)
	case MethodReportSharedData:
		go protocol.HandleReportSharedData(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStoreSmash4 method ID: %#v\n", request.MethodID())
	}
}

// NewDataStoreSuperSmashBros4Protocol returns a new DataStoreSuperSmashBros4Protocol
func NewDataStoreSuperSmashBros4Protocol(server *nex.Server) *DataStoreSuperSmashBros4Protocol {
	protocol := &DataStoreSuperSmashBros4Protocol{Server: server}
	protocol.DataStoreProtocol.Server = server

	return protocol
}
