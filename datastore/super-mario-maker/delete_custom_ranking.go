// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCustomRanking sets the DeleteCustomRanking handler function
func (protocol *Protocol) DeleteCustomRanking(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDList []uint64) uint32) {
	protocol.deleteCustomRankingHandler = handler
}

func (protocol *Protocol) handleDeleteCustomRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteCustomRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteCustomRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.deleteCustomRankingHandler(fmt.Errorf("Failed to read dataIDList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteCustomRankingHandler(nil, packet, callID, dataIDList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
