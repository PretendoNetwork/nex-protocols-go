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
	Server                  nex.ServerInterface
	PutScore                func(err error, packet nex.PacketInterface, callID uint32, scoreDataList []*ranking2_types.Ranking2ScoreData, nexUniqueID uint64) (*nex.RMCMessage, uint32)
	GetCommonData           func(err error, packet nex.PacketInterface, callID uint32, optionFlags uint32, principalID *nex.PID, nexUniqueID uint64) (*nex.RMCMessage, uint32)
	PutCommonData           func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID uint64) (*nex.RMCMessage, uint32)
	DeleteCommonData        func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID uint64) (*nex.RMCMessage, uint32)
	GetRanking              func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) (*nex.RMCMessage, uint32)
	GetRankingByPrincipalID func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList []*nex.PID) (*nex.RMCMessage, uint32)
	GetCategorySetting      func(err error, packet nex.PacketInterface, callID uint32, category uint32) (*nex.RMCMessage, uint32)
	GetRankingChart         func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, uint32)
	GetRankingCharts        func(err error, packet nex.PacketInterface, callID uint32, infoArray []*ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, uint32)
	GetEstimateScoreRank    func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			switch request.MethodID {
			case MethodPutScore:
				protocol.handlePutScore(packet)
			case MethodGetCommonData:
				protocol.handleGetCommonData(packet)
			case MethodPutCommonData:
				protocol.handlePutCommonData(packet)
			case MethodDeleteCommonData:
				protocol.handleDeleteCommonData(packet)
			case MethodGetRanking:
				protocol.handleGetRanking(packet)
			case MethodGetRankingByPrincipalID:
				protocol.handleGetRankingByPrincipalID(packet)
			case MethodGetCategorySetting:
				protocol.handleGetCategorySetting(packet)
			case MethodGetRankingChart:
				protocol.handleGetRankingChart(packet)
			case MethodGetRankingCharts:
				protocol.handleGetRankingCharts(packet)
			case MethodGetEstimateScoreRank:
				protocol.handleGetEstimateScoreRank(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Ranking2 method ID: %#v\n", request.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Ranking2 protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
