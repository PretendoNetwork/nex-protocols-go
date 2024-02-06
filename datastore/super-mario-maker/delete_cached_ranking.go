// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCachedRanking(packet nex.PacketInterface) {
	var err error

	if protocol.DeleteCachedRanking == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::DeleteCachedRanking not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	rankingType := types.NewString("")
	err = rankingType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteCachedRanking(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rankingArgs := types.NewList[*types.String]()
	rankingArgs.Type = types.NewString("")
	err = rankingArgs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteCachedRanking(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteCachedRanking(nil, packet, callID, rankingType, rankingArgs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
