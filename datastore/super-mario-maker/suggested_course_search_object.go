// Package protocol implements the Super Mario Maker DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SuggestedCourseSearchObject sets the SuggestedCourseSearchObject handler function
func (protocol *Protocol) SuggestedCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)) {
	protocol.suggestedCourseSearchObjectHandler = handler
}

func (protocol *Protocol) handleSuggestedCourseSearchObject(packet nex.PacketInterface) {
	if protocol.suggestedCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSMM::SuggestedCourseSearchObject not implemented")
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
		go protocol.suggestedCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.suggestedCourseSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.suggestedCourseSearchObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreSearchParam), extraData)
}
