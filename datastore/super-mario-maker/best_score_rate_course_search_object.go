// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// BestScoreRateCourseSearchObject sets the BestScoreRateCourseSearchObject handler function
func (protocol *Protocol) BestScoreRateCourseSearchObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string) uint32) {
	protocol.bestScoreRateCourseSearchObjectHandler = handler
}

func (protocol *Protocol) handleBestScoreRateCourseSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.bestScoreRateCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::BestScoreRateCourseSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		errorCode = protocol.bestScoreRateCourseSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		errorCode = protocol.bestScoreRateCourseSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.bestScoreRateCourseSearchObjectHandler(nil, packet, callID, param.(*datastore_types.DataStoreSearchParam), extraData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
