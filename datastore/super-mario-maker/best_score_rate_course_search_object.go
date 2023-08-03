// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// BestScoreRateCourseSearchObject sets the BestScoreRateCourseSearchObject handler function
func (protocol *Protocol) BestScoreRateCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)) {
	protocol.bestScoreRateCourseSearchObjectHandler = handler
}

func (protocol *Protocol) handleBestScoreRateCourseSearchObject(packet nex.PacketInterface) {
	if protocol.bestScoreRateCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::BestScoreRateCourseSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		go protocol.bestScoreRateCourseSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.bestScoreRateCourseSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.bestScoreRateCourseSearchObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreSearchParam), extraData)
}
