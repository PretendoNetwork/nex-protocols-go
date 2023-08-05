// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

const (
	// ProtocolID is the protocol ID for the Ranking2 protocol
	ProtocolID = 0x7A

	// MethodPutScore is the method ID for the method PutScore
	MethodPutScore = 0x1

	// MethodGetCommonData is the method ID for the method GetCommonData
	MethodGetCommonData = 0x2

	// MethodPutCommonData is the method ID for the method PutCommonData
	MethodPutCommonData = 0x3

	// MethodDeleteCommonData is the method ID for the method DeleteCommonData
	MethodDeleteCommonData = 0x4

	// MethodGetRanking is the method ID for the method GetRanking
	MethodGetRanking = 0x5

	// MethodGetRankingByPrincipalID is the method ID for the method GetRankingByPrincipalID
	MethodGetRankingByPrincipalID = 0x6

	// MethodGetCategorySetting is the method ID for the method GetCategorySetting
	MethodGetCategorySetting = 0x7

	// MethodGetRankingChart is the method ID for the method GetRankingChart
	MethodGetRankingChart = 0x8

	// MethodGetRankingCharts is the method ID for the method GetRankingCharts
	MethodGetRankingCharts = 0x9

	// MethodGetEstimateScoreRank is the method ID for the method GetEstimateScoreRank
	MethodGetEstimateScoreRank = 0xA
)

// Protocol stores all the RMC method handlers for the Ranking2 protocol and listens for requests
type Protocol struct {
	Server                         *nex.Server
	putScoreHandler                func(err error, client *nex.Client, callID uint32, scoreDataList []*ranking2_types.Ranking2ScoreData, nexUniqueID uint64) uint32
	getCommonDataHandler           func(err error, client *nex.Client, callID uint32, optionFlags uint32, principalID uint32, nexUniqueID uint64) uint32
	putCommonDataHandler           func(err error, client *nex.Client, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID uint64) uint32
	deleteCommonDataHandler        func(err error, client *nex.Client, callID uint32, nexUniqueID uint64) uint32
	getRankingHandler              func(err error, client *nex.Client, callID uint32, getParam *ranking2_types.Ranking2GetParam) uint32
	getRankingByPrincipalIDHandler func(err error, client *nex.Client, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList []uint32) uint32
	getCategorySettingHandler      func(err error, client *nex.Client, callID uint32, category uint32) uint32
	getRankingChartHandler         func(err error, client *nex.Client, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) uint32
	getRankingChartsHandler        func(err error, client *nex.Client, callID uint32, infoArray []*ranking2_types.Ranking2ChartInfoInput) uint32
	getEstimateScoreRankHandler    func(err error, client *nex.Client, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodPutScore:
				go protocol.handlePutScore(packet)
			case MethodGetCommonData:
				go protocol.handleGetCommonData(packet)
			case MethodPutCommonData:
				go protocol.handlePutCommonData(packet)
			case MethodDeleteCommonData:
				go protocol.handleDeleteCommonData(packet)
			case MethodGetRanking:
				go protocol.handleGetRanking(packet)
			case MethodGetRankingByPrincipalID:
				go protocol.handleGetRankingByPrincipalID(packet)
			case MethodGetCategorySetting:
				go protocol.handleGetCategorySetting(packet)
			case MethodGetRankingChart:
				go protocol.handleGetRankingChart(packet)
			case MethodGetRankingCharts:
				go protocol.handleGetRankingCharts(packet)
			case MethodGetEstimateScoreRank:
				go protocol.handleGetEstimateScoreRank(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Ranking2 method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Ranking2 protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
