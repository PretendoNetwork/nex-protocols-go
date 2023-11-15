// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCachedRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeleteCachedRanking == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteCachedRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	rankingType, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.DeleteCachedRanking(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingArgs, err := parametersStream.ReadListString()
	if err != nil {
		_, errorCode = protocol.DeleteCachedRanking(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteCachedRanking(nil, packet, callID, rankingType, rankingArgs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
