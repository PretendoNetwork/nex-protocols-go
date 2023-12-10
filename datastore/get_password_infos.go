// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPasswordInfos(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetPasswordInfos == nil {
		globals.Logger.Warning("DataStore::GetPasswordInfos not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetPasswordInfos(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPasswordInfos(nil, packet, callID, dataIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
