package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// RankingProtocolID is the protocol ID for the Ranking protocol
	RankingProtocolID = 0x70
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
		request := packet.GetRMCRequest()

		if RankingProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			default:
				fmt.Printf("Unsupported Ranking method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

// UploadCommonData sets the UploadCommonData handler function
func (rankingProtocol *RankingProtocol) UploadCommonData(handler func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueId uint64)) {
	rankingProtocol.UploadCommonDataHandler = handler
}

// NewRankingProtocol returns a new RankingProtocol
func NewRankingProtocol(server *nex.Server) *RankingProtocol {
	RankingProtocol := &RankingProtocol{server: server}

	RankingProtocol.Setup()

	return RankingProtocol
}
