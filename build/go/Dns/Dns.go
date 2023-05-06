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
	_ = abi.ConvertType
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Pos\",\"type\":\"uint64\"}],\"internalType\":\"structChangeInitPosParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"AddInitPos\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"ApproveDNSCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"CreateDefaultUrl\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"DNSNodeReg\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"DelDNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"DelHeader\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"DnsInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllDnsNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetDNSNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"GetHeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"HeaderOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structHeaderInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"GetName\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structNameInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"GetPeerPoolItem\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"WalletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"TotalInitPos\",\"type\":\"uint64\"}],\"internalType\":\"structPeerPoolItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPeerPoolMap\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"WalletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"TotalInitPos\",\"type\":\"uint64\"}],\"internalType\":\"structPeerPoolItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPluginList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structNameInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"GetUrl\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"internalType\":\"structQuitNodeParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"QuitNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Pos\",\"type\":\"uint64\"}],\"internalType\":\"structChangeInitPosParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"ReduceInitPos\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestHeader\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"RegisterHeader\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestName\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"RegisterName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"RejectDNSCandidate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"}],\"internalType\":\"structTransferInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"TransferHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"}],\"internalType\":\"structTransferInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"TransferName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"internalType\":\"structUnRegisterCandidateParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"UnRegDNSNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"}],\"internalType\":\"structUpdateNodeParam\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"UpdateDNSNodesInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestName\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"UpdateName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b2\",\"type\":\"bytes\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"toBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615391806100206000396000f3fe6080604052600436106101b75760003560e01c80636b9ceb53116100ec578063ba47d33d1161008a578063d72b6be611610064578063d72b6be6146104b5578063dbc8ef9c146104d5578063eafafaf7146104e8578063fd11c4c6146104fb57600080fd5b8063ba47d33d14610453578063ccfd72f614610473578063d68414701461049557600080fd5b80638129fc1c116100c65780638129fc1c146103c457806387dacef2146103d95780638b79782b14610406578063919064eb1461042657600080fd5b80636b9ceb53146103715780636e91c5681461039e578063809bae57146103b157600080fd5b8063370d7104116101595780634f2ffbba116101335780634f2ffbba14610309578063513728191461031c5780635f3376f31461032f5780636839a7c61461034f57600080fd5b8063370d7104146102c35780633e4e5e54146102d6578063465efc9a146102f657600080fd5b8063240ce9f911610195578063240ce9f9146102295780632a5da31e146102565780632c261f4214610283578063326cf61c146102a357600080fd5b80630e580e79146101bc5780631654a894146101e7578063228a09b214610209575b600080fd5b3480156101c857600080fd5b506101d161050e565b6040516101de9190614f2f565b60405180910390f35b3480156101f357600080fd5b50610207610202366004614729565b610601565b005b34801561021557600080fd5b50610207610224366004614522565b6106fb565b34801561023557600080fd5b50610249610244366004614522565b610775565b6040516101de919061513d565b34801561026257600080fd5b50610276610271366004614556565b6107aa565b6040516101de9190614f40565b34801561028f57600080fd5b5061020761029e366004614625565b610804565b3480156102af57600080fd5b506102766102be3660046144e6565b610941565b6102076102d13660046146c1565b61096a565b3480156102e257600080fd5b506102076102f13660046146f5565b610ab5565b6102076103043660046145f1565b610bc2565b610207610317366004614659565b610cf2565b61020761032a3660046145bd565b610d5b565b34801561033b57600080fd5b5061027661034a366004614556565b610e99565b34801561035b57600080fd5b5061036461105d565b6040516101de9190614f0d565b34801561037d57600080fd5b5061039161038c366004614659565b611152565b6040516101de919061512c565b6102076103ac366004614659565b6111dd565b6102076103bf36600461468d565b611409565b3480156103d057600080fd5b5061020761156d565b3480156103e557600080fd5b506103f96103f43660046144c8565b61172e565b6040516101de919061510a565b34801561041257600080fd5b506102076104213660046146f5565b611763565b34801561043257600080fd5b50610446610441366004614659565b611ab0565b6040516101de919061511b565b34801561045f57600080fd5b5061020761046e3660046144c8565b611c7d565b34801561047f57600080fd5b50610488611e62565b6040516101de9190614f1e565b3480156104a157600080fd5b506102076104b0366004614625565b611f90565b3480156104c157600080fd5b506102766104d0366004614522565b612088565b6102076104e3366004614522565b6120dd565b6102076104f63660046145bd565b61217a565b6102076105093660046146c1565b61226f565b60606000600a600201546001600160401b0381111561053d57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561059057816020015b6040805160808101825260608082526000602080840182905293830181905290820152825260001990920191018161055b5790505b509050600061059f600a61272d565b90505b600b548110156105fb5760006105b9600a83612748565b915050808383815181106105dd57634e487b7160e01b600052603260045260246000fd5b6020908102919091010152506105f4600a82612931565b90506105a2565b50919050565b60008160200151511161062f5760405162461bcd60e51b815260040161062690614fb0565b60405180910390fd5b6000816040015151116106545760405162461bcd60e51b815260040161062690614fc0565b6000610661600d336129a1565b80519091506001600160a01b0316331461068d5760405162461bcd60e51b815260040161062690614fe0565b60208083015190820152604080830151908201526106ad600d3383612bd4565b507f602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf826020015183604001513384606001516040516106ef9493929190614f51565b60405180910390a15050565b6000610708600a83612d49565b9050600181606001516001600160401b031610156107385760405162461bcd60e51b8152600401610626906150ea565b604081015160ff161561075d5760405162461bcd60e51b815260040161062690614fe0565b60026040820152610770600a8383612e69565b505050565b6040805160808101825260608082526000602083018190529282018390528101919091526107a4600a83612d49565b92915050565b606060006107ed846040518060400160405280600381526020017f3a2f2f0000000000000000000000000000000000000000000000000000000000815250610e99565b905060006107fb8285610e99565b95945050505050565b60208101516001600160a01b031633146108305760405162461bcd60e51b8152600401610626906150da565b805160009061084190600a90612d49565b604081015190915060ff16156108695760405162461bcd60e51b815260040161062690614fe0565b81602001516001600160a01b031681602001516001600160a01b0316146108a25760405162461bcd60e51b8152600401610626906150da565b81602001516001600160a01b03166108fc82606001516001600160401b03169081150290604051600060405180830381858888f193505050501580156108ec573d6000803e3d6000fd5b5081516108fb90600a90613007565b50602082015161090d90600d90613107565b507f7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d82602001516040516106ef9190614e7e565b6060816040516020016109549190614e51565b6040516020818303038152906040529050919050565b600061097e82602001518360400151610e99565b9050600061098d60068361321e565b90508060c00151600014156109d6577fd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e356040516109c990615074565b60405180910390a1505050565b82608001516001600160a01b031681608001516001600160a01b031614610a24577fd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e356040516109c990615097565b82516001600160401b0390811682526060808501519083015260a0808501519083015260c08401511660e0820152610a5d436001615194565b60c0820152610a6e60068383613544565b507f0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe081508360800151610aa7856020015186604001516107aa565b6040516109c9929190614eb9565b6000610ac982600001518360200151610e99565b90506000610ad860068361321e565b905082604001516001600160a01b031681608001516001600160a01b031614610b135760405162461bcd60e51b8152600401610626906150da565b60608301516001600160a01b03166080820152610b31436001615194565b60c0820181905260e08201514391610b51916001600160401b0316615194565b610b5b91906151df565b6001600160401b031660e0820152610b7560068383613544565b507fa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f83604001518460600151610bb3866000015187602001516107aa565b6040516109c993929190614e8c565b600081606001516001600160401b031611610bef5760405162461bcd60e51b8152600401610626906150ca565b80606001516001600160401b0316341015610c1c5760405162461bcd60e51b815260040161062690614fa0565b80516001600160a01b03163314610c455760405162461bcd60e51b8152600401610626906150da565b604080516080808201835260608083526000602084018181529484018190528184019081529185018051845285516001600160a01b03169094528401516001600160401b031690529051610c9c90600a9083612e69565b508151610cac90600d9084612bd4565b507f602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf82602001518360400151846000015185606001516040516106ef9493929190614f51565b6000610d0682600001518360200151610e99565b90506000610d1560068361321e565b60045460808201519192506001600160a01b03918216911614610d4a5760405162461bcd60e51b8152600401610626906150ba565b610d55600683613684565b50505050565b600081604001516001600160401b031611610d885760405162461bcd60e51b8152600401610626906150fa565b80604001516001600160401b0316341015610db55760405162461bcd60e51b815260040161062690614fa0565b60208101516001600160a01b03163314610de15760405162461bcd60e51b8152600401610626906150da565b8051600090610df290600a90612d49565b905081602001516001600160a01b031681602001516001600160a01b031614610e2d5760405162461bcd60e51b8152600401610626906150da565b604081015160ff16600214801590610e4b5750604081015160ff1615155b15610e685760405162461bcd60e51b815260040161062690614fd0565b816040015181606001818151610e7e91906151ac565b6001600160401b0316905250815161077090600a9083612e69565b81518151606091906000610ead8284615194565b6001600160401b03811115610ed257634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015610efc576020820181803683370190505b50905060005b83811015610f9657868181518110610f2a57634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b828281518110610f5557634e487b7160e01b600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080610f8e816152d5565b915050610f02565b5060005b8281101561105357858181518110610fc257634e487b7160e01b600052603260045260246000fd5b01602001517fff000000000000000000000000000000000000000000000000000000000000001682610ff48684615194565b8151811061101257634e487b7160e01b600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508061104b816152d5565b915050610f9a565b5095945050505050565b60606000600d600201546001600160401b0381111561108c57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156110e757816020015b6040805160a08101825260008082526060602083018190529282018390528282015260808101919091528152602001906001900390816110aa5790505b50905060006110f6600d613767565b90505b600e548110156105fb576000611110600d83613775565b9150508083838151811061113457634e487b7160e01b600052603260045260246000fd5b60209081029190910101525061114b600d826139f4565b90506110f9565b6111b560405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b60006111c983600001518460200151610e99565b90506111d660068261321e565b9392505050565b6000600582600001516040516111f39190614e66565b90815260200160405180910390206040518060a001604052908160008201805461121c90615282565b80601f016020809104026020016040519081016040528092919081815260200182805461124890615282565b80156112955780601f1061126a57610100808354040283529160200191611295565b820191906000526020600020905b81548152906001019060200180831161127857829003601f168201915b505050918352505060018201546001600160a01b031660208201526002820180546040909201916112c590615282565b80601f01602080910402602001604051908101604052809291908181526020018280546112f190615282565b801561133e5780601f106113135761010080835404028352916020019161133e565b820191906000526020600020905b81548152906001019060200180831161132157829003601f168201915b505050918352505060038201546020808301919091526004928301546001600160401b03166040909201919091529054908201519192506001600160a01b0391821691161461139f5760405162461bcd60e51b8152600401610626906150ba565b81516040516005916113b091614e66565b90815260405190819003602001902060006113cb8282613ea8565b6001820180546001600160a01b03191690556113eb600283016000613ea8565b5060006003820155600401805467ffffffffffffffff191690555050565b61144d6040518060a001604052806060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b815181526020808301516001600160a01b03169082015260408083015190820152611479436001615194565b606082015260006080820152815160405182916005916114999190614e66565b908152602001604051809103902060008201518160000190805190602001906114c3929190613ee2565b506020828101516001830180546001600160a01b0319166001600160a01b03909216919091179055604083015180516115029260028501920190613ee2565b50606082015160038201556080909101516004909101805467ffffffffffffffff19166001600160401b03909216919091179055602082015182516040517f356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869926106ef929091614eb9565b600054610100900460ff161580801561158d5750600054600160ff909116105b806115a75750303b1580156115a7575060005460ff166001145b6115c35760405162461bcd60e51b815260040161062690614ff0565b6000805460ff1916600117905580156115e6576000805461ff0019166101001790555b600080547fffffffffffff000000000000000000000000000000000000000000000000ffff167202000000000000000100000000000000000000179055600180547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016680800000000000000041790556040805180820190915260038082527f6473700000000000000000000000000000000000000000000000000000000000602090920191825261169a91600291613ee2565b5060408051808201909152600a8082527f6473702d706c7567696e0000000000000000000000000000000000000000000060209092019182526116df91600391613ee2565b50801561172b576000805461ff00191690556040517f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989061172290600190614f92565b60405180910390a15b50565b6040805160a08101825260008082526060602083018190529282018390528282015260808101919091526107a4600d836129a1565b6000600582600001516040516117799190614e66565b90815260200160405180910390206040518060a00160405290816000820180546117a290615282565b80601f01602080910402602001604051908101604052809291908181526020018280546117ce90615282565b801561181b5780601f106117f05761010080835404028352916020019161181b565b820191906000526020600020905b8154815290600101906020018083116117fe57829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161184b90615282565b80601f016020809104026020016040519081016040528092919081815260200182805461187790615282565b80156118c45780601f10611899576101008083540402835291602001916118c4565b820191906000526020600020905b8154815290600101906020018083116118a757829003601f168201915b505050918352505060038201546020808301919091526004909201546001600160401b0316604091820152840151908201519192506001600160a01b039182169116146119235760405162461bcd60e51b8152600401610626906150da565b600043826060015183608001516001600160401b03166119439190615194565b116119505750600061197b565b43826060015183608001516001600160401b031661196e9190615194565b61197891906151df565b90505b825160405160059061198e908390614e66565b908152602001604051809103902060000190805190602001906119b2929190613ee2565b50606083015183516040516005916119c991614e66565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550436001611a0b9190615194565b8351604051600591611a1c91614e66565b9081526040519081900360200181206003019190915583518291600591611a4291614e66565b9081526040805191829003602001822060040180546001600160401b039490941667ffffffffffffffff199094169390931790925590840151606085015185517f178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca936109c993929190614e8c565b611af46040518060a001604052806060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b8151604051600591611b0591614e66565b90815260200160405180910390206040518060a0016040529081600082018054611b2e90615282565b80601f0160208091040260200160405190810160405280929190818152602001828054611b5a90615282565b8015611ba75780601f10611b7c57610100808354040283529160200191611ba7565b820191906000526020600020905b815481529060010190602001808311611b8a57829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191611bd790615282565b80601f0160208091040260200160405190810160405280929190818152602001828054611c0390615282565b8015611c505780601f10611c2557610100808354040283529160200191611c50565b820191906000526020600020905b815481529060010190602001808311611c3357829003601f168201915b5050509183525050600382015460208201526004909101546001600160401b031660409091015292915050565b600480546001600160a01b0319166001600160a01b0383161790556040805160a08101825260608082526000602083018190529282018190528101829052608081019190915260028054611cd090615282565b80601f0160208091040260200160405190810160405280929190818152602001828054611cfc90615282565b8015611d495780601f10611d1e57610100808354040283529160200191611d49565b820191906000526020600020905b815481529060010190602001808311611d2c57829003601f168201915b50505091835250506001600160a01b038216602080830191909152604080518082018252601581527f7265736572766564206473702070726f746f636f6c000000000000000000000092810192909252808301919091526000606083018190526080830152815190518291600591611dc19190614e66565b90815260200160405180910390206000820151816000019080519060200190611deb929190613ee2565b506020828101516001830180546001600160a01b0319166001600160a01b0390921691909117905560408301518051611e2a9260028501920190613ee2565b50606082015160038201556080909101516004909101805467ffffffffffffffff19166001600160401b039092169190911790555050565b606060006006600201546001600160401b03811115611e9157634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611f2557816020015b611f1260405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b815260200190600190039081611eaf5790505b5090506000611f346006613a63565b90505b6007548110156105fb576000611f4e600683613a71565b91505080838381518110611f7257634e487b7160e01b600052603260045260246000fd5b602090810291909101015250611f89600682613e38565b9050611f37565b60208101516001600160a01b03163314611fbc5760405162461bcd60e51b8152600401610626906150da565b8051600090611fcd90600a90612d49565b905081602001516001600160a01b031681602001516001600160a01b0316146120085760405162461bcd60e51b8152600401610626906150da565b604081015160ff166002148015906120265750604081015160ff1615155b156120435760405162461bcd60e51b815260040161062690614fd0565b604081015160ff166002141561205f5760036040820152612067565b600460408201525b815161207690600a9083612e69565b50602082015161077090600d90613107565b60606107a460028360405161209d9190614e66565b602060405180830381855afa1580156120ba573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906102be9190614504565b60006120ea600a83612d49565b604081015190915060ff16156121125760405162461bcd60e51b815260040161062690614fe0565b80602001516001600160a01b03166108fc82606001516001600160401b03169081150290604051600060405180830381858888f1935050505015801561215c573d6000803e3d6000fd5b50612168600a83613007565b50602081015161077090600d90613107565b600081604001516001600160401b0316116121a75760405162461bcd60e51b8152600401610626906150fa565b60208101516001600160a01b031633146121d35760405162461bcd60e51b8152600401610626906150da565b80516000906121e490600a90612d49565b905081602001516001600160a01b031681602001516001600160a01b03161461221f5760405162461bcd60e51b8152600401610626906150da565b81604001516001600160401b031681606001516001600160401b031610156122595760405162461bcd60e51b8152600401610626906150ea565b816040015181606001818151610e7e91906151fa565b600481606001515110156122aa577fd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e3560405161172290615051565b61230d60405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b60005482516001600160401b039081166201000090920416141561242157600081526002805461233c90615282565b80601f016020809104026020016040519081016040528092919081815260200182805461236890615282565b80156123b55780601f1061238a576101008083540402835291602001916123b5565b820191906000526020600020905b81548152906001019060200180831161239857829003601f168201915b505050505081602001819052506123cf8260600151612088565b6040820152606080830151908201526080808301516001600160a01b03169082015260a08083015190820152612406436001615194565b60c0808301919091528201516001600160401b031660e08201525b60005482516001600160401b039081166a01000000000000000000009092041614156124b5576000815260208281015190820152606082015161246390612088565b6040820152606080830151908201526080808301516001600160a01b03169082015260a0808301519082015261249a436001615194565b60c0808301919091528201516001600160401b031660e08201525b60005482516001600160401b0390811672010000000000000000000000000000000000009092041614156125ce5760008152600280546124f490615282565b80601f016020809104026020016040519081016040528092919081815260200182805461252090615282565b801561256d5780601f106125425761010080835404028352916020019161256d565b820191906000526020600020905b81548152906001019060200180831161255057829003601f168201915b5050505050602082015260408281015190820152606080830151908201526080808301516001600160a01b03169082015260a080830151908201526125b3436001615194565b60c0808301919091528201516001600160401b031660e08201525b60015482516001600160401b039081169116141561264b57600081526020828101519082015260408083015190820152606080830151908201526080808301516001600160a01b03169082015260a08083015190820152612630436001615194565b60c0808301919091528201516001600160401b031660e08201525b600061265f83602001518460400151610e99565b905061266d60068284613544565b5060015483516001600160401b0390811691161480156126ad575060036040516126979190614e72565b6040518091039020836020015180519060200120145b156126e55760016009826040516126c49190614e66565b908152604051908190036020019020805491151560ff199092169190911790555b7f517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41836080015161271d846020015185604001516107aa565b846040516109c993929190614ed9565b60008061273b836000612931565b90506111d66001826151df565b6040805160808101825260608082526000602083018190529282018390528181019290925283600101838154811061279057634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160000180546127ac90615282565b80601f01602080910402602001604051908101604052809291908181526020018280546127d890615282565b80156128255780601f106127fa57610100808354040283529160200191612825565b820191906000526020600020905b81548152906001019060200180831161280857829003601f168201915b50505050509150836000018260405161283e9190614e66565b908152602001604051809103902060010160405180608001604052908160008201805461286a90615282565b80601f016020809104026020016040519081016040528092919081815260200182805461289690615282565b80156128e35780601f106128b8576101008083540402835291602001916128e3565b820191906000526020600020905b8154815290600101906020018083116128c657829003601f168201915b5050509183525050600191909101546001600160a01b0381166020830152600160a01b810460ff166040830152600160a81b90046001600160401b0316606090910152919491935090915050565b60008161293d816152d5565b9250505b60018301548210801561298a575082600101828154811061297257634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b156105fb5781612999816152d5565b925050612941565b6040805160a08101825260008082526060602083018190529282018390528282015260808101919091526001600160a01b0380831660009081526020858152604091829020825160a081019093526001810180549094168352600201805492939291840191612a0f90615282565b80601f0160208091040260200160405190810160405280929190818152602001828054612a3b90615282565b8015612a885780601f10612a5d57610100808354040283529160200191612a88565b820191906000526020600020905b815481529060010190602001808311612a6b57829003601f168201915b50505050508152602001600282018054612aa190615282565b80601f0160208091040260200160405190810160405280929190818152602001828054612acd90615282565b8015612b1a5780601f10612aef57610100808354040283529160200191612b1a565b820191906000526020600020905b815481529060010190602001808311612afd57829003601f168201915b505050918352505060038201546001600160401b03166020820152600482018054604090920191612b4a90615282565b80601f0160208091040260200160405190810160405280929190818152602001828054612b7690615282565b8015612bc35780601f10612b9857610100808354040283529160200191612bc3565b820191906000526020600020905b815481529060010190602001808311612ba657829003601f168201915b505050505081525050905092915050565b6001600160a01b038281166000908152602085815260408220805485516001830180546001600160a01b031916919096161785558583015180519495919487949293612c27936002909101920190613ee2565b5060408201518051612c43916002840191602090910190613ee2565b50606082015160038201805467ffffffffffffffff19166001600160401b0390921691909117905560808201518051612c86916004840191602090910190613ee2565b505081159050612c9a5760019150506111d6565b5060018085018054808301825560009190915290612cb9908290615194565b6001600160a01b03851660009081526020879052604090205560018501805485919083908110612cf957634e487b7160e01b600052603260045260246000fd5b6000918252602082200180546001600160a01b0319166001600160a01b03939093169290921790915560028601805491612d32836152d5565b919050555060009150506111d6565b509392505050565b6040805160808101825260608082526000602083018190528284018190529082015290518390612d7a908490614e66565b9081526020016040518091039020600101604051806080016040529081600082018054612da690615282565b80601f0160208091040260200160405190810160405280929190818152602001828054612dd290615282565b8015612e1f5780601f10612df457610100808354040283529160200191612e1f565b820191906000526020600020905b815481529060010190602001808311612e0257829003601f168201915b5050509183525050600191909101546001600160a01b0381166020830152600160a01b810460ff166040830152600160a81b90046001600160401b03166060909101529392505050565b6000808460000184604051612e7e9190614e66565b90815260405190819003602001812054915083908690612e9f908790614e66565b90815260200160405180910390206001016000820151816000019080519060200190612ecc929190613ee2565b5060208201516001909101805460408401516060909401516001600160401b0316600160a81b027fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff60ff909516600160a01b027fffffffffffffffffffffff0000000000000000000000000000000000000000009092166001600160a01b039094169390931717929092161790558015612f6a5760019150506111d6565b5060018085018054808301825560009190915290612f89908290615194565b6040518690612f99908790614e66565b9081526040519081900360200190205560018501805485919083908110612fd057634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016000019080519060200190612ff4929190613ee2565b50600285018054906000612d32836152d5565b600080836000018360405161301c9190614e66565b9081526040519081900360200190205490508061303d5760009150506107a4565b604051849061304d908590614e66565b908152604051908190036020019020600080825560018201816130708282613ea8565b50600190810180547fffffff00000000000000000000000000000000000000000000000000000000001690559150508481016130ac82846151df565b815481106130ca57634e487b7160e01b600052603260045260246000fd5b600091825260208220600291820201600101805460ff1916931515939093179092559085018054916130fb8361526b565b91905055505092915050565b6001600160a01b0381166000908152602083905260408120548061312f5760009150506107a4565b6001600160a01b03831660009081526020859052604081208181556001810180546001600160a01b03191681559091908161316d6002850182613ea8565b61317b600283016000613ea8565b60038201805467ffffffffffffffff1916905561319c600483016000613ea8565b506001925050508481016131b082846151df565b815481106131ce57634e487b7160e01b600052603260045260246000fd5b600091825260208220018054921515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff90931692909217909155600285018054916130fb8361526b565b61328160405180610100016040528060006001600160401b0316815260200160608152602001606081526020016060815260200160006001600160a01b03168152602001606081526020016000815260200160006001600160401b031681525090565b6040518390613291908490614e66565b90815260408051918290036020908101832061010084019092526001820180546001600160401b0316845260029092018054918401916132d090615282565b80601f01602080910402602001604051908101604052809291908181526020018280546132fc90615282565b80156133495780601f1061331e57610100808354040283529160200191613349565b820191906000526020600020905b81548152906001019060200180831161332c57829003601f168201915b5050505050815260200160028201805461336290615282565b80601f016020809104026020016040519081016040528092919081815260200182805461338e90615282565b80156133db5780601f106133b0576101008083540402835291602001916133db565b820191906000526020600020905b8154815290600101906020018083116133be57829003601f168201915b505050505081526020016003820180546133f490615282565b80601f016020809104026020016040519081016040528092919081815260200182805461342090615282565b801561346d5780601f106134425761010080835404028352916020019161346d565b820191906000526020600020905b81548152906001019060200180831161345057829003601f168201915b505050918352505060048201546001600160a01b0316602082015260058201805460409092019161349d90615282565b80601f01602080910402602001604051908101604052809291908181526020018280546134c990615282565b80156135165780601f106134eb57610100808354040283529160200191613516565b820191906000526020600020905b8154815290600101906020018083116134f957829003601f168201915b5050509183525050600682015460208201526007909101546001600160401b03166040909101529392505050565b60008084600001846040516135599190614e66565b9081526040519081900360200181205491508390869061357a908790614e66565b90815260405160209181900382019020825160018201805467ffffffffffffffff19166001600160401b0390921691909117815583830151805191936135c69360020192910190613ee2565b50604082015180516135e2916002840191602090910190613ee2565b50606082015180516135fe916003840191602090910190613ee2565b5060808201516004820180546001600160a01b0319166001600160a01b0390921691909117905560a08201518051613640916005840191602090910190613ee2565b5060c0820151600682015560e0909101516007909101805467ffffffffffffffff19166001600160401b039092169190911790558015612f6a5760019150506111d6565b60008083600001836040516136999190614e66565b908152604051908190036020019020549050806136ba5760009150506107a4565b60405184906136ca908590614e66565b908152604051908190036020019020600080825560018201805467ffffffffffffffff19168155816136ff6002850182613ea8565b61370d600283016000613ea8565b61371b600383016000613ea8565b6004820180546001600160a01b031916905561373b600583016000613ea8565b5060006006820155600701805467ffffffffffffffff1916905550600190508481016130ac82846151df565b60008061273b8360006139f4565b60006137bb6040518060a0016040528060006001600160a01b03168152602001606081526020016060815260200160006001600160401b03168152602001606081525090565b8360010183815481106137de57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03908116808452878352604093849020845160a081019095526001810180549093168552600201805491965091928401919061382e90615282565b80601f016020809104026020016040519081016040528092919081815260200182805461385a90615282565b80156138a75780601f1061387c576101008083540402835291602001916138a7565b820191906000526020600020905b81548152906001019060200180831161388a57829003601f168201915b505050505081526020016002820180546138c090615282565b80601f01602080910402602001604051908101604052809291908181526020018280546138ec90615282565b80156139395780601f1061390e57610100808354040283529160200191613939565b820191906000526020600020905b81548152906001019060200180831161391c57829003601f168201915b505050918352505060038201546001600160401b0316602082015260048201805460409092019161396990615282565b80601f016020809104026020016040519081016040528092919081815260200182805461399590615282565b80156139e25780601f106139b7576101008083540402835291602001916139e2565b820191906000526020600020905b8154815290600101906020018083116139c557829003601f168201915b50505050508152505090509250929050565b600081613a00816152d5565b9250505b600183015482108015613a4c5750826001018281548110613a3557634e487b7160e01b600052603260045260246000fd5b600091825260209091200154600160a01b900460ff165b156105fb5781613a5b816152d5565b925050613a04565b60008061273b836000613e38565b604080516101008101825260008082526060602083018190529282018390528183018390526080820181905260a0820183905260c0820181905260e0820152836001018381548110613ad357634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016000018054613aef90615282565b80601f0160208091040260200160405190810160405280929190818152602001828054613b1b90615282565b8015613b685780601f10613b3d57610100808354040283529160200191613b68565b820191906000526020600020905b815481529060010190602001808311613b4b57829003601f168201915b505050505091508360000182604051613b819190614e66565b90815260408051918290036020908101832061010084019092526001820180546001600160401b031684526002909201805491840191613bc090615282565b80601f0160208091040260200160405190810160405280929190818152602001828054613bec90615282565b8015613c395780601f10613c0e57610100808354040283529160200191613c39565b820191906000526020600020905b815481529060010190602001808311613c1c57829003601f168201915b50505050508152602001600282018054613c5290615282565b80601f0160208091040260200160405190810160405280929190818152602001828054613c7e90615282565b8015613ccb5780601f10613ca057610100808354040283529160200191613ccb565b820191906000526020600020905b815481529060010190602001808311613cae57829003601f168201915b50505050508152602001600382018054613ce490615282565b80601f0160208091040260200160405190810160405280929190818152602001828054613d1090615282565b8015613d5d5780601f10613d3257610100808354040283529160200191613d5d565b820191906000526020600020905b815481529060010190602001808311613d4057829003601f168201915b505050918352505060048201546001600160a01b03166020820152600582018054604090920191613d8d90615282565b80601f0160208091040260200160405190810160405280929190818152602001828054613db990615282565b8015613e065780601f10613ddb57610100808354040283529160200191613e06565b820191906000526020600020905b815481529060010190602001808311613de957829003601f168201915b5050509183525050600682015460208201526007909101546001600160401b0316604090910152919491935090915050565b600081613e44816152d5565b9250505b600183015482108015613e915750826001018281548110613e7957634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b156105fb5781613ea0816152d5565b925050613e48565b508054613eb490615282565b6000825580601f10613ec4575050565b601f01602090049060005260206000209081019061172b9190613f66565b828054613eee90615282565b90600052602060002090601f016020900481019282613f105760008555613f56565b82601f10613f2957805160ff1916838001178555613f56565b82800160010185558215613f56579182015b82811115613f56578251825591602001919060010190613f3b565b50613f62929150613f66565b5090565b5b80821115613f625760008155600101613f67565b6000613f8e613f898461516a565b61514e565b905082815260208101848484011115613fa657600080fd5b612d41848285615233565b80356107a481615332565b80356107a481615346565b80516107a481615346565b600082601f830112613fe357600080fd5b8135613ff3848260208601613f7b565b949350505050565b60006060828403121561400d57600080fd5b614017606061514e565b905081356001600160401b0381111561402f57600080fd5b61403b84828501613fd2565b825250602061404c84848301613fb1565b6020830152506040614060848285016144bd565b60408301525092915050565b600060a0828403121561407e57600080fd5b61408860a061514e565b905060006140968484613fb1565b82525060208201356001600160401b038111156140b257600080fd5b6140be84828501613fd2565b60208301525060408201356001600160401b038111156140dd57600080fd5b6140e984828501613fd2565b60408301525060606140fd848285016144bd565b60608301525060808201356001600160401b0381111561411c57600080fd5b61412884828501613fd2565b60808301525092915050565b60006040828403121561414657600080fd5b614150604061514e565b905081356001600160401b0381111561416857600080fd5b61417484828501613fd2565b825250602061418584848301613fb1565b60208301525092915050565b6000606082840312156141a357600080fd5b6141ad606061514e565b905081356001600160401b038111156141c557600080fd5b6141d184828501613fd2565b82525060208201356001600160401b038111156141ed57600080fd5b6141f984828501613fd2565b602083015250604061406084828501613fb1565b60006080828403121561421f57600080fd5b614229608061514e565b905081356001600160401b0381111561424157600080fd5b61424d84828501613fd2565b825250602061425e84848301613fb1565b60208301525060408201356001600160401b0381111561427d57600080fd5b61428984828501613fd2565b604083015250606061429d848285016144bd565b60608301525092915050565b600060e082840312156142bb57600080fd5b6142c560e061514e565b905060006142d384846144bd565b82525060208201356001600160401b038111156142ef57600080fd5b6142fb84828501613fd2565b60208301525060408201356001600160401b0381111561431a57600080fd5b61432684828501613fd2565b60408301525060608201356001600160401b0381111561434557600080fd5b61435184828501613fd2565b606083015250608061436584828501613fb1565b60808301525060a08201356001600160401b0381111561438457600080fd5b61439084828501613fd2565b60a08301525060c06143a4848285016144bd565b60c08301525092915050565b6000608082840312156143c257600080fd5b6143cc608061514e565b905081356001600160401b038111156143e457600080fd5b6143f084828501613fd2565b82525060208201356001600160401b0381111561440c57600080fd5b61441884828501613fd2565b602083015250604061442c84828501613fb1565b604083015250606061429d84828501613fb1565b60006060828403121561445257600080fd5b61445c606061514e565b9050600061446a8484613fb1565b82525060208201356001600160401b0381111561448657600080fd5b61449284828501613fd2565b60208301525060408201356001600160401b038111156144b157600080fd5b61406084828501613fd2565b80356107a48161534c565b6000602082840312156144da57600080fd5b6000613ff38484613fb1565b6000602082840312156144f857600080fd5b6000613ff38484613fbc565b60006020828403121561451657600080fd5b6000613ff38484613fc7565b60006020828403121561453457600080fd5b81356001600160401b0381111561454a57600080fd5b613ff384828501613fd2565b6000806040838503121561456957600080fd5b82356001600160401b0381111561457f57600080fd5b61458b85828601613fd2565b92505060208301356001600160401b038111156145a757600080fd5b6145b385828601613fd2565b9150509250929050565b6000602082840312156145cf57600080fd5b81356001600160401b038111156145e557600080fd5b613ff384828501613ffb565b60006020828403121561460357600080fd5b81356001600160401b0381111561461957600080fd5b613ff38482850161406c565b60006020828403121561463757600080fd5b81356001600160401b0381111561464d57600080fd5b613ff384828501614134565b60006020828403121561466b57600080fd5b81356001600160401b0381111561468157600080fd5b613ff384828501614191565b60006020828403121561469f57600080fd5b81356001600160401b038111156146b557600080fd5b613ff38482850161420d565b6000602082840312156146d357600080fd5b81356001600160401b038111156146e957600080fd5b613ff3848285016142a9565b60006020828403121561470757600080fd5b81356001600160401b0381111561471d57600080fd5b613ff3848285016143b0565b60006020828403121561473b57600080fd5b81356001600160401b0381111561475157600080fd5b613ff384828501614440565b60006111d68383614c54565b60006111d68383614d34565b60006111d68383614dea565b61478a81615217565b82525050565b600061479a825190565b808452602084019350836020820285016147b48560200190565b8060005b858110156147e957848403895281516147d1858261475d565b94506020830160209a909a01999250506001016147b8565b5091979650505050505050565b6000614800825190565b8084526020840193508360208202850161481a8560200190565b8060005b858110156147e957848403895281516148378582614769565b94506020830160209a909a019992505060010161481e565b6000614859825190565b808452602084019350836020820285016148738560200190565b8060005b858110156147e957848403895281516148908582614775565b94506020830160209a909a0199925050600101614877565b8061478a565b60006148b8825190565b8084526020840193506148cf81856020860161523f565b601f01601f19169290920192915050565b60006148ea825190565b6148f881856020860161523f565b9290920192915050565b6000815461490f81615282565b600182168015614926576001811461493757614967565b60ff19831686528186019350614967565b60008581526020902060005b8381101561495f57815488820152600190910190602001614943565b838801955050505b50505092915050565b61478a81615228565b601081526000602082017f6465706f736974206d757374203e203000000000000000000000000000000000815291505b5060200190565b601181526000602082017f6970206d757374206e6f7420656d707479000000000000000000000000000000815291506149a9565b600d81526000602082017f706f7274206d757374203e203000000000000000000000000000000000000000815291506149a9565b600d81526000602082017f6e6f7420636f6e73656e73757300000000000000000000000000000000000000815291506149a9565b600c81526000602082017f6e6f742072656769737465720000000000000000000000000000000000000000815291506149a9565b600e81526000602082017f6e616d65206e6f74206578697374000000000000000000000000000000000000815291506149a9565b600c81526000602082017f52656769737465724e616d650000000000000000000000000000000000000000815291506149a9565b601581526000602082017f6e616d65206c656e677468206d757374203e3d20340000000000000000000000815291506149a9565b600a81526000602082017f5570646174654e616d6500000000000000000000000000000000000000000000815291506149a9565b600981526000602082017f6e6f742061646d696e0000000000000000000000000000000000000000000000815291506149a9565b600e81526000602082017f696e646578206d757374203e2030000000000000000000000000000000000000815291506149a9565b600981526000602082017f6e6f74206f776e65720000000000000000000000000000000000000000000000815291506149a9565b601781526000602082017f6e6f7420656e6f75676820696e6974206465706f736974000000000000000000815291506149a9565b600c81526000602082017f706f73206d757374203e20300000000000000000000000000000000000000000815291506149a9565b805160009060a0840190614c688582614781565b5060208301518482036020860152614c8082826148ae565b91505060408301518482036040860152614c9a82826148ae565b9150506060830151614caf6060860182614e39565b50608083015184820360808601526107fb82826148ae565b805160a080845260009190840190614cdf82826148ae565b9150506020830151614cf46020860182614781565b5060408301518482036040860152614d0c82826148ae565b9150506060830151614d2160608601826148a8565b506080830151612d416080860182614e39565b8051600090610100840190614d498582614e39565b5060208301518482036020860152614d6182826148ae565b91505060408301518482036040860152614d7b82826148ae565b91505060608301518482036060860152614d9582826148ae565b9150506080830151614daa6080860182614781565b5060a083015184820360a0860152614dc282826148ae565b91505060c0830151614dd760c08601826148a8565b5060e0830151612d4160e0860182614e39565b8051608080845260009190840190614e0282826148ae565b9150506020830151614e176020860182614781565b506040830151614e2a6040860182614e48565b506060830151612d4160608601825b6001600160401b03811661478a565b60ff811661478a565b6000614e5d82846148a8565b50602001919050565b60006111d682846148e0565b60006111d68284614902565b602081016107a48284614781565b60608101614e9a8286614781565b614ea76020830185614781565b81810360408301526107fb81846148ae565b60408101614ec78285614781565b8181036020830152613ff381846148ae565b60608101614ee78286614781565b8181036020830152614ef981856148ae565b905081810360408301526107fb8184614d34565b602080825281016111d68184614790565b602080825281016111d681846147f6565b602080825281016111d6818461484f565b602080825281016111d681846148ae565b60808082528101614f6281876148ae565b90508181036020830152614f7681866148ae565b9050614f856040830185614781565b6107fb6060830184614e39565b602081016107a48284614970565b602080825281016107a481614979565b602080825281016107a4816149b0565b602080825281016107a4816149e4565b602080825281016107a481614a18565b602080825281016107a481614a4c565b602080825281016107a481602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b6040808252810161506181614ab4565b905081810360208301526107a481614ae8565b6040808252810161508481614b1c565b905081810360208301526107a481614a80565b604080825281016150a781614b1c565b905081810360208301526107a481614bb8565b602080825281016107a481614b50565b602080825281016107a481614b84565b602080825281016107a481614bb8565b602080825281016107a481614bec565b602080825281016107a481614c20565b602080825281016111d68184614c54565b602080825281016111d68184614cc7565b602080825281016111d68184614d34565b602080825281016111d68184614dea565b600061515960405190565b905061516582826152a9565b919050565b60006001600160401b038211156151835761518361531c565b601f19601f83011660200192915050565b600082198211156151a7576151a76152f0565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b03038211156151a7576151a76152f0565b6000825b9250828210156151f5576151f56152f0565b500390565b60006001600160401b03821691506001600160401b0383166151e3565b60006001600160a01b0382166107a4565b600060ff82166107a4565b82818337506000910152565b60005b8381101561525a578181015183820152602001615242565b83811115610d555750506000910152565b60008161527a5761527a6152f0565b506000190190565b60028104600182168061529657607f821691505b602082108114156105fb576105fb615306565b601f19601f83011681018181106001600160401b03821117156152ce576152ce61531c565b6040525050565b60006000198214156152e9576152e96152f0565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b61533b81615217565b811461172b57600080fd5b8061533b565b6001600160401b03811661533b56fea264697066735822122063f95aae9ff4f16f87561a8621e2a0abfe8c16dced1770f4ee8437a4d79476b364736f6c63430008040033",
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
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	ProveLevel  uint8
	Size        uint64
	IsPlots     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateSectorEvent is a free log retrieval operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) FilterCreateSectorEvent(opts *bind.FilterOpts) (*StoreCreateSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreCreateSectorEventIterator{contract: _Store.contract, event: "CreateSectorEvent", logs: logs, sub: sub}, nil
}

