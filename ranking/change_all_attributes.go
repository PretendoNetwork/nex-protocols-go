// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// ChangeAllAttributes sets the ChangeAllAttributes handler function
func (protocol *Protocol) ChangeAllAttributes(handler func(err error, client *nex.Client, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID uint64) uint32) {
	protocol.changeAllAttributesHandler = handler
}

func (protocol *Protocol) handleChangeAllAttributes(packet nex.PacketInterface) {
	if protocol.changeAllAttributesHandler == nil {
		globals.Logger.Warning("Ranking::ChangeAllAttributes not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	changeParam, err := parametersStream.ReadStructure(ranking_types.NewRankingChangeAttributesParam())
	if err != nil {
		go protocol.changeAllAttributesHandler(fmt.Errorf("Failed to read changeParam from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.changeAllAttributesHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.changeAllAttributesHandler(nil, client, callID, changeParam.(*ranking_types.RankingChangeAttributesParam), uniqueID)
}
