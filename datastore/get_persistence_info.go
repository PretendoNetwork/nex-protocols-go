// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPersistenceInfo sets the GetPersistenceInfo handler function
func (protocol *Protocol) GetPersistenceInfo(handler func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16) uint32) {
	protocol.getPersistenceInfoHandler = handler
}

func (protocol *Protocol) handleGetPersistenceInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPersistenceInfoHandler == nil {
		globals.Logger.Warning("DataStore::GetPersistenceInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ownerID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getPersistenceInfoHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), client, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.getPersistenceInfoHandler(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), client, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPersistenceInfoHandler(nil, client, callID, ownerID, persistenceSlotID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