// WatchCreateSectorEvent is a free log subscription operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
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
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
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

// StoreDnsErrorIterator is returned from FilterDnsError and is used to iterate over the raw logs and unpacked data for DnsError events raised by the Store contract.
type StoreDnsErrorIterator struct {
	Event *StoreDnsError // Event containing the contract specifics and raw log

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
func (it *StoreDnsErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDnsError)
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
		it.Event = new(StoreDnsError)
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
func (it *StoreDnsErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDnsErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDnsError represents a DnsError event raised by the Store contract.
type StoreDnsError struct {
	Method string
	Msg    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDnsError is a free log retrieval operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) FilterDnsError(opts *bind.FilterOpts) (*StoreDnsErrorIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DnsError")
	if err != nil {
		return nil, err
	}
	return &StoreDnsErrorIterator{contract: _Store.contract, event: "DnsError", logs: logs, sub: sub}, nil
}

// WatchDnsError is a free log subscription operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) WatchDnsError(opts *bind.WatchOpts, sink chan<- *StoreDnsError) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DnsError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDnsError)
				if err := _Store.contract.UnpackLog(event, "DnsError", log); err != nil {
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

// ParseDnsError is a log parse operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) ParseDnsError(log types.Log) (*StoreDnsError, error) {
	event := new(StoreDnsError)
	if err := _Store.contract.UnpackLog(event, "DnsError", log); err != nil {
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

// StoreFsErrorIterator is returned from FilterFsError and is used to iterate over the raw logs and unpacked data for FsError events raised by the Store contract.
type StoreFsErrorIterator struct {
	Event *StoreFsError // Event containing the contract specifics and raw log

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
func (it *StoreFsErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFsError)
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
		it.Event = new(StoreFsError)
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
func (it *StoreFsErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFsErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFsError represents a FsError event raised by the Store contract.
type StoreFsError struct {
	Method string
	Msg    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFsError is a free log retrieval operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) FilterFsError(opts *bind.FilterOpts) (*StoreFsErrorIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FsError")
	if err != nil {
		return nil, err
	}
	return &StoreFsErrorIterator{contract: _Store.contract, event: "FsError", logs: logs, sub: sub}, nil
}

// WatchFsError is a free log subscription operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) WatchFsError(opts *bind.WatchOpts, sink chan<- *StoreFsError) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FsError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFsError)
				if err := _Store.contract.UnpackLog(event, "FsError", log); err != nil {
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

// ParseFsError is a log parse operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) ParseFsError(log types.Log) (*StoreFsError, error) {
	event := new(StoreFsError)
	if err := _Store.contract.UnpackLog(event, "FsError", log); err != nil {
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

// StoreInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Store contract.
type StoreInitializedIterator struct {
	Event *StoreInitialized // Event containing the contract specifics and raw log

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
func (it *StoreInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreInitialized)
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
		it.Event = new(StoreInitialized)
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
func (it *StoreInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreInitialized represents a Initialized event raised by the Store contract.
type StoreInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Store *StoreFilterer) FilterInitialized(opts *bind.FilterOpts) (*StoreInitializedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StoreInitializedIterator{contract: _Store.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Store *StoreFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StoreInitialized) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreInitialized)
				if err := _Store.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Store *StoreFilterer) ParseInitialized(log types.Log) (*StoreInitialized, error) {
	event := new(StoreInitialized)
	if err := _Store.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// StorePDPVerifyEventIterator is returned from FilterPDPVerifyEvent and is used to iterate over the raw logs and unpacked data for PDPVerifyEvent events raised by the Store contract.
type StorePDPVerifyEventIterator struct {
	Event *StorePDPVerifyEvent // Event containing the contract specifics and raw log

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
func (it *StorePDPVerifyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePDPVerifyEvent)
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
		it.Event = new(StorePDPVerifyEvent)
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
func (it *StorePDPVerifyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePDPVerifyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePDPVerifyEvent represents a PDPVerifyEvent event raised by the Store contract.
type StorePDPVerifyEvent struct {
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPDPVerifyEvent is a free log retrieval operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) FilterPDPVerifyEvent(opts *bind.FilterOpts) (*StorePDPVerifyEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "PDPVerifyEvent")
	if err != nil {
		return nil, err
	}
	return &StorePDPVerifyEventIterator{contract: _Store.contract, event: "PDPVerifyEvent", logs: logs, sub: sub}, nil
}

// WatchPDPVerifyEvent is a free log subscription operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) WatchPDPVerifyEvent(opts *bind.WatchOpts, sink chan<- *StorePDPVerifyEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "PDPVerifyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePDPVerifyEvent)
				if err := _Store.contract.UnpackLog(event, "PDPVerifyEvent", log); err != nil {
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

// ParsePDPVerifyEvent is a log parse operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) ParsePDPVerifyEvent(log types.Log) (*StorePDPVerifyEvent, error) {
	event := new(StorePDPVerifyEvent)
	if err := _Store.contract.UnpackLog(event, "PDPVerifyEvent", log); err != nil {
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
