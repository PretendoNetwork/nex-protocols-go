// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// ChangeAttributes sets the ChangeAttributes handler function
func (protocol *Protocol) ChangeAttributes(handler func(err error, client *nex.Client, callID uint32, category uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID uint64)) {
	protocol.changeAttributesHandler = handler
}

func (protocol *Protocol) handleChangeAttributes(packet nex.PacketInterface) {
	if protocol.changeAttributesHandler == nil {
		globals.Logger.Warning("Ranking::ChangeAttributes not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.changeAttributesHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), client, callID, 0, nil, 0)
		return
	}

	changeParam, err := parametersStream.ReadStructure(ranking_types.NewRankingChangeAttributesParam())
	if err != nil {
		go protocol.changeAttributesHandler(fmt.Errorf("Failed to read changeParam from parameters. %s", err.Error()), client, callID, 0, nil, 0)
		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.changeAttributesHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, 0, nil, 0)
		return
	}

	go protocol.changeAttributesHandler(nil, client, callID, category, changeParam.(*ranking_types.RankingChangeAttributesParam), uniqueID)
}
