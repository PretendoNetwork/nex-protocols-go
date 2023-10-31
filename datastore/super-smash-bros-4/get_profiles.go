// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetProfiles sets the GetProfiles handler function
func (protocol *Protocol) GetProfiles(handler func(err error, packet nex.PacketInterface, callID uint32, pidList []uint32) uint32) {
	protocol.getProfilesHandler = handler
}

func (protocol *Protocol) handleGetProfiles(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getProfilesHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetProfiles not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pidList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getProfilesHandler(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getProfilesHandler(nil, packet, callID, pidList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
