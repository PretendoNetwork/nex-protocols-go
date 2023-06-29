package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPlayingSession sets the GetPlayingSession handler function
func (protocol *MatchmakeExtensionProtocol) GetPlayingSession(handler func(err error, client *nex.Client, callID uint32, lstPID []uint32)) {
	protocol.GetPlayingSessionHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleGetPlayingSession(packet nex.PacketInterface) {
	if protocol.GetPlayingSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetPlayingSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.GetPlayingSessionHandler(fmt.Errorf("Failed to read lstPID from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetPlayingSessionHandler(nil, client, callID, lstPID)
}
