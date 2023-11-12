// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetDeletionReason sets the SetDeletionReason handler function
func (protocol *Protocol) SetDeletionReason(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDLst []uint64, deletionReason uint32) uint32) {
	protocol.setDeletionReasonHandler = handler
}

func (protocol *Protocol) handleSetDeletionReason(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.setDeletionReasonHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetDeletionReason not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.setDeletionReasonHandler(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	deletionReason, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.setDeletionReasonHandler(fmt.Errorf("Failed to read deletionReason from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.setDeletionReasonHandler(nil, packet, callID, dataIDLst, deletionReason)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
