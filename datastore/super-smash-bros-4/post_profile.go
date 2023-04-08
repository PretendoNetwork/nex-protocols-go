package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostProfile sets the PostProfile handler function
func (protocol *DataStoreSuperSmashBros4Protocol) PostProfile(handler func(err error, client *nex.Client, callID uint32, param *DataStorePostProfileParam)) {
	protocol.PostProfileHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandlePostProfile(packet nex.PacketInterface) {
	if protocol.PostProfileHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::PostProfile not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStorePostProfileParam())
	if err != nil {
		go protocol.PostProfileHandler(err, client, callID, nil)
		return
	}

	go protocol.PostProfileHandler(nil, client, callID, param.(*DataStorePostProfileParam))
}
