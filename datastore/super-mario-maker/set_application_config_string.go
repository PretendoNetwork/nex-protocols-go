// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetApplicationConfigString sets the SetApplicationConfigString handler function
func (protocol *Protocol) SetApplicationConfigString(handler func(err error, client *nex.Client, callID uint32, applicationID uint32, key uint32, value string)) {
	protocol.setApplicationConfigStringHandler = handler
}

func (protocol *Protocol) handleSetApplicationConfigString(packet nex.PacketInterface) {
	if protocol.setApplicationConfigStringHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetApplicationConfigString not implemented")
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
		go protocol.setApplicationConfigStringHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	key, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.setApplicationConfigStringHandler(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	value, err := parametersStream.ReadString()
	if err != nil {
		go protocol.setApplicationConfigStringHandler(fmt.Errorf("Failed to read value from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	go protocol.setApplicationConfigStringHandler(nil, client, callID, applicationID, key, value)
}
