package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RecommendedCourseSearchObject sets the RecommendedCourseSearchObject handler function
func (protocol *DataStoreSuperMarioMakerProtocol) RecommendedCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore.DataStoreSearchParam, extraData []string)) {
	protocol.RecommendedCourseSearchObjectHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleRecommendedCourseSearchObject(packet nex.PacketInterface) {
	if protocol.RecommendedCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSMM::RecommendedCourseSearchObject not implemented")
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
		go protocol.RecommendedCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData := parametersStream.ReadListString()

	go protocol.RecommendedCourseSearchObjectHandler(nil, client, callID, param.(*datastore.DataStoreSearchParam), extraData)
}
