// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
)

// UpdateSimpleSearchObject sets the UpdateSimpleSearchObject handler function
func (protocol *Protocol) UpdateSimpleSearchObject(handler func(err error, client *nex.Client, callID uint32, objectID uint32, newObject *matchmake_extension_mario_kart8_types.SimpleSearchObject)) {
	protocol.updateSimpleSearchObjectHandler = handler
}

func (protocol *Protocol) handleUpdateSimpleSearchObject(packet nex.PacketInterface) {
	if protocol.updateSimpleSearchObjectHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::UpdateSimpleSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	objectID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateSimpleSearchObjectHandler(fmt.Errorf("Failed to read objectID from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	newObject, err := parametersStream.ReadStructure(matchmake_extension_mario_kart8_types.NewSimpleSearchObject())
	if err != nil {
		go protocol.updateSimpleSearchObjectHandler(fmt.Errorf("Failed to read newObject from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.updateSimpleSearchObjectHandler(nil, client, callID, objectID, newObject.(*matchmake_extension_mario_kart8_types.SimpleSearchObject))
}
