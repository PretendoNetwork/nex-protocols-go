// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

const (
	// ProtocolID is the protocol ID for the Ranking protocol
	ProtocolID = 0x70

	// MethodUploadScore is the method ID for method UploadScore
	MethodUploadScore = 0x1

	// MethodDeleteScore is the method ID for method DeleteScore
	MethodDeleteScore = 0x2

	// MethodDeleteAllScores is the method ID for method DeleteAllScores
	MethodDeleteAllScores = 0x3

	// MethodUploadCommonData is the method ID for method UploadCommonData
	MethodUploadCommonData = 0x4

	// MethodDeleteCommonData is the method ID for method DeleteCommonData
	MethodDeleteCommonData = 0x5

	// MethodGetCommonData is the method ID for method GetCommonData
	MethodGetCommonData = 0x6

	// MethodChangeAttributes is the method ID for method ChangeAttributes
	MethodChangeAttributes = 0x7

	// MethodChangeAllAttributes is the method ID for method ChangeAllAttributes
	MethodChangeAllAttributes = 0x8

	// MethodGetRanking is the method ID for method GetRanking
	MethodGetRanking = 0x9

	// MethodGetApproxOrder is the method ID for method GetApproxOrder
	MethodGetApproxOrder = 0xA

	// MethodGetStats is the method ID for method GetStats
	MethodGetStats = 0xB

	// MethodGetRankingByPIDList is the method ID for method GetRankingByPIDList
	MethodGetRankingByPIDList = 0xC

	// MethodGetRankingByUniqueIDList is the method ID for method GetRankingByUniqueIDList
	MethodGetRankingByUniqueIDList = 0xD

	// MethodGetCachedTopXRanking is the method ID for method GetCachedTopXRanking
	MethodGetCachedTopXRanking = 0xE

	// MethodGetCachedTopXRankings is the method ID for method GetCachedTopXRankings
	MethodGetCachedTopXRankings = 0xF
)

// Protocol stores all the RMC method handlers for the Ranking protocol and listens for requests
type Protocol struct {
	Server                  *nex.Server
	UploadScoreHandler      func(err error, client *nex.Client, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID uint64)
	UploadCommonDataHandler func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueID uint64)
	GetRankingHandler       func(err error, client *nex.Client, callID uint32, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64, principalID uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodUploadScore:
				go protocol.handleUploadScore(packet)
			case MethodUploadCommonData:
				go protocol.handleUploadCommonData(packet)
			case MethodGetRanking:
				go protocol.handleGetRanking(packet)
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported Ranking method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Ranking protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
