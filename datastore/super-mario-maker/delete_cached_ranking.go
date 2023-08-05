// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCachedRanking sets the DeleteCachedRanking handler function
func (protocol *Protocol) DeleteCachedRanking(handler func(err error, client *nex.Client, callID uint32, rankingType string, rankingArgs []string) uint32) {
	protocol.deleteCachedRankingHandler = handler
}

func (protocol *Protocol) handleDeleteCachedRanking(packet nex.PacketInterface) {
	if protocol.deleteCachedRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteCachedRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	rankingType, err := parametersStream.ReadString()
	if err != nil {
		go protocol.deleteCachedRankingHandler(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	rankingArgs, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.deleteCachedRankingHandler(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	go protocol.deleteCachedRankingHandler(nil, client, callID, rankingType, rankingArgs)
}
