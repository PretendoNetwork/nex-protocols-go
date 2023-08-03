// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetCachedRanking sets the SetCachedRanking handler function
func (protocol *Protocol) SetCachedRanking(handler func(err error, client *nex.Client, callID uint32, rankingType string, rankingArgs []string, dataIDLst []uint64)) {
	protocol.setCachedRankingHandler = handler
}

func (protocol *Protocol) handleSetCachedRanking(packet nex.PacketInterface) {
	if protocol.setCachedRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetCachedRanking not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	rankingType, err := parametersStream.ReadString()
	if err != nil {
		go protocol.setCachedRankingHandler(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), client, callID, "", nil, nil)
		return
	}

	rankingArgs, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.setCachedRankingHandler(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), client, callID, "", nil, nil)
		return
	}

	dataIDLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.setCachedRankingHandler(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), client, callID, "", nil, nil)
		return
	}

	go protocol.setCachedRankingHandler(nil, client, callID, rankingType, rankingArgs, dataIDLst)
}
