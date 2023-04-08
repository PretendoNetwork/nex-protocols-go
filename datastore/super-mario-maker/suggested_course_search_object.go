package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SuggestedCourseSearchObject sets the SuggestedCourseSearchObject handler function
func (protocol *DataStoreSuperMarioMakerProtocol) SuggestedCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore.DataStoreSearchParam, extraData []string)) {
	protocol.SuggestedCourseSearchObjectHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleSuggestedCourseSearchObject(packet nex.PacketInterface) {
	if protocol.SuggestedCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSMM::SuggestedCourseSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore.NewDataStoreSearchParam())
	if err != nil {
		go protocol.SuggestedCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData := parametersStream.ReadListString()

	go protocol.SuggestedCourseSearchObjectHandler(nil, client, callID, param.(*datastore.DataStoreSearchParam), extraData)
}
