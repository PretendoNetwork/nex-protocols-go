package datastore

import (
	"errors"

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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 6 {
		err := errors.New("[DataStore::GetPersistenceInfo] Data length too small")
		go protocol.GetPersistenceInfoHandler(err, client, callID, 0, 0)
		return
	}

	ownerID := parametersStream.ReadUInt32LE()
	persistenceSlotID := parametersStream.ReadUInt16LE()

	go protocol.GetPersistenceInfoHandler(nil, client, callID, ownerID, persistenceSlotID)
}
