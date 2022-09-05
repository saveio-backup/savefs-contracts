// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// ChangeInitPosParam is an auto generated low-level Go binding around an user-defined struct.
type ChangeInitPosParam struct {
	PeerPubKey string
	Address    common.Address
	Pos        uint64
}

// DNSNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type DNSNodeInfo struct {
	WalletAddr  common.Address
	IP          []byte
	Port        []byte
	InitDeposit uint64
	PeerPubKey  string
}

// HeaderInfo is an auto generated low-level Go binding around an user-defined struct.
type HeaderInfo struct {
	Header      []byte
	HeaderOwner common.Address
	Desc        []byte
	BlockHeight *big.Int
	TTL         uint64
}

// NameInfo is an auto generated low-level Go binding around an user-defined struct.
type NameInfo struct {
	Type        uint64
	Header      []byte
	URL         []byte
	Name        []byte
	NameOwner   common.Address
	Desc        []byte
	BlockHeight *big.Int
	TTL         uint64
}

// PeerPoolItem is an auto generated low-level Go binding around an user-defined struct.
type PeerPoolItem struct {
	PeerPubKey    string
	WalletAddress common.Address
	Status        uint8
	TotalInitPos  uint64
}

// QuitNodeParam is an auto generated low-level Go binding around an user-defined struct.
type QuitNodeParam struct {
	PeerPubKey string
	Address    common.Address
}

// ReqInfo is an auto generated low-level Go binding around an user-defined struct.
type ReqInfo struct {
	Header []byte
	URL    []byte
	Owner  common.Address
}

// RequestHeader is an auto generated low-level Go binding around an user-defined struct.
type RequestHeader struct {
	Header    []byte
	NameOwner common.Address
	Desc      []byte
	DesireTTL uint64
}

// RequestName is an auto generated low-level Go binding around an user-defined struct.
type RequestName struct {
	Type      uint64
	Header    []byte
	URL       []byte
	Name      []byte
	NameOwner common.Address
	Desc      []byte
	DesireTTL uint64
}

// TransferInfo is an auto generated low-level Go binding around an user-defined struct.
type TransferInfo struct {
	Header []byte
	URL    []byte
	From   common.Address
	To     common.Address
}

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// UnRegisterCandidateParam is an auto generated low-level Go binding around an user-defined struct.
type UnRegisterCandidateParam struct {
	PeerPubKey string
	Address    common.Address
}

