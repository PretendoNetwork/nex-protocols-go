// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangePlayablePlatform sets the ChangePlayablePlatform handler function
func (protocol *Protocol) ChangePlayablePlatform(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam)) {
	protocol.changePlayablePlatformHandler = handler
}

func (protocol *Protocol) handleChangePlayablePlatform(packet nex.PacketInterface) {
	if protocol.changePlayablePlatformHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ChangePlayablePlatform not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewDataStoreChangePlayablePlatformParam())
	if err != nil {
		go protocol.changePlayablePlatformHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.changePlayablePlatformHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam))
}
