// Package ranking implements the Ranking NEX protocol
package ranking

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// UploadScore sets the UploadScore handler function
func (protocol *RankingProtocol) UploadScore(handler func(err error, client *nex.Client, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID uint64)) {
	protocol.UploadScoreHandler = handler
}

func (protocol *RankingProtocol) handleUploadScore(packet nex.PacketInterface) {
	if protocol.UploadScoreHandler == nil {
		globals.Logger.Warning("Ranking::UploadScore not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	scoreData, err := parametersStream.ReadStructure(ranking_types.NewRankingScoreData())
	if err != nil {
		go protocol.UploadScoreHandler(fmt.Errorf("Failed to read scoreData from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.UploadScoreHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.UploadScoreHandler(nil, client, callID, scoreData.(*ranking_types.RankingScoreData), uniqueID)
}
