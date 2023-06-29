package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AutoMatchmake_Postpone sets the AutoMatchmake_Postpone handler function
func (protocol *MatchmakeExtensionProtocol) AutoMatchmake_Postpone(handler func(err error, client *nex.Client, callID uint32, anyGathering *nex.DataHolder, message string)) {
	protocol.AutoMatchmake_PostponeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleAutoMatchmake_Postpone(packet nex.PacketInterface) {
	if protocol.AutoMatchmake_PostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmake_Postpone not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.AutoMatchmake_PostponeHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	message, err := parametersStream.ReadString()
	if err != nil {
		go protocol.AutoMatchmake_PostponeHandler(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), client, callID, nil, "")
		return
	}

	go protocol.AutoMatchmake_PostponeHandler(nil, client, callID, anyGathering, message)
}