// UpdateNodeParam is an auto generated low-level Go binding around an user-defined struct.
type UpdateNodeParam struct {
	WalletAddr common.Address
	IP         []byte
	Port       []byte
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstUserSpaceOperationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorOpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UserspaceChangeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserspaceDeleteError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Pos\",\"type\":\"uint64\"}],\"internalType\":\"structChangeInitPosParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"AddInitPos\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"ApproveDNSCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"CreateDefaultUrl\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"DNSNodeReg\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"DelDNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"DelHeader\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"DnsInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllDnsNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetDNSNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"GetHeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"HeaderOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structHeaderInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"GetName\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structNameInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"GetPeerPoolItem\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"WalletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"TotalInitPos\",\"type\":\"uint64\"}],\"internalType\":\"structPeerPoolItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPeerPoolMap\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"WalletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"TotalInitPos\",\"type\":\"uint64\"}],\"internalType\":\"structPeerPoolItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPluginList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structNameInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"GetUrl\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"internalType\":\"structQuitNodeParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"QuitNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Pos\",\"type\":\"uint64\"}],\"internalType\":\"structChangeInitPosParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"ReduceInitPos\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestHeader\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"RegisterHeader\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestName\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"RegisterName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"RejectDNSCandidate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"}],\"internalType\":\"structTransferInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"TransferHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"}],\"internalType\":\"structTransferInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"TransferName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"internalType\":\"structUnRegisterCandidateParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"UnRegDNSNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"}],\"internalType\":\"structUpdateNodeParam\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"UpdateDNSNodesInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestName\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"UpdateName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b2\",\"type\":\"bytes\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"toBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615010806100206000396000f3fe6080604052600436106101b75760003560e01c80636b9ceb53116100ec578063ba47d33d1161008a578063d72b6be611610064578063d72b6be6146104b5578063dbc8ef9c146104d5578063eafafaf7146104e8578063fd11c4c6146104fb57600080fd5b8063ba47d33d14610453578063ccfd72f614610473578063d68414701461049557600080fd5b80638129fc1c116100c65780638129fc1c146103c457806387dacef2146103d95780638b79782b14610406578063919064eb1461042657600080fd5b80636b9ceb53146103715780636e91c5681461039e578063809bae57146103b157600080fd5b8063370d7104116101595780634f2ffbba116101335780634f2ffbba14610309578063513728191461031c5780635f3376f31461032f5780636839a7c61461034f57600080fd5b8063370d7104146102c35780633e4e5e54146102d6578063465efc9a146102f657600080fd5b8063240ce9f911610195578063240ce9f9146102295780632a5da31e146102565780632c261f4214610283578063326cf61c146102a357600080fd5b80630e580e79146101bc5780631654a894146101e7578063228a09b214610209575b600080fd5b3480156101c857600080fd5b506101d161050e565b6040516101de9190614c30565b60405180910390f35b3480156101f357600080fd5b50610207610202366004614503565b610601565b005b34801561021557600080fd5b506102076102243660046142fc565b6106fb565b34801561023557600080fd5b506102496102443660046142fc565b610775565b6040516101de9190614dc7565b34801561026257600080fd5b50610276610271366004614330565b6107aa565b6040516101de9190614c41565b34801561028f57600080fd5b5061020761029e3660046143ff565b610804565b3480156102af57600080fd5b506102766102be3660046142c0565b610941565b6102076102d136600461449b565b61096a565b3480156102e257600080fd5b506102076102f13660046144cf565b610a5a565b6102076103043660046143cb565b610b67565b610207610317366004614433565b610c97565b61020761032a366004614397565b610d00565b34801561033b57600080fd5b5061027661034a366004614330565b610e3e565b34801561035b57600080fd5b50610364610eb5565b6040516101de9190614c0e565b34801561037d57600080fd5b5061039161038c366004614433565b610faa565b6040516101de9190614db6565b6102076103ac366004614433565b611035565b6102076103bf366004614467565b611261565b3480156103d057600080fd5b506102076113c5565b3480156103e557600080fd5b506103f96103f43660046142a2565b611531565b6040516101de9190614d94565b34801561041257600080fd5b506102076104213660046144cf565b611566565b34801561043257600080fd5b50610446610441366004614433565b6118b3565b6040516101de9190614da5565b34801561045f57600080fd5b5061020761046e3660046142a2565b611a80565b34801561047f57600080fd5b50610488611c65565b6040516101de9190614c1f565b3480156104a157600080fd5b506102076104b03660046143ff565b611d93565b3480156104c157600080fd5b506102766104d03660046142fc565b611e8b565b6102076104e33660046142fc565b611ee0565b6102076104f6366004614397565b611f7d565b61020761050936600461449b565b612072565b60606000600a600201546001600160401b0381111561053d57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561059057816020015b6040805160808101825260608082526000602080840182905293830181905290820152825260001990920191018161055b5790505b509050600061059f600a612507565b90505b600b548110156105fb5760006105b9600a83612522565b915050808383815181106105dd57634e487b7160e01b600052603260045260246000fd5b6020908102919091010152506105f4600a8261270b565b90506105a2565b50919050565b60008160200151511161062f5760405162461bcd60e51b815260040161062690614ca3565b60405180910390fd5b6000816040015151116106545760405162461bcd60e51b815260040161062690614cb3565b6000610661600d3361277b565b80519091506001600160a01b0316331461068d5760405162461bcd60e51b815260040161062690614cd3565b60208083015190820152604080830151908201526106ad600d33836129ae565b507f602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf826020015183604001513384606001516040516106ef9493929190614c52565b60405180910390a15050565b6000610708600a83612b23565b9050600181606001516001600160401b031610156107385760405162461bcd60e51b815260040161062690614d74565b604081015160ff161561075d5760405162461bcd60e51b815260040161062690614cd3565b60026040820152610770600a8383612c43565b505050565b6040805160808101825260608082526000602083018190529282018390528101919091526107a4600a83612b23565b92915050565b606060006107ed846040518060400160405280600381526020017f3a2f2f0000000000000000000000000000000000000000000000000000000000815250610e3e565b905060006107fb8285610e3e565b95945050505050565b60208101516001600160a01b031633146108305760405162461bcd60e51b815260040161062690614d64565b805160009061084190600a90612b23565b604081015190915060ff16156108695760405162461bcd60e51b815260040161062690614cd3565b81602001516001600160a01b031681602001516001600160a01b0316146108a25760405162461bcd60e51b815260040161062690614d64565b81602001516001600160a01b03166108fc82606001516001600160401b03169081150290604051600060405180830381858888f193505050501580156108ec573d6000803e3d6000fd5b5081516108fb90600a90612de1565b50602082015161090d90600d90612ee1565b507f7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d82602001516040516106ef9190614b7f565b6060816040516020016109549190614b52565b6040516020818303038152906040529050919050565b600061097e82602001518360400151610e3e565b9050600061098d600683612ff8565b60808101519091506001600160a01b031633146109bc5760405162461bcd60e51b815260040161062690614d64565b82516001600160401b0390811682526060808501519083015260a0808501519083015260c08401511660e08201526109f5436001614e1e565b60c0820152610a066006838361331e565b507f0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe081508360800151610a3f856020015186604001516107aa565b604051610a4d929190614bba565b60405180910390a1505050565b6000610a6e82600001518360200151610e3e565b90506000610a7d600683612ff8565b905082604001516001600160a01b031681608001516001600160a01b031614610ab85760405162461bcd60e51b815260040161062690614d64565b60608301516001600160a01b03166080820152610ad6436001614e1e565b60c0820181905260e08201514391610af6916001600160401b0316614e1e565b610b009190614e69565b6001600160401b031660e0820152610b1a6006838361331e565b507fa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f83604001518460600151610b58866000015187602001516107aa565b604051610a4d93929190614b8d565b600081606001516001600160401b031611610b945760405162461bcd60e51b815260040161062690614d54565b80606001516001600160401b0316341015610bc15760405162461bcd60e51b815260040161062690614c93565b80516001600160a01b03163314610bea5760405162461bcd60e51b815260040161062690614d64565b604080516080808201835260608083526000602084018181529484018190528184019081529185018051845285516001600160a01b03169094528401516001600160401b031690529051610c4190600a9083612c43565b508151610c5190600d90846129ae565b507f602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf82602001518360400151846000015185606001516040516106ef9493929190614c52565b6000610cab82600001518360200151610e3e565b90506000610cba600683612ff8565b60045460808201519192506001600160a01b03918216911614610cef5760405162461bcd60e51b815260040161062690614d44565b610cfa60068361345e565b50505050565b600081604001516001600160401b031611610d2d5760405162461bcd60e51b815260040161062690614d84565b80604001516001600160401b0316341015610d5a5760405162461bcd60e51b815260040161062690614c93565b60208101516001600160a01b03163314610d865760405162461bcd60e51b815260040161062690614d64565b8051600090610d9790600a90612b23565b905081602001516001600160a01b031681602001516001600160a01b031614610dd25760405162461bcd60e51b815260040161062690614d64565b604081015160ff16600214801590610df05750604081015160ff1615155b15610e0d5760405162461bcd60e51b815260040161062690614cc3565b816040015181606001818151610e239190614e36565b6001600160401b0316905250815161077090600a9083612c43565b81518151606091906000610e528284614e1e565b6001600160401b03811115610e7757634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015610ea1576020820181803683370190505b509283019590955281019290925250919050565b60606000600d600201546001600160401b03811115610ee457634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610f3f57816020015b6040805160a0810182526000808252606060208301819052928201839052828201526080810191909152815260200190600190039081610f025790505b5090506000610f4e600d613541565b90505b600e548110156105fb576000610f68600d8361354f565b91505080838381518110610f8c57634e487b7160e01b600052603260045260246000fd5b602090810291909101015250610fa3600d826137ce565b9050610f51565b61100d60405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b600061102183600001518460200151610e3e565b905061102e600682612ff8565b9392505050565b60006005826000015160405161104b9190614b67565b90815260200160405180910390206040518060a001604052908160008201805461107490614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546110a090614f01565b80156110ed5780601f106110c2576101008083540402835291602001916110ed565b820191906000526020600020905b8154815290600101906020018083116110d057829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161111d90614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461114990614f01565b80156111965780601f1061116b57610100808354040283529160200191611196565b820191906000526020600020905b81548152906001019060200180831161117957829003601f168201915b505050918352505060038201546020808301919091526004928301546001600160401b03166040909201919091529054908201519192506001600160a01b039182169116146111f75760405162461bcd60e51b815260040161062690614d44565b815160405160059161120891614b67565b90815260405190819003602001902060006112238282613c82565b6001820180546001600160a01b0319169055611243600283016000613c82565b5060006003820155600401805467ffffffffffffffff191690555050565b6112a56040518060a001604052806060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b815181526020808301516001600160a01b031690820152604080830151908201526112d1436001614e1e565b606082015260006080820152815160405182916005916112f19190614b67565b9081526020016040518091039020600082015181600001908051906020019061131b929190613cbc565b506020828101516001830180546001600160a01b0319166001600160a01b039092169190911790556040830151805161135a9260028501920190613cbc565b50606082015160038201556080909101516004909101805467ffffffffffffffff19166001600160401b03909216919091179055602082015182516040517f356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869926106ef929091614bba565b600054610100900460ff166113e05760005460ff16156113e4565b303b155b6114005760405162461bcd60e51b815260040161062690614ce3565b600054610100900460ff16158015611422576000805461ffff19166101011790555b600080547fffffffffffff000000000000000000000000000000000000000000000000ffff167202000000000000000100000000000000000000179055600180547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016680800000000000000041790556040805180820190915260038082527f647370000000000000000000000000000000000000000000000000000000000060209092019182526114d691600291613cbc565b5060408051808201909152600a8082527f6473702d706c7567696e00000000000000000000000000000000000000000000602090920191825261151b91600391613cbc565b50801561152e576000805461ff00191690555b50565b6040805160a08101825260008082526060602083018190529282018390528282015260808101919091526107a4600d8361277b565b60006005826000015160405161157c9190614b67565b90815260200160405180910390206040518060a00160405290816000820180546115a590614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546115d190614f01565b801561161e5780601f106115f35761010080835404028352916020019161161e565b820191906000526020600020905b81548152906001019060200180831161160157829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161164e90614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461167a90614f01565b80156116c75780601f1061169c576101008083540402835291602001916116c7565b820191906000526020600020905b8154815290600101906020018083116116aa57829003601f168201915b505050918352505060038201546020808301919091526004909201546001600160401b0316604091820152840151908201519192506001600160a01b039182169116146117265760405162461bcd60e51b815260040161062690614d64565b600043826060015183608001516001600160401b03166117469190614e1e565b116117535750600061177e565b43826060015183608001516001600160401b03166117719190614e1e565b61177b9190614e69565b90505b8251604051600590611791908390614b67565b908152602001604051809103902060000190805190602001906117b5929190613cbc565b50606083015183516040516005916117cc91614b67565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555043600161180e9190614e1e565b835160405160059161181f91614b67565b908152604051908190036020018120600301919091558351829160059161184591614b67565b9081526040805191829003602001822060040180546001600160401b039490941667ffffffffffffffff199094169390931790925590840151606085015185517f178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca93610a4d93929190614b8d565b6118f76040518060a001604052806060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b815160405160059161190891614b67565b90815260200160405180910390206040518060a001604052908160008201805461193190614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461195d90614f01565b80156119aa5780601f1061197f576101008083540402835291602001916119aa565b820191906000526020600020905b81548152906001019060200180831161198d57829003601f168201915b505050918352505060018201546001600160a01b031660208201526002820180546040909201916119da90614f01565b80601f0160208091040260200160405190810160405280929190818152602001828054611a0690614f01565b8015611a535780601f10611a2857610100808354040283529160200191611a53565b820191906000526020600020905b815481529060010190602001808311611a3657829003601f168201915b5050509183525050600382015460208201526004909101546001600160401b031660409091015292915050565b600480546001600160a01b0319166001600160a01b0383161790556040805160a08101825260608082526000602083018190529282018190528101829052608081019190915260028054611ad390614f01565b80601f0160208091040260200160405190810160405280929190818152602001828054611aff90614f01565b8015611b4c5780601f10611b2157610100808354040283529160200191611b4c565b820191906000526020600020905b815481529060010190602001808311611b2f57829003601f168201915b50505091835250506001600160a01b038216602080830191909152604080518082018252601581527f7265736572766564206473702070726f746f636f6c000000000000000000000092810192909252808301919091526000606083018190526080830152815190518291600591611bc49190614b67565b90815260200160405180910390206000820151816000019080519060200190611bee929190613cbc565b506020828101516001830180546001600160a01b0319166001600160a01b0390921691909117905560408301518051611c2d9260028501920190613cbc565b50606082015160038201556080909101516004909101805467ffffffffffffffff19166001600160401b039092169190911790555050565b606060006006600201546001600160401b03811115611c9457634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611d2857816020015b611d1560405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b815260200190600190039081611cb25790505b5090506000611d37600661383d565b90505b6007548110156105fb576000611d5160068361384b565b91505080838381518110611d7557634e487b7160e01b600052603260045260246000fd5b602090810291909101015250611d8c600682613c12565b9050611d3a565b60208101516001600160a01b03163314611dbf5760405162461bcd60e51b815260040161062690614d64565b8051600090611dd090600a90612b23565b905081602001516001600160a01b031681602001516001600160a01b031614611e0b5760405162461bcd60e51b815260040161062690614d64565b604081015160ff16600214801590611e295750604081015160ff1615155b15611e465760405162461bcd60e51b815260040161062690614cc3565b604081015160ff1660021415611e625760036040820152611e6a565b600460408201525b8151611e7990600a9083612c43565b50602082015161077090600d90612ee1565b60606107a4600283604051611ea09190614b67565b602060405180830381855afa158015611ebd573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906102be91906142de565b6000611eed600a83612b23565b604081015190915060ff1615611f155760405162461bcd60e51b815260040161062690614cd3565b80602001516001600160a01b03166108fc82606001516001600160401b03169081150290604051600060405180830381858888f19350505050158015611f5f573d6000803e3d6000fd5b50611f6b600a83612de1565b50602081015161077090600d90612ee1565b600081604001516001600160401b031611611faa5760405162461bcd60e51b815260040161062690614d84565b60208101516001600160a01b03163314611fd65760405162461bcd60e51b815260040161062690614d64565b8051600090611fe790600a90612b23565b905081602001516001600160a01b031681602001516001600160a01b0316146120225760405162461bcd60e51b815260040161062690614d64565b81604001516001600160401b031681606001516001600160401b0316101561205c5760405162461bcd60e51b815260040161062690614d74565b816040015181606001818151610e239190614e84565b60048160600151511161208457600080fd5b6120e760405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b60005482516001600160401b03908116620100009092041614156121fb57600081526002805461211690614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461214290614f01565b801561218f5780601f106121645761010080835404028352916020019161218f565b820191906000526020600020905b81548152906001019060200180831161217257829003601f168201915b505050505081602001819052506121a98260600151611e8b565b6040820152606080830151908201526080808301516001600160a01b03169082015260a080830151908201526121e0436001614e1e565b60c0808301919091528201516001600160401b031660e08201525b60005482516001600160401b039081166a010000000000000000000090920416141561228f576000815260208281015190820152606082015161223d90611e8b565b6040820152606080830151908201526080808301516001600160a01b03169082015260a08083015190820152612274436001614e1e565b60c0808301919091528201516001600160401b031660e08201525b60005482516001600160401b0390811672010000000000000000000000000000000000009092041614156123a85760008152600280546122ce90614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546122fa90614f01565b80156123475780601f1061231c57610100808354040283529160200191612347565b820191906000526020600020905b81548152906001019060200180831161232a57829003601f168201915b5050505050602082015260408281015190820152606080830151908201526080808301516001600160a01b03169082015260a0808301519082015261238d436001614e1e565b60c0808301919091528201516001600160401b031660e08201525b60015482516001600160401b039081169116141561242557600081526020828101519082015260408083015190820152606080830151908201526080808301516001600160a01b03169082015260a0808301519082015261240a436001614e1e565b60c0808301919091528201516001600160401b031660e08201525b600061243982602001518360400151610e3e565b90506124476006828461331e565b5060015483516001600160401b039081169116148015612487575060036040516124719190614b73565b6040518091039020836020015180519060200120145b156124bf57600160098260405161249e9190614b67565b908152604051908190036020019020805491151560ff199092169190911790555b7f517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e4183608001516124f7846020015185604001516107aa565b84604051610a4d93929190614bda565b60008061251583600061270b565b905061102e600182614e69565b6040805160808101825260608082526000602083018190529282018390528181019290925283600101838154811061256a57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600001805461258690614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546125b290614f01565b80156125ff5780601f106125d4576101008083540402835291602001916125ff565b820191906000526020600020905b8154815290600101906020018083116125e257829003601f168201915b5050505050915083600001826040516126189190614b67565b908152602001604051809103902060010160405180608001604052908160008201805461264490614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461267090614f01565b80156126bd5780601f10612692576101008083540402835291602001916126bd565b820191906000526020600020905b8154815290600101906020018083116126a057829003601f168201915b5050509183525050600191909101546001600160a01b0381166020830152600160a01b810460ff166040830152600160a81b90046001600160401b0316606090910152919491935090915050565b60008161271781614f54565b9250505b600183015482108015612764575082600101828154811061274c57634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b156105fb578161277381614f54565b92505061271b565b6040805160a08101825260008082526060602083018190529282018390528282015260808101919091526001600160a01b0380831660009081526020858152604091829020825160a0810190935260018101805490941683526002018054929392918401916127e990614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461281590614f01565b80156128625780601f1061283757610100808354040283529160200191612862565b820191906000526020600020905b81548152906001019060200180831161284557829003601f168201915b5050505050815260200160028201805461287b90614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546128a790614f01565b80156128f45780601f106128c9576101008083540402835291602001916128f4565b820191906000526020600020905b8154815290600101906020018083116128d757829003601f168201915b505050918352505060038201546001600160401b0316602082015260048201805460409092019161292490614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461295090614f01565b801561299d5780601f106129725761010080835404028352916020019161299d565b820191906000526020600020905b81548152906001019060200180831161298057829003601f168201915b505050505081525050905092915050565b6001600160a01b038281166000908152602085815260408220805485516001830180546001600160a01b031916919096161785558583015180519495919487949293612a01936002909101920190613cbc565b5060408201518051612a1d916002840191602090910190613cbc565b50606082015160038201805467ffffffffffffffff19166001600160401b0390921691909117905560808201518051612a60916004840191602090910190613cbc565b505081159050612a7457600191505061102e565b5060018085018054808301825560009190915290612a93908290614e1e565b6001600160a01b03851660009081526020879052604090205560018501805485919083908110612ad357634e487b7160e01b600052603260045260246000fd5b6000918252602082200180546001600160a01b0319166001600160a01b03939093169290921790915560028601805491612b0c83614f54565b9190505550600091505061102e565b509392505050565b6040805160808101825260608082526000602083018190528284018190529082015290518390612b54908490614b67565b9081526020016040518091039020600101604051806080016040529081600082018054612b8090614f01565b80601f0160208091040260200160405190810160405280929190818152602001828054612bac90614f01565b8015612bf95780601f10612bce57610100808354040283529160200191612bf9565b820191906000526020600020905b815481529060010190602001808311612bdc57829003601f168201915b5050509183525050600191909101546001600160a01b0381166020830152600160a01b810460ff166040830152600160a81b90046001600160401b03166060909101529392505050565b6000808460000184604051612c589190614b67565b90815260405190819003602001812054915083908690612c79908790614b67565b90815260200160405180910390206001016000820151816000019080519060200190612ca6929190613cbc565b5060208201516001909101805460408401516060909401516001600160401b0316600160a81b027fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff60ff909516600160a01b027fffffffffffffffffffffff0000000000000000000000000000000000000000009092166001600160a01b039094169390931717929092161790558015612d4457600191505061102e565b5060018085018054808301825560009190915290612d63908290614e1e565b6040518690612d73908790614b67565b9081526040519081900360200190205560018501805485919083908110612daa57634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016000019080519060200190612dce929190613cbc565b50600285018054906000612b0c83614f54565b6000808360000183604051612df69190614b67565b90815260405190819003602001902054905080612e175760009150506107a4565b6040518490612e27908590614b67565b90815260405190819003602001902060008082556001820181612e4a8282613c82565b50600190810180547fffffff0000000000000000000000000000000000000000000000000000000000169055915050848101612e868284614e69565b81548110612ea457634e487b7160e01b600052603260045260246000fd5b600091825260208220600291820201600101805460ff191693151593909317909255908501805491612ed583614eea565b91905055505092915050565b6001600160a01b03811660009081526020839052604081205480612f095760009150506107a4565b6001600160a01b03831660009081526020859052604081208181556001810180546001600160a01b031916815590919081612f476002850182613c82565b612f55600283016000613c82565b60038201805467ffffffffffffffff19169055612f76600483016000613c82565b50600192505050848101612f8a8284614e69565b81548110612fa857634e487b7160e01b600052603260045260246000fd5b600091825260208220018054921515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9093169290921790915560028501805491612ed583614eea565b61305b60405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b604051839061306b908490614b67565b90815260408051918290036020908101832061010084019092526001820180546001600160401b0316845260029092018054918401916130aa90614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546130d690614f01565b80156131235780601f106130f857610100808354040283529160200191613123565b820191906000526020600020905b81548152906001019060200180831161310657829003601f168201915b5050505050815260200160028201805461313c90614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461316890614f01565b80156131b55780601f1061318a576101008083540402835291602001916131b5565b820191906000526020600020905b81548152906001019060200180831161319857829003601f168201915b505050505081526020016003820180546131ce90614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546131fa90614f01565b80156132475780601f1061321c57610100808354040283529160200191613247565b820191906000526020600020905b81548152906001019060200180831161322a57829003601f168201915b505050918352505060048201546001600160a01b0316602082015260058201805460409092019161327790614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546132a390614f01565b80156132f05780601f106132c5576101008083540402835291602001916132f0565b820191906000526020600020905b8154815290600101906020018083116132d357829003601f168201915b5050509183525050600682015460208201526007909101546001600160401b03166040909101529392505050565b60008084600001846040516133339190614b67565b90815260405190819003602001812054915083908690613354908790614b67565b90815260405160209181900382019020825160018201805467ffffffffffffffff19166001600160401b0390921691909117815583830151805191936133a09360020192910190613cbc565b50604082015180516133bc916002840191602090910190613cbc565b50606082015180516133d8916003840191602090910190613cbc565b5060808201516004820180546001600160a01b0319166001600160a01b0390921691909117905560a0820151805161341a916005840191602090910190613cbc565b5060c0820151600682015560e0909101516007909101805467ffffffffffffffff19166001600160401b039092169190911790558015612d4457600191505061102e565b60008083600001836040516134739190614b67565b908152604051908190036020019020549050806134945760009150506107a4565b60405184906134a4908590614b67565b908152604051908190036020019020600080825560018201805467ffffffffffffffff19168155816134d96002850182613c82565b6134e7600283016000613c82565b6134f5600383016000613c82565b6004820180546001600160a01b0319169055613515600583016000613c82565b5060006006820155600701805467ffffffffffffffff191690555060019050848101612e868284614e69565b6000806125158360006137ce565b60006135956040518060a0016040528060006001600160a01b03168152602001606081526020016060815260200160006001600160401b03168152602001606081525090565b8360010183815481106135b857634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03908116808452878352604093849020845160a081019095526001810180549093168552600201805491965091928401919061360890614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461363490614f01565b80156136815780601f1061365657610100808354040283529160200191613681565b820191906000526020600020905b81548152906001019060200180831161366457829003601f168201915b5050505050815260200160028201805461369a90614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546136c690614f01565b80156137135780601f106136e857610100808354040283529160200191613713565b820191906000526020600020905b8154815290600101906020018083116136f657829003601f168201915b505050918352505060038201546001600160401b0316602082015260048201805460409092019161374390614f01565b80601f016020809104026020016040519081016040528092919081815260200182805461376f90614f01565b80156137bc5780601f10613791576101008083540402835291602001916137bc565b820191906000526020600020905b81548152906001019060200180831161379f57829003601f168201915b50505050508152505090509250929050565b6000816137da81614f54565b9250505b600183015482108015613826575082600101828154811061380f57634e487b7160e01b600052603260045260246000fd5b600091825260209091200154600160a01b900460ff165b156105fb578161383581614f54565b9250506137de565b600080612515836000613c12565b604080516101008101825260008082526060602083018190529282018390528183018390526080820181905260a0820183905260c0820181905260e08201528360010183815481106138ad57634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160000180546138c990614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546138f590614f01565b80156139425780601f1061391757610100808354040283529160200191613942565b820191906000526020600020905b81548152906001019060200180831161392557829003601f168201915b50505050509150836000018260405161395b9190614b67565b90815260408051918290036020908101832061010084019092526001820180546001600160401b03168452600290920180549184019161399a90614f01565b80601f01602080910402602001604051908101604052809291908181526020018280546139c690614f01565b8015613a135780601f106139e857610100808354040283529160200191613a13565b820191906000526020600020905b8154815290600101906020018083116139f657829003601f168201915b50505050508152602001600282018054613a2c90614f01565b80601f0160208091040260200160405190810160405280929190818152602001828054613a5890614f01565b8015613aa55780601f10613a7a57610100808354040283529160200191613aa5565b820191906000526020600020905b815481529060010190602001808311613a8857829003601f168201915b50505050508152602001600382018054613abe90614f01565b80601f0160208091040260200160405190810160405280929190818152602001828054613aea90614f01565b8015613b375780601f10613b0c57610100808354040283529160200191613b37565b820191906000526020600020905b815481529060010190602001808311613b1a57829003601f168201915b505050918352505060048201546001600160a01b03166020820152600582018054604090920191613b6790614f01565b80601f0160208091040260200160405190810160405280929190818152602001828054613b9390614f01565b8015613be05780601f10613bb557610100808354040283529160200191613be0565b820191906000526020600020905b815481529060010190602001808311613bc357829003601f168201915b5050509183525050600682015460208201526007909101546001600160401b0316604090910152919491935090915050565b600081613c1e81614f54565b9250505b600183015482108015613c6b5750826001018281548110613c5357634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b156105fb5781613c7a81614f54565b925050613c22565b508054613c8e90614f01565b6000825580601f10613c9e575050565b601f01602090049060005260206000209081019061152e9190613d40565b828054613cc890614f01565b90600052602060002090601f016020900481019282613cea5760008555613d30565b82601f10613d0357805160ff1916838001178555613d30565b82800160010185558215613d30579182015b82811115613d30578251825591602001919060010190613d15565b50613d3c929150613d40565b5090565b5b80821115613d3c5760008155600101613d41565b6000613d68613d6384614df4565b614dd8565b905082815260208101848484011115613d8057600080fd5b612b1b848285614eb2565b80356107a481614fb1565b80356107a481614fc5565b80516107a481614fc5565b600082601f830112613dbd57600080fd5b8135613dcd848260208601613d55565b949350505050565b600060608284031215613de757600080fd5b613df16060614dd8565b905081356001600160401b03811115613e0957600080fd5b613e1584828501613dac565b8252506020613e2684848301613d8b565b6020830152506040613e3a84828501614297565b60408301525092915050565b600060a08284031215613e5857600080fd5b613e6260a0614dd8565b90506000613e708484613d8b565b82525060208201356001600160401b03811115613e8c57600080fd5b613e9884828501613dac565b60208301525060408201356001600160401b03811115613eb757600080fd5b613ec384828501613dac565b6040830152506060613ed784828501614297565b60608301525060808201356001600160401b03811115613ef657600080fd5b613f0284828501613dac565b60808301525092915050565b600060408284031215613f2057600080fd5b613f2a6040614dd8565b905081356001600160401b03811115613f4257600080fd5b613f4e84828501613dac565b8252506020613f5f84848301613d8b565b60208301525092915050565b600060608284031215613f7d57600080fd5b613f876060614dd8565b905081356001600160401b03811115613f9f57600080fd5b613fab84828501613dac565b82525060208201356001600160401b03811115613fc757600080fd5b613fd384828501613dac565b6020830152506040613e3a84828501613d8b565b600060808284031215613ff957600080fd5b6140036080614dd8565b905081356001600160401b0381111561401b57600080fd5b61402784828501613dac565b825250602061403884848301613d8b565b60208301525060408201356001600160401b0381111561405757600080fd5b61406384828501613dac565b604083015250606061407784828501614297565b60608301525092915050565b600060e0828403121561409557600080fd5b61409f60e0614dd8565b905060006140ad8484614297565b82525060208201356001600160401b038111156140c957600080fd5b6140d584828501613dac565b60208301525060408201356001600160401b038111156140f457600080fd5b61410084828501613dac565b60408301525060608201356001600160401b0381111561411f57600080fd5b61412b84828501613dac565b606083015250608061413f84828501613d8b565b60808301525060a08201356001600160401b0381111561415e57600080fd5b61416a84828501613dac565b60a08301525060c061417e84828501614297565b60c08301525092915050565b60006080828403121561419c57600080fd5b6141a66080614dd8565b905081356001600160401b038111156141be57600080fd5b6141ca84828501613dac565b82525060208201356001600160401b038111156141e657600080fd5b6141f284828501613dac565b602083015250604061420684828501613d8b565b604083015250606061407784828501613d8b565b60006060828403121561422c57600080fd5b6142366060614dd8565b905060006142448484613d8b565b82525060208201356001600160401b0381111561426057600080fd5b61426c84828501613dac565b60208301525060408201356001600160401b0381111561428b57600080fd5b613e3a84828501613dac565b80356107a481614fcb565b6000602082840312156142b457600080fd5b6000613dcd8484613d8b565b6000602082840312156142d257600080fd5b6000613dcd8484613d96565b6000602082840312156142f057600080fd5b6000613dcd8484613da1565b60006020828403121561430e57600080fd5b81356001600160401b0381111561432457600080fd5b613dcd84828501613dac565b6000806040838503121561434357600080fd5b82356001600160401b0381111561435957600080fd5b61436585828601613dac565b92505060208301356001600160401b0381111561438157600080fd5b61438d85828601613dac565b9150509250929050565b6000602082840312156143a957600080fd5b81356001600160401b038111156143bf57600080fd5b613dcd84828501613dd5565b6000602082840312156143dd57600080fd5b81356001600160401b038111156143f357600080fd5b613dcd84828501613e46565b60006020828403121561441157600080fd5b81356001600160401b0381111561442757600080fd5b613dcd84828501613f0e565b60006020828403121561444557600080fd5b81356001600160401b0381111561445b57600080fd5b613dcd84828501613f6b565b60006020828403121561447957600080fd5b81356001600160401b0381111561448f57600080fd5b613dcd84828501613fe7565b6000602082840312156144ad57600080fd5b81356001600160401b038111156144c357600080fd5b613dcd84828501614083565b6000602082840312156144e157600080fd5b81356001600160401b038111156144f757600080fd5b613dcd8482850161418a565b60006020828403121561451557600080fd5b81356001600160401b0381111561452b57600080fd5b613dcd8482850161421a565b600061102e8383614955565b600061102e8383614a35565b600061102e8383614aeb565b61456481614ea1565b82525050565b6000614574825190565b8084526020840193508360208202850161458e8560200190565b8060005b858110156145c357848403895281516145ab8582614537565b94506020830160209a909a0199925050600101614592565b5091979650505050505050565b60006145da825190565b808452602084019350836020820285016145f48560200190565b8060005b858110156145c357848403895281516146118582614543565b94506020830160209a909a01999250506001016145f8565b6000614633825190565b8084526020840193508360208202850161464d8560200190565b8060005b858110156145c3578484038952815161466a858261454f565b94506020830160209a909a0199925050600101614651565b80614564565b6000614692825190565b8084526020840193506146a9818560208601614ebe565b601f01601f19169290920192915050565b60006146c4825190565b6146d2818560208601614ebe565b9290920192915050565b600081546146e981614f01565b600182168015614700576001811461471157614741565b60ff19831686528186019350614741565b60008581526020902060005b838110156147395781548882015260019091019060200161471d565b838801955050505b50505092915050565b601081526000602082017f6465706f736974206d757374203e203000000000000000000000000000000000815291505b5060200190565b601181526000602082017f6970206d757374206e6f7420656d7074790000000000000000000000000000008152915061477a565b600d81526000602082017f706f7274206d757374203e2030000000000000000000000000000000000000008152915061477a565b600d81526000602082017f6e6f7420636f6e73656e737573000000000000000000000000000000000000008152915061477a565b600c81526000602082017f6e6f7420726567697374657200000000000000000000000000000000000000008152915061477a565b600981526000602082017f6e6f742061646d696e00000000000000000000000000000000000000000000008152915061477a565b600e81526000602082017f696e646578206d757374203e20300000000000000000000000000000000000008152915061477a565b600981526000602082017f6e6f74206f776e657200000000000000000000000000000000000000000000008152915061477a565b601781526000602082017f6e6f7420656e6f75676820696e6974206465706f7369740000000000000000008152915061477a565b600c81526000602082017f706f73206d757374203e203000000000000000000000000000000000000000008152915061477a565b805160009060a0840190614969858261455b565b50602083015184820360208601526149818282614688565b9150506040830151848203604086015261499b8282614688565b91505060608301516149b06060860182614b3a565b50608083015184820360808601526107fb8282614688565b805160a0808452600091908401906149e08282614688565b91505060208301516149f5602086018261455b565b5060408301518482036040860152614a0d8282614688565b9150506060830151614a226060860182614682565b506080830151612b1b6080860182614b3a565b8051600090610100840190614a4a8582614b3a565b5060208301518482036020860152614a628282614688565b91505060408301518482036040860152614a7c8282614688565b91505060608301518482036060860152614a968282614688565b9150506080830151614aab608086018261455b565b5060a083015184820360a0860152614ac38282614688565b91505060c0830151614ad860c0860182614682565b5060e0830151612b1b60e0860182614b3a565b8051608080845260009190840190614b038282614688565b9150506020830151614b18602086018261455b565b506040830151614b2b6040860182614b49565b506060830151612b1b60608601825b6001600160401b038116614564565b60ff8116614564565b6000614b5e8284614682565b50602001919050565b600061102e82846146ba565b600061102e82846146dc565b602081016107a4828461455b565b60608101614b9b828661455b565b614ba8602083018561455b565b81810360408301526107fb8184614688565b60408101614bc8828561455b565b8181036020830152613dcd8184614688565b60608101614be8828661455b565b8181036020830152614bfa8185614688565b905081810360408301526107fb8184614a35565b6020808252810161102e818461456a565b6020808252810161102e81846145d0565b6020808252810161102e8184614629565b6020808252810161102e8184614688565b60808082528101614c638187614688565b90508181036020830152614c778186614688565b9050614c86604083018561455b565b6107fb6060830184614b3a565b602080825281016107a48161474a565b602080825281016107a481614781565b602080825281016107a4816147b5565b602080825281016107a4816147e9565b602080825281016107a48161481d565b602080825281016107a481602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b602080825281016107a481614851565b602080825281016107a481614885565b602080825281016107a4816148b9565b602080825281016107a4816148ed565b602080825281016107a481614921565b6020808252810161102e8184614955565b6020808252810161102e81846149c8565b6020808252810161102e8184614a35565b6020808252810161102e8184614aeb565b6000614de360405190565b9050614def8282614f28565b919050565b60006001600160401b03821115614e0d57614e0d614f9b565b601f19601f83011660200192915050565b60008219821115614e3157614e31614f6f565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115614e3157614e31614f6f565b6000825b925082821015614e7f57614e7f614f6f565b500390565b60006001600160401b03821691506001600160401b038316614e6d565b60006001600160a01b0382166107a4565b82818337506000910152565b60005b83811015614ed9578181015183820152602001614ec1565b83811115610cfa5750506000910152565b600081614ef957614ef9614f6f565b506000190190565b600281046001821680614f1557607f821691505b602082108114156105fb576105fb614f85565b601f19601f83011681018181106001600160401b0382111715614f4d57614f4d614f9b565b6040525050565b6000600019821415614f6857614f68614f6f565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b614fba81614ea1565b811461152e57600080fd5b80614fba565b6001600160401b038116614fba56fea26469706673582212205896fab2a77fc7638bdbdbf7990c8049d99f90b7d3253644323df003069df2c464736f6c63430008040033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// CreateDefaultUrl is a free data retrieval call binding the contract method 0xd72b6be6.
//
// Solidity: function CreateDefaultUrl(bytes name) pure returns(bytes)
func (_Store *StoreCaller) CreateDefaultUrl(opts *bind.CallOpts, name []byte) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CreateDefaultUrl", name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CreateDefaultUrl is a free data retrieval call binding the contract method 0xd72b6be6.
//
// Solidity: function CreateDefaultUrl(bytes name) pure returns(bytes)
func (_Store *StoreSession) CreateDefaultUrl(name []byte) ([]byte, error) {
	return _Store.Contract.CreateDefaultUrl(&_Store.CallOpts, name)
}

