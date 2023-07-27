// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPersistenceInfo sets the GetPersistenceInfo handler function
func (protocol *Protocol) GetPersistenceInfo(handler func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16)) {
	protocol.getPersistenceInfoHandler = handler
}

func (protocol *Protocol) handleGetPersistenceInfo(packet nex.PacketInterface) {
	if protocol.getPersistenceInfoHandler == nil {
		globals.Logger.Warning("DataStorePotocol::GetPersistenceInfo not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ownerID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getPersistenceInfoHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.getPersistenceInfoHandler(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.getPersistenceInfoHandler(nil, client, callID, ownerID, persistenceSlotID)
}
