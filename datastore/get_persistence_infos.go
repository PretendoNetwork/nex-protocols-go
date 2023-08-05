// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPersistenceInfos sets the GetPersistenceInfos handler function
func (protocol *Protocol) GetPersistenceInfos(handler func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotIDs []uint16) uint32) {
	protocol.getPersistenceInfosHandler = handler
}

func (protocol *Protocol) handleGetPersistenceInfos(packet nex.PacketInterface) {
	if protocol.getPersistenceInfosHandler == nil {
		globals.Logger.Warning("DataStore::GetPersistenceInfos not implemented")
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
		go protocol.getPersistenceInfosHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	persistenceSlotIDs, err := parametersStream.ReadListUInt16LE()
	if err != nil {
		go protocol.getPersistenceInfosHandler(fmt.Errorf("Failed to read persistenceSlotIDs from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.getPersistenceInfosHandler(nil, client, callID, ownerID, persistenceSlotIDs)
}
