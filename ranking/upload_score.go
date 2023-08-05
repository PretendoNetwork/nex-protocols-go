// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// UploadScore sets the UploadScore handler function
func (protocol *Protocol) UploadScore(handler func(err error, client *nex.Client, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID uint64) uint32) {
	protocol.uploadScoreHandler = handler
}

func (protocol *Protocol) handleUploadScore(packet nex.PacketInterface) {
	if protocol.uploadScoreHandler == nil {
		globals.Logger.Warning("Ranking::UploadScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	scoreData, err := parametersStream.ReadStructure(ranking_types.NewRankingScoreData())
	if err != nil {
		go protocol.uploadScoreHandler(fmt.Errorf("Failed to read scoreData from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.uploadScoreHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.uploadScoreHandler(nil, client, callID, scoreData.(*ranking_types.RankingScoreData), uniqueID)
}
