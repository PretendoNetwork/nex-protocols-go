package debug

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApiCalls sets the GetApiCalls handler function
func (protocol *DebugProtocol) GetApiCalls(handler func(err error, client *nex.Client, callID uint32, pids []uint32, unknown *nex.DateTime, unknown2 *nex.DateTime)) {
	protocol.GetApiCallsHandler = handler
}

func (protocol *DebugProtocol) HandleGetApiCalls(packet nex.PacketInterface) {
	if protocol.GetApiCallsHandler == nil {
		globals.Logger.Warning("Debug::GetApiCalls not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.GetApiCallsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	unknown, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.GetApiCallsHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	unknown2, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.GetApiCallsHandler(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	go protocol.GetApiCallsHandler(nil, client, callID, pids, unknown, unknown2)
}
