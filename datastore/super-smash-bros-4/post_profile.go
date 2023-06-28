package datastore_super_smash_bros_4

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostProfile sets the PostProfile handler function
func (protocol *DataStoreSuperSmashBros4Protocol) PostProfile(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePostProfileParam)) {
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

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStorePostProfileParam())
	if err != nil {
		go protocol.PostProfileHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PostProfileHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStorePostProfileParam))
}
