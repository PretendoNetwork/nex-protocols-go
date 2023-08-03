// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ConditionalSearchObject sets the ConditionalSearchObject handler function
func (protocol *Protocol) ConditionalSearchObject(handler func(err error, client *nex.Client, callID uint32, condition uint32, param *datastore_types.DataStoreSearchParam, extraData []string)) {
	protocol.conditionalSearchObjectHandler = handler
}

func (protocol *Protocol) handleConditionalSearchObject(packet nex.PacketInterface) {
	if protocol.conditionalSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ConditionalSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	condition, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.conditionalSearchObjectHandler(fmt.Errorf("Failed to read condition from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		go protocol.conditionalSearchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.conditionalSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		return
	}

	go protocol.conditionalSearchObjectHandler(nil, client, callID, condition, param.(*datastore_types.DataStoreSearchParam), extraData)
}
