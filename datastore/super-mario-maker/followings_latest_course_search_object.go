package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FollowingsLatestCourseSearchObject sets the FollowingsLatestCourseSearchObject handler function
func (protocol *DataStoreSuperMarioMakerProtocol) FollowingsLatestCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore.DataStoreSearchParam, extraData []string)) {
	protocol.FollowingsLatestCourseSearchObjectHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleFollowingsLatestCourseSearchObject(packet nex.PacketInterface) {
	if protocol.FollowingsLatestCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSMM::FollowingsLatestCourseSearchObject not implemented")
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
		go protocol.FollowingsLatestCourseSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.FollowingsLatestCourseSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.FollowingsLatestCourseSearchObjectHandler(nil, client, callID, param.(*datastore.DataStoreSearchParam), extraData)
}
