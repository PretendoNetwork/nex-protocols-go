// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCachedRanking sets the DeleteCachedRanking handler function
func (protocol *Protocol) DeleteCachedRanking(handler func(err error, packet nex.PacketInterface, callID uint32, rankingType string, rankingArgs []string) uint32) {
	protocol.deleteCachedRankingHandler = handler
}

func (protocol *Protocol) handleDeleteCachedRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteCachedRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteCachedRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	rankingType, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.deleteCachedRankingHandler(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingArgs, err := parametersStream.ReadListString()
	if err != nil {
		errorCode = protocol.deleteCachedRankingHandler(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteCachedRankingHandler(nil, packet, callID, rankingType, rankingArgs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
