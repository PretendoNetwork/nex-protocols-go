// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetCachedRanking sets the SetCachedRanking handler function
func (protocol *Protocol) SetCachedRanking(handler func(err error, packet nex.PacketInterface, callID uint32, rankingType string, rankingArgs []string, dataIDLst []uint64) uint32) {
	protocol.setCachedRankingHandler = handler
}

func (protocol *Protocol) handleSetCachedRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.setCachedRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetCachedRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	rankingType, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.setCachedRankingHandler(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingArgs, err := parametersStream.ReadListString()
	if err != nil {
		errorCode = protocol.setCachedRankingHandler(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataIDLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.setCachedRankingHandler(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.setCachedRankingHandler(nil, packet, callID, rankingType, rankingArgs, dataIDLst)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
