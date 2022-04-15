package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// RankingProtocolID is the protocol ID for the Ranking protocol
	RankingProtocolID = 0x70

	// RankingMethodUploadScore is the method ID for method UploadScore
	RankingMethodUploadScore = 0x1

	// RankingMethodDeleteScore is the method ID for method DeleteScore
	RankingMethodDeleteScore = 0x2

	// RankingMethodDeleteAllScores is the method ID for method DeleteAllScores
	RankingMethodDeleteAllScores = 0x3

	// RankingMethodUploadCommonData is the method ID for method UploadCommonData
	RankingMethodUploadCommonData = 0x4

	// RankingMethodDeleteCommonData is the method ID for method DeleteCommonData
	RankingMethodDeleteCommonData = 0x5

	// RankingMethodGetCommonData is the method ID for method GetCommonData
	RankingMethodGetCommonData = 0x6

	// RankingMethodChangeAttributes is the method ID for method ChangeAttributes
	RankingMethodChangeAttributes = 0x7

	// RankingMethodChangeAllAttributes is the method ID for method ChangeAllAttributes
	RankingMethodChangeAllAttributes = 0x8

	// RankingMethodGetRanking is the method ID for method GetRanking
	RankingMethodGetRanking = 0x9

	// RankingMethodGetApproxOrder is the method ID for method GetApproxOrder
	RankingMethodGetApproxOrder = 0xA

	// RankingMethodGetStats is the method ID for method GetStats
	RankingMethodGetStats = 0xB

	// RankingMethodGetRankingByPIDList is the method ID for method GetRankingByPIDList
	RankingMethodGetRankingByPIDList = 0xC

	// RankingMethodGetRankingByUniqueIDList is the method ID for method GetRankingByUniqueIdList
	RankingMethodGetRankingByUniqueIDList = 0xD

	// RankingMethodGetCachedTopXRanking is the method ID for method GetCachedTopXRanking
	RankingMethodGetCachedTopXRanking = 0xE

	// RankingMethodGetCachedTopXRankings is the method ID for method GetCachedTopXRankings
	RankingMethodGetCachedTopXRankings = 0xF
)

// RankingProtocol handles the Ranking nex protocol
type RankingProtocol struct {
	server                  *nex.Server
	UploadCommonDataHandler func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueId uint64)
}

// Setup initializes the protocol
func (rankingProtocol *RankingProtocol) Setup() {
	nexServer := rankingProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if RankingProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case RankingMethodUploadCommonData:
				go rankingProtocol.handleUploadCommonData(packet)
			default:
				go respondNotImplemented(packet, RankingProtocolID)
				fmt.Printf("Unsupported Ranking method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// UploadCommonData sets the UploadCommonData handler function
func (rankingProtocol *RankingProtocol) UploadCommonData(handler func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueID uint64)) {
	rankingProtocol.UploadCommonDataHandler = handler
}

func (rankingProtocol *RankingProtocol) handleUploadCommonData(packet nex.PacketInterface) {
	if rankingProtocol.UploadCommonDataHandler == nil {
		logger.Warning("RankingProtocol::UploadCommonData not implemented")
		go respondNotImplemented(packet, RankingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, rankingProtocol.server)

	commonData, err := parametersStream.ReadBuffer()
	if err != nil {
		go rankingProtocol.UploadCommonDataHandler(err, client, callID, nil, 0)
		return
	}

	uniqueID := parametersStream.ReadUInt64LE()

	go rankingProtocol.UploadCommonDataHandler(nil, client, callID, commonData, uniqueID)
}

// NewRankingProtocol returns a new RankingProtocol
func NewRankingProtocol(server *nex.Server) *RankingProtocol {
	rankingProtocol := &RankingProtocol{server: server}

	rankingProtocol.Setup()

	return rankingProtocol
}
