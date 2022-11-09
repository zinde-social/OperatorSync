// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataTypesCharacter is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCharacter struct {
	CharacterId *big.Int
	Handle      string
	Uri         string
	NoteCount   *big.Int
	SocialToken common.Address
	LinkModule  common.Address
}

// DataTypesCreateCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateCharacterData struct {
	To                 common.Address
	Handle             string
	Uri                string
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypesERC721Struct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesERC721Struct struct {
	TokenAddress  common.Address
	Erc721TokenId *big.Int
}

// DataTypesMintNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMintNoteData struct {
	CharacterId    *big.Int
	NoteId         *big.Int
	To             common.Address
	MintModuleData []byte
}

// DataTypesNote is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNote struct {
	LinkItemType [32]byte
	LinkKey      [32]byte
	ContentUri   string
	LinkModule   common.Address
	MintModule   common.Address
	MintNFT      common.Address
	Deleted      bool
	Locked       bool
}

// DataTypesNoteStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNoteStruct struct {
	CharacterId *big.Int
	NoteId      *big.Int
}

// DataTypesPostNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostNoteData struct {
	CharacterId        *big.Int
	ContentUri         string
	LinkModule         common.Address
	LinkModuleInitData []byte
	MintModule         common.Address
	MintModuleInitData []byte
	Locked             bool
}

// DataTypescreateThenLinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypescreateThenLinkCharacterData struct {
	FromCharacterId *big.Int
	To              common.Address
	LinkType        [32]byte
}

// DataTypeslinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypessetLinkModule4AddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4AddressData struct {
	Account            common.Address
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4CharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4CharacterData struct {
	CharacterId        *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4ERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4ERC721Data struct {
	TokenAddress       common.Address
	TokenId            *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4LinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4LinklistData struct {
	LinklistId         *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4NoteData struct {
	CharacterId        *big.Int
	NoteId             *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetMintModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetMintModule4NoteData struct {
	CharacterId        *big.Int
	NoteId             *big.Int
	MintModule         common.Address
	MintModuleInitData []byte
}

// DataTypesunlinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
}

// DataTypesunlinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
}

// DataTypesunlinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.createThenLinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createThenLinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"deleteNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacter\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getCharacterByHandle\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacterUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLinkModule4Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4ERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4Linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinklistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinklistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkListId\",\"type\":\"uint256\"}],\"name\":\"getLinklistType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinklistUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.Note\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrimaryCharacterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_linklistContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mintNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_periphery\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"isPrimaryCharacter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"lockNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MintNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mintNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"postNote4Address\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"postNote4AnyUri\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"}],\"name\":\"postNote4Character\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"erc721\",\"type\":\"tuple\"}],\"name\":\"postNote4ERC721\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NoteStruct\",\"name\":\"note\",\"type\":\"tuple\"}],\"name\":\"postNote4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setCharacterUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"setHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4CharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Character\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4ERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4ERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Linklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setLinklistUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setMintModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setNoteUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryCharacterId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setSocialToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615d4280620000216000396000f3fe608060405234801561001057600080fd5b50600436106104545760003560e01c8063867884e611610241578063c87b56dd1161013b578063e56f2fe4116100c3578063f2ad807511610087578063f2ad807514610a85578063f316bacd14610a98578063f6479d7714610aab578063fd2d866f14610abe578063fe9299fb14610ad157600080fd5b8063e56f2fe4146109fd578063e7a1c1c014610a10578063e985e9c514610a23578063ec81d19414610a5f578063ef0828ab14610a7257600080fd5b8063d70e10c61161010a578063d70e10c614610979578063dabb0531146109a4578063db491e80146109b7578063db8c198d146109d7578063dca27135146109ea57600080fd5b8063c87b56dd1461092d578063cb8e757e14610940578063cd69fe6114610953578063d23b320b1461096657600080fd5b80639a50248d116101c9578063af90b1121161018d578063af90b112146108d0578063b88d4fde146108e3578063b9d32845146108f6578063c053f6b814610909578063c2a6fe3b1461091a57600080fd5b80639a50248d14610864578063a22cb46514610884578063a48047ba14610897578063a6e6178d146108aa578063a7ccb4bf146108bd57600080fd5b806393f057e51161021057806393f057e51461081057806395d89b411461082357806395d9fa7d1461082b5780639864c3071461083e5780639a4dec181461085157600080fd5b8063867884e6146107c45780638734bbfc146107d75780638b4ca06a146107ea57806392f7070b146107fd57600080fd5b806331b9d08c116103525780634f6ccce7116102da57806369492c971161029e57806369492c97146107585780636bf55d5f1461076b57806370a082311461078b57806374f345cf1461079e5780637d46a713146107b157600080fd5b80634f6ccce7146106f95780635a936d101461070c5780635fb881831461071f578063628b644a146107325780636352211e1461074557600080fd5b806340ad34d81161032157806340ad34d81461069a57806342842e0e146106ad57806342966c68146106c057806344b82a24146106d357806347f94de7146106e657600080fd5b806331b9d08c14610635578063327b2a031461066157806333f06ee614610674578063388f50831461068757600080fd5b80630ff98244116103e05780632209d145116103a45780632209d1451461059d57806323b872dd146105d357806329c301c2146105e65780632abc6bf6146105f95780632f745c591461062257600080fd5b80630ff982441461054b5780631316529d1461055e578063144a3e831461056f57806318160ddd14610582578063188b04b31461058a57600080fd5b8063081812fc11610427578063081812fc146104ea57806308cb68ff146104fd578063095ea7b3146105125780630c4dd5f2146105255780630d18d07a1461053857600080fd5b806301ffc9a71461045957806304f3bcec1461048157806305f63c8a146104ac57806306fdde03146104d5575b600080fd5b61046c610467366004614788565b610afa565b60405190151581526020015b60405180910390f35b601754610494906001600160a01b031681565b6040516001600160a01b039091168152602001610478565b6104946104ba3660046147a5565b6000908152601660205260409020546001600160a01b031690565b6104dd610b25565b6040516104789190614816565b6104946104f83660046147a5565b610bb7565b61051061050b366004614841565b610c51565b005b61051061052036600461489a565b610cea565b610510610533366004614841565b610e00565b6105106105463660046148c6565b610e65565b6105106105593660046148f6565b610e7c565b60045b604051908152602001610478565b6104dd61057d3660046147a5565b610f7b565b600854610561565b610510610598366004614924565b610f86565b6104946105ab36600461489a565b6001600160a01b03918216600090815260106020908152604080832093835292905220541690565b6105106105e1366004614958565b611083565b6105616105f43660046149ab565b6110b4565b6105616106073660046149df565b6001600160a01b03166000908152600c602052604090205490565b61056161063036600461489a565b611160565b6104946106433660046149df565b6001600160a01b039081166000908152601160205260409020541690565b61056161066f366004614a0e565b6111f6565b610510610682366004614aa4565b6113b9565b610510610695366004614924565b61142d565b6105106106a8366004614aef565b611480565b6105106106bb366004614958565b6114fa565b6105106106ce3660046147a5565b611515565b6105616106e1366004614b0b565b6115ae565b6105106106f4366004614aa4565b611624565b6105616107073660046147a5565b611664565b61051061071a3660046148f6565b6116f7565b61051061072d366004614924565b611787565b610510610740366004614b4f565b6117da565b6104946107533660046147a5565b611866565b610510610766366004614924565b6118dd565b61077e6107793660046147a5565b611a15565b6040516104789190614ba1565b6105616107993660046149df565b611a2f565b6105106107ac366004614bee565b611ab6565b6105106107bf3660046148c6565b611b2e565b6105106107d2366004614aef565b611b41565b61046c6107e53660046147a5565b611bb1565b6105616107f83660046147a5565b611bdf565b61056161080b366004614c10565b611c4b565b61051061081e3660046148f6565b611cdf565b6104dd611d4e565b6105106108393660046148c6565b611d5d565b61051061084c366004614924565b611de3565b61056161085f366004614a0e565b611e36565b610877610872366004614c56565b611f3e565b6040516104789190614c97565b610510610892366004614d23565b6120ed565b61046c6108a53660046148c6565b6120f8565b6105106108b8366004614aa4565b612143565b6105616108cb366004614924565b6121b9565b6105616108de366004614b0b565b61225d565b6105106108f1366004614e3a565b61229b565b610510610904366004614eb7565b6122d3565b6013546001600160a01b0316610494565b610510610928366004614bee565b6123a2565b6104dd61093b3660046147a5565b612412565b61051061094e366004614eb7565b6124b7565b610510610961366004614eb7565b612527565b610510610974366004614924565b61253b565b610561610987366004614bee565b6000918252600d6020908152604080842092845291905290205490565b6108776109b23660046147a5565b6125b5565b6109ca6109c5366004614bee565b61273c565b6040516104789190614eeb565b6105106109e5366004614924565b61289d565b6104dd6109f83660046147a5565b6128e7565b610510610a0b366004614f81565b612959565b610510610a1e3660046148c6565b612b0e565b61046c610a3136600461503b565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6104dd610a6d3660046147a5565b612b21565b610510610a80366004614841565b612b41565b610510610a933660046147a5565b612bb0565b610561610aa6366004615069565b612bff565b610510610ab93660046148f6565b612d5b565b610510610acc366004614841565b612d95565b610494610adf3660046147a5565b6000908152600f60205260409020546001600160a01b031690565b60006001600160e01b0319821663780e9d6360e01b1480610b1f5750610b1f82612e0e565b92915050565b606060008054610b34906150c4565b80601f0160208091040260200160405190810160405280929190818152602001828054610b60906150c4565b8015610bad5780601f10610b8257610100808354040283529160200191610bad565b820191906000526020600020905b815481529060010190602001808311610b9057829003601f168201915b5050505050905090565b6000818152600260205260408120546001600160a01b0316610c355760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b73__$69db18565ecbefece480e92cc5f8fb1274$__63dfc34f25610c7860208401846149df565b610c8860408501602086016149df565b610c9560408601866150f9565b60116040518663ffffffff1660e01b8152600401610cb7959493929190615168565b60006040518083038186803b158015610ccf57600080fd5b505af4158015610ce3573d6000803e3d6000fd5b5050505050565b6000610cf582611866565b9050806001600160a01b0316836001600160a01b03161415610d635760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610c2c565b336001600160a01b0382161480610d7f5750610d7f8133610a31565b610df15760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610c2c565b610dfb8383612e5e565b505050565b610e0a8135612ecc565b73__$69db18565ecbefece480e92cc5f8fb1274$__636252159e8235610e3660408501602086016149df565b610e4360408601866150f9565b600f6040518663ffffffff1660e01b8152600401610cb79594939291906151a7565b610e6e82612f8c565b610e78828261300d565b5050565b610e86813561306e565b6040516331a9108f60e11b81528135600482015273__$e2c9a8f399964af47bb173600dd9e5f662$__90633019fadd9083903090636352211e90602401602060405180830381865afa158015610ee0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0491906151d2565b60135485356000908152600d60209081526040808320818a0135845282529182902054825160e088901b6001600160e01b031916815286356004820152918601356024830152949091013560448201526001600160a01b0392831660648201529116608482015260a481019190915260c401610cb7565b6060610b1f82612412565b610f90813561306e565b610f9d816020013561313c565b73__$e2c9a8f399964af47bb173600dd9e5f662$__63957f8a98823560208401356040850135610fd060608701876150f9565b6040516331a9108f60e11b8152883560048201523090636352211e90602401602060405180830381865afa15801561100c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103091906151d2565b6013546020808b01356000908152600a9091526040908190206005015490516001600160e01b031960e08b901b168152610cb7989796959493926001600160a01b03908116921690600d906004016151ef565b61108d3382613195565b6110a95760405162461bcd60e51b8152600401610c2c90615249565b610dfb83838361328c565b60006110c0823561306e565b81356000908152600a60205260408120600301805482906110e0906152b0565b91829055506040516342a34a5360e01b815290915073__$2873850e5b3f5697aafc2f5cd9f27cc884$__906342a34a539061112990869085906000908190600e906004016153d5565b60006040518083038186803b15801561114157600080fd5b505af4158015611155573d6000803e3d6000fd5b509295945050505050565b600061116b83611a2f565b82106111cd5760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610c2c565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b6000611202833561306e565b61121c61121260208401846149df565b8360200135613433565b6013546545524337323160d01b906000906001600160a01b0316632ea24efc8261124960208801886149df565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b03166024820152602087013560448201526064016020604051808303816000875af115801561129f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c3919061541a565b85356000908152600a602052604081206003018054929350909182906112e8906152b0565b9182905550905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a538783868661131a60208c018c6149df565b8b6020013560405160200161134d92919060609290921b6bffffffffffffffffffffffff19168252601482015260340190565b604051602081830303815290604052600e6040518763ffffffff1660e01b815260040161137f96959493929190615433565b60006040518083038186803b15801561139757600080fd5b505af41580156113ab573d6000803e3d6000fd5b509298975050505050505050565b6113c283612ecc565b601354604051633c17845760e11b81526001600160a01b039091169063782f08ae906113f69086908690869060040161547e565b600060405180830381600087803b15801561141057600080fd5b505af1158015611424573d6000803e3d6000fd5b50505050505050565b611437813561306e565b60135460405163e4bb4f2360e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9163e4bb4f2391610cb79185916001600160a01b0390911690600d90600401615498565b61148a813561306e565b60135460408051635cebf08d60e11b81528335600482015260208401356024820152908301356044820152606083013560648201526001600160a01b039091166084820152600d60a482015273__$e2c9a8f399964af47bb173600dd9e5f662$__9063b9d7e11a9060c401610cb7565b610dfb8383836040518060200160405280600081525061229b565b6000818152600a602052604080822090516115339160010190615506565b60408051918290039091206000818152600b6020908152838220829055858252600a90529182208281559092509061156e60018301826145e3565b61157c6002830160006145e3565b50600060038201556004810180546001600160a01b0319908116909155600590910180549091169055610e78826134e3565b60006115ba833561306e565b82356000908152600a60205260408120600301805467131a5b9adb1a5cdd60c21b928592909182906115eb906152b0565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a60405160200161134d91815260200190565b610dfb8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061354292505050565b600061166f60085490565b82106116d25760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610c2c565b600882815481106116e5576116e56155a2565b90600052602060002001549050919050565b611701813561306e565b60135481356000818152600d60209081526040808320818701358085529083529281902054905163d862986560e01b8152600481019490945290850135602484015260448301919091526001600160a01b039092166064820152608481019190915273__$e2c9a8f399964af47bb173600dd9e5f662$__9063d86298659060a401610cb7565b611791813561306e565b6013546040516332f107bf60e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__916332f107bf91610cb79185916001600160a01b0390911690600d906004016155b8565b6117e38461306e565b6117ed84846135a1565b6040516001626802bf60e01b0319815273__$2873850e5b3f5697aafc2f5cd9f27cc884$__9063ff97fd4190611830908790879087908790600e90600401615635565b60006040518083038186803b15801561184857600080fd5b505af415801561185c573d6000803e3d6000fd5b5050505050505050565b6000818152600260205260408120546001600160a01b031680610b1f5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610c2c565b6118ea60208201826149df565b6040516331a9108f60e11b8152602083013560048201526001600160a01b039190911690636352211e90602401602060405180830381865afa158015611934573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061195891906151d2565b6001600160a01b0316336001600160a01b0316146119a95760405162461bcd60e51b815260206004820152600e60248201526d2737ba22a9219b9918a7bbb732b960911b6044820152606401610c2c565b73__$69db18565ecbefece480e92cc5f8fb1274$__63384afeb36119d060208401846149df565b60208401356119e560608601604087016149df565b6119f260608701876150f9565b60106040518763ffffffff1660e01b8152600401610cb796959493929190615655565b6000818152601860205260409020606090610b1f90613656565b60006001600160a01b038216611a9a5760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610c2c565b506001600160a01b031660009081526003602052604090205490565b611abf8261306e565b611ac982826135a1565b6000828152600e60209081526040808320848452825291829020600501805460ff60a81b1916600160a81b179055905182815283917f036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd3091015b60405180910390a25050565b611b3782612f8c565b610e78828261366a565b611b4b813561306e565b60135481356000908152600d6020908152604080832060608601358452909152908190205490516317e9a8b160e31b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9263bf4d458892610cb79286926001600160a01b0316919060040161569c565b600080611bbd83611866565b6001600160a01b03166000908152600c60205260409020549290921492915050565b60135460405162fba02760e01b8152600481018390526000916001600160a01b03169062fba02790602401602060405180830381865afa158015611c27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b1f919061541a565b6000611c57833561306e565b82356000908152600a602052604081206003018054664164647265737360c81b926001600160a01b0386169290918290611c90906152b0565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a60405160200161134d919060609190911b6bffffffffffffffffffffffff1916815260140190565b611ce9813561306e565b60135481356000908152600d6020908152604080832081860135845290915290819020549051633d7f9b3d60e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__92637aff367a92610cb79286926001600160a01b031691906004016156ee565b606060018054610b34906150c4565b611d6682612f8c565b6040516384b44a2f60e01b8152600481018390526001600160a01b0382166024820152600a604482015273__$d8529f0216812a2dec6b96adf1943c3537$__906384b44a2f9060640160006040518083038186803b158015611dc757600080fd5b505af4158015611ddb573d6000803e3d6000fd5b505050505050565b611ded813561306e565b601354604051635fe5df1d60e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9163bfcbbe3a91610cb79185916001600160a01b0390911690600d90600401615736565b6000611e42833561306e565b601354604051635cb46be760e01b81526000600482018190528435602483015260208501356044830152634e6f746560e01b9290916001600160a01b0390911690635cb46be7906064016020604051808303816000875af1158015611eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ecf919061541a565b85356000908152600a60205260408120600301805492935090918290611ef4906152b0565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a600001358b6020013560405160200161134d929190918252602082015260400190565b611f4661461d565b60008383604051611f58929190615796565b604080519182900382206000818152600b602090815283822054808352600a82529184902060c0860190945283548552600184018054939650919493929084019190611fa3906150c4565b80601f0160208091040260200160405190810160405280929190818152602001828054611fcf906150c4565b801561201c5780601f10611ff15761010080835404028352916020019161201c565b820191906000526020600020905b815481529060010190602001808311611fff57829003601f168201915b50505050508152602001600282018054612035906150c4565b80601f0160208091040260200160405190810160405280929190818152602001828054612061906150c4565b80156120ae5780601f10612083576101008083540402835291602001916120ae565b820191906000526020600020905b81548152906001019060200180831161209157829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015295945050505050565b610e783383836136bf565b600082815260166020908152604080832054601890925282206001600160a01b0384811692169190911490829061212f908561378e565b9050818061213a5750805b95945050505050565b61214c83612f8c565b60405163130f361d60e01b815273__$d8529f0216812a2dec6b96adf1943c3537$__9063130f361d9061218d90869086908690600b90600a906004016157a6565b60006040518083038186803b1580156121a557600080fd5b505af4158015611424573d6000803e3d6000fd5b60006121ca823560208401356135a1565b73__$2873850e5b3f5697aafc2f5cd9f27cc884$__635be69415833560208501356121fb60608701604088016149df565b61220860608801886150f9565b6014546040516001600160e01b031960e089901b1681526122409695949392916001600160a01b031690600a90600e906004016157d4565b602060405180830381865af4158015611c27573d6000803e3d6000fd5b6000612269833561306e565b82356000908152600a6020526040812060030180546821b430b930b1ba32b960b91b928592909182906115eb906152b0565b6122a53383613195565b6122c15760405162461bcd60e51b8152600401610c2c90615249565b6122cd848484846137b0565b50505050565b6122dd813561306e565b6122ef816020013582604001356135a1565b6040516331a9108f60e11b81528135600482015273__$e2c9a8f399964af47bb173600dd9e5f662$__9063018212d19083903090636352211e90602401602060405180830381865afa158015612349573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061236d91906151d2565b6013546040516001600160e01b031960e086901b168152610cb79392916001600160a01b031690600e90600d90600401615822565b6123ab82612f8c565b6123b582826135a1565b6000828152600e60209081526040808320848452825291829020600501805460ff60a01b1916600160a01b179055905182815283917f4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf939101611b22565b6000818152600a60205260409020600201805460609190612432906150c4565b80601f016020809104026020016040519081016040528092919081815260200182805461245e906150c4565b80156124ab5780601f10612480576101008083540402835291602001916124ab565b820191906000526020600020905b81548152906001019060200180831161248e57829003601f168201915b50505050509050919050565b6124c1813561306e565b6124de6124d460408301602084016149df565b8260400135613433565b601354604051638f3334ff60e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__91638f3334ff91610cb79185916001600160a01b0390911690600d906004016158a1565b61253861253382615904565b6137e3565b50565b612545813561306e565b612554813560208301356135a1565b73__$69db18565ecbefece480e92cc5f8fb1274$__6320828a028235602084013561258560608601604087016149df565b61259260608701876150f9565b600e6040518763ffffffff1660e01b8152600401610cb7969594939291906159b7565b6125bd61461d565b600a60008381526020019081526020016000206040518060c0016040529081600082015481526020016001820180546125f5906150c4565b80601f0160208091040260200160405190810160405280929190818152602001828054612621906150c4565b801561266e5780601f106126435761010080835404028352916020019161266e565b820191906000526020600020905b81548152906001019060200180831161265157829003601f168201915b50505050508152602001600282018054612687906150c4565b80601f01602080910402602001604051908101604052809291908181526020018280546126b3906150c4565b80156127005780601f106126d557610100808354040283529160200191612700565b820191906000526020600020905b8154815290600101906020018083116126e357829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015292915050565b60408051610100808201835260008083526020808401829052606084860181905284018290526080840182905260a0840182905260c0840182905260e08401829052868252600e815284822086835281529084902084519283018552805483526001810154918301919091526002810180549394929391928401916127c0906150c4565b80601f01602080910402602001604051908101604052809291908181526020018280546127ec906150c4565b80156128395780601f1061280e57610100808354040283529160200191612839565b820191906000526020600020905b81548152906001019060200180831161281c57829003601f168201915b505050918352505060038201546001600160a01b039081166020830152600483015481166040830152600590920154918216606082015260ff600160a01b8304811615156080830152600160a81b909204909116151560a090910152905092915050565b6128a7813561306e565b6128b6813560208301356135a1565b73__$69db18565ecbefece480e92cc5f8fb1274$__631f2ffb698235602084013561258560608601604087016149df565b601354604051632b05429560e21b8152600481018390526060916001600160a01b03169063ac150a5490602401600060405180830381865afa158015612931573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b1f91908101906159e5565b601454600160a81b900460ff161580801561298157506014546001600160a01b90910460ff16105b806129a25750303b1580156129a25750601454600160a01b900460ff166001145b612a055760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610c2c565b6014805460ff60a01b1916600160a01b1790558015612a32576014805460ff60a81b1916600160a81b1790555b612a3e898989896138b8565b601380546001600160a01b03199081166001600160a01b0388811691909117909255601480548216878416179055601580548216868416179055601780549091169184169190911790556040514281527f400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf9060200160405180910390a18015612b03576014805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b612b1782612f8c565b610e788282613909565b6000818152600a60205260409020600101805460609190612432906150c4565b612b4b813561306e565b60135481356000908152600d602090815260408083208186013584529091529081902054905163bf5c00c160e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9263bf5c00c192610cb79286926001600160a01b03169190600401615a52565b612bb981612f8c565b336000818152600c602052604080822080549085905590519092839285927fce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de9190a45050565b6000612c0b843561306e565b601354604051633610bf0960e11b815265416e7955726960d01b916000916001600160a01b0390911690636c217e1290612c4d9084908990899060040161547e565b6020604051808303816000875af1158015612c6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c90919061541a565b86356000908152600a60205260408120600301805492935090918290612cb5906152b0565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53888386868b8b604051602001612cee929190615796565b604051602081830303815290604052600e6040518763ffffffff1660e01b8152600401612d2096959493929190615433565b60006040518083038186803b158015612d3857600080fd5b505af4158015612d4c573d6000803e3d6000fd5b50929998505050505050505050565b6125388135612d7060408401602085016149df565b836040013560405180604001604052806002815260200161060f60f31b815250613966565b612d9f813561306e565b73__$d8529f0216812a2dec6b96adf1943c3537$__631dc831338235612dcb60408501602086016149df565b612dd860408601866150f9565b86356000908152600a60205260409081902090516001600160e01b031960e088901b168152610cb79594939291906004016151a7565b60006001600160e01b031982166380ac58cd60e01b1480612e3f57506001600160e01b03198216635b5e139f60e01b145b80610b1f57506301ffc9a760e01b6001600160e01b0319831614610b1f565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190612e9382611866565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6013546040516331a9108f60e11b8152600481018390526001600160a01b0390911690636352211e90602401602060405180830381865afa158015612f15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f3991906151d2565b6001600160a01b0316336001600160a01b0316146125385760405162461bcd60e51b815260206004820152601060248201526f2737ba2634b735a634b9ba27bbb732b960811b6044820152606401610c2c565b6000612f9782611866565b9050336001600160a01b0382161480612fcd5750326001600160a01b038216148015612fcd57506015546001600160a01b031633145b610e785760405162461bcd60e51b81526020600482015260116024820152702737ba21b430b930b1ba32b927bbb732b960791b6044820152606401610c2c565b60008281526018602052604090206130259082613bec565b50806001600160a01b0316827faa9506a57073a80893a2d2fdd53f4bd49d281a8548f483ad230f2d5da7f6710c4260405161306291815260200190565b60405180910390a35050565b600061307982611866565b6000838152601860205260409020909150613094903361378e565b806130a75750336001600160a01b038216145b806130c857506000828152601660205260409020546001600160a01b031633145b806130f05750326001600160a01b0382161480156130f057506015546001600160a01b031633145b610e785760405162461bcd60e51b815260206004820152601c60248201527f4e6f744368617261637465724f776e65724e6f724f70657261746f72000000006044820152606401610c2c565b6000818152600260205260409020546001600160a01b03166125385760405162461bcd60e51b81526020600482015260126024820152714368617261637465724e6f7445786973747360701b6044820152606401610c2c565b6000818152600260205260408120546001600160a01b031661320e5760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610c2c565b600061321983611866565b9050806001600160a01b0316846001600160a01b031614806132545750836001600160a01b031661324984610bb7565b6001600160a01b0316145b8061328457506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b949350505050565b826001600160a01b031661329f82611866565b6001600160a01b0316146133035760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610c2c565b6001600160a01b0382166133655760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610c2c565b613370838383613c01565b61337b600082612e5e565b6001600160a01b03831660009081526003602052604081208054600192906133a4908490615aaa565b90915550506001600160a01b03821660009081526003602052604081208054600192906133d2908490615ac1565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6040516331a9108f60e11b8152600481018290526001600160a01b03831690636352211e90602401602060405180830381865afa158015613478573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061349c91906151d2565b6001600160a01b0316610e785760405162461bcd60e51b815260206004820152600f60248201526e5245433732314e6f7445786973747360881b6044820152606401610c2c565b6134ed3382613195565b6135395760405162461bcd60e51b815260206004820152601b60248201527f4e4654426173653a204e6f744f776e65724f72417070726f76656400000000006044820152606401610c2c565b61253881613c8a565b61354b8261306e565b6000828152600a60209081526040909120825161357092600290920191840190614665565b50817f17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c663199782604051611b229190614816565b6000828152600e60209081526040808320848452909152902060050154600160a01b900460ff16156136055760405162461bcd60e51b815260206004820152600d60248201526c139bdd19525cd1195b195d1959609a1b6044820152606401610c2c565b6000828152600a6020526040902060030154811115610e785760405162461bcd60e51b815260206004820152600d60248201526c4e6f74654e6f7445786973747360981b6044820152606401610c2c565b6060600061366383613d31565b9392505050565b60008281526018602052604090206136829082613d8c565b50806001600160a01b0316827f58f51b5bb567864de85c6a422b33491f2418924a44613d7b9341f58657bdd8334260405161306291815260200190565b816001600160a01b0316836001600160a01b031614156137215760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610c2c565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b03811660009081526001830160205260408120541515613663565b6137bb84848461328c565b6137c784848484613da1565b6122cd5760405162461bcd60e51b8152600401610c2c90615ad9565b60006012600081546137f4906152b0565b918290555082519091506138089082613e9f565b604051632902741560e01b815273__$d8529f0216812a2dec6b96adf1943c3537$__9063290274159061384a9085906001908690600b90600a90600401615b2b565b60006040518083038186803b15801561386257600080fd5b505af4158015613876573d6000803e3d6000fd5b505083516001600160a01b03166000908152600c602052604090205415159150610e7890505790516001600160a01b03166000908152600c6020526040902055565b6138c484848484613eb9565b7f414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c3084848484426040516138fb959493929190615bd4565b60405180910390a150505050565b60008281526016602090815260409182902080546001600160a01b0319166001600160a01b038516908117909155915142815284917f691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c4149319101613062565b61396f84612f8c565b6001600160a01b0383166000908152600c6020526040902054156139eb5760405162461bcd60e51b815260206004820152602d60248201527f546172676574206164647265737320616c726561647920686173207072696d6160448201526c393c9031b430b930b1ba32b91760991b6064820152608401610c2c565b60006012600081546139fc906152b0565b91829055509050613a0d8482613e9f565b73__$d8529f0216812a2dec6b96adf1943c3537$__63290274156040518060a00160405280876001600160a01b03168152602001613a55886001600160a01b03166014613ed2565b815260200160405180602001604052806000815250815260200160006001600160a01b0316815260200160405180602001604052806000815250815250600084600b600a6040518663ffffffff1660e01b8152600401613ab9959493929190615b2b565b60006040518083038186803b158015613ad157600080fd5b505af4158015613ae5573d6000803e3d6000fd5b505050506001600160a01b0384166000908152600c602052604090819020829055516331a9108f60e11b81526004810186905273__$e2c9a8f399964af47bb173600dd9e5f662$__9063957f8a989087908490879087903090636352211e90602401602060405180830381865afa158015613b64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b8891906151d2565b6013546040516001600160e01b031960e089901b168152613bc09695949392916001600160a01b031690600090600d90600401615c0e565b60006040518083038186803b158015613bd857600080fd5b505af4158015612b03573d6000803e3d6000fd5b6000613663836001600160a01b03841661406d565b6000818152601860205260408120613c1890613656565b905060005b8151811015613c7e576000828281518110613c3a57613c3a6155a2565b60200260200101519050613c698160186000878152602001908152602001600020613bec90919063ffffffff16565b50508080613c76906152b0565b915050613c1d565b506122cd848484614160565b6000613c9582611866565b9050613ca381600084613c01565b613cae600083612e5e565b6001600160a01b0381166000908152600360205260408120805460019290613cd7908490615aaa565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6060816000018054806020026020016040519081016040528092919081815260200182805480156124ab57602002820191906000526020600020905b815481526020019060010190808311613d6d5750505050509050919050565b6000613663836001600160a01b0384166141cb565b60006001600160a01b0384163b15613e9457604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290613de5903390899088908890600401615c66565b6020604051808303816000875af1925050508015613e20575060408051601f3d908101601f19168201909252613e1d91810190615ca3565b60015b613e7a573d808015613e4e576040519150601f19603f3d011682016040523d82523d6000602084013e613e53565b606091505b508051613e725760405162461bcd60e51b8152600401610c2c90615ad9565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050613284565b506001949350505050565b610e7882826040518060200160405280600081525061421a565b613ec5600085856146e9565b50610ce3600183836146e9565b60606000613ee1836002615cc0565b613eec906002615ac1565b6001600160401b03811115613f0357613f03614d4f565b6040519080825280601f01601f191660200182016040528015613f2d576020820181803683370190505b509050600360fc1b81600081518110613f4857613f486155a2565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110613f7757613f776155a2565b60200101906001600160f81b031916908160001a9053506000613f9b846002615cc0565b613fa6906001615ac1565b90505b600181111561401e576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110613fda57613fda6155a2565b1a60f81b828281518110613ff057613ff06155a2565b60200101906001600160f81b031916908160001a90535060049490941c9361401781615cdf565b9050613fa9565b5083156136635760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610c2c565b60008181526001830160205260408120548015614156576000614091600183615aaa565b85549091506000906140a590600190615aaa565b905081811461410a5760008660000182815481106140c5576140c56155a2565b90600052602060002001549050808760000184815481106140e8576140e86155a2565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061411b5761411b615cf6565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610b1f565b6000915050610b1f565b6000818152601660205260409020546001600160a01b03161561418857614188816000613909565b6001600160a01b0383166000908152600c6020526040902054156141c0576001600160a01b0383166000908152600c60205260408120555b610dfb83838361424d565b600081815260018301602052604081205461421257508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610b1f565b506000610b1f565b6142248383614305565b6142316000848484613da1565b610dfb5760405162461bcd60e51b8152600401610c2c90615ad9565b6001600160a01b0383166142a8576142a381600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b6142cb565b816001600160a01b0316836001600160a01b0316146142cb576142cb8382614453565b6001600160a01b0382166142e257610dfb816144f0565b826001600160a01b0316826001600160a01b031614610dfb57610dfb828261459f565b6001600160a01b03821661435b5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610c2c565b6000818152600260205260409020546001600160a01b0316156143c05760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610c2c565b6143cc60008383613c01565b6001600160a01b03821660009081526003602052604081208054600192906143f5908490615ac1565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000600161446084611a2f565b61446a9190615aaa565b6000838152600760205260409020549091508082146144bd576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b60085460009061450290600190615aaa565b6000838152600960205260408120546008805493945090928490811061452a5761452a6155a2565b90600052602060002001549050806008838154811061454b5761454b6155a2565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061458357614583615cf6565b6001900381819060005260206000200160009055905550505050565b60006145aa83611a2f565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b5080546145ef906150c4565b6000825580601f106145ff575050565b601f016020900490600052602060002090810190612538919061475d565b6040518060c001604052806000815260200160608152602001606081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b828054614671906150c4565b90600052602060002090601f01602090048101928261469357600085556146d9565b82601f106146ac57805160ff19168380011785556146d9565b828001600101855582156146d9579182015b828111156146d95782518255916020019190600101906146be565b506146e592915061475d565b5090565b8280546146f5906150c4565b90600052602060002090601f01602090048101928261471757600085556146d9565b82601f106147305782800160ff198235161785556146d9565b828001600101855582156146d9579182015b828111156146d9578235825591602001919060010190614742565b5b808211156146e5576000815560010161475e565b6001600160e01b03198116811461253857600080fd5b60006020828403121561479a57600080fd5b813561366381614772565b6000602082840312156147b757600080fd5b5035919050565b60005b838110156147d95781810151838201526020016147c1565b838111156122cd5750506000910152565b600081518084526148028160208601602086016147be565b601f01601f19169290920160200192915050565b60208152600061366360208301846147ea565b60006060828403121561483b57600080fd5b50919050565b60006020828403121561485357600080fd5b81356001600160401b0381111561486957600080fd5b61328484828501614829565b6001600160a01b038116811461253857600080fd5b803561489581614875565b919050565b600080604083850312156148ad57600080fd5b82356148b881614875565b946020939093013593505050565b600080604083850312156148d957600080fd5b8235915060208301356148eb81614875565b809150509250929050565b60006060828403121561490857600080fd5b6136638383614829565b60006080828403121561483b57600080fd5b60006020828403121561493657600080fd5b81356001600160401b0381111561494c57600080fd5b61328484828501614912565b60008060006060848603121561496d57600080fd5b833561497881614875565b9250602084013561498881614875565b929592945050506040919091013590565b600060e0828403121561483b57600080fd5b6000602082840312156149bd57600080fd5b81356001600160401b038111156149d357600080fd5b61328484828501614999565b6000602082840312156149f157600080fd5b813561366381614875565b60006040828403121561483b57600080fd5b60008060608385031215614a2157600080fd5b82356001600160401b03811115614a3757600080fd5b614a4385828601614999565b925050614a5384602085016149fc565b90509250929050565b60008083601f840112614a6e57600080fd5b5081356001600160401b03811115614a8557600080fd5b602083019150836020828501011115614a9d57600080fd5b9250929050565b600080600060408486031215614ab957600080fd5b8335925060208401356001600160401b03811115614ad657600080fd5b614ae286828701614a5c565b9497909650939450505050565b600060808284031215614b0157600080fd5b6136638383614912565b60008060408385031215614b1e57600080fd5b82356001600160401b03811115614b3457600080fd5b614b4085828601614999565b95602094909401359450505050565b60008060008060608587031215614b6557600080fd5b843593506020850135925060408501356001600160401b03811115614b8957600080fd5b614b9587828801614a5c565b95989497509550505050565b6020808252825182820181905260009190848201906040850190845b81811015614be25783516001600160a01b031683529284019291840191600101614bbd565b50909695505050505050565b60008060408385031215614c0157600080fd5b50508035926020909101359150565b60008060408385031215614c2357600080fd5b82356001600160401b03811115614c3957600080fd5b614c4585828601614999565b92505060208301356148eb81614875565b60008060208385031215614c6957600080fd5b82356001600160401b03811115614c7f57600080fd5b614c8b85828601614a5c565b90969095509350505050565b60208152815160208201526000602083015160c06040840152614cbd60e08401826147ea565b90506040840151601f19848303016060850152614cda82826147ea565b91505060608401516080840152608084015160018060a01b0380821660a08601528060a08701511660c086015250508091505092915050565b8035801515811461489557600080fd5b60008060408385031215614d3657600080fd5b8235614d4181614875565b9150614a5360208401614d13565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715614d8757614d87614d4f565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614db557614db5614d4f565b604052919050565b60006001600160401b03821115614dd657614dd6614d4f565b50601f01601f191660200190565b600082601f830112614df557600080fd5b8135614e08614e0382614dbd565b614d8d565b818152846020838601011115614e1d57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215614e5057600080fd5b8435614e5b81614875565b93506020850135614e6b81614875565b92506040850135915060608501356001600160401b03811115614e8d57600080fd5b614e9987828801614de4565b91505092959194509250565b600060a0828403121561483b57600080fd5b600060208284031215614ec957600080fd5b81356001600160401b03811115614edf57600080fd5b61328484828501614ea5565b60208152815160208201526020820151604082015260006040830151610100806060850152614f1e6101208501836147ea565b9150606085015160018060a01b0380821660808701528060808801511660a0870152505060a0850151614f5c60c08601826001600160a01b03169052565b5060c085015180151560e08601525060e0850151801515858301525090949350505050565b60008060008060008060008060c0898b031215614f9d57600080fd5b88356001600160401b0380821115614fb457600080fd5b614fc08c838d01614a5c565b909a50985060208b0135915080821115614fd957600080fd5b50614fe68b828c01614a5c565b9097509550506040890135614ffa81614875565b9350606089013561500a81614875565b9250608089013561501a81614875565b915060a089013561502a81614875565b809150509295985092959890939650565b6000806040838503121561504e57600080fd5b823561505981614875565b915060208301356148eb81614875565b60008060006040848603121561507e57600080fd5b83356001600160401b038082111561509557600080fd5b6150a187838801614999565b945060208601359150808211156150b757600080fd5b50614ae286828701614a5c565b600181811c908216806150d857607f821691505b6020821081141561483b57634e487b7160e01b600052602260045260246000fd5b6000808335601e1984360301811261511057600080fd5b8301803591506001600160401b0382111561512a57600080fd5b602001915036819003821315614a9d57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03868116825285166020820152608060408201819052600090615195908301858761513f565b90508260608301529695505050505050565b8581526001600160a01b0385166020820152608060408201819052600090615195908301858761513f565b6000602082840312156151e457600080fd5b815161366381614875565b60006101008b83528a6020840152896040840152806060840152615216818401898b61513f565b6001600160a01b03978816608085015295871660a084015250509190931660c082015260e0019190915295945050505050565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60006000198214156152c4576152c461529a565b5060010190565b6000808335601e198436030181126152e257600080fd5b83016020810192503590506001600160401b0381111561530157600080fd5b803603831315614a9d57600080fd5b80358252600061532360208301836152cb565b60e0602086015261533860e08601828461513f565b915050604083013561534981614875565b6001600160a01b03818116604087015261536660608601866152cb565b9250868403606088015261537b84848361513f565b9350506080850135915061538e82614875565b1660808501526153a160a08401846152cb565b85830360a08701526153b483828461513f565b925050506153c460c08401614d13565b151560c08501528091505092915050565b60c0815260006153e860c0830188615310565b602083810197909752604083019590955250606081019290925281830360808301526000835260a09091015201919050565b60006020828403121561542c57600080fd5b5051919050565b60c08152600061544660c0830189615310565b876020840152866040840152856060840152828103608084015261546a81866147ea565b9150508260a0830152979650505050505050565b83815260406020820152600061213a60408301848661513f565b6060815283356060820152600060208501356154b381614875565b6001600160a01b038181166080850152604087013560a08501526154da60608801886152cb565b9250608060c08601526154f160e08601848361513f565b96909116602085015250505060400152919050565b600080835481600182811c91508083168061552257607f831692505b602080841082141561554257634e487b7160e01b86526022600452602486fd5b818015615556576001811461556757615594565b60ff19861689528489019650615594565b60008a81526020902060005b8681101561558c5781548b820152908501908301615573565b505084890196505b509498975050505050505050565b634e487b7160e01b600052603260045260246000fd5b606081528335606082015260006155d260208601866152cb565b6080808501526155e660e08501828461513f565b915050604086013560a084015261560060608701876152cb565b848303605f190160c086015261561783828461513f565b6001600160a01b039790971660208601525050505060400152919050565b85815284602082015260806040820152600061519560808301858761513f565b6001600160a01b038781168252602082018790528516604082015260a060608201819052600090615689908301858761513f565b9050826080830152979650505050505050565b8335815260c0810160208501356156b281614875565b6001600160a01b039081166020840152604086810135908401526060958601359583019590955292909316608084015260a09092019190915290565b8335815260a08101602085013561570481614875565b6001600160a01b0390811660208401526040958601359583019590955292909316606084015260809092019190915290565b606081528335606082015260208401356080820152604084013560a0820152600061576460608601866152cb565b608060c085015261577960e08501828461513f565b6001600160a01b0396909616602085015250505060400152919050565b8183823760009101908152919050565b8581526080602082015260006157c060808301868861513f565b604083019490945250606001529392505050565b888152876020820152600060018060a01b03808916604084015260e0606084015261580360e08401888a61513f565b951660808301525060a081019290925260c09091015295945050505050565b60a08152853560a0820152602086013560c0820152604086013560e08201526060860135610100820152600061585b60808801886152cb565b60a06101208501526158726101408501828461513f565b6001600160a01b0398891660208601529690971660408401525050606081019290925260809091015292915050565b6060815283356060820152600060208501356158bc81614875565b60018060a01b038082166080850152604087013560a0850152606087013560c08501526158ec60808801886152cb565b925060a060e08601526154f16101008601848361513f565b600060a0823603121561591657600080fd5b61591e614d65565b6159278361488a565b815260208301356001600160401b038082111561594357600080fd5b61594f36838701614de4565b6020840152604085013591508082111561596857600080fd5b61597436838701614de4565b60408401526159856060860161488a565b6060840152608085013591508082111561599e57600080fd5b506159ab36828601614de4565b60808301525092915050565b86815285602082015260018060a01b038516604082015260a06060820152600061568960a08301858761513f565b6000602082840312156159f757600080fd5b81516001600160401b03811115615a0d57600080fd5b8201601f81018413615a1e57600080fd5b8051615a2c614e0382614dbd565b818152856020838501011115615a4157600080fd5b61213a8260208301602086016147be565b60608152833560608201526000615a6c60208601866152cb565b60606080850152615a8160c08501828461513f565b60409788013560a08601526001600160a01b039690961660208501525050509092019190915290565b600082821015615abc57615abc61529a565b500390565b60008219821115615ad457615ad461529a565b500190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60a08152600060018060a01b038088511660a0840152602088015160a060c0850152615b5b6101408501826147ea565b90506040890151609f19808684030160e0870152615b7983836147ea565b92508360608c01511661010087015260808b0151935080868403016101208701525050615ba681836147ea565b92505050615bb8602083018715159052565b8460408301528360608301528260808301529695505050505050565b606081526000615be860608301878961513f565b8281036020840152615bfb81868861513f565b9150508260408301529695505050505050565b60006101008a8352896020840152886040840152806060840152615c34818401896147ea565b6001600160a01b03978816608085015295871660a084015250509190931660c082015260e00191909152949350505050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090615c99908301846147ea565b9695505050505050565b600060208284031215615cb557600080fd5b815161366381614772565b6000816000190483118215151615615cda57615cda61529a565b500290565b600081615cee57615cee61529a565b506000190190565b634e487b7160e01b600052603160045260246000fdfea264697066735822122049a2aac7ed482b326f8730c3f913b716a6cc3a5c6af76cb8b9fc89fada0a57b764736f6c634300080a0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCaller) GetCharacter(opts *bind.CallOpts, characterId *big.Int) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCharacter", characterId)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacter(&_Contract.CallOpts, characterId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCallerSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacter(&_Contract.CallOpts, characterId)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCaller) GetCharacterByHandle(opts *bind.CallOpts, handle string) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCharacterByHandle", handle)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacterByHandle(&_Contract.CallOpts, handle)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCallerSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacterByHandle(&_Contract.CallOpts, handle)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contract *ContractCaller) GetCharacterUri(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCharacterUri", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contract *ContractSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetCharacterUri(&_Contract.CallOpts, characterId)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contract *ContractCallerSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetCharacterUri(&_Contract.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contract *ContractCaller) GetHandle(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getHandle", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contract *ContractSession) GetHandle(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetHandle(&_Contract.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contract *ContractCallerSession) GetHandle(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetHandle(&_Contract.CallOpts, characterId)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contract *ContractCaller) GetLinkModule4Address(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinkModule4Address", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contract *ContractSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Address(&_Contract.CallOpts, account)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contract *ContractCallerSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Address(&_Contract.CallOpts, account)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetLinkModule4ERC721(opts *bind.CallOpts, tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinkModule4ERC721", tokenAddress, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4ERC721(&_Contract.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4ERC721(&_Contract.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetLinkModule4Linklist(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinkModule4Linklist", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Linklist(&_Contract.CallOpts, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Linklist(&_Contract.CallOpts, tokenId)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contract *ContractCaller) GetLinklistContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contract *ContractSession) GetLinklistContract() (common.Address, error) {
	return _Contract.Contract.GetLinklistContract(&_Contract.CallOpts)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contract *ContractCallerSession) GetLinklistContract() (common.Address, error) {
	return _Contract.Contract.GetLinklistContract(&_Contract.CallOpts)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contract *ContractCaller) GetLinklistId(opts *bind.CallOpts, characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistId", characterId, linkType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contract *ContractSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Contract.Contract.GetLinklistId(&_Contract.CallOpts, characterId, linkType)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contract *ContractCallerSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Contract.Contract.GetLinklistId(&_Contract.CallOpts, characterId, linkType)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contract *ContractCaller) GetLinklistType(opts *bind.CallOpts, linkListId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistType", linkListId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contract *ContractSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Contract.Contract.GetLinklistType(&_Contract.CallOpts, linkListId)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contract *ContractCallerSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Contract.Contract.GetLinklistType(&_Contract.CallOpts, linkListId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) GetLinklistUri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistUri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contract *ContractSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetLinklistUri(&_Contract.CallOpts, tokenId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetLinklistUri(&_Contract.CallOpts, tokenId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contract *ContractCaller) GetNote(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNote", characterId, noteId)

	if err != nil {
		return *new(DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNote)).(*DataTypesNote)

	return out0, err

}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contract *ContractSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Contract.Contract.GetNote(&_Contract.CallOpts, characterId, noteId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contract *ContractCallerSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Contract.Contract.GetNote(&_Contract.CallOpts, characterId, noteId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Contract *ContractCaller) GetOperator(opts *bind.CallOpts, characterId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperator", characterId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Contract *ContractSession) GetOperator(characterId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetOperator(&_Contract.CallOpts, characterId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Contract *ContractCallerSession) GetOperator(characterId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetOperator(&_Contract.CallOpts, characterId)
}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contract *ContractCaller) GetOperators(opts *bind.CallOpts, characterId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperators", characterId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contract *ContractSession) GetOperators(characterId *big.Int) ([]common.Address, error) {
	return _Contract.Contract.GetOperators(&_Contract.CallOpts, characterId)
}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contract *ContractCallerSession) GetOperators(characterId *big.Int) ([]common.Address, error) {
	return _Contract.Contract.GetOperators(&_Contract.CallOpts, characterId)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contract *ContractCaller) GetPrimaryCharacterId(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPrimaryCharacterId", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contract *ContractSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetPrimaryCharacterId(&_Contract.CallOpts, account)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contract *ContractCallerSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetPrimaryCharacterId(&_Contract.CallOpts, account)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contract *ContractCaller) GetRevision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRevision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contract *ContractSession) GetRevision() (*big.Int, error) {
	return _Contract.Contract.GetRevision(&_Contract.CallOpts)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contract *ContractCallerSession) GetRevision() (*big.Int, error) {
	return _Contract.Contract.GetRevision(&_Contract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0xa48047ba.
//
// Solidity: function isOperator(uint256 characterId, address operator) view returns(bool)
func (_Contract *ContractCaller) IsOperator(opts *bind.CallOpts, characterId *big.Int, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOperator", characterId, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xa48047ba.
//
// Solidity: function isOperator(uint256 characterId, address operator) view returns(bool)
func (_Contract *ContractSession) IsOperator(characterId *big.Int, operator common.Address) (bool, error) {
	return _Contract.Contract.IsOperator(&_Contract.CallOpts, characterId, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0xa48047ba.
//
// Solidity: function isOperator(uint256 characterId, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsOperator(characterId *big.Int, operator common.Address) (bool, error) {
	return _Contract.Contract.IsOperator(&_Contract.CallOpts, characterId, operator)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contract *ContractCaller) IsPrimaryCharacter(opts *bind.CallOpts, characterId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isPrimaryCharacter", characterId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contract *ContractSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Contract.Contract.IsPrimaryCharacter(&_Contract.CallOpts, characterId)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contract *ContractCallerSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Contract.Contract.IsPrimaryCharacter(&_Contract.CallOpts, characterId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractSession) Resolver() (common.Address, error) {
	return _Contract.Contract.Resolver(&_Contract.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractCallerSession) Resolver() (common.Address, error) {
	return _Contract.Contract.Resolver(&_Contract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contract *ContractCaller) TokenURI(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenURI", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contract *ContractSession) TokenURI(characterId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, characterId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contract *ContractCallerSession) TokenURI(characterId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, characterId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// AddOperator is a paid mutator transaction binding the contract method 0x7d46a713.
//
// Solidity: function addOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactor) AddOperator(opts *bind.TransactOpts, characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addOperator", characterId, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x7d46a713.
//
// Solidity: function addOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractSession) AddOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOperator(&_Contract.TransactOpts, characterId, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x7d46a713.
//
// Solidity: function addOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactorSession) AddOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOperator(&_Contract.TransactOpts, characterId, operator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, tokenId)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractTransactor) CreateCharacter(opts *bind.TransactOpts, vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createCharacter", vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateCharacter(&_Contract.TransactOpts, vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateCharacter(&_Contract.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactor) CreateThenLinkCharacter(opts *bind.TransactOpts, vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createThenLinkCharacter", vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Contract *ContractSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateThenLinkCharacter(&_Contract.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateThenLinkCharacter(&_Contract.TransactOpts, vars)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactor) DeleteNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deleteNote", characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeleteNote(&_Contract.TransactOpts, characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactorSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeleteNote(&_Contract.TransactOpts, characterId, noteId)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Contract *ContractSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Contract *ContractTransactorSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkAddress(opts *bind.TransactOpts, vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkAddress", vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAddress(&_Contract.TransactOpts, vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAddress(&_Contract.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkAnyUri(opts *bind.TransactOpts, vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkAnyUri", vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAnyUri(&_Contract.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAnyUri(&_Contract.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkCharacter(opts *bind.TransactOpts, vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkCharacter", vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.LinkCharacter(&_Contract.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.LinkCharacter(&_Contract.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkERC721(opts *bind.TransactOpts, vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkERC721", vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.LinkERC721(&_Contract.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.LinkERC721(&_Contract.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkLinklist(opts *bind.TransactOpts, vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkLinklist", vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.LinkLinklist(&_Contract.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.LinkLinklist(&_Contract.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkNote(opts *bind.TransactOpts, vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkNote", vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.LinkNote(&_Contract.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.LinkNote(&_Contract.TransactOpts, vars)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactor) LockNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockNote", characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockNote(&_Contract.TransactOpts, characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactorSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockNote(&_Contract.TransactOpts, characterId, noteId)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Contract *ContractTransactor) MintNote(opts *bind.TransactOpts, vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintNote", vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Contract *ContractSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contract.Contract.MintNote(&_Contract.TransactOpts, vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Contract *ContractTransactorSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contract.Contract.MintNote(&_Contract.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Contract *ContractTransactor) PostNote(opts *bind.TransactOpts, vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote", vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Contract *ContractSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contract.Contract.PostNote(&_Contract.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contract.Contract.PostNote(&_Contract.TransactOpts, vars)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Address(opts *bind.TransactOpts, noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Address", noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Contract *ContractSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Address(&_Contract.TransactOpts, noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Address(&_Contract.TransactOpts, noteData, ethAddress)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Contract *ContractTransactor) PostNote4AnyUri(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4AnyUri", postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Contract *ContractSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4AnyUri(&_Contract.TransactOpts, postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4AnyUri(&_Contract.TransactOpts, postNoteData, uri)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Character(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Character", postNoteData, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Contract *ContractSession) PostNote4Character(postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Character(&_Contract.TransactOpts, postNoteData, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Character(postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Character(&_Contract.TransactOpts, postNoteData, toCharacterId)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Contract *ContractTransactor) PostNote4ERC721(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4ERC721", postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Contract *ContractSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4ERC721(&_Contract.TransactOpts, postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4ERC721(&_Contract.TransactOpts, postNoteData, erc721)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Linklist(opts *bind.TransactOpts, noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Linklist", noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Contract *ContractSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Linklist(&_Contract.TransactOpts, noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Linklist(&_Contract.TransactOpts, noteData, toLinklistId)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Note(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Note", postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Contract *ContractSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Note(&_Contract.TransactOpts, postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Note(&_Contract.TransactOpts, postNoteData, note)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x0d18d07a.
//
// Solidity: function removeOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactor) RemoveOperator(opts *bind.TransactOpts, characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeOperator", characterId, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x0d18d07a.
//
// Solidity: function removeOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractSession) RemoveOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, characterId, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x0d18d07a.
//
// Solidity: function removeOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactorSession) RemoveOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, characterId, operator)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contract *ContractTransactor) SetCharacterUri(opts *bind.TransactOpts, characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCharacterUri", characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contract *ContractSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetCharacterUri(&_Contract.TransactOpts, characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contract *ContractTransactorSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetCharacterUri(&_Contract.TransactOpts, characterId, newUri)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contract *ContractTransactor) SetHandle(opts *bind.TransactOpts, characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setHandle", characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contract *ContractSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contract.Contract.SetHandle(&_Contract.TransactOpts, characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contract *ContractTransactorSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contract.Contract.SetHandle(&_Contract.TransactOpts, characterId, newHandle)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Address(opts *bind.TransactOpts, vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Address", vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Address(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Address(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Character(opts *bind.TransactOpts, vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Character", vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Character(vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Character(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Character(vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Character(&_Contract.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4ERC721(opts *bind.TransactOpts, vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4ERC721", vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4ERC721(&_Contract.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4ERC721(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Linklist(opts *bind.TransactOpts, vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Linklist", vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Linklist(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Linklist(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Note(opts *bind.TransactOpts, vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Note", vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Note(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Note(&_Contract.TransactOpts, vars)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contract *ContractTransactor) SetLinklistUri(opts *bind.TransactOpts, linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinklistUri", linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contract *ContractSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.SetLinklistUri(&_Contract.TransactOpts, linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contract *ContractTransactorSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.SetLinklistUri(&_Contract.TransactOpts, linklistId, uri)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetMintModule4Note(opts *bind.TransactOpts, vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMintModule4Note", vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetMintModule4Note(&_Contract.TransactOpts, vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetMintModule4Note(&_Contract.TransactOpts, vars)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contract *ContractTransactor) SetNoteUri(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setNoteUri", characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contract *ContractSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetNoteUri(&_Contract.TransactOpts, characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contract *ContractTransactorSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetNoteUri(&_Contract.TransactOpts, characterId, noteId, newUri)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactor) SetOperator(opts *bind.TransactOpts, characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOperator", characterId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractSession) SetOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperator(&_Contract.TransactOpts, characterId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactorSession) SetOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperator(&_Contract.TransactOpts, characterId, operator)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contract *ContractTransactor) SetPrimaryCharacterId(opts *bind.TransactOpts, characterId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPrimaryCharacterId", characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contract *ContractSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPrimaryCharacterId(&_Contract.TransactOpts, characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contract *ContractTransactorSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPrimaryCharacterId(&_Contract.TransactOpts, characterId)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contract *ContractTransactor) SetSocialToken(opts *bind.TransactOpts, characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSocialToken", characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contract *ContractSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSocialToken(&_Contract.TransactOpts, characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contract *ContractTransactorSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSocialToken(&_Contract.TransactOpts, characterId, tokenAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkAddress(opts *bind.TransactOpts, vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkAddress", vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAddress(&_Contract.TransactOpts, vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAddress(&_Contract.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkAnyUri(opts *bind.TransactOpts, vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkAnyUri", vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAnyUri(&_Contract.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAnyUri(&_Contract.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkCharacter(opts *bind.TransactOpts, vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkCharacter", vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkCharacter(&_Contract.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkCharacter(&_Contract.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkERC721(opts *bind.TransactOpts, vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkERC721", vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkERC721(&_Contract.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkERC721(&_Contract.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkLinklist(opts *bind.TransactOpts, vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkLinklist", vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkLinklist(&_Contract.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkLinklist(&_Contract.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkNote(opts *bind.TransactOpts, vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkNote", vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkNote(&_Contract.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkNote(&_Contract.TransactOpts, vars)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contract contract.
type ContractApprovalForAllIterator struct {
	Event *ContractApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApprovalForAll represents a ApprovalForAll event raised by the Contract contract.
type ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalForAllIterator{contract: _Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApprovalForAll)
				if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) ParseApprovalForAll(log types.Log) (*ContractApprovalForAll, error) {
	event := new(ContractApprovalForAll)
	if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
