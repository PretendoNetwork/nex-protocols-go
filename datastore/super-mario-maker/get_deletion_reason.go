// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetDeletionReason sets the GetDeletionReason handler function
func (protocol *Protocol) GetDeletionReason(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDLst []uint64) uint32) {
	protocol.getDeletionReasonHandler = handler
}

func (protocol *Protocol) handleGetDeletionReason(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getDeletionReasonHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetDeletionReason not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.getDeletionReasonHandler(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getDeletionReasonHandler(nil, packet, callID, dataIDLst)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
