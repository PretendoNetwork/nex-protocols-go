// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetDeletionReason(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SetDeletionReason == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetDeletionReason not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	dataIDLst := types.NewList[*types.PrimitiveU64]()
	dataIDLst.Type = types.NewPrimitiveU64(0)
	err = dataIDLst.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetDeletionReason(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	deletionReason := types.NewPrimitiveU32(0)
	err = deletionReason.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetDeletionReason(fmt.Errorf("Failed to read deletionReason from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SetDeletionReason(nil, packet, callID, dataIDLst, deletionReason)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