// CreateDefaultUrl is a free data retrieval call binding the contract method 0xd72b6be6.
//
// Solidity: function CreateDefaultUrl(bytes name) pure returns(bytes)
func (_Store *StoreCallerSession) CreateDefaultUrl(name []byte) ([]byte, error) {
	return _Store.Contract.CreateDefaultUrl(&_Store.CallOpts, name)
}

// GetAllDnsNodes is a free data retrieval call binding the contract method 0x6839a7c6.
//
// Solidity: function GetAllDnsNodes() view returns((address,bytes,bytes,uint64,string)[])
func (_Store *StoreCaller) GetAllDnsNodes(opts *bind.CallOpts) ([]DNSNodeInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetAllDnsNodes")

	if err != nil {
		return *new([]DNSNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DNSNodeInfo)).(*[]DNSNodeInfo)

	return out0, err

}

// GetAllDnsNodes is a free data retrieval call binding the contract method 0x6839a7c6.
//
// Solidity: function GetAllDnsNodes() view returns((address,bytes,bytes,uint64,string)[])
func (_Store *StoreSession) GetAllDnsNodes() ([]DNSNodeInfo, error) {
	return _Store.Contract.GetAllDnsNodes(&_Store.CallOpts)
}

// GetAllDnsNodes is a free data retrieval call binding the contract method 0x6839a7c6.
//
// Solidity: function GetAllDnsNodes() view returns((address,bytes,bytes,uint64,string)[])
func (_Store *StoreCallerSession) GetAllDnsNodes() ([]DNSNodeInfo, error) {
	return _Store.Contract.GetAllDnsNodes(&_Store.CallOpts)
}

