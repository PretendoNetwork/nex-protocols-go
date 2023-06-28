package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// CreateMatchmakeSessionWithParam sets the CreateMatchmakeSessionWithParam handler function
func (protocol *MatchmakeExtensionProtocol) CreateMatchmakeSessionWithParam(handler func(err error, client *nex.Client, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam)) {
	protocol.CreateMatchmakeSessionWithParamHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleCreateMatchmakeSessionWithParam(packet nex.PacketInterface) {
	if protocol.CreateMatchmakeSessionWithParamHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSessionWithParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	createMatchmakeSessionParam, err := parametersStream.ReadStructure(match_making_types.NewCreateMatchmakeSessionParam())
	if err != nil {
		go protocol.CreateMatchmakeSessionWithParamHandler(fmt.Errorf("Failed to read createMatchmakeSessionParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CreateMatchmakeSessionWithParamHandler(nil, client, callID, createMatchmakeSessionParam.(*match_making_types.CreateMatchmakeSessionParam))
}
