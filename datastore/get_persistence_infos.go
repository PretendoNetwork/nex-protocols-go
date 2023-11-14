// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPersistenceInfos sets the GetPersistenceInfos handler function
func (protocol *Protocol) GetPersistenceInfos(handler func(err error, packet nex.PacketInterface, callID uint32, ownerID *nex.PID, persistenceSlotIDs []uint16) uint32) {
	protocol.getPersistenceInfosHandler = handler
}

func (protocol *Protocol) handleGetPersistenceInfos(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPersistenceInfosHandler == nil {
		globals.Logger.Warning("DataStore::GetPersistenceInfos not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ownerID, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.getPersistenceInfosHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	persistenceSlotIDs, err := parametersStream.ReadListUInt16LE()
	if err != nil {
		errorCode = protocol.getPersistenceInfosHandler(fmt.Errorf("Failed to read persistenceSlotIDs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPersistenceInfosHandler(nil, packet, callID, ownerID, persistenceSlotIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
