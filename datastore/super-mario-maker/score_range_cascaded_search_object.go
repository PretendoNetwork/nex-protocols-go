// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ScoreRangeCascadedSearchObject sets the ScoreRangeCascadedSearchObject handler function
func (protocol *Protocol) ScoreRangeCascadedSearchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)) {
	protocol.scoreRangeCascadedSearchObjectHandler = handler
}

func (protocol *Protocol) handleScoreRangeCascadedSearchObject(packet nex.PacketInterface) {
	if protocol.scoreRangeCascadedSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ScoreRangeCascadedSearchObject not implemented")
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
		go protocol.scoreRangeCascadedSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.scoreRangeCascadedSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.scoreRangeCascadedSearchObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreSearchParam), extraData)
}
