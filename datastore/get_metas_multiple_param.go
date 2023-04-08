package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetasMultipleParam sets the GetMetasMultipleParam handler function
func (protocol *DataStoreProtocol) GetMetasMultipleParam(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParams []*DataStoreGetMetaParam)) {
	protocol.GetMetasMultipleParamHandler = handler
}

func (protocol *DataStoreProtocol) HandleGetMetasMultipleParam(packet nex.PacketInterface) {
	if protocol.GetMetasMultipleParamHandler == nil {
		globals.Logger.Warning("DataStore::GetMetasMultipleParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.GetMetasMultipleParamHandler(err, client, callID, nil)
		return
	}

	go protocol.GetMetasMultipleParamHandler(nil, client, callID, params.([]*DataStoreGetMetaParam))
}
