// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetCachedRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SetCachedRanking == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetCachedRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	rankingType, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.SetCachedRanking(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingArgs, err := parametersStream.ReadListString()
	if err != nil {
		errorCode = protocol.SetCachedRanking(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataIDLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.SetCachedRanking(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.SetCachedRanking(nil, packet, callID, rankingType, rankingArgs, dataIDLst)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
