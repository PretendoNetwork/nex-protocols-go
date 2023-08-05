// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// PutScore sets the PutScore handler function
func (protocol *Protocol) PutScore(handler func(err error, client *nex.Client, callID uint32, scoreDataList []*ranking2_types.Ranking2ScoreData, nexUniqueID uint64) uint32) {
	protocol.putScoreHandler = handler
}

func (protocol *Protocol) handlePutScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.putScoreHandler == nil {
		globals.Logger.Warning("Ranking2::PutScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	scoreDataList, err := parametersStream.ReadListStructure(ranking2_types.NewRanking2ScoreData())
	if err != nil {
		errorCode = protocol.putScoreHandler(fmt.Errorf("Failed to read scoreDataList from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.putScoreHandler(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.putScoreHandler(nil, client, callID, scoreDataList.([]*ranking2_types.Ranking2ScoreData), nexUniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
