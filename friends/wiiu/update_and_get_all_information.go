package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAndGetAllInformation sets the UpdateAndGetAllInformation handler function
func (protocol *FriendsWiiUProtocol) UpdateAndGetAllInformation(handler func(err error, client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)) {
	protocol.UpdateAndGetAllInformationHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	if protocol.UpdateAndGetAllInformationHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateAndGetAllInformation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nnaInfo, err := parametersStream.ReadStructure(NewNNAInfo())
	if err != nil {
		go protocol.UpdateAndGetAllInformationHandler(fmt.Errorf("Failed to read nnaInfo from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	presence, err := parametersStream.ReadStructure(NewNintendoPresenceV2())
	if err != nil {
		go protocol.UpdateAndGetAllInformationHandler(fmt.Errorf("Failed to read presence from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	birthday, err := parametersStream.ReadDateTime()
	if err != nil {
		go protocol.UpdateAndGetAllInformationHandler(fmt.Errorf("Failed to read birthday from parameters. %s", err.Error()), client, callID, nil, nil, nil)
		return
	}

	go protocol.UpdateAndGetAllInformationHandler(nil, client, callID, nnaInfo.(*NNAInfo), presence.(*NintendoPresenceV2), birthday)
}