// GetDNSNodeByAddress is a free data retrieval call binding the contract method 0x87dacef2.
//
// Solidity: function GetDNSNodeByAddress(address addr) view returns((address,bytes,bytes,uint64,string))
func (_Store *StoreCaller) GetDNSNodeByAddress(opts *bind.CallOpts, addr common.Address) (DNSNodeInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetDNSNodeByAddress", addr)

	if err != nil {
		return *new(DNSNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DNSNodeInfo)).(*DNSNodeInfo)

	return out0, err

}

// GetDNSNodeByAddress is a free data retrieval call binding the contract method 0x87dacef2.
//
// Solidity: function GetDNSNodeByAddress(address addr) view returns((address,bytes,bytes,uint64,string))
func (_Store *StoreSession) GetDNSNodeByAddress(addr common.Address) (DNSNodeInfo, error) {
	return _Store.Contract.GetDNSNodeByAddress(&_Store.CallOpts, addr)
}

// GetDNSNodeByAddress is a free data retrieval call binding the contract method 0x87dacef2.
//
// Solidity: function GetDNSNodeByAddress(address addr) view returns((address,bytes,bytes,uint64,string))
func (_Store *StoreCallerSession) GetDNSNodeByAddress(addr common.Address) (DNSNodeInfo, error) {
	return _Store.Contract.GetDNSNodeByAddress(&_Store.CallOpts, addr)
}

