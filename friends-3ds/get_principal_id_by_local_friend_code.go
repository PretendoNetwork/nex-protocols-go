// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrincipalIDByLocalFriendCode sets the GetPrincipalIDByLocalFriendCode handler function
func (protocol *Protocol) GetPrincipalIDByLocalFriendCode(handler func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, lfcList []uint64) uint32) {
	protocol.getPrincipalIDByLocalFriendCodeHandler = handler
}

func (protocol *Protocol) handleGetPrincipalIDByLocalFriendCode(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPrincipalIDByLocalFriendCodeHandler == nil {
		globals.Logger.Warning("Friends3DS::GetPrincipalIDByLocalFriendCode not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getPrincipalIDByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lfcList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.getPrincipalIDByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPrincipalIDByLocalFriendCodeHandler(nil, packet, callID, lfc, lfcList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
