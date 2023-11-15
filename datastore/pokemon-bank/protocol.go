// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Pokemon Bank) protocol
	ProtocolID = 0x73

	// MethodUploadPokemon is the method ID for the UploadPokemon method
	MethodUploadPokemon = 0x29

	// MethodSearchPokemon is the method ID for the SearchPokemon method
	MethodSearchPokemon = 0x2A

	// MethodPrepareTradePokemon is the method ID for the PrepareTradePokemon method
	MethodPrepareTradePokemon = 0x2B

	// MethodTradePokemon is the method ID for the TradePokemon method
	MethodTradePokemon = 0x2C

	// MethodDownloadOtherPokemon is the method ID for the DownloadOtherPokemon method
	MethodDownloadOtherPokemon = 0x2D

	// MethodDownloadMyPokemon is the method ID for the DownloadMyPokemon method
	MethodDownloadMyPokemon = 0x2E

	// MethodDeletePokemon is the method ID for the DeletePokemon method
	MethodDeletePokemon = 0x2F

	// MethodGetTransactionParam is the method ID for the GetTransactionParam method
	MethodGetTransactionParam = 0x30

	// MethodPreparePostBankObject is the method ID for the PreparePostBankObject method
	MethodPreparePostBankObject = 0x31

	// MethodCompletePostBankObject is the method ID for the CompletePostBankObject method
	MethodCompletePostBankObject = 0x32

	// MethodPrepareGetBankObject is the method ID for the PrepareGetBankObject method
	MethodPrepareGetBankObject = 0x33

	// MethodPrepareUpdateBankObject is the method ID for the PrepareUpdateBankObject method
	MethodPrepareUpdateBankObject = 0x34

	// MethodCompleteUpdateBankObject is the method ID for the CompleteUpdateBankObject method
	MethodCompleteUpdateBankObject = 0x35

	// MethodRollbackBankObject is the method ID for the RollbackBankObject method
	MethodRollbackBankObject = 0x36

	// MethodGetUnlockKey is the method ID for the GetUnlockKey method
	MethodGetUnlockKey = 0x37

	// MethodRequestMigration is the method ID for the RequestMigration method
	MethodRequestMigration = 0x38
)

var patchedMethods = []uint32{
	MethodUploadPokemon,
	MethodSearchPokemon,
	MethodPrepareTradePokemon,
	MethodTradePokemon,
	MethodDownloadOtherPokemon,
	MethodDownloadMyPokemon,
	MethodDeletePokemon,
	MethodGetTransactionParam,
	MethodPreparePostBankObject,
	MethodCompletePostBankObject,
	MethodPrepareGetBankObject,
	MethodPrepareUpdateBankObject,
	MethodCompleteUpdateBankObject,
	MethodRollbackBankObject,
	MethodGetUnlockKey,
	MethodRequestMigration,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Pokemon Bank) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	Server nex.ServerInterface
	dataStoreProtocol
	UploadPokemon            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationUploadPokemonParam) (*nex.RMCMessage, uint32)
	SearchPokemon            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationSearchPokemonParam) (*nex.RMCMessage, uint32)
	PrepareTradePokemon      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationPrepareTradePokemonParam) (*nex.RMCMessage, uint32)
	TradePokemon             func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationTradePokemonParam) (*nex.RMCMessage, uint32)
	DownloadOtherPokemon     func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDownloadOtherPokemonParam) (*nex.RMCMessage, uint32)
	DownloadMyPokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDownloadMyPokemonParam) (*nex.RMCMessage, uint32)
	DeletePokemon            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDeletePokemonParam) (*nex.RMCMessage, uint32)
	GetTransactionParam      func(err error, packet nex.PacketInterface, callID uint32, slotID uint16) (*nex.RMCMessage, uint32)
	PreparePostBankObject    func(err error, packet nex.PacketInterface, callID uint32, slotID uint16, size uint32) (*nex.RMCMessage, uint32)
	CompletePostBankObject   func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, uint32)
	PrepareGetBankObject     func(err error, packet nex.PacketInterface, callID uint32, slotID uint16, applicationID uint16) (*nex.RMCMessage, uint32)
	PrepareUpdateBankObject  func(err error, packet nex.PacketInterface, callID uint32, transactionParam *datastore_pokemon_bank_types.BankTransactionParam) (*nex.RMCMessage, uint32)
	CompleteUpdateBankObject func(err error, packet nex.PacketInterface, callID uint32, slotID uint16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce bool) (*nex.RMCMessage, uint32)
	RollbackBankObject       func(err error, packet nex.PacketInterface, callID uint32, slotID uint16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce bool) (*nex.RMCMessage, uint32)
	GetUnlockKey             func(err error, packet nex.PacketInterface, callID uint32, challengeValue uint32) (*nex.RMCMessage, uint32)
	RequestMigration         func(err error, packet nex.PacketInterface, callID uint32, oneTimePassword string, boxes []uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID) {
				protocol.HandlePacket(packet)
			} else {
				protocol.dataStoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodUploadPokemon:
		protocol.handleUploadPokemon(packet)
	case MethodSearchPokemon:
		protocol.handleSearchPokemon(packet)
	case MethodPrepareTradePokemon:
		protocol.handlePrepareTradePokemon(packet)
	case MethodTradePokemon:
		protocol.handleTradePokemon(packet)
	case MethodDownloadOtherPokemon:
		protocol.handleDownloadOtherPokemon(packet)
	case MethodDownloadMyPokemon:
		protocol.handleDownloadMyPokemon(packet)
	case MethodDeletePokemon:
		protocol.handleDeletePokemon(packet)
	case MethodGetTransactionParam:
		protocol.handleGetTransactionParam(packet)
	case MethodPreparePostBankObject:
		protocol.handlePreparePostBankObject(packet)
	case MethodCompletePostBankObject:
		protocol.handleCompletePostBankObject(packet)
	case MethodPrepareGetBankObject:
		protocol.handlePrepareGetBankObject(packet)
	case MethodPrepareUpdateBankObject:
		protocol.handlePrepareUpdateBankObject(packet)
	case MethodCompleteUpdateBankObject:
		protocol.handleCompleteUpdateBankObject(packet)
	case MethodRollbackBankObject:
		protocol.handleRollbackBankObject(packet)
	case MethodGetUnlockKey:
		protocol.handleGetUnlockKey(packet)
	case MethodRequestMigration:
		protocol.handleRequestMigration(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported DataStore (Pokemon Bank) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new DataStorePokemonBank protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.dataStoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
