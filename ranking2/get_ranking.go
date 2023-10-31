// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// GetRanking sets the GetRanking handler function
func (protocol *Protocol) GetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) uint32) {
	protocol.getRankingHandler = handler
}

func (protocol *Protocol) handleGetRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRankingHandler == nil {
		globals.Logger.Warning("Ranking2::GetRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getParam, err := parametersStream.ReadStructure(ranking2_types.NewRanking2GetParam())
	if err != nil {
		errorCode = protocol.getRankingHandler(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRankingHandler(nil, packet, callID, getParam.(*ranking2_types.Ranking2GetParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
