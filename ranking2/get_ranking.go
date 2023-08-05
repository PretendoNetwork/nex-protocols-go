// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// GetRanking sets the GetRanking handler function
func (protocol *Protocol) GetRanking(handler func(err error, client *nex.Client, callID uint32, getParam *ranking2_types.Ranking2GetParam) uint32) {
	protocol.getRankingHandler = handler
}

func (protocol *Protocol) handleGetRanking(packet nex.PacketInterface) {
	if protocol.getRankingHandler == nil {
		globals.Logger.Warning("Ranking2::GetRanking not implemented")
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
		go protocol.getRankingHandler(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getRankingHandler(nil, client, callID, getParam.(*ranking2_types.Ranking2GetParam))
}
