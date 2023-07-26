// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAPICalls sets the GetAPICalls handler function
func (protocol *Protocol) GetAPICalls(handler func(err error, client *nex.Client, callID uint32, pids []uint32, unknown *nex.DateTime, unknown2 *nex.DateTime)) {
	protocol.GetAPICallsHandler = handler
}

func (protocol *Protocol) handleGetAPICalls(packet nex.PacketInterface) {
	if protocol.GetAPICallsHandler == nil {
		globals.Logger.Warning("Debug::GetAPICalls not implemented")
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
		go protocol.GetAPICallsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	unknown, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.GetAPICallsHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	unknown2, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.GetAPICallsHandler(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	go protocol.GetAPICallsHandler(nil, client, callID, pids, unknown, unknown2)
}
