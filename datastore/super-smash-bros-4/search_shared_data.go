package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchSharedData sets the SearchSharedData handler function
func (protocol *DataStoreSuperSmashBros4Protocol) SearchSharedData(handler func(err error, client *nex.Client, callID uint32, param *DataStoreSearchSharedDataParam)) {
	protocol.SearchSharedDataHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleSearchSharedData(packet nex.PacketInterface) {
	if protocol.SearchSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::SearchSharedData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreSearchSharedDataParam())
	if err != nil {
		go protocol.SearchSharedDataHandler(err, client, callID, nil)
		return
	}

	go protocol.SearchSharedDataHandler(nil, client, callID, param.(*DataStoreSearchSharedDataParam))
}