// GetHeader is a free data retrieval call binding the contract method 0x919064eb.
//
// Solidity: function GetHeader((bytes,bytes,address) req) view returns((bytes,address,bytes,uint256,uint64))
func (_Store *StoreCaller) GetHeader(opts *bind.CallOpts, req ReqInfo) (HeaderInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetHeader", req)

	if err != nil {
		return *new(HeaderInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(HeaderInfo)).(*HeaderInfo)

	return out0, err

}

// GetHeader is a free data retrieval call binding the contract method 0x919064eb.
//
// Solidity: function GetHeader((bytes,bytes,address) req) view returns((bytes,address,bytes,uint256,uint64))
func (_Store *StoreSession) GetHeader(req ReqInfo) (HeaderInfo, error) {
	return _Store.Contract.GetHeader(&_Store.CallOpts, req)
}

// GetHeader is a free data retrieval call binding the contract method 0x919064eb.
//
// Solidity: function GetHeader((bytes,bytes,address) req) view returns((bytes,address,bytes,uint256,uint64))
func (_Store *StoreCallerSession) GetHeader(req ReqInfo) (HeaderInfo, error) {
	return _Store.Contract.GetHeader(&_Store.CallOpts, req)
}

// GetName is a free data retrieval call binding the contract method 0x6b9ceb53.
//
// Solidity: function GetName((bytes,bytes,address) req) view returns((uint64,bytes,bytes,bytes,address,bytes,uint256,uint64))
func (_Store *StoreCaller) GetName(opts *bind.CallOpts, req ReqInfo) (NameInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetName", req)

	if err != nil {
		return *new(NameInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(NameInfo)).(*NameInfo)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x6b9ceb53.
//
// Solidity: function GetName((bytes,bytes,address) req) view returns((uint64,bytes,bytes,bytes,address,bytes,uint256,uint64))
func (_Store *StoreSession) GetName(req ReqInfo) (NameInfo, error) {
	return _Store.Contract.GetName(&_Store.CallOpts, req)
}

// GetName is a free data retrieval call binding the contract method 0x6b9ceb53.
//
// Solidity: function GetName((bytes,bytes,address) req) view returns((uint64,bytes,bytes,bytes,address,bytes,uint256,uint64))
func (_Store *StoreCallerSession) GetName(req ReqInfo) (NameInfo, error) {
	return _Store.Contract.GetName(&_Store.CallOpts, req)
}

// GetPeerPoolItem is a free data retrieval call binding the contract method 0x240ce9f9.
//
// Solidity: function GetPeerPoolItem(string peerPubKey) view returns((string,address,uint8,uint64))
func (_Store *StoreCaller) GetPeerPoolItem(opts *bind.CallOpts, peerPubKey string) (PeerPoolItem, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetPeerPoolItem", peerPubKey)

	if err != nil {
		return *new(PeerPoolItem), err
	}

	out0 := *abi.ConvertType(out[0], new(PeerPoolItem)).(*PeerPoolItem)

	return out0, err

}

// GetPeerPoolItem is a free data retrieval call binding the contract method 0x240ce9f9.
//
// Solidity: function GetPeerPoolItem(string peerPubKey) view returns((string,address,uint8,uint64))
func (_Store *StoreSession) GetPeerPoolItem(peerPubKey string) (PeerPoolItem, error) {
	return _Store.Contract.GetPeerPoolItem(&_Store.CallOpts, peerPubKey)
}

// GetPeerPoolItem is a free data retrieval call binding the contract method 0x240ce9f9.
//
// Solidity: function GetPeerPoolItem(string peerPubKey) view returns((string,address,uint8,uint64))
func (_Store *StoreCallerSession) GetPeerPoolItem(peerPubKey string) (PeerPoolItem, error) {
	return _Store.Contract.GetPeerPoolItem(&_Store.CallOpts, peerPubKey)
}

// GetPeerPoolMap is a free data retrieval call binding the contract method 0x0e580e79.
//
// Solidity: function GetPeerPoolMap() view returns((string,address,uint8,uint64)[])
func (_Store *StoreCaller) GetPeerPoolMap(opts *bind.CallOpts) ([]PeerPoolItem, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetPeerPoolMap")

	if err != nil {
		return *new([]PeerPoolItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]PeerPoolItem)).(*[]PeerPoolItem)

	return out0, err

}

// GetPeerPoolMap is a free data retrieval call binding the contract method 0x0e580e79.
//
// Solidity: function GetPeerPoolMap() view returns((string,address,uint8,uint64)[])
func (_Store *StoreSession) GetPeerPoolMap() ([]PeerPoolItem, error) {
	return _Store.Contract.GetPeerPoolMap(&_Store.CallOpts)
}

// GetPeerPoolMap is a free data retrieval call binding the contract method 0x0e580e79.
//
// Solidity: function GetPeerPoolMap() view returns((string,address,uint8,uint64)[])
func (_Store *StoreCallerSession) GetPeerPoolMap() ([]PeerPoolItem, error) {
	return _Store.Contract.GetPeerPoolMap(&_Store.CallOpts)
}

// GetPluginList is a free data retrieval call binding the contract method 0xccfd72f6.
//
// Solidity: function GetPluginList() view returns((uint64,bytes,bytes,bytes,address,bytes,uint256,uint64)[])
func (_Store *StoreCaller) GetPluginList(opts *bind.CallOpts) ([]NameInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetPluginList")

	if err != nil {
		return *new([]NameInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]NameInfo)).(*[]NameInfo)

	return out0, err

}

// GetPluginList is a free data retrieval call binding the contract method 0xccfd72f6.
//
// Solidity: function GetPluginList() view returns((uint64,bytes,bytes,bytes,address,bytes,uint256,uint64)[])
func (_Store *StoreSession) GetPluginList() ([]NameInfo, error) {
	return _Store.Contract.GetPluginList(&_Store.CallOpts)
}

// GetPluginList is a free data retrieval call binding the contract method 0xccfd72f6.
//
// Solidity: function GetPluginList() view returns((uint64,bytes,bytes,bytes,address,bytes,uint256,uint64)[])
func (_Store *StoreCallerSession) GetPluginList() ([]NameInfo, error) {
	return _Store.Contract.GetPluginList(&_Store.CallOpts)
}

// GetUrl is a free data retrieval call binding the contract method 0x2a5da31e.
//
// Solidity: function GetUrl(bytes header, bytes url) pure returns(bytes)
func (_Store *StoreCaller) GetUrl(opts *bind.CallOpts, header []byte, url []byte) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUrl", header, url)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetUrl is a free data retrieval call binding the contract method 0x2a5da31e.
//
// Solidity: function GetUrl(bytes header, bytes url) pure returns(bytes)
func (_Store *StoreSession) GetUrl(header []byte, url []byte) ([]byte, error) {
	return _Store.Contract.GetUrl(&_Store.CallOpts, header, url)
}

// GetUrl is a free data retrieval call binding the contract method 0x2a5da31e.
//
// Solidity: function GetUrl(bytes header, bytes url) pure returns(bytes)
func (_Store *StoreCallerSession) GetUrl(header []byte, url []byte) ([]byte, error) {
	return _Store.Contract.GetUrl(&_Store.CallOpts, header, url)
}

// Concat is a free data retrieval call binding the contract method 0x5f3376f3.
//
// Solidity: function concat(bytes b1, bytes b2) pure returns(bytes)
func (_Store *StoreCaller) Concat(opts *bind.CallOpts, b1 []byte, b2 []byte) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "concat", b1, b2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Concat is a free data retrieval call binding the contract method 0x5f3376f3.
//
// Solidity: function concat(bytes b1, bytes b2) pure returns(bytes)
func (_Store *StoreSession) Concat(b1 []byte, b2 []byte) ([]byte, error) {
	return _Store.Contract.Concat(&_Store.CallOpts, b1, b2)
}

// Concat is a free data retrieval call binding the contract method 0x5f3376f3.
//
// Solidity: function concat(bytes b1, bytes b2) pure returns(bytes)
func (_Store *StoreCallerSession) Concat(b1 []byte, b2 []byte) ([]byte, error) {
	return _Store.Contract.Concat(&_Store.CallOpts, b1, b2)
}

// ToBytes is a free data retrieval call binding the contract method 0x326cf61c.
//
// Solidity: function toBytes(bytes32 _data) pure returns(bytes)
func (_Store *StoreCaller) ToBytes(opts *bind.CallOpts, _data [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "toBytes", _data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ToBytes is a free data retrieval call binding the contract method 0x326cf61c.
//
// Solidity: function toBytes(bytes32 _data) pure returns(bytes)
func (_Store *StoreSession) ToBytes(_data [32]byte) ([]byte, error) {
	return _Store.Contract.ToBytes(&_Store.CallOpts, _data)
}

// ToBytes is a free data retrieval call binding the contract method 0x326cf61c.
//
// Solidity: function toBytes(bytes32 _data) pure returns(bytes)
func (_Store *StoreCallerSession) ToBytes(_data [32]byte) ([]byte, error) {
	return _Store.Contract.ToBytes(&_Store.CallOpts, _data)
}

// AddInitPos is a paid mutator transaction binding the contract method 0x51372819.
//
// Solidity: function AddInitPos((string,address,uint64) req) payable returns()
func (_Store *StoreTransactor) AddInitPos(opts *bind.TransactOpts, req ChangeInitPosParam) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddInitPos", req)
}

// AddInitPos is a paid mutator transaction binding the contract method 0x51372819.
//
// Solidity: function AddInitPos((string,address,uint64) req) payable returns()
func (_Store *StoreSession) AddInitPos(req ChangeInitPosParam) (*types.Transaction, error) {
	return _Store.Contract.AddInitPos(&_Store.TransactOpts, req)
}

// AddInitPos is a paid mutator transaction binding the contract method 0x51372819.
//
// Solidity: function AddInitPos((string,address,uint64) req) payable returns()
func (_Store *StoreTransactorSession) AddInitPos(req ChangeInitPosParam) (*types.Transaction, error) {
	return _Store.Contract.AddInitPos(&_Store.TransactOpts, req)
}

// ApproveDNSCandidate is a paid mutator transaction binding the contract method 0x228a09b2.
//
// Solidity: function ApproveDNSCandidate(string peerPubKey) returns()
func (_Store *StoreTransactor) ApproveDNSCandidate(opts *bind.TransactOpts, peerPubKey string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ApproveDNSCandidate", peerPubKey)
}

// ApproveDNSCandidate is a paid mutator transaction binding the contract method 0x228a09b2.
//
// Solidity: function ApproveDNSCandidate(string peerPubKey) returns()
func (_Store *StoreSession) ApproveDNSCandidate(peerPubKey string) (*types.Transaction, error) {
	return _Store.Contract.ApproveDNSCandidate(&_Store.TransactOpts, peerPubKey)
}

// ApproveDNSCandidate is a paid mutator transaction binding the contract method 0x228a09b2.
//
// Solidity: function ApproveDNSCandidate(string peerPubKey) returns()
func (_Store *StoreTransactorSession) ApproveDNSCandidate(peerPubKey string) (*types.Transaction, error) {
	return _Store.Contract.ApproveDNSCandidate(&_Store.TransactOpts, peerPubKey)
}

// DNSNodeReg is a paid mutator transaction binding the contract method 0x465efc9a.
//
// Solidity: function DNSNodeReg((address,bytes,bytes,uint64,string) info) payable returns()
func (_Store *StoreTransactor) DNSNodeReg(opts *bind.TransactOpts, info DNSNodeInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DNSNodeReg", info)
}

// DNSNodeReg is a paid mutator transaction binding the contract method 0x465efc9a.
//
// Solidity: function DNSNodeReg((address,bytes,bytes,uint64,string) info) payable returns()
func (_Store *StoreSession) DNSNodeReg(info DNSNodeInfo) (*types.Transaction, error) {
	return _Store.Contract.DNSNodeReg(&_Store.TransactOpts, info)
}

// DNSNodeReg is a paid mutator transaction binding the contract method 0x465efc9a.
//
// Solidity: function DNSNodeReg((address,bytes,bytes,uint64,string) info) payable returns()
func (_Store *StoreTransactorSession) DNSNodeReg(info DNSNodeInfo) (*types.Transaction, error) {
	return _Store.Contract.DNSNodeReg(&_Store.TransactOpts, info)
}

// DelDNS is a paid mutator transaction binding the contract method 0x4f2ffbba.
//
// Solidity: function DelDNS((bytes,bytes,address) req) payable returns()
func (_Store *StoreTransactor) DelDNS(opts *bind.TransactOpts, req ReqInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelDNS", req)
}

// DelDNS is a paid mutator transaction binding the contract method 0x4f2ffbba.
//
// Solidity: function DelDNS((bytes,bytes,address) req) payable returns()
func (_Store *StoreSession) DelDNS(req ReqInfo) (*types.Transaction, error) {
	return _Store.Contract.DelDNS(&_Store.TransactOpts, req)
}

// DelDNS is a paid mutator transaction binding the contract method 0x4f2ffbba.
//
// Solidity: function DelDNS((bytes,bytes,address) req) payable returns()
func (_Store *StoreTransactorSession) DelDNS(req ReqInfo) (*types.Transaction, error) {
	return _Store.Contract.DelDNS(&_Store.TransactOpts, req)
}

// DelHeader is a paid mutator transaction binding the contract method 0x6e91c568.
//
// Solidity: function DelHeader((bytes,bytes,address) req) payable returns()
func (_Store *StoreTransactor) DelHeader(opts *bind.TransactOpts, req ReqInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelHeader", req)
}

// DelHeader is a paid mutator transaction binding the contract method 0x6e91c568.
//
// Solidity: function DelHeader((bytes,bytes,address) req) payable returns()
func (_Store *StoreSession) DelHeader(req ReqInfo) (*types.Transaction, error) {
	return _Store.Contract.DelHeader(&_Store.TransactOpts, req)
}

// DelHeader is a paid mutator transaction binding the contract method 0x6e91c568.
//
// Solidity: function DelHeader((bytes,bytes,address) req) payable returns()
func (_Store *StoreTransactorSession) DelHeader(req ReqInfo) (*types.Transaction, error) {
	return _Store.Contract.DelHeader(&_Store.TransactOpts, req)
}

// DnsInit is a paid mutator transaction binding the contract method 0xba47d33d.
//
// Solidity: function DnsInit(address _admin) returns()
func (_Store *StoreTransactor) DnsInit(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DnsInit", _admin)
}

// DnsInit is a paid mutator transaction binding the contract method 0xba47d33d.
//
// Solidity: function DnsInit(address _admin) returns()
func (_Store *StoreSession) DnsInit(_admin common.Address) (*types.Transaction, error) {
	return _Store.Contract.DnsInit(&_Store.TransactOpts, _admin)
}

// DnsInit is a paid mutator transaction binding the contract method 0xba47d33d.
//
// Solidity: function DnsInit(address _admin) returns()
func (_Store *StoreTransactorSession) DnsInit(_admin common.Address) (*types.Transaction, error) {
	return _Store.Contract.DnsInit(&_Store.TransactOpts, _admin)
}

// QuitNode is a paid mutator transaction binding the contract method 0xd6841470.
//
// Solidity: function QuitNode((string,address) req) returns()
func (_Store *StoreTransactor) QuitNode(opts *bind.TransactOpts, req QuitNodeParam) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "QuitNode", req)
}

