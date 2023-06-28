package utility

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

// AssociateNexUniqueIDWithMyPrincipalID sets the AssociateNexUniqueIDWithMyPrincipalID handler function
func (protocol *UtilityProtocol) AssociateNexUniqueIDWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo *utility_types.UniqueIDInfo)) {
	protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler = handler
}

func (protocol *UtilityProtocol) HandleAssociateNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::AssociateNexUniqueIDWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueIDInfo, err := parametersStream.ReadStructure(utility_types.NewUniqueIDInfo())
	if err != nil {
		go protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo.(*utility_types.UniqueIDInfo))
}
