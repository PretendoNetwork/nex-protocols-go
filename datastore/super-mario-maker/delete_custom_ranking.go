// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCustomRanking sets the DeleteCustomRanking handler function
func (protocol *Protocol) DeleteCustomRanking(handler func(err error, client *nex.Client, callID uint32, dataIDList []uint64) uint32) {
	protocol.deleteCustomRankingHandler = handler
}

func (protocol *Protocol) handleDeleteCustomRanking(packet nex.PacketInterface) {
	if protocol.deleteCustomRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteCustomRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.deleteCustomRankingHandler(fmt.Errorf("Failed to read dataIDList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.deleteCustomRankingHandler(nil, client, callID, dataIDList)
}
