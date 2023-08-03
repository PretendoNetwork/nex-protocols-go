// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApplicationConfigString sets the GetApplicationConfigString handler function
func (protocol *Protocol) GetApplicationConfigString(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	protocol.getApplicationConfigStringHandler = handler
}

func (protocol *Protocol) handleGetApplicationConfigString(packet nex.PacketInterface) {
	if protocol.getApplicationConfigStringHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetApplicationConfigString not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getApplicationConfigStringHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getApplicationConfigStringHandler(nil, client, callID, applicationID)
}