// QuitNode is a paid mutator transaction binding the contract method 0xd6841470.
//
// Solidity: function QuitNode((string,address) req) returns()
func (_Store *StoreSession) QuitNode(req QuitNodeParam) (*types.Transaction, error) {
	return _Store.Contract.QuitNode(&_Store.TransactOpts, req)
}

// QuitNode is a paid mutator transaction binding the contract method 0xd6841470.
//
// Solidity: function QuitNode((string,address) req) returns()
func (_Store *StoreTransactorSession) QuitNode(req QuitNodeParam) (*types.Transaction, error) {
	return _Store.Contract.QuitNode(&_Store.TransactOpts, req)
}

// ReduceInitPos is a paid mutator transaction binding the contract method 0xeafafaf7.
//
// Solidity: function ReduceInitPos((string,address,uint64) req) payable returns()
func (_Store *StoreTransactor) ReduceInitPos(opts *bind.TransactOpts, req ChangeInitPosParam) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ReduceInitPos", req)
}

// ReduceInitPos is a paid mutator transaction binding the contract method 0xeafafaf7.
//
// Solidity: function ReduceInitPos((string,address,uint64) req) payable returns()
func (_Store *StoreSession) ReduceInitPos(req ChangeInitPosParam) (*types.Transaction, error) {
	return _Store.Contract.ReduceInitPos(&_Store.TransactOpts, req)
}

// ReduceInitPos is a paid mutator transaction binding the contract method 0xeafafaf7.
//
// Solidity: function ReduceInitPos((string,address,uint64) req) payable returns()
func (_Store *StoreTransactorSession) ReduceInitPos(req ChangeInitPosParam) (*types.Transaction, error) {
	return _Store.Contract.ReduceInitPos(&_Store.TransactOpts, req)
}

// RegisterHeader is a paid mutator transaction binding the contract method 0x809bae57.
//
// Solidity: function RegisterHeader((bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreTransactor) RegisterHeader(opts *bind.TransactOpts, req RequestHeader) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "RegisterHeader", req)
}

// RegisterHeader is a paid mutator transaction binding the contract method 0x809bae57.
//
// Solidity: function RegisterHeader((bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreSession) RegisterHeader(req RequestHeader) (*types.Transaction, error) {
	return _Store.Contract.RegisterHeader(&_Store.TransactOpts, req)
}

// RegisterHeader is a paid mutator transaction binding the contract method 0x809bae57.
//
// Solidity: function RegisterHeader((bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreTransactorSession) RegisterHeader(req RequestHeader) (*types.Transaction, error) {
	return _Store.Contract.RegisterHeader(&_Store.TransactOpts, req)
}

// RegisterName is a paid mutator transaction binding the contract method 0xfd11c4c6.
//
// Solidity: function RegisterName((uint64,bytes,bytes,bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreTransactor) RegisterName(opts *bind.TransactOpts, req RequestName) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "RegisterName", req)
}

// RegisterName is a paid mutator transaction binding the contract method 0xfd11c4c6.
//
// Solidity: function RegisterName((uint64,bytes,bytes,bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreSession) RegisterName(req RequestName) (*types.Transaction, error) {
	return _Store.Contract.RegisterName(&_Store.TransactOpts, req)
}

// RegisterName is a paid mutator transaction binding the contract method 0xfd11c4c6.
//
// Solidity: function RegisterName((uint64,bytes,bytes,bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreTransactorSession) RegisterName(req RequestName) (*types.Transaction, error) {
	return _Store.Contract.RegisterName(&_Store.TransactOpts, req)
}

// RejectDNSCandidate is a paid mutator transaction binding the contract method 0xdbc8ef9c.
//
// Solidity: function RejectDNSCandidate(string peerPubKey) payable returns()
func (_Store *StoreTransactor) RejectDNSCandidate(opts *bind.TransactOpts, peerPubKey string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "RejectDNSCandidate", peerPubKey)
}

// RejectDNSCandidate is a paid mutator transaction binding the contract method 0xdbc8ef9c.
//
// Solidity: function RejectDNSCandidate(string peerPubKey) payable returns()
func (_Store *StoreSession) RejectDNSCandidate(peerPubKey string) (*types.Transaction, error) {
	return _Store.Contract.RejectDNSCandidate(&_Store.TransactOpts, peerPubKey)
}

// RejectDNSCandidate is a paid mutator transaction binding the contract method 0xdbc8ef9c.
//
// Solidity: function RejectDNSCandidate(string peerPubKey) payable returns()
func (_Store *StoreTransactorSession) RejectDNSCandidate(peerPubKey string) (*types.Transaction, error) {
	return _Store.Contract.RejectDNSCandidate(&_Store.TransactOpts, peerPubKey)
}

// TransferHeader is a paid mutator transaction binding the contract method 0x8b79782b.
//
// Solidity: function TransferHeader((bytes,bytes,address,address) info) returns()
func (_Store *StoreTransactor) TransferHeader(opts *bind.TransactOpts, info TransferInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "TransferHeader", info)
}

// TransferHeader is a paid mutator transaction binding the contract method 0x8b79782b.
//
// Solidity: function TransferHeader((bytes,bytes,address,address) info) returns()
func (_Store *StoreSession) TransferHeader(info TransferInfo) (*types.Transaction, error) {
	return _Store.Contract.TransferHeader(&_Store.TransactOpts, info)
}

// TransferHeader is a paid mutator transaction binding the contract method 0x8b79782b.
//
// Solidity: function TransferHeader((bytes,bytes,address,address) info) returns()
func (_Store *StoreTransactorSession) TransferHeader(info TransferInfo) (*types.Transaction, error) {
	return _Store.Contract.TransferHeader(&_Store.TransactOpts, info)
}

// TransferName is a paid mutator transaction binding the contract method 0x3e4e5e54.
//
// Solidity: function TransferName((bytes,bytes,address,address) info) returns()
func (_Store *StoreTransactor) TransferName(opts *bind.TransactOpts, info TransferInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "TransferName", info)
}

// TransferName is a paid mutator transaction binding the contract method 0x3e4e5e54.
//
// Solidity: function TransferName((bytes,bytes,address,address) info) returns()
func (_Store *StoreSession) TransferName(info TransferInfo) (*types.Transaction, error) {
	return _Store.Contract.TransferName(&_Store.TransactOpts, info)
}

// TransferName is a paid mutator transaction binding the contract method 0x3e4e5e54.
//
// Solidity: function TransferName((bytes,bytes,address,address) info) returns()
func (_Store *StoreTransactorSession) TransferName(info TransferInfo) (*types.Transaction, error) {
	return _Store.Contract.TransferName(&_Store.TransactOpts, info)
}

// UnRegDNSNode is a paid mutator transaction binding the contract method 0x2c261f42.
//
// Solidity: function UnRegDNSNode((string,address) req) returns()
func (_Store *StoreTransactor) UnRegDNSNode(opts *bind.TransactOpts, req UnRegisterCandidateParam) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UnRegDNSNode", req)
}

// UnRegDNSNode is a paid mutator transaction binding the contract method 0x2c261f42.
//
// Solidity: function UnRegDNSNode((string,address) req) returns()
func (_Store *StoreSession) UnRegDNSNode(req UnRegisterCandidateParam) (*types.Transaction, error) {
	return _Store.Contract.UnRegDNSNode(&_Store.TransactOpts, req)
}

// UnRegDNSNode is a paid mutator transaction binding the contract method 0x2c261f42.
//
// Solidity: function UnRegDNSNode((string,address) req) returns()
func (_Store *StoreTransactorSession) UnRegDNSNode(req UnRegisterCandidateParam) (*types.Transaction, error) {
	return _Store.Contract.UnRegDNSNode(&_Store.TransactOpts, req)
}

// UpdateDNSNodesInfo is a paid mutator transaction binding the contract method 0x1654a894.
//
// Solidity: function UpdateDNSNodesInfo((address,bytes,bytes) info) returns()
func (_Store *StoreTransactor) UpdateDNSNodesInfo(opts *bind.TransactOpts, info UpdateNodeParam) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateDNSNodesInfo", info)
}

// UpdateDNSNodesInfo is a paid mutator transaction binding the contract method 0x1654a894.
//
// Solidity: function UpdateDNSNodesInfo((address,bytes,bytes) info) returns()
func (_Store *StoreSession) UpdateDNSNodesInfo(info UpdateNodeParam) (*types.Transaction, error) {
	return _Store.Contract.UpdateDNSNodesInfo(&_Store.TransactOpts, info)
}

// UpdateDNSNodesInfo is a paid mutator transaction binding the contract method 0x1654a894.
//
// Solidity: function UpdateDNSNodesInfo((address,bytes,bytes) info) returns()
func (_Store *StoreTransactorSession) UpdateDNSNodesInfo(info UpdateNodeParam) (*types.Transaction, error) {
	return _Store.Contract.UpdateDNSNodesInfo(&_Store.TransactOpts, info)
}

// UpdateName is a paid mutator transaction binding the contract method 0x370d7104.
//
// Solidity: function UpdateName((uint64,bytes,bytes,bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreTransactor) UpdateName(opts *bind.TransactOpts, req RequestName) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateName", req)
}

// UpdateName is a paid mutator transaction binding the contract method 0x370d7104.
//
// Solidity: function UpdateName((uint64,bytes,bytes,bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreSession) UpdateName(req RequestName) (*types.Transaction, error) {
	return _Store.Contract.UpdateName(&_Store.TransactOpts, req)
}

// UpdateName is a paid mutator transaction binding the contract method 0x370d7104.
//
// Solidity: function UpdateName((uint64,bytes,bytes,bytes,address,bytes,uint64) req) payable returns()
func (_Store *StoreTransactorSession) UpdateName(req RequestName) (*types.Transaction, error) {
	return _Store.Contract.UpdateName(&_Store.TransactOpts, req)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactorSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
}

