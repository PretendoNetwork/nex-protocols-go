// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// GetRankingByPrincipalID sets the GetRankingByPrincipalID handler function
func (protocol *Protocol) GetRankingByPrincipalID(handler func(err error, client *nex.Client, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList []uint32) uint32) {
	protocol.getRankingByPrincipalIDHandler = handler
}

func (protocol *Protocol) handleGetRankingByPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRankingByPrincipalIDHandler == nil {
		globals.Logger.Warning("Ranking2::GetRankingByPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getParam, err := parametersStream.ReadStructure(ranking2_types.NewRanking2GetParam())
	if err != nil {
		errorCode = protocol.getRankingByPrincipalIDHandler(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	principalIDList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getRankingByPrincipalIDHandler(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRankingByPrincipalIDHandler(nil, client, callID, getParam.(*ranking2_types.Ranking2GetParam), principalIDList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
