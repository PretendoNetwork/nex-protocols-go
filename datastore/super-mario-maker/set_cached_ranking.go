// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetCachedRanking(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SetCachedRanking == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetCachedRanking not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	rankingType := types.NewString("")
	err = rankingType.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetCachedRanking(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingArgs := types.NewList[*types.String]()
	rankingArgs.Type = types.NewString("")
	err = rankingArgs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetCachedRanking(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataIDLst := types.NewList[*types.PrimitiveU64]()
	dataIDLst.Type = types.NewPrimitiveU64(0)
	err = dataIDLst.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetCachedRanking(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SetCachedRanking(nil, packet, callID, rankingType, rankingArgs, dataIDLst)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
