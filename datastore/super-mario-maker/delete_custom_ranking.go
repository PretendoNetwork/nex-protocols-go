// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCustomRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeleteCustomRanking == nil {
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
		errorCode = protocol.DeleteCustomRanking(fmt.Errorf("Failed to read dataIDList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.DeleteCustomRanking(nil, packet, callID, dataIDList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
