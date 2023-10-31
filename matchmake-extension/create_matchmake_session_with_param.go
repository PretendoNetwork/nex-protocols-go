// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// CreateMatchmakeSessionWithParam sets the CreateMatchmakeSessionWithParam handler function
func (protocol *Protocol) CreateMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) uint32) {
	protocol.createMatchmakeSessionWithParamHandler = handler
}

func (protocol *Protocol) handleCreateMatchmakeSessionWithParam(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.createMatchmakeSessionWithParamHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSessionWithParam not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	createMatchmakeSessionParam, err := parametersStream.ReadStructure(match_making_types.NewCreateMatchmakeSessionParam())
	if err != nil {
		errorCode = protocol.createMatchmakeSessionWithParamHandler(fmt.Errorf("Failed to read createMatchmakeSessionParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.createMatchmakeSessionWithParamHandler(nil, packet, callID, createMatchmakeSessionParam.(*match_making_types.CreateMatchmakeSessionParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