// StoreCreateSectorEventIterator is returned from FilterCreateSectorEvent and is used to iterate over the raw logs and unpacked data for CreateSectorEvent events raised by the Store contract.
type StoreCreateSectorEventIterator struct {
	Event *StoreCreateSectorEvent // Event containing the contract specifics and raw log

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
func (it *StoreCreateSectorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCreateSectorEvent)
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
		it.Event = new(StoreCreateSectorEvent)
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
func (it *StoreCreateSectorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCreateSectorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCreateSectorEvent represents a CreateSectorEvent event raised by the Store contract.
type StoreCreateSectorEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SectorId    uint64
	ProvLevel   uint8
	Size        uint64
	IsPlots     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateSectorEvent is a free log retrieval operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) FilterCreateSectorEvent(opts *bind.FilterOpts) (*StoreCreateSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreCreateSectorEventIterator{contract: _Store.contract, event: "CreateSectorEvent", logs: logs, sub: sub}, nil
}

// WatchCreateSectorEvent is a free log subscription operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) WatchCreateSectorEvent(opts *bind.WatchOpts, sink chan<- *StoreCreateSectorEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCreateSectorEvent)
				if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
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

// ParseCreateSectorEvent is a log parse operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) ParseCreateSectorEvent(log types.Log) (*StoreCreateSectorEvent, error) {
	event := new(StoreCreateSectorEvent)
	if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDNSNodeRegisterIterator is returned from FilterDNSNodeRegister and is used to iterate over the raw logs and unpacked data for DNSNodeRegister events raised by the Store contract.
type StoreDNSNodeRegisterIterator struct {
	Event *StoreDNSNodeRegister // Event containing the contract specifics and raw log

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
func (it *StoreDNSNodeRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDNSNodeRegister)
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
		it.Event = new(StoreDNSNodeRegister)
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
func (it *StoreDNSNodeRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDNSNodeRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDNSNodeRegister represents a DNSNodeRegister event raised by the Store contract.
type StoreDNSNodeRegister struct {
	Ip         []byte
	Port       []byte
	WalletAddr common.Address
	Deposit    uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDNSNodeRegister is a free log retrieval operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) FilterDNSNodeRegister(opts *bind.FilterOpts) (*StoreDNSNodeRegisterIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DNSNodeRegister")
	if err != nil {
		return nil, err
	}
	return &StoreDNSNodeRegisterIterator{contract: _Store.contract, event: "DNSNodeRegister", logs: logs, sub: sub}, nil
}

// WatchDNSNodeRegister is a free log subscription operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) WatchDNSNodeRegister(opts *bind.WatchOpts, sink chan<- *StoreDNSNodeRegister) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DNSNodeRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDNSNodeRegister)
				if err := _Store.contract.UnpackLog(event, "DNSNodeRegister", log); err != nil {
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

// ParseDNSNodeRegister is a log parse operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) ParseDNSNodeRegister(log types.Log) (*StoreDNSNodeRegister, error) {
	event := new(StoreDNSNodeRegister)
	if err := _Store.contract.UnpackLog(event, "DNSNodeRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDNSNodeUnRegIterator is returned from FilterDNSNodeUnReg and is used to iterate over the raw logs and unpacked data for DNSNodeUnReg events raised by the Store contract.
type StoreDNSNodeUnRegIterator struct {
	Event *StoreDNSNodeUnReg // Event containing the contract specifics and raw log

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
func (it *StoreDNSNodeUnRegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDNSNodeUnReg)
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
		it.Event = new(StoreDNSNodeUnReg)
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
func (it *StoreDNSNodeUnRegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDNSNodeUnRegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDNSNodeUnReg represents a DNSNodeUnReg event raised by the Store contract.
type StoreDNSNodeUnReg struct {
	WalletAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDNSNodeUnReg is a free log retrieval operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) FilterDNSNodeUnReg(opts *bind.FilterOpts) (*StoreDNSNodeUnRegIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DNSNodeUnReg")
	if err != nil {
		return nil, err
	}
	return &StoreDNSNodeUnRegIterator{contract: _Store.contract, event: "DNSNodeUnReg", logs: logs, sub: sub}, nil
}

// WatchDNSNodeUnReg is a free log subscription operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) WatchDNSNodeUnReg(opts *bind.WatchOpts, sink chan<- *StoreDNSNodeUnReg) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DNSNodeUnReg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDNSNodeUnReg)
				if err := _Store.contract.UnpackLog(event, "DNSNodeUnReg", log); err != nil {
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

// ParseDNSNodeUnReg is a log parse operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) ParseDNSNodeUnReg(log types.Log) (*StoreDNSNodeUnReg, error) {
	event := new(StoreDNSNodeUnReg)
	if err := _Store.contract.UnpackLog(event, "DNSNodeUnReg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteFileEventIterator is returned from FilterDeleteFileEvent and is used to iterate over the raw logs and unpacked data for DeleteFileEvent events raised by the Store contract.
type StoreDeleteFileEventIterator struct {
	Event *StoreDeleteFileEvent // Event containing the contract specifics and raw log

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
func (it *StoreDeleteFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteFileEvent)
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
		it.Event = new(StoreDeleteFileEvent)
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
func (it *StoreDeleteFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteFileEvent represents a DeleteFileEvent event raised by the Store contract.
type StoreDeleteFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteFileEvent is a free log retrieval operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) FilterDeleteFileEvent(opts *bind.FilterOpts) (*StoreDeleteFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteFileEventIterator{contract: _Store.contract, event: "DeleteFileEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteFileEvent is a free log subscription operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) WatchDeleteFileEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteFileEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteFileEvent", log); err != nil {
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

// ParseDeleteFileEvent is a log parse operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) ParseDeleteFileEvent(log types.Log) (*StoreDeleteFileEvent, error) {
	event := new(StoreDeleteFileEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteFilesEventIterator is returned from FilterDeleteFilesEvent and is used to iterate over the raw logs and unpacked data for DeleteFilesEvent events raised by the Store contract.
type StoreDeleteFilesEventIterator struct {
	Event *StoreDeleteFilesEvent // Event containing the contract specifics and raw log

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
func (it *StoreDeleteFilesEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteFilesEvent)
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
		it.Event = new(StoreDeleteFilesEvent)
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
func (it *StoreDeleteFilesEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteFilesEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteFilesEvent represents a DeleteFilesEvent event raised by the Store contract.
type StoreDeleteFilesEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHashs   [][]byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteFilesEvent is a free log retrieval operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) FilterDeleteFilesEvent(opts *bind.FilterOpts) (*StoreDeleteFilesEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteFilesEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteFilesEventIterator{contract: _Store.contract, event: "DeleteFilesEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteFilesEvent is a free log subscription operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) WatchDeleteFilesEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteFilesEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteFilesEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteFilesEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteFilesEvent", log); err != nil {
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

// ParseDeleteFilesEvent is a log parse operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) ParseDeleteFilesEvent(log types.Log) (*StoreDeleteFilesEvent, error) {
	event := new(StoreDeleteFilesEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteFilesEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteSectorEventIterator is returned from FilterDeleteSectorEvent and is used to iterate over the raw logs and unpacked data for DeleteSectorEvent events raised by the Store contract.
type StoreDeleteSectorEventIterator struct {
	Event *StoreDeleteSectorEvent // Event containing the contract specifics and raw log

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
func (it *StoreDeleteSectorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteSectorEvent)
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
		it.Event = new(StoreDeleteSectorEvent)
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
func (it *StoreDeleteSectorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteSectorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteSectorEvent represents a DeleteSectorEvent event raised by the Store contract.
type StoreDeleteSectorEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SectorId    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteSectorEvent is a free log retrieval operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) FilterDeleteSectorEvent(opts *bind.FilterOpts) (*StoreDeleteSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteSectorEventIterator{contract: _Store.contract, event: "DeleteSectorEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteSectorEvent is a free log subscription operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) WatchDeleteSectorEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteSectorEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteSectorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteSectorEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteSectorEvent", log); err != nil {
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

// ParseDeleteSectorEvent is a log parse operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) ParseDeleteSectorEvent(log types.Log) (*StoreDeleteSectorEvent, error) {
	event := new(StoreDeleteSectorEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteSectorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreFilePDPSuccessEventIterator is returned from FilterFilePDPSuccessEvent and is used to iterate over the raw logs and unpacked data for FilePDPSuccessEvent events raised by the Store contract.
type StoreFilePDPSuccessEventIterator struct {
	Event *StoreFilePDPSuccessEvent // Event containing the contract specifics and raw log

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
func (it *StoreFilePDPSuccessEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFilePDPSuccessEvent)
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
		it.Event = new(StoreFilePDPSuccessEvent)
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
func (it *StoreFilePDPSuccessEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFilePDPSuccessEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFilePDPSuccessEvent represents a FilePDPSuccessEvent event raised by the Store contract.
type StoreFilePDPSuccessEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFilePDPSuccessEvent is a free log retrieval operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) FilterFilePDPSuccessEvent(opts *bind.FilterOpts) (*StoreFilePDPSuccessEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FilePDPSuccessEvent")
	if err != nil {
		return nil, err
	}
	return &StoreFilePDPSuccessEventIterator{contract: _Store.contract, event: "FilePDPSuccessEvent", logs: logs, sub: sub}, nil
}

// WatchFilePDPSuccessEvent is a free log subscription operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) WatchFilePDPSuccessEvent(opts *bind.WatchOpts, sink chan<- *StoreFilePDPSuccessEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FilePDPSuccessEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFilePDPSuccessEvent)
				if err := _Store.contract.UnpackLog(event, "FilePDPSuccessEvent", log); err != nil {
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

// ParseFilePDPSuccessEvent is a log parse operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) ParseFilePDPSuccessEvent(log types.Log) (*StoreFilePDPSuccessEvent, error) {
	event := new(StoreFilePDPSuccessEvent)
	if err := _Store.contract.UnpackLog(event, "FilePDPSuccessEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreGetUpdateCostEventIterator is returned from FilterGetUpdateCostEvent and is used to iterate over the raw logs and unpacked data for GetUpdateCostEvent events raised by the Store contract.
type StoreGetUpdateCostEventIterator struct {
	Event *StoreGetUpdateCostEvent // Event containing the contract specifics and raw log

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
func (it *StoreGetUpdateCostEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreGetUpdateCostEvent)
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
		it.Event = new(StoreGetUpdateCostEvent)
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
func (it *StoreGetUpdateCostEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreGetUpdateCostEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreGetUpdateCostEvent represents a GetUpdateCostEvent event raised by the Store contract.
type StoreGetUpdateCostEvent struct {
	State TransferState
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGetUpdateCostEvent is a free log retrieval operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) FilterGetUpdateCostEvent(opts *bind.FilterOpts) (*StoreGetUpdateCostEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "GetUpdateCostEvent")
	if err != nil {
		return nil, err
	}
	return &StoreGetUpdateCostEventIterator{contract: _Store.contract, event: "GetUpdateCostEvent", logs: logs, sub: sub}, nil
}

// WatchGetUpdateCostEvent is a free log subscription operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) WatchGetUpdateCostEvent(opts *bind.WatchOpts, sink chan<- *StoreGetUpdateCostEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "GetUpdateCostEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreGetUpdateCostEvent)
				if err := _Store.contract.UnpackLog(event, "GetUpdateCostEvent", log); err != nil {
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

// ParseGetUpdateCostEvent is a log parse operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) ParseGetUpdateCostEvent(log types.Log) (*StoreGetUpdateCostEvent, error) {
	event := new(StoreGetUpdateCostEvent)
	if err := _Store.contract.UnpackLog(event, "GetUpdateCostEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyHeaderAddIterator is returned from FilterNotifyHeaderAdd and is used to iterate over the raw logs and unpacked data for NotifyHeaderAdd events raised by the Store contract.
type StoreNotifyHeaderAddIterator struct {
	Event *StoreNotifyHeaderAdd // Event containing the contract specifics and raw log

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
func (it *StoreNotifyHeaderAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyHeaderAdd)
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
		it.Event = new(StoreNotifyHeaderAdd)
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
func (it *StoreNotifyHeaderAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyHeaderAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyHeaderAdd represents a NotifyHeaderAdd event raised by the Store contract.
type StoreNotifyHeaderAdd struct {
	Owner  common.Address
	Header []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyHeaderAdd is a free log retrieval operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) FilterNotifyHeaderAdd(opts *bind.FilterOpts) (*StoreNotifyHeaderAddIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyHeaderAdd")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyHeaderAddIterator{contract: _Store.contract, event: "NotifyHeaderAdd", logs: logs, sub: sub}, nil
}

// WatchNotifyHeaderAdd is a free log subscription operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) WatchNotifyHeaderAdd(opts *bind.WatchOpts, sink chan<- *StoreNotifyHeaderAdd) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyHeaderAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyHeaderAdd)
				if err := _Store.contract.UnpackLog(event, "NotifyHeaderAdd", log); err != nil {
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

// ParseNotifyHeaderAdd is a log parse operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) ParseNotifyHeaderAdd(log types.Log) (*StoreNotifyHeaderAdd, error) {
	event := new(StoreNotifyHeaderAdd)
	if err := _Store.contract.UnpackLog(event, "NotifyHeaderAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyHeaderTransferIterator is returned from FilterNotifyHeaderTransfer and is used to iterate over the raw logs and unpacked data for NotifyHeaderTransfer events raised by the Store contract.
type StoreNotifyHeaderTransferIterator struct {
	Event *StoreNotifyHeaderTransfer // Event containing the contract specifics and raw log

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
func (it *StoreNotifyHeaderTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyHeaderTransfer)
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
		it.Event = new(StoreNotifyHeaderTransfer)
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
func (it *StoreNotifyHeaderTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyHeaderTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyHeaderTransfer represents a NotifyHeaderTransfer event raised by the Store contract.
type StoreNotifyHeaderTransfer struct {
	From   common.Address
	To     common.Address
	Header []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyHeaderTransfer is a free log retrieval operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) FilterNotifyHeaderTransfer(opts *bind.FilterOpts) (*StoreNotifyHeaderTransferIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyHeaderTransfer")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyHeaderTransferIterator{contract: _Store.contract, event: "NotifyHeaderTransfer", logs: logs, sub: sub}, nil
}

// WatchNotifyHeaderTransfer is a free log subscription operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) WatchNotifyHeaderTransfer(opts *bind.WatchOpts, sink chan<- *StoreNotifyHeaderTransfer) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyHeaderTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyHeaderTransfer)
				if err := _Store.contract.UnpackLog(event, "NotifyHeaderTransfer", log); err != nil {
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

// ParseNotifyHeaderTransfer is a log parse operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) ParseNotifyHeaderTransfer(log types.Log) (*StoreNotifyHeaderTransfer, error) {
	event := new(StoreNotifyHeaderTransfer)
	if err := _Store.contract.UnpackLog(event, "NotifyHeaderTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoAddIterator is returned from FilterNotifyNameInfoAdd and is used to iterate over the raw logs and unpacked data for NotifyNameInfoAdd events raised by the Store contract.
type StoreNotifyNameInfoAddIterator struct {
	Event *StoreNotifyNameInfoAdd // Event containing the contract specifics and raw log

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
func (it *StoreNotifyNameInfoAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoAdd)
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
		it.Event = new(StoreNotifyNameInfoAdd)
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
func (it *StoreNotifyNameInfoAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoAdd represents a NotifyNameInfoAdd event raised by the Store contract.
type StoreNotifyNameInfoAdd struct {
	Owner common.Address
	Url   []byte
	Newer NameInfo
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoAdd is a free log retrieval operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) FilterNotifyNameInfoAdd(opts *bind.FilterOpts) (*StoreNotifyNameInfoAddIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoAdd")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoAddIterator{contract: _Store.contract, event: "NotifyNameInfoAdd", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoAdd is a free log subscription operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) WatchNotifyNameInfoAdd(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoAdd) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoAdd)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoAdd", log); err != nil {
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

// ParseNotifyNameInfoAdd is a log parse operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) ParseNotifyNameInfoAdd(log types.Log) (*StoreNotifyNameInfoAdd, error) {
	event := new(StoreNotifyNameInfoAdd)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoChangeIterator is returned from FilterNotifyNameInfoChange and is used to iterate over the raw logs and unpacked data for NotifyNameInfoChange events raised by the Store contract.
type StoreNotifyNameInfoChangeIterator struct {
	Event *StoreNotifyNameInfoChange // Event containing the contract specifics and raw log

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
func (it *StoreNotifyNameInfoChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoChange)
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
		it.Event = new(StoreNotifyNameInfoChange)
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
func (it *StoreNotifyNameInfoChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoChange represents a NotifyNameInfoChange event raised by the Store contract.
type StoreNotifyNameInfoChange struct {
	Owner common.Address
	Url   []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoChange is a free log retrieval operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) FilterNotifyNameInfoChange(opts *bind.FilterOpts) (*StoreNotifyNameInfoChangeIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoChange")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoChangeIterator{contract: _Store.contract, event: "NotifyNameInfoChange", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoChange is a free log subscription operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) WatchNotifyNameInfoChange(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoChange) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoChange)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoChange", log); err != nil {
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

// ParseNotifyNameInfoChange is a log parse operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) ParseNotifyNameInfoChange(log types.Log) (*StoreNotifyNameInfoChange, error) {
	event := new(StoreNotifyNameInfoChange)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoTransferIterator is returned from FilterNotifyNameInfoTransfer and is used to iterate over the raw logs and unpacked data for NotifyNameInfoTransfer events raised by the Store contract.
type StoreNotifyNameInfoTransferIterator struct {
	Event *StoreNotifyNameInfoTransfer // Event containing the contract specifics and raw log

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
func (it *StoreNotifyNameInfoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoTransfer)
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
		it.Event = new(StoreNotifyNameInfoTransfer)
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
func (it *StoreNotifyNameInfoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoTransfer represents a NotifyNameInfoTransfer event raised by the Store contract.
type StoreNotifyNameInfoTransfer struct {
	From common.Address
	To   common.Address
	Url  []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoTransfer is a free log retrieval operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) FilterNotifyNameInfoTransfer(opts *bind.FilterOpts) (*StoreNotifyNameInfoTransferIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoTransfer")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoTransferIterator{contract: _Store.contract, event: "NotifyNameInfoTransfer", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoTransfer is a free log subscription operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) WatchNotifyNameInfoTransfer(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoTransfer) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoTransfer)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoTransfer", log); err != nil {
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

// ParseNotifyNameInfoTransfer is a log parse operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) ParseNotifyNameInfoTransfer(log types.Log) (*StoreNotifyNameInfoTransfer, error) {
	event := new(StoreNotifyNameInfoTransfer)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreProveFileEventIterator is returned from FilterProveFileEvent and is used to iterate over the raw logs and unpacked data for ProveFileEvent events raised by the Store contract.
type StoreProveFileEventIterator struct {
	Event *StoreProveFileEvent // Event containing the contract specifics and raw log

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
func (it *StoreProveFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreProveFileEvent)
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
		it.Event = new(StoreProveFileEvent)
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
func (it *StoreProveFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreProveFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreProveFileEvent represents a ProveFileEvent event raised by the Store contract.
type StoreProveFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	Profit      uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProveFileEvent is a free log retrieval operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) FilterProveFileEvent(opts *bind.FilterOpts) (*StoreProveFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ProveFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreProveFileEventIterator{contract: _Store.contract, event: "ProveFileEvent", logs: logs, sub: sub}, nil
}

// WatchProveFileEvent is a free log subscription operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) WatchProveFileEvent(opts *bind.WatchOpts, sink chan<- *StoreProveFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ProveFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreProveFileEvent)
				if err := _Store.contract.UnpackLog(event, "ProveFileEvent", log); err != nil {
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

// ParseProveFileEvent is a log parse operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) ParseProveFileEvent(log types.Log) (*StoreProveFileEvent, error) {
	event := new(StoreProveFileEvent)
	if err := _Store.contract.UnpackLog(event, "ProveFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRegisterNodeEventIterator is returned from FilterRegisterNodeEvent and is used to iterate over the raw logs and unpacked data for RegisterNodeEvent events raised by the Store contract.
type StoreRegisterNodeEventIterator struct {
	Event *StoreRegisterNodeEvent // Event containing the contract specifics and raw log

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
func (it *StoreRegisterNodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRegisterNodeEvent)
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
		it.Event = new(StoreRegisterNodeEvent)
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
func (it *StoreRegisterNodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRegisterNodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRegisterNodeEvent represents a RegisterNodeEvent event raised by the Store contract.
type StoreRegisterNodeEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	NodeAddr    []byte
	Volume      uint64
	ServiceTime uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterNodeEvent is a free log retrieval operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) FilterRegisterNodeEvent(opts *bind.FilterOpts) (*StoreRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreRegisterNodeEventIterator{contract: _Store.contract, event: "RegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchRegisterNodeEvent is a free log subscription operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) WatchRegisterNodeEvent(opts *bind.WatchOpts, sink chan<- *StoreRegisterNodeEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRegisterNodeEvent)
				if err := _Store.contract.UnpackLog(event, "RegisterNodeEvent", log); err != nil {
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

// ParseRegisterNodeEvent is a log parse operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) ParseRegisterNodeEvent(log types.Log) (*StoreRegisterNodeEvent, error) {
	event := new(StoreRegisterNodeEvent)
	if err := _Store.contract.UnpackLog(event, "RegisterNodeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSetUserSpaceEventIterator is returned from FilterSetUserSpaceEvent and is used to iterate over the raw logs and unpacked data for SetUserSpaceEvent events raised by the Store contract.
type StoreSetUserSpaceEventIterator struct {
	Event *StoreSetUserSpaceEvent // Event containing the contract specifics and raw log

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
func (it *StoreSetUserSpaceEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSetUserSpaceEvent)
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
		it.Event = new(StoreSetUserSpaceEvent)
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
func (it *StoreSetUserSpaceEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetUserSpaceEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSetUserSpaceEvent represents a SetUserSpaceEvent event raised by the Store contract.
type StoreSetUserSpaceEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SizeType    uint8
	Size        uint64
	CountType   uint8
	Count       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetUserSpaceEvent is a free log retrieval operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) FilterSetUserSpaceEvent(opts *bind.FilterOpts) (*StoreSetUserSpaceEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SetUserSpaceEvent")
	if err != nil {
		return nil, err
	}
	return &StoreSetUserSpaceEventIterator{contract: _Store.contract, event: "SetUserSpaceEvent", logs: logs, sub: sub}, nil
}

// WatchSetUserSpaceEvent is a free log subscription operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) WatchSetUserSpaceEvent(opts *bind.WatchOpts, sink chan<- *StoreSetUserSpaceEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SetUserSpaceEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSetUserSpaceEvent)
				if err := _Store.contract.UnpackLog(event, "SetUserSpaceEvent", log); err != nil {
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

// ParseSetUserSpaceEvent is a log parse operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) ParseSetUserSpaceEvent(log types.Log) (*StoreSetUserSpaceEvent, error) {
	event := new(StoreSetUserSpaceEvent)
	if err := _Store.contract.UnpackLog(event, "SetUserSpaceEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreStoreFileEventIterator is returned from FilterStoreFileEvent and is used to iterate over the raw logs and unpacked data for StoreFileEvent events raised by the Store contract.
type StoreStoreFileEventIterator struct {
	Event *StoreStoreFileEvent // Event containing the contract specifics and raw log

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
func (it *StoreStoreFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStoreFileEvent)
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
		it.Event = new(StoreStoreFileEvent)
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
func (it *StoreStoreFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStoreFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStoreFileEvent represents a StoreFileEvent event raised by the Store contract.
type StoreStoreFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	FileSize    uint64
	WalletAddr  common.Address
	Cost        uint64
	IsPlotFile  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStoreFileEvent is a free log retrieval operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) FilterStoreFileEvent(opts *bind.FilterOpts) (*StoreStoreFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "StoreFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreStoreFileEventIterator{contract: _Store.contract, event: "StoreFileEvent", logs: logs, sub: sub}, nil
}

// WatchStoreFileEvent is a free log subscription operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) WatchStoreFileEvent(opts *bind.WatchOpts, sink chan<- *StoreStoreFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "StoreFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStoreFileEvent)
				if err := _Store.contract.UnpackLog(event, "StoreFileEvent", log); err != nil {
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

// ParseStoreFileEvent is a log parse operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) ParseStoreFileEvent(log types.Log) (*StoreStoreFileEvent, error) {
	event := new(StoreStoreFileEvent)
	if err := _Store.contract.UnpackLog(event, "StoreFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreUnRegisterNodeEventIterator is returned from FilterUnRegisterNodeEvent and is used to iterate over the raw logs and unpacked data for UnRegisterNodeEvent events raised by the Store contract.
type StoreUnRegisterNodeEventIterator struct {
	Event *StoreUnRegisterNodeEvent // Event containing the contract specifics and raw log

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
func (it *StoreUnRegisterNodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreUnRegisterNodeEvent)
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
		it.Event = new(StoreUnRegisterNodeEvent)
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
func (it *StoreUnRegisterNodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreUnRegisterNodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreUnRegisterNodeEvent represents a UnRegisterNodeEvent event raised by the Store contract.
type StoreUnRegisterNodeEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnRegisterNodeEvent is a free log retrieval operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) FilterUnRegisterNodeEvent(opts *bind.FilterOpts) (*StoreUnRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "UnRegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreUnRegisterNodeEventIterator{contract: _Store.contract, event: "UnRegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchUnRegisterNodeEvent is a free log subscription operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) WatchUnRegisterNodeEvent(opts *bind.WatchOpts, sink chan<- *StoreUnRegisterNodeEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "UnRegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreUnRegisterNodeEvent)
				if err := _Store.contract.UnpackLog(event, "UnRegisterNodeEvent", log); err != nil {
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

// ParseUnRegisterNodeEvent is a log parse operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) ParseUnRegisterNodeEvent(log types.Log) (*StoreUnRegisterNodeEvent, error) {
	event := new(StoreUnRegisterNodeEvent)
	if err := _Store.contract.UnpackLog(event, "UnRegisterNodeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
