// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetDeletionReason(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetDeletionReason == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetDeletionReason not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	dataIDLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetDeletionReason(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetDeletionReason(nil, packet, callID, dataIDLst)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
