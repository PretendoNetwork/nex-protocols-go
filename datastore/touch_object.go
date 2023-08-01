// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TouchObject sets the TouchObject handler function
func (protocol *Protocol) TouchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreTouchObjectParam)) {
	protocol.touchObjectHandler = handler
}

func (protocol *Protocol) handleTouchObject(packet nex.PacketInterface) {
	if protocol.touchObjectHandler == nil {
		globals.Logger.Warning("DataStore::TouchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreTouchObjectParam())
	if err != nil {
		go protocol.touchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.touchObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreTouchObjectParam))
}