// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareAttachFile sets the PrepareAttachFile handler function
func (protocol *Protocol) PrepareAttachFile(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreAttachFileParam)) {
	protocol.prepareAttachFileHandler = handler
}

func (protocol *Protocol) handlePrepareAttachFile(packet nex.PacketInterface) {
	if protocol.prepareAttachFileHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::PrepareAttachFile not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreAttachFileParam())
	if err != nil {
		go protocol.prepareAttachFileHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.prepareAttachFileHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreAttachFileParam))
}
