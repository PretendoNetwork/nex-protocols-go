// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

// AssociateNexUniqueIDWithMyPrincipalID sets the AssociateNexUniqueIDWithMyPrincipalID handler function
func (protocol *Protocol) AssociateNexUniqueIDWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo *utility_types.UniqueIDInfo) uint32) {
	protocol.associateNexUniqueIDWithMyPrincipalIDHandler = handler
}

func (protocol *Protocol) handleAssociateNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.associateNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::AssociateNexUniqueIDWithMyPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueIDInfo, err := parametersStream.ReadStructure(utility_types.NewUniqueIDInfo())
	if err != nil {
		errorCode = protocol.associateNexUniqueIDWithMyPrincipalIDHandler(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.associateNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo.(*utility_types.UniqueIDInfo))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
