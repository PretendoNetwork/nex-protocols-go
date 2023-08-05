// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FollowingsLatestCourseSearchObject sets the FollowingsLatestCourseSearchObject handler function
func (protocol *Protocol) FollowingsLatestCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) uint32) {
	protocol.followingsLatestCourseSearchObjectHandler = handler
}

func (protocol *Protocol) handleFollowingsLatestCourseSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.followingsLatestCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::FollowingsLatestCourseSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		errorCode = protocol.followingsLatestCourseSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		errorCode = protocol.followingsLatestCourseSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.followingsLatestCourseSearchObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreSearchParam), extraData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
