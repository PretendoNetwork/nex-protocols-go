// Package datastore_super_mario_maker implements the Super Mario Maker DataStore NEX protocol
package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApplicationConfigString sets the GetApplicationConfigString handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetApplicationConfigString(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	protocol.GetApplicationConfigStringHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) handleGetApplicationConfigString(packet nex.PacketInterface) {
	if protocol.GetApplicationConfigStringHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetApplicationConfigString not implemented")
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
		go protocol.GetApplicationConfigStringHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.GetApplicationConfigStringHandler(nil, client, callID, applicationID)
}
