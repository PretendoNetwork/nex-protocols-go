// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetEnvironment sets the GetEnvironment handler function
func (protocol *Protocol) GetEnvironment(handler func(err error, client *nex.Client, callID uint32, uniqueID string, platform uint8) uint32) {
	protocol.getEnvironmentHandler = handler
}

func (protocol *Protocol) handleGetEnvironment(packet nex.PacketInterface) {
	if protocol.getEnvironmentHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetEnvironment not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueID, err := parametersStream.ReadString()
	if err != nil {
		go protocol.getEnvironmentHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, "", 0)
		return
	}

	platform, err := parametersStream.ReadUInt8()
	if err != nil {
		go protocol.getEnvironmentHandler(fmt.Errorf("Failed to read platform from parameters. %s", err.Error()), client, callID, "", 0)
		return
	}

	go protocol.getEnvironmentHandler(nil, client, callID, uniqueID, platform)
}
