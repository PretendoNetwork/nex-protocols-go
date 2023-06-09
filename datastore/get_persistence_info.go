package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPersistenceInfo sets the GetPersistenceInfo handler function
func (protocol *DataStoreProtocol) GetPersistenceInfo(handler func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16)) {
	protocol.GetPersistenceInfoHandler = handler
}

func (protocol *DataStoreProtocol) HandleGetPersistenceInfo(packet nex.PacketInterface) {
	if protocol.GetPersistenceInfoHandler == nil {
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
		go protocol.GetPersistenceInfoHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.GetPersistenceInfoHandler(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.GetPersistenceInfoHandler(nil, client, callID, ownerID, persistenceSlotID)
}
