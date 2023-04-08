package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetas sets the GetMetas handler function
func (protocol *DataStoreProtocol) GetMetas(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64, param *DataStoreGetMetaParam)) {
	protocol.GetMetasHandler = handler
}

func (protocol *DataStoreProtocol) HandleGetMetas(packet nex.PacketInterface) {
	if protocol.GetMetasHandler == nil {
		globals.Logger.Warning("DataStore::GetMetas not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs := parametersStream.ReadListUInt64LE()

	param, err := parametersStream.ReadStructure(NewDataStoreGetMetaParam())

	if err != nil {
		go protocol.GetMetasHandler(err, client, callID, nil, nil)
		return
	}

	go protocol.GetMetasHandler(nil, client, callID, dataIDs, param.(*DataStoreGetMetaParam))
}
