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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Pos\",\"type\":\"uint64\"}],\"internalType\":\"structChangeInitPosParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"AddInitPos\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"ApproveDNSCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"CreateDefaultUrl\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"DNSNodeReg\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"DelDNS\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"DelHeader\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"DnsInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllDnsNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetDNSNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"InitDeposit\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"}],\"internalType\":\"structDNSNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"GetHeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"HeaderOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structHeaderInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"internalType\":\"structReqInfo\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"GetName\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structNameInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"GetPeerPoolItem\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"WalletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"TotalInitPos\",\"type\":\"uint64\"}],\"internalType\":\"structPeerPoolItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPeerPoolMap\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"WalletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"TotalInitPos\",\"type\":\"uint64\"}],\"internalType\":\"structPeerPoolItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPluginList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"internalType\":\"structNameInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"GetUrl\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"internalType\":\"structQuitNodeParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"QuitNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Pos\",\"type\":\"uint64\"}],\"internalType\":\"structChangeInitPosParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"ReduceInitPos\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestHeader\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"RegisterHeader\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestName\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"RegisterName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peerPubKey\",\"type\":\"string\"}],\"name\":\"RejectDNSCandidate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"}],\"internalType\":\"structTransferInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"TransferHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"}],\"internalType\":\"structTransferInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"TransferName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"PeerPubKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Address\",\"type\":\"address\"}],\"internalType\":\"structUnRegisterCandidateParam\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"UnRegDNSNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IP\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Port\",\"type\":\"bytes\"}],\"internalType\":\"structUpdateNodeParam\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"UpdateDNSNodesInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"DesireTTL\",\"type\":\"uint64\"}],\"internalType\":\"structRequestName\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"UpdateName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x95c39a6f\",\"type\":\"bytes32\"}],\"name\":\"c_0x95c39a6f\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b2\",\"type\":\"bytes\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"toBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061f88d80620000226000396000f3fe6080604052600436106101c25760003560e01c80636839a7c6116100f7578063919064eb11610095578063d72b6be611610064578063d72b6be6146105fa578063dbc8ef9c14610637578063eafafaf714610653578063fd11c4c61461066f576101c2565b8063919064eb14610540578063ba47d33d1461057d578063ccfd72f6146105a6578063d6841470146105d1576101c2565b8063809bae57116100d1578063809bae57146104a75780638129fc1c146104c357806387dacef2146104da5780638b79782b14610517576101c2565b80636839a7c6146104235780636b9ceb531461044e5780636e91c5681461048b576101c2565b8063326cf61c11610164578063465efc9a1161013e578063465efc9a146103925780634f2ffbba146103ae57806351372819146103ca5780635f3376f3146103e6576101c2565b8063326cf61c14610310578063370d71041461034d5780633e4e5e5414610369576101c2565b8063228a09b2116101a0578063228a09b214610244578063240ce9f91461026d5780632a5da31e146102aa5780632c261f42146102e7576101c2565b8063084cbe86146101c75780630e580e79146101f05780631654a8941461021b575b600080fd5b3480156101d357600080fd5b506101ee60048036038101906101e9919061df1e565b61068b565b005b3480156101fc57600080fd5b5061020561068e565b604051610212919061ed93565b60405180910390f35b34801561022757600080fd5b50610242600480360381019061023d919061e266565b6109a1565b005b34801561025057600080fd5b5061026b6004803603810190610266919061e01d565b610f84565b005b34801561027957600080fd5b50610294600480360381019061028f919061e01d565b61141d565b6040516102a1919061f089565b60405180910390f35b3480156102b657600080fd5b506102d160048036038101906102cc919061dfb1565b6114c4565b6040516102de919061edb5565b60405180910390f35b3480156102f357600080fd5b5061030e6004803603810190610309919061e225565b611656565b005b34801561031c57600080fd5b506103376004803603810190610332919061df1e565b611da5565b604051610344919061edb5565b60405180910390f35b6103676004803603810190610362919061e1a3565b611e52565b005b34801561037557600080fd5b50610390600480360381019061038b919061e1e4565b612627565b005b6103ac60048036038101906103a7919061e09f565b612b71565b005b6103c860048036038101906103c3919061e121565b613320565b005b6103e460048036038101906103df919061e05e565b613638565b005b3480156103f257600080fd5b5061040d6004803603810190610408919061dfb1565b613eb2565b60405161041a919061edb5565b60405180910390f35b34801561042f57600080fd5b506104386143dd565b604051610445919061ed4f565b60405180910390f35b34801561045a57600080fd5b506104756004803603810190610470919061e121565b6146f0565b604051610482919061f067565b60405180910390f35b6104a560048036038101906104a0919061e121565b614806565b005b6104c160048036038101906104bc919061e162565b614ccc565b005b3480156104cf57600080fd5b506104d8615172565b005b3480156104e657600080fd5b5061050160048036038101906104fc919061def5565b615653565b60405161050e919061f023565b60405180910390f35b34801561052357600080fd5b5061053e6004803603810190610539919061e1e4565b6156fa565b005b34801561054c57600080fd5b506105676004803603810190610562919061e121565b616022565b604051610574919061f045565b60405180910390f35b34801561058957600080fd5b506105a4600480360381019061059f919061def5565b616298565b005b3480156105b257600080fd5b506105bb6167e6565b6040516105c8919061ed71565b60405180910390f35b3480156105dd57600080fd5b506105f860048036038101906105f3919061e0e0565b616af9565b005b34801561060657600080fd5b50610621600480360381019061061c919061df70565b61739e565b60405161062e919061edb5565b60405180910390f35b610651600480360381019061064c919061e01d565b617483565b005b61066d6004803603810190610668919061e05e565b61783f565b005b6106896004803603810190610684919061e1a3565b617f41565b005b50565b60606106bc7f2cd842034c90466aa7addf67a7cf83e08cec4ec383b349481d06ab56b2ac0f8060001b61068b565b6106e87f773834520d425517d308170070398a3584c254d2df6181e59c34a4b1b547decb60001b61068b565b6107147fd1d131f4836c7b5c57a11216cec85217103f5437139a738c4e3cb982cf28296260001b61068b565b6000600a6002015467ffffffffffffffff81111561075b577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561079457816020015b61078161d4c2565b8152602001906001900390816107795790505b5090506107c37f3dd9960cd73a5338de7dcbe8e55e39fc3fcf7212915b4151f9d654bd94d150d760001b61068b565b6107ef7f98526d6e39cc73a0a8f907265497064c8d8f5ab5f17b8ec40ec70b594a86213960001b61068b565b60006107fb600a61987f565b90505b61081281600a61998190919063ffffffff16565b15610941576108437f856ffaf5c30ae647f7d2258d69dc449f322cae51c82f2f4b32670255c5c1e1a660001b61068b565b61086f7fa9671f4977ace2da205c8034864b1e73b692c94b053f8e26d6b7eea005e73b7f60001b61068b565b600061088582600a619a1990919063ffffffff16565b9150506108b47f685eeb9f81061bf27f6827a7c0a227ef37392a106478ac949092b5086db6dc7f60001b61068b565b6108e07f4d54308a5c14b24a85a37ea337cc97fd6fddd93b721f1d291876400674a49a1f60001b61068b565b8083838151811061091a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052505061093a81600a619d4590919063ffffffff16565b90506107fe565b5061096e7f6fb260d0fd5bdc3208f10cb7681429f42c9bbcb11594d36192b806e7330841d460001b61068b565b61099a7f5df934ab0a34342a7c6b9b7a828c37ba8ed790beca1ab1db655c817f81f121a560001b61068b565b8091505090565b6109cd7f7bba85067e59ee9e572bea813bbff5526f21de8459206442995c7ca48a90110760001b61068b565b6109f97f7e3be24c9940df67c7acbc86231bb941c2fbe4e9ac8bbe552e9be563e62e252f60001b61068b565b610a257f258a8e0a4984dad25aa8b383709cf0ed23a5fb73c2c4508309d364d35f96930e60001b61068b565b610a517f6daa55a99e690e9d53a339dc283020adaa15acfa4640c85117441e9e5c40549460001b61068b565b600081602001515111610a99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a909061ee4a565b60405180910390fd5b610ac57f8faad9c6b716d601af7ab4137ff1ff505b7d367aa5ab7a594a540dce013cac1360001b61068b565b610af17fc3983b438be2819a966bacd1c9a08ff652ce59b8785907d1fe9b7819e297ae3c60001b61068b565b610b1d7f4a28192e0fb4687dc3038e51aa5dd34726de8d908f515427ffde44d0fcc6cd1c60001b61068b565b610b497f01eb1490c8aa0f6c344b287e2d1fde5ef817306c32d59c4ff0036360f2c753b260001b61068b565b600081604001515111610b91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b889061ee6a565b60405180910390fd5b610bbd7f122756e739e3897779aed0a837963f1b6b79866e53a7832d323773b8f636bbff60001b61068b565b610be97f682cb5e1154f8d3461f0edfaf17d3448b4d83f7eaef40073444e705f4980c1bd60001b61068b565b610c157f2fda7bbbf71b88e37c5714edec6c38226eb06399eca5223fafc6f0f1698e15e760001b61068b565b6000610c2b33600d619f1890919063ffffffff16565b9050610c597fabd94c316c41af999b45c395256385c7380f51656523e78b3562a16ba14de98860001b61068b565b610c857f637c06b20a8dfb5c845f9c359c90169751e168689fc419d061cadab456773a8560001b61068b565b3373ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610d7b57610ce87f915cb89fcbd37480680bb7a718e255e3289797020d79449724c4aec289b99f3e60001b61068b565b610d147f389a09ca278fb32c5a803c66a188d0f7fcc019bf4a5a1dea9e2eeaa28a8f11e460001b61068b565b610d407f252d3dec376a9cc8ca1ff6f0cc95e93eeef6f848062dadf7f54761c14f496cde60001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d729061eeaa565b60405180910390fd5b610da77fda13c28c38ef5e861aaab117609a0492b6621d82a18f68c11892b423cc54b88560001b61068b565b610dd37febafa063a34402e92fb6cd374ae0d1173463b8dff06b672efeeb4bfd4334674160001b61068b565b610dff7f35c1fb3cbe64273183cb62d6ec83f7831171d90dcc0f69666d64121ac175cc1660001b61068b565b81602001518160200181905250610e387f79f10dfcca6a303da1b3cbe881351345feb1b7d6ba025d32761963f8e00ad1c660001b61068b565b610e637e96eaa35401e826e7db960a511f20e0ab58a211ac84f6aa5edc560f265f877160001b61068b565b81604001518160400181905250610e9c7f163ad8f3b13e5e0a1211d21f477ec1b501c28b38e73c813f1458f4b869d73e6160001b61068b565b610ec87ff3c1177a3eed721e7b34d507955923b70696afbd32d6c6b6efee3537dbe5742f60001b61068b565b610ede3382600d61a2399092919063ffffffff16565b50610f0b7fe96e82b8b2f0c6019c607b0c4064215e82886891ddc2c80c3904c0ec291c82a260001b61068b565b610f377feab631d5e524f861b36ae5c1119759295325501ddb671062009028df3ee21ad760001b61068b565b7f602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf82602001518360400151338460600151604051610f78949392919061edd7565b60405180910390a15050565b610fb07fdf7cc45cdb2790ce91ee566f97b592c7e343874b5dc72a6ab839e5080a46a8eb60001b61068b565b610fdc7fdaec78302e480e573254d5108ec992bccc6f06326908b375cdde1794c2da784560001b61068b565b6110087f31c5e9ccfa71c2e067c8ddb8d07989cac7639b9b1132a671d8fbb63734ff48a660001b61068b565b600061101e82600a61a87590919063ffffffff16565b905061104c7f9c73e06780a30be6ae92db682eee31d5a36ef200955794bc8119bd3727f08e3060001b61068b565b6110787ff93126a5df79108ba958a752a07c50cc43905f933426f8766b2ea78616713ed960001b61068b565b6001816060015167ffffffffffffffff16101561114e576110bb7f0f2d1ba6a2e7e3975ee50f1027c09f76abbeeea2564afc04307b418d3a16349960001b61068b565b6110e77fbe95f8bb38d55715e92b311714426e688b9c223bef8759fcbedc32928f05ed5c60001b61068b565b6111137f04a2b758588441dc43c1900bc0fd27e1ed021e14c7f3b218189fd64a4b3751b360001b61068b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111459061efe3565b60405180910390fd5b61117a7f368be697915c459443fa4159b36f88683967db32b67c0dc0abe7cf44373b097060001b61068b565b6111a67f6c86ca8aef670c6c4b35d24ce9c4f93f506a92e6124d496a09fc426a57dd58b460001b61068b565b6111d27fac0649480d159452e57db66b8b5bd6c570c965fec701c0654d772cf07e74d0be60001b61068b565b6000600581111561120c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff16146112db576112487f3fec986af600b5a066cafaf37a6c95d5ad3f5ee329becd7a7334b8db3334d15960001b61068b565b6112747f0653456e9c84709b6e942e439371523a7273a2dd3472697b7dd595b869cf753260001b61068b565b6112a07ffb915e5cca0cc7118a1a72d8864a583099b35fa6f5610085523f5d2cd42169dd60001b61068b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112d29061eeaa565b60405180910390fd5b6113077f4f6457ac0f54397116a0f92f3890f3046f3f98ed7baefe58961dd03d5a06678560001b61068b565b6113337f1f6c287d96ced5fd7f7462bb063764c3d6e118081e7e1b4c846e72dd88d29ec460001b61068b565b61135f7f16d186bcb28579338abb0020ce8466497a5f727e4e809994e4f8d23a5ddaeee660001b61068b565b60026005811115611399577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816040019060ff16908160ff16815250506113d67f1d74b1c3991a00f4408ec85b76c606cf1a9d390682f64ac9baf15a3f800dfe7160001b61068b565b6114027f2005e06a8a3a606a1df14686e6b544d6c9041a523e6840776e293d60baf3415e60001b61068b565b6114188282600a61aa6e9092919063ffffffff16565b505050565b61142561d4c2565b6114517f5a3c8847c912133c076e0b0b15ed8e6d920e1238bc1a83cbea1f5c179b06d0b460001b61068b565b61147d7fdc62026e95cdba2cec0b556417cddc7002784da655e8da764a62c031b76ad27660001b61068b565b6114a97f926a0c2bf22339c47028184ba0df4f97606703ee8cc34285f2116c55187614fb60001b61068b565b6114bd82600a61a87590919063ffffffff16565b9050919050565b60606114f27f374ec67693d8c00e0a6ecc37ca7ac09173f84bbcc99266eafd7c4b65421a209760001b61068b565b61151e7fc2bbf758b29c1e259316e26bf544d9c40045a34b3d82dfe68ffe21b9b0b9fd9860001b61068b565b61154a7f6df3688148b16b0902e2308f95e2670daf4bf81f49cef72ed0ed6ecc5206567360001b61068b565b600061158b846040518060400160405280600381526020017f3a2f2f0000000000000000000000000000000000000000000000000000000000815250613eb2565b90506115b97fd07b6be52e05cccdfcad00347db155e19569d14bc3a0ba7a936f64ed86f196a760001b61068b565b6115e57fff63a456454e5f6c96ca1f6dc1db50adf5e5cba78291af54b3bf47fbae2e51ee60001b61068b565b60006115f18285613eb2565b905061161f7f7dc36c891a42bae720b013eff7f2c7b6b4dd9d61bfbf2131c4112b8113a2909860001b61068b565b61164b7ff0338e2137a11b891a2012acae7094fd23d99989b7c0f149db37193d6dcf158c60001b61068b565b809250505092915050565b6116827f4b3fbb4e124abc3aaf2147931d7dca7cf09a042ec9a460c643a48d7521ad5bec60001b61068b565b6116ae7fb84dcc3fff69a8bb2cb9027942a01f75f1bc4c15f65722cba2f6e5ec3c1adc2460001b61068b565b6116da7f097cdd548c5e3a7bb62eed9d1a93698e30f7fc6f6e99e5fdf1227ad343dbf46a60001b61068b565b3373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16146117d05761173d7f6566c0d031d9061d5dbd24935ef1cf4d10421a5c02ba5dbace5255ab0879f1b060001b61068b565b6117697f6a8494f1bab168872bcfe9f89a79b6214dcce1efd870efa5366ece738d1eefd260001b61068b565b6117957f1e950262ff03706bf860989e2442c332450941f21559bcf017c8e0f0c723e7c160001b61068b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117c79061efc3565b60405180910390fd5b6117fc7f92dad41c7eb96d81925a924cd4d0d63a5c3b20c93e720ee2df78203bcc68d4dd60001b61068b565b6118287f686a4fd9882a7b9cba29d78b4c94b7d81004b2a0356d77a00f845b31a98a67a760001b61068b565b6118547f060e9cb4176a919ba6d108260c07a06ab59c5f6b0a2beb37ca24c5e88066269e60001b61068b565b600061186e8260000151600a61a87590919063ffffffff16565b905061189c7f16e54df5cbb8626ee6d3da61f95d4b524dc001be25fb2608e755669cf57a892c60001b61068b565b6118c87f6f59666739ee8b0f42c0ba8b66fca12332c48f8404322d9614f9da020bef982a60001b61068b565b60006005811115611902577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff16146119d15761193e7f6236f72f77fe59ed3af5d06054633f407c535afba2c00e7a1082dc12ff076bc760001b61068b565b61196a7fac61c66bf5e620231bf334403dd970e845449feb5b1a766d3097bfe0fb3ccb3760001b61068b565b6119967f6187cd0e91693b063db84d91b640b53e0fa9e23c345e77d8f370285d7e5edbf960001b61068b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119c89061eeaa565b60405180910390fd5b6119fd7fe78b10ab7bd530d6352800e25398e91e63c1775857be77e66b0d22f628b4ed6a60001b61068b565b611a297fd628c26c2e1e245a0cf71aafb2cc3f2d23048fdf1d59cc59d2f6548d1fd9c5ec60001b61068b565b611a557f0f53b96bc8a3c55b6e75e1cb130b81c5f0cbd07b7ffabb49651382f7962a164d60001b61068b565b816020015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614611b4f57611abc7f22719fac400fed8dc3a37f02abcda75b622f6dd61ca8ee165ab7583dc9c1894860001b61068b565b611ae87f328f3efae98524274b08073542168350717a605746c8aa4414b7601010dbb2a360001b61068b565b611b147f512c2c0468d5faaf5190ba22b4bec6f80a91c1db3e53842c05986bb0486230b460001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b469061efc3565b60405180910390fd5b611b7b7f7f336bbdb611726301fbfad1484c793ad2ef77674734d299ee37d914a2f1f83460001b61068b565b611ba77fdc47ff9f1cfe8ae04828ea8b35ea7dc7d375906188d2afb2b1636e130fae623460001b61068b565b611bd37f99e03199b7132f050d98b709cc60d1ddfcb799f98b127430c4b594023d6d9a2b60001b61068b565b816020015173ffffffffffffffffffffffffffffffffffffffff166108fc826060015167ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015611c2b573d6000803e3d6000fd5b50611c587f4c1393ea31056f454942070c5bd4936d66ae24395e4e3bdff614f2b49063692860001b61068b565b611c847f759d06f6008d80a7482bd044002b2eb020161c0e130a2e5cf34243bd8e198e1660001b61068b565b611c9c8260000151600a61b00890919063ffffffff16565b50611cc97f2069cdb5deb84d596af724190736218cb2cb184e9f0846ea7ce60ed395baa8f460001b61068b565b611cf57f6aa771a8e3731e21f556bb331982b1d207d725a3e12d74a226276ad520bb5dda60001b61068b565b611d0d8260200151600d61b38190919063ffffffff16565b50611d3a7fd2e6376092ef50b610541905bf0cdcb7c2a1eba3b50366218ebaae800eef31a360001b61068b565b611d667fea62aa9fbb675939c8621e6ab698dbd757655a7583c2eef78efa59b25ce7830160001b61068b565b7f7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d8260200151604051611d99919061ec81565b60405180910390a15050565b6060611dd37f70214c7ef2c21177300729e5167f6887584bdb5d99b23d9c5be620854a9bc02a60001b61068b565b611dff7fecb0cb14150a84e9ed814b1b48036222c15d1b8e416842a366476e6ecead341c60001b61068b565b611e2b7f3b79c76bb1a6853f4d3a556fe20e81c8c8c1120fca6756c11cc64aa0b2c4397d60001b61068b565b81604051602001611e3c919061ec21565b6040516020818303038152906040529050919050565b611e7e7fe58ae019a69500b592ba03312db5223d94610e2126c0af272933a7d7b0a3faa960001b61068b565b611eaa7fbb6e10f660738583c32bb03241ff8253c1b7d5262b7476cfa943e523208cfd5e60001b61068b565b611ed67fcd8ab5c2cf1a871429cd2b591c31acc60c8276900b9b9df02b44e2f2cb72233160001b61068b565b6000611eea82602001518360400151613eb2565b9050611f187f72dcaccbad8b22414b5044d7f4867e530861c9c338f7f4a3871deaee9616245e60001b61068b565b611f447f9dba534651256c8b599782ad7fb9396b4476ac8de90b02461b3765649554b4b960001b61068b565b6000611f5a82600661b74490919063ffffffff16565b9050611f887f9aeb76f3078d81b098f4d07a64d00c16ad26d71eb05d2b6a71c30d580a3e6e3360001b61068b565b611fb47f9533532f6a3e22ffb2508e639f0b58139a59dd89cc0794803628ec7a0584dce360001b61068b565b60008160c0015114156120d957611fed7f817650e08fde040942409404e95530b7f1187fd4686723223cef5064b70949cd60001b61068b565b6120197f9146acf9589c465895bf5b811caa85eb82f1a8ae20d0ef305a857f825533a70260001b61068b565b6120457f1097c0c9b6eddc73f5d53e85d050dea995d633d63d293f196b1eb2fe08abecb060001b61068b565b7fd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e356040516120729061ef1d565b60405180910390a16120a67fd148ba4f0a8908e55db573da87f8cd14197988b7819055441b73f32cf52bb45660001b61068b565b6120d27ff8d035209d7da6479677b7495c80771055bed7eb9df73412fae14814ba9f4c4560001b61068b565b5050612624565b6121057f1e01990cd638a79eecd504b49d768e9d53fe4ae271db0441f7b738d11fdd996460001b61068b565b6121317f3c39081e6c3e854ab1b386451ed26b15f2a1d7b075974617aa5d8e93ce4a511560001b61068b565b61215d7fcb3e2167a4b991fbca0476ee9e0a55a561ccaa303b20c2263bed677cd6fe51a360001b61068b565b826080015173ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16146122b0576121c47f0a33b4fdb8cc7ccc05d8c969eada80c93930b0735148e9d70f60e7dcc04480ad60001b61068b565b6121f07fced7ea805725e2f500416d3e2855f145eef15f2d814016fe99e2778d816252f360001b61068b565b61221c7fad9bb0216a4b2d5776f0d3fb9145ae16218e9412c59f209958ab17ac929eef8d60001b61068b565b7fd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e356040516122499061ef50565b60405180910390a161227d7f1d655a7c1769ee6e87b51400d3815e37fbf1c7eb5c86203819e585c1f03dc68f60001b61068b565b6122a97f441fcdf0e70b9856a0f3affd8bdef7bf08c0575071f65f08b36d0596c05bc07160001b61068b565b5050612624565b6122dc7f7e6a84d3224aff553259390db998bb78923ff70bf3b3f63ea8c0feea9033db1560001b61068b565b6123087f48d098edeee0b0d5c8b1ce1083320f7418fc037ec05d42cb9c47da3b84142fcf60001b61068b565b6123347fa5696470d5480deae6fdcf2996d758192f9fdc52c3a737554632c1b87531e89f60001b61068b565b8260000151816000019067ffffffffffffffff16908167ffffffffffffffff16815250506123847f54fff5e470cab055705d2d4848826b0b3c8080c987e1415a87a16570622f28ed60001b61068b565b6123b07fa5f76580f535d9aed103c5ad1d97dc1097350e2231531b4fef3006f90c4173c760001b61068b565b826060015181606001819052506123e97fb1c19cca1157b67ca8d38f7981d4e86ad8f993ac1ccc19f5d2f699c45d80dc3760001b61068b565b6124157fc6cef6d43743a2cded63944da333679d860e821484838750959edd2069b3f15160001b61068b565b8260a001518160a0018190525061244e7fd8d869d94e685cc032757cd5b04f3f133120a635052508ec59aec16e89f090aa60001b61068b565b61247a7f6fa35c207f05086d609c9152a5b1edb403988a1fcad0919f882f2b51945dd6e960001b61068b565b8260c001518160e0019067ffffffffffffffff16908167ffffffffffffffff16815250506124ca7f50ed44626649685394b30a122c1384e1c2fff21941183e300427da2658f663cc60001b61068b565b6124f67f212125d3e0584c3fb17f2634316d0797153919c3133dae369a24b09582eacab460001b61068b565b600143612503919061f262565b8160c00181815250506125387f955908e383fbf7316cf218bf6ef3c7da635911f64de77a57ee06edcee3dac86b60001b61068b565b6125647f8e2f0b134bb4f4a7f10577617926f3313a793152a6c574fade68f613ff690fa960001b61068b565b61257a8282600661bb139092919063ffffffff16565b506125a77f6af1d0f749883a608ce28b8e15efdbee3ab19ab544daa6eb97da32d61039c77e60001b61068b565b6125d37fd43193d87e42ccbd89af6e7b8fd4247b20e4d081a69b6d42e62a3c2c0297c3f460001b61068b565b7f0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150836080015161260b856020015186604001516114c4565b60405161261992919061ecda565b60405180910390a150505b50565b6126537fd82faf353da5f146f752f7a7e6166550103ea1beb65ad07b4c002193e34fbe0060001b61068b565b61267f7faf526b086154db4017d3531a6f1efd40c14156822474fd169a547f122b9cf04a60001b61068b565b6126ab7f2fb15cf1501d85bcf6ed8d73bd6e27b360b421dc25fb5c5af08a95d8538c0fd960001b61068b565b60006126bf82600001518360200151613eb2565b90506126ed7ffee8c05e1bdca5336085d71085a7a656d2f343da85fef56a8cf1c9e998bf6eb360001b61068b565b6127197f33b3a4d508625254afbb0e64028b2a0212916eab65910e468f7e5a3172edf48960001b61068b565b600061272f82600661b74490919063ffffffff16565b905061275d7f3f37e085f7266c43b3daa9cbfa54cd287031f515aeda3e60e8bc2162452c9ce060001b61068b565b6127897f028700f7ac948957320e468dc3b35b12cbeaa39fe9682ed9487c28d733d4416460001b61068b565b826040015173ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff1614612883576127f07f7066bc96aec59dedf190fd136a64f87a68719cfd00ba3f115b65d0595ee8ed7d60001b61068b565b61281c7f7570127045d598cda3682e924204c13eb1f0a83370b46e7c4cbbe88a0b1fb1bc60001b61068b565b6128487f554e3509efcf2bf8ff097720b16a0728b6a29ac4e7083434554e81638e930de460001b61068b565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161287a9061efc3565b60405180910390fd5b6128af7f977b75718ac3bcd131e8a4559c05b4a3ad1da543a93a236fabb604330c9f3c8e60001b61068b565b6128db7fffaa7c72909bd6a4588ae7a7dd715ea7c3161036f3105148e6647d3c02103b6c60001b61068b565b6129077fe53f66b52512fce9f3d8b14eb470d5bd8e07648e2ca4542b518010f98f14886e60001b61068b565b8260600151816080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061296f7f4e069850960422822e03076149ee33007be8a31cf4b4e78df2f8a6cff033596060001b61068b565b61299b7f8337de8ee8d11d6c24ccaa8f8d6e6ded6dd12d48aa7a3815df4fd7d8e39d0be460001b61068b565b6001436129a8919061f262565b8160c00181815250506129dd7f2cc4d77eeb63d72f4f19d1a9dbcd08c76363b5d17bd8a6115bb6052e2cb2cfbd60001b61068b565b612a097f69a1ecc45846fbe0661b0744cb0c1ee23925d6910f23d6416653b685bb4de78f60001b61068b565b438160c001518260e0015167ffffffffffffffff16612a28919061f262565b612a32919061f2f6565b8160e0019067ffffffffffffffff16908167ffffffffffffffff1681525050612a7d7f0212025efe9a8358bcab00a05303ccf9c6a62574fc568f5c56be6014b947277060001b61068b565b612aa97f01bc0d4f0a772e88604f01902c8e2517155d4fcde091c10c3e14d9787e8ab81960001b61068b565b612abf8282600661bb139092919063ffffffff16565b50612aec7fb226d2cb2169342028243f387eb29e3e3d6115cb6eb3d346284f6ab527c0003560001b61068b565b612b187f52904161ff748ea15e313f409733b82b727c52d09651a31d263751c8bb7a854e60001b61068b565b7fa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f83604001518460600151612b55866000015187602001516114c4565b604051612b649392919061ec9c565b60405180910390a1505050565b612b9d7f2a031a373c056b5efd57cd3614dc883e87978b011b345041b042ad7aeed3101260001b61068b565b612bc97fc4a8eb9d22a0cecb522eac18315e1059ccc17d01714d1f7fb3743f959a244aeb60001b61068b565b612bf57f9645a6a326b6935348eecafa948b75c8dfb981341facf85fb3c721d085484c7f60001b61068b565b612c217ff995c0d512adf99db25a288059b541bfa75160e32faa5b0da968d54f8ffba59560001b61068b565b6000816060015167ffffffffffffffff1611612c72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c699061efa3565b60405180910390fd5b612c9e7f113b79b5fc1889068e42b074d8d9b46be7aa7fb7a52c6e2ac2391e8552dc0c9760001b61068b565b612cca7fbf096d2cab6191f6b149472f906d9707e88c007910a28dd7f9ac4bafdeb2040c60001b61068b565b612cf67fe1fde579af267ea03ad7b0d4cc0bdc22e4ee9a861e2c3da0d72370caa1dd121360001b61068b565b612d227f4de046787746da4ef07d205a6b52caf8ad0778dd7b57fcb41cbb6ca9e1718bbb60001b61068b565b806060015167ffffffffffffffff16341015612d73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d6a9061ee2a565b60405180910390fd5b612d9f7f81700aa145b652c858406f7db145a2faf7247603f91dd3c008b1fa6244c11b3260001b61068b565b612dcb7f06ecdd519028d233e75438d5acfc6aea5d3fa668210511fe0857028d3f0f2f8560001b61068b565b612df77fbdd585cc84267ae924867f0c8185133d0ad8d981ca20d3170f4e5585abbf557f60001b61068b565b3373ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614612eed57612e5a7fc67a68dbae5cf74efc59293c0e1c64bfcab4e161d09fd8d773b9b10032058bd560001b61068b565b612e867f9f1a63a956e42f7d04d169c86a8277447a01caa4e90fe3c5d96e6e9cda4da4f360001b61068b565b612eb27fed03a316a598ca1512148fcb4a78754d51aa938681a78a864c2a1d25f5ac397e60001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ee49061efc3565b60405180910390fd5b612f197f982ba9d43b1514fc255de444a6a4d257bdf8e4ca3e1654936ee1adcfeca5591660001b61068b565b612f457ff6ec9558b874699a2b3eb0c244c475eae801bbb557aec3405668d004d65e532260001b61068b565b612f717f2859eb719248adc92832d83c70461687d3f5305da47fb9a1443edf31f3410e7560001b61068b565b612f7961d4c2565b612fa57f23bb5b13465016cf09d4803d5f6f1e38bf57082ada2b0e38687d1d2c20f3c20e60001b61068b565b612fd17fc16b673bfe0a90779a0a5d0fd5c7c760ec6a18123c9c8fd1d9f0bad0c5f6521460001b61068b565b8160800151816000018190525061300a7fa332beecd433dbc696b63f38a449afc219d7ac6da98d1f54aa00f54433aa5e3660001b61068b565b6130367fe8753920b41ffb6cb969c2e37c08f992f4921d011fb7cebb7f029b3a24fa949d60001b61068b565b8160000151816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061309e7f4bbf36786fb01157786512c9540ae802fdb7c9a7bc149cc8681b629a8e74010960001b61068b565b6130ca7f05fadbbf1776bb140a55f7b08db14f68daf73e3402f1fe321df9ab3be6f9640a60001b61068b565b60006005811115613104577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816040019060ff16908160ff16815250506131417f19df1846102978029437327408ba65b4ebcd0043b9f3b14fb04a420f64fe764660001b61068b565b61316d7f113c7dc847eb8ce296df2132134689aa7415cef3bc43a1665c41b2a848f8ef0660001b61068b565b8160600151816060019067ffffffffffffffff16908167ffffffffffffffff16815250506131bd7fcdfecbe8adccae6bd0bbb158775320e3334e680922f1d1a2e3fb91bbe4989ba260001b61068b565b6131e97f7b858a38b1bbb6fcf1e46f44b1be218dabe6542c243fb2e177e65af398cc67d760001b61068b565b613203826080015182600a61aa6e9092919063ffffffff16565b506132307f3e527b085ddc77c69d9056852efc84fd3e8fe2055a30abae626ae6da6ce8551f60001b61068b565b61325c7fc15df0ba390bba36281574b41287c0af87211f1f5561d5df5e89e823e02e8e0760001b61068b565b613276826000015183600d61a2399092919063ffffffff16565b506132a37f57e39d1af411375fe94814f92337a53029576eec89930688d272996cbc49449f60001b61068b565b6132cf7f59a0533d9d6ba262cfe8345b6aafdf5f2f0d443a693dcbb3cf35d7fe2485d1c860001b61068b565b7f602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf8260200151836040015184600001518560600151604051613314949392919061edd7565b60405180910390a15050565b61334c7fa3e1d8d4ceccc6a1aff81c1e6eb27cb964c2d20130df989f9b056d649c3dab1c60001b61068b565b6133787f5045f7d19e0bcb3e2fab09850b34b0356e793ab3ec425599ea1da2630bee792d60001b61068b565b6133a47f677b9ac86b151474ecf2ce81c5275fd498d827e1a89ef2041ecafefd34b4f71a60001b61068b565b60006133b882600001518360200151613eb2565b90506133e67f7985e3df2336fa82bcbd86842c34fc43da0d16268cbd3fb50f5c8b0a6423a0bf60001b61068b565b6134127f0c3187968906542df8f18e164628f58a72f2648ac59e4c91fd0ddd7d8e002dda60001b61068b565b600061342882600661b74490919063ffffffff16565b90506134567ff01fdfcced43c8d06701374b01be407d4c985af644b20a6dd2fee0897e8c5fff60001b61068b565b6134827f2118f9df819af7839bd7cd34b1caead22d06e866f7b2208d07f94421c9b2d29560001b61068b565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff161461359a576135077f043afa3b32fac210f928d296c80130754064782b8dd82cca0f91e9c80fe04a2860001b61068b565b6135337fdc43d6abb7ac9a664c4b45c944f49691c00d03055ee07ef6c087cb91b6a2b34b60001b61068b565b61355f7f55831c520310dcee6c1a599c71445b9312ea1388bada095a2a51f152d11841fc60001b61068b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135919061ef83565b60405180910390fd5b6135c67f453b3fedbc773cb06d3f5f2aad58f2a3baac70377c7fe90a694b77240068c9f660001b61068b565b6135f27fc2751f35bc248bfd5c33e7435d3fee3bbb49398a76286c77a38fad3583e79a0d60001b61068b565b61361e7fd11e928339d24b493c6a008d4c386578180807dab38001e1f95a7085daa2b35560001b61068b565b61363282600661c11d90919063ffffffff16565b50505050565b6136647f13f6dca6cf0dbd1417b47907c420126c87e8d02eba85289afdb0033fd0bb6beb60001b61068b565b6136907f10753607e00b7c308690678f33f2ef37727629b698a5b7b027d9749ba2e908f560001b61068b565b6136bc7f9c51eb6bb010a1b0240e0d553432426c21fba83cdf29315522c41d564053ef9460001b61068b565b6136e87f192cb6755f1ded106db197f644aa84e7deb1689768babe52e2fb14e32a212bda60001b61068b565b6000816040015167ffffffffffffffff1611613739576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137309061f003565b60405180910390fd5b6137657fa45b8a0d031969c014d50bbc8d0cddcb9721f53dba13df982002fab3e0b9514260001b61068b565b6137917f8d27555d28b4e6e18e45be525a6ca6af5c0c58785eab1c92b4c100ad0b0655b460001b61068b565b6137bd7f414b945d4bb767fe3d754d3055784b6975ca1c26f66e0699ad9d4c42e063d0c560001b61068b565b6137e97fbcf56f12011023fbf0f46c50703d145c133285eedb7ddb72d30ccb7bad569f5660001b61068b565b806040015167ffffffffffffffff1634101561383a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138319061ee2a565b60405180910390fd5b6138667fc49f0b6fb3a6a742dc7d9f5dd331fca10a8d23785f67cdb227000249db0049b260001b61068b565b6138927fcdef8e204f03da1aafb0665ecbd7aa149017a3697364b5ef89b25e13d083d02b60001b61068b565b6138be7f58ac1f0aae57421d6d7fc36239496432760245e5ec9cf6cb189c8307967b2d2f60001b61068b565b3373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16146139b4576139217f05c7d5dfacc47ac67dceb4e2ea0ed227676ae13eba06c6e73b15a8b388958b2e60001b61068b565b61394d7f2a5439fbc4ba0677ca4de60628d35b9f29414b04743cfe112cfaa99e0966a6a460001b61068b565b6139797fa56e7eb37c1b5a420605fa280e0f6627fe37c7aa6b4dc830647cfad23627a63c60001b61068b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139ab9061efc3565b60405180910390fd5b6139e07f8b5b39598b28334599b95140014f64815fe02c45c047e5131d40ae6294d7eb1e60001b61068b565b613a0c7fcbbd2cf27dfbbe66f108035f9b1cc4615e40134dea855677054d2df70a1ea09160001b61068b565b613a387fee94cf85764a5e15ffa9e97a46091337dbe206f1e1814630785d98ecbb651d2960001b61068b565b6000613a528260000151600a61a87590919063ffffffff16565b9050613a807f34a5b433d2916b0f7b6d37c28099d06c32006085b82b01e7782c89198300cc4060001b61068b565b613aac7fdeafcc19658efad9ec2f722145cffcb8522dbe8afdf93ccb821f3b14bbc67a9460001b61068b565b816020015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614613ba657613b137f152212c8e281d35deb9d72a75e46cc0f39681709f5a0acb2d3db698e9522279c60001b61068b565b613b3f7fa3e23be1dfbfca8c090bebe71da27c3c8e89fc41594858f671d36251dbec45c160001b61068b565b613b6b7f32d68a2a5b41392edc473bda0a58979586785672a90186bc8c165add089aa7ee60001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b9d9061efc3565b60405180910390fd5b613bd27f505e4383806505c60418ceb38bb1f3f06a066afc37efb591faaf2159fd11de5d60001b61068b565b613bfe7fb64e472d12854d5a40b47686812d56aaae5060b729ae50b8f5ee60a6ddbb49d360001b61068b565b613c2a7f49fc26b4c189e785547013ec57f7598422f131ecdc8b15e972eb144cc4873e7960001b61068b565b60026005811115613c64577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff1614158015613cc0575060006005811115613cb2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff1614155b15613d8457613cf17f02b41388cbfc41dbc6bfdac3f1a6b4d5eb36860c501950f8009f5c602a8f4bc560001b61068b565b613d1d7f6467361e32ef8a11fe1ef107e4730c444732f0ae816675f5a57c5b29106349ba60001b61068b565b613d497f4aecf80b2da567f4af0c9d1824d1ed11acbe5e535b9700420ebd72e609c0b9bc60001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d7b9061ee8a565b60405180910390fd5b613db07fc046d8f42d0d9f48082ae84729daacfa9eae52286ccc5deca1df5c961afb7fb960001b61068b565b613ddc7fe81fc44ace13dec1d68edf43a84ec1e86da60ab9a21d6959dd6392eb233b617e60001b61068b565b613e087fbf8ac3d1b9a23acd2d1a492f21a2e82a6cdcb67d9813984ead56d8b3acadc8a860001b61068b565b816040015181606001818151613e1e919061f2b8565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050613e677fe072ee4394e019627e50c371b2d75f2b9810cb424fcc898d7ea99cded9853a1960001b61068b565b613e937fe01a19f786a03a73263399b29e7e4f48b2b86a8fb226f301e5e04d7d0964c4f960001b61068b565b613ead826000015182600a61aa6e9092919063ffffffff16565b505050565b6060613ee07f3e1a9dc80d49653b0081f83857efdbf315980951d36df77265c441478a59f66160001b61068b565b613f0c7faf5fa98953aed4bfefb5c808bf5b585b270b76350eab09eee71b3ac6421c2efa60001b61068b565b613f387f2b492de5c4113b0deb081c84622b920658b1191f45a3ac303ac75f3adae1d71460001b61068b565b600083519050613f6a7f5fd5a651011ff9fb1d6afe60741dbc928e32401eac1dc5c2d4f3630c8bc2a88f60001b61068b565b613f967f96f279be1561381477a936d12501e8fc5ad92d7c2b85f0310e860098ce9cabf760001b61068b565b600083519050613fc87fdda72b5e77c6fdb1cec3faf5eedaf8c7106b687008639e568eb9deadea5bcede60001b61068b565b613ff47f3510a1442a1155ddf87cce9154b52ad00b40a0e1cd19ec989727ccec5f610e7560001b61068b565b600084518651614004919061f262565b67ffffffffffffffff811115614043577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156140755781602001600182028036833780820191505090505b5090506140a47f09e9b8d8ee2d08196077cfccc4b6d0b8a679968066fc995b69224e96428b703560001b61068b565b6140d07f0f2b28d52aee7d32c5e42c277fc6dcc26364017dda4710c50c25abc48139287160001b61068b565b60005b838110156141f2576141077f6272d9e243c6179b8b6e957e8685a0ca588ed015d9647dab029c47b054574a6c60001b61068b565b6141337f8b090e33ae9db6e3b4f855da865ec2935a2a510baf7472ee86f5048c64a14da360001b61068b565b86818151811061416c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b8282815181106141b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806141ea9061f494565b9150506140d3565b5061421f7f7cff8c2cac48d8b38b9388156c4beb47a9ba866f4547a79cedb5abb249ca0a2b60001b61068b565b61424b7f6584f7761ddaa5691eca7dba46e03fd6b87cbc173153c63a969c2d4fa32428e260001b61068b565b60005b82811015614378576142827f0a8ea68653499d9f9733f39d9a115be0e5a3ad0c94b3a62dadd814cb3f0e3c2360001b61068b565b6142ae7f390cdad0e6cc3f10fbb075942c230cf4ad7bb779a7911ee549029bd05c28f0ad60001b61068b565b8581815181106142e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b8285836142ff919061f262565b81518110614336577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806143709061f494565b91505061424e565b506143a57fa556f11b7e62a468a5942b633bbceafa73d6bbc1e6c21311f9f28d897ea6436160001b61068b565b6143d17f4e9d951ea6c2d02faa67ad7116bcea8645d0f7e5c2b4a6e590c0da18406ae80860001b61068b565b80935050505092915050565b606061440b7f344655eb199bc8fd1e4760c3356d1c3f445623523736605f6286b9f6d2eeb73b60001b61068b565b6144377fab62f58819087a3e11364f083471707b0bf1621774579e76a1f26c7e3a3ca3f760001b61068b565b6144637f2c916d70e83794f1aa601d87f1ca5575b09aa5a1e83aa26cc3baa064afad78ea60001b61068b565b6000600d6002015467ffffffffffffffff8111156144aa577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156144e357816020015b6144d061d50d565b8152602001906001900390816144c85790505b5090506145127ff65135a07464f5db22189c1a628be3bb1cf15e14142dcdeee464bc5c78dcef1260001b61068b565b61453e7f62c4445af46093fe6772022a652c7368a53c4f2fd123ad1c0d033421d136bc8b60001b61068b565b600061454a600d61c4d5565b90505b61456181600d61c5d790919063ffffffff16565b15614690576145927f7f68f86cb956297543c427a059924b592509fc1f77eab79c56763d6dede7dd7c60001b61068b565b6145be7f6121daa06d874abec5fe7a9d101c2cb959e1b80b6e418d1c1d286bc49eb477c860001b61068b565b60006145d482600d61c66f90919063ffffffff16565b9150506146037f7b479fa41360201d3ffe6d787c3939d71c1aa91597728baba9c4c3e0d71fe5c760001b61068b565b61462f7fb4c3f0032ec22c6dc15f2d4e1c5a3ed42d419a4f95ce9cfc07f887c86075b69960001b61068b565b80838381518110614669577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052505061468981600d61ca5790919063ffffffff16565b905061454d565b506146bd7f7eb930f8a08f81e860ba364b072db33d32c99e9942fe6d6bf7c48d01e39bc77d60001b61068b565b6146e97f2a79a84d0cecc5440c57c3f601d8514271efc2bc6f345e221aaefd4f90c2059a60001b61068b565b8091505090565b6146f861d55c565b6147247fd776cec038db24669a8500ab37f307139494ab5f10ed7cea728244671d44780960001b61068b565b6147507f1cfff1c34604bb688442108be6e6949956c13f8bd9458d2c25215f2e1956886160001b61068b565b61477c7f1e34cc16ae3f52a99caba50d21d3fec80c824fdec295ed437d741193aac8086260001b61068b565b600061479083600001518460200151613eb2565b90506147be7f91bb7e9804b966231b7c01a20936573a2a896ce25211b115786fd825492b049a60001b61068b565b6147ea7fe7307edc4e5e1e03aa7a9d7d019887a1bb690d28789f3c579c62264e0f7cff3a60001b61068b565b6147fe81600661b74490919063ffffffff16565b915050919050565b6148327f16daae697ffc08921c5ad34e0eba43cda2d2d6ab8ac8bbb323b534cde95b029b60001b61068b565b61485e7fd86766e421c2e6d26c7ffea561f41dda26602831765d35c92ba44bbc064fe50f60001b61068b565b61488a7f836732f5b631ebd540b78c7911cf0e9ebbb67f0ab633ef79a1dcadae476f82b160001b61068b565b6000600582600001516040516148a0919061ec3c565b90815260200160405180910390206040518060a00160405290816000820180546148c99061f431565b80601f01602080910402602001604051908101604052809291908181526020018280546148f59061f431565b80156149425780601f1061491757610100808354040283529160200191614942565b820191906000526020600020905b81548152906001019060200180831161492557829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546149b19061f431565b80601f01602080910402602001604051908101604052809291908181526020018280546149dd9061f431565b8015614a2a5780601f106149ff57610100808354040283529160200191614a2a565b820191906000526020600020905b815481529060010190602001808311614a0d57829003601f168201915b50505050508152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250509050614a9d7fd66b9e240f1f29fe8df4aea3019530849f1f10df5202763ed48d4f70c4e2302160001b61068b565b614ac97faadd1f1b2e0bf5e8852ae5fb50532a65d5f64f42dd8812c7fb4b09aeda61533160001b61068b565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614614be157614b4e7f6f6f6647c72b48f0578b3d8db0659d6c221720304f8a8db3b362099cb7b9cf2760001b61068b565b614b7a7f6cc87ed15886ce01aa91aa3ce4be9ac73f274b5774cff39e3b507f61b2d28a4c60001b61068b565b614ba67f959cf646a0dc61f0f7394818722bca95e2038827f34459dd531acbe0968d193f60001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614bd89061ef83565b60405180910390fd5b614c0d7fabe2381e281a728f389ce41ed4f98265a905774b3528681ed272dc2eae3ec90e60001b61068b565b614c397fc27dadc21e3872b04c62bb952793a31f512c70dc6f9d88921b5aaca87dca09eb60001b61068b565b60058260000151604051614c4d919061ec3c565b908152602001604051809103902060008082016000614c6c919061d5cb565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600282016000614ca3919061d5cb565b60038201600090556004820160006101000a81549067ffffffffffffffff021916905550505050565b614cf87f4ae9cc49cee0134fa42de71d8651c396e7188bdaa7f556117082f4203702c19d60001b61068b565b614d247f51a406ddd17f0b4d313a1d2cca0ba8b3bd0265451a38f506dfcead650895a3eb60001b61068b565b614d507f07937e85f3be8af01726a62609b8791f0ed664fb75f638433ec44b8d1354917a60001b61068b565b614d5861d60b565b614d847f4ac3f09114ab6ffa8fd45a5f3e5c29ed347ad5179836bc9fd15336b10c9c76c860001b61068b565b614db07f78d693bde40f0b660f84f9c9369dc34a8bbac51b785acd0af2af71ecf61c0be160001b61068b565b81600001518160000181905250614de97f2e18eaa14add0b32a6268bd47f43c0d23c9d1093e14aff89eb919db12aab9ad260001b61068b565b614e157fe17d72c3d021c168aec0710024497cf652afabd87a9c434f85af0b026ce775d660001b61068b565b8160200151816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050614e7d7f3001b8742785b3952bf3eb9ae7fdd88486b05a4cf87e1430195cc99bdc1285c160001b61068b565b614ea97f33555bbbf56373c632024783b3c318653d984f9687362cea3fd950722672dac160001b61068b565b81604001518160400181905250614ee27f1bdea87e3421faff376a69e7a742f006afc11cc1383ab8ca011504b336ff842460001b61068b565b614f0e7f728f2490f71ec41b76d4cb78de098faff1253eafe08dd22a809012c8a47bdb9460001b61068b565b600143614f1b919061f262565b816060018181525050614f507f7c04d3ac81bde6b054628f4f907a9aaff109eb5273ce05775047c37a8c385e0060001b61068b565b614f7c7f56785914dd894da0f458d84bddb271f4003e1e8b5b02fc206056b4cdae11250260001b61068b565b6000816080019067ffffffffffffffff16908167ffffffffffffffff1681525050614fc97faa1c2384150d69195d2071b1b5f6eae2f4078f5921542611373645df146db15060001b61068b565b614ff57faba4ca7550453f6965e80621e11570b620f463b066f320ccf39ac27b8beb7ef860001b61068b565b806005836000015160405161500a919061ec3c565b9081526020016040518091039020600082015181600001908051906020019061503492919061d65a565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201908051906020019061509892919061d65a565b506060820151816003015560808201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050506151017fe5276ec2788da65111c6512e9a390e26ed1df11d83e9184c232445758ebb5ab560001b61068b565b61512d7fd4fac880a7afd1d7d241b955659f47e48ce0c0d9881d07004872407561f1359b60001b61068b565b7f356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc8698260200151836000015160405161516692919061ecda565b60405180910390a15050565b600060019054906101000a900460ff1661519a5760008054906101000a900460ff16156151a3565b6151a261cc26565b5b6151e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016151d99061eeca565b60405180910390fd5b60008060019054906101000a900460ff161590508015615232576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b61525e7f14b27ce75693e09338461b2b5f4aaf177061a7e15001c750b0ca5605459da8d560001b61068b565b61528a7fe2234e3b8cfb774695b740370c0bf649f3a4eea0b263534c58a66e4b036c720360001b61068b565b6152b67f6e01c49a71878bc955ae1177bbbd493d7bec43014f826ae01323da6f768ec5af60001b61068b565b60008060026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061530b7fb79850d34cf93dca64b9a045f3d6eaec53591ca331d2815a704fa38497f0bd8960001b61068b565b6153377f22cb3b870894a54bfd42b67ee74ffe82955366c0cc7e8fea739a708d7140688e60001b61068b565b60016000600a6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061538d7f68ad1e3f9a25f169f95e19e53c5c9c1e5216a5fde33cefd3d4051f5e4e0b82ad60001b61068b565b6153b97f19d672f9ccc61a9ea57f85bd9af9d0ef81f2f78633e83b9815fdf92e1f20433f60001b61068b565b6002600060126101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061540f7f3fbdfbd8934e97858e6c3ac42480ac78334d4771bad796584860f6ef3fbd313660001b61068b565b61543b7f32a4ee39bd75755d6ea3499f981ca4e49f26f3e1a42035cbc684d2d7a77b38ba60001b61068b565b6004600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506154917f92765d0f6f98fa4c09847ba64995beb588d6c777b6fd0738a82ff5a5ae9825e160001b61068b565b6154bd7f10ca0599a129cb91d0219f70dc3d2231e653be69f9d29f0970f94056f95164ac60001b61068b565b6008600160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506155137f4e253d455668a2ab38299d80a19c1a4665e58d4650e4036eb225e3273a476d5760001b61068b565b61553f7ff15c33b57dca5e7966b9533832f05d4953f512ad19252c73423113050a6327ae60001b61068b565b6040518060400160405280600381526020017f64737000000000000000000000000000000000000000000000000000000000008152506002908051906020019061558a92919061d6e0565b506155b77f086d272314240923c0d9db90cf4c45da47f3a96aadf7271d6fae51e4a66e19d060001b61068b565b6155e37f82dda4ffd678e1a201a70fe5d5b46ff0f0adad229ebe8593c48bee8e32b3c3d160001b61068b565b6040518060400160405280600a81526020017f6473702d706c7567696e000000000000000000000000000000000000000000008152506003908051906020019061562e92919061d6e0565b5080156156505760008060016101000a81548160ff0219169083151502179055505b50565b61565b61d50d565b6156877fc846e79c75e0ec409d3b873b0fa1ff35f7dcc564a4953d465b8fa250730cb52160001b61068b565b6156b37f873470904cfa37409e4a07cbf7bc8a740462f33ab49053326eba5c336f09844f60001b61068b565b6156df7f61220b560a88d71149a9e9b78e18e1f860bc7fdfec8132681d0b8768dc0c63a360001b61068b565b6156f382600d619f1890919063ffffffff16565b9050919050565b6157267f3abfeee79971dbaf7100fdf38849ef21385a8ccb4572ebd9699631d4374cf8da60001b61068b565b6157527f9c96215ddb5bd816d1ac0d4126a27d85fed3fc4a3c10c99e151d3e2b28adde6660001b61068b565b61577e7f794fb0b9d453b7b6c95086104c1336f9e45e29be24fc85a94a2ab91b3b3c142060001b61068b565b600060058260000151604051615794919061ec3c565b90815260200160405180910390206040518060a00160405290816000820180546157bd9061f431565b80601f01602080910402602001604051908101604052809291908181526020018280546157e99061f431565b80156158365780601f1061580b57610100808354040283529160200191615836565b820191906000526020600020905b81548152906001019060200180831161581957829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546158a59061f431565b80601f01602080910402602001604051908101604052809291908181526020018280546158d19061f431565b801561591e5780601f106158f35761010080835404028352916020019161591e565b820191906000526020600020905b81548152906001019060200180831161590157829003601f168201915b50505050508152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090506159917f791591df9d26bfe71f7e0efb25c9c60e58cf1ddad4a0c8a72af2bdd3d154a26560001b61068b565b6159bd7f1dff7f9ca9b6bfb1952ed0e5846b45f0336069612841bb36e3d95b3124d6025460001b61068b565b816040015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614615ab757615a247f78a3572ec4a6565838ac53e2cda6c29616f0051daa35913e2abb50d99647d73e60001b61068b565b615a507f4da3f679edb0d7465e6fe55234caded55d07fdae773aa473ce3420bbc0af306960001b61068b565b615a7c7fc3ad2c63732d5dc6b763eb158714c0671be5cab1f41f51cf4b997350df24449160001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615aae9061efc3565b60405180910390fd5b615ae37fdf6bca1b13979f1ad2270240484a7cab4b6cc3ae4f4d9be9a06d3def1a9a61ce60001b61068b565b615b0f7f6eda4c9517c0f4238b7b866c612234219355d7195fcf019fa437c04e714b459960001b61068b565b615b3b7f21c6bab6b69919be81a761002377f70132603e0a3faffe796042d76bda2a243060001b61068b565b6000615b697fe1cc38a2fc853fbfc5a24abc208d35d16448829ff91c1d52959a82fe0247970f60001b61068b565b615b957f0b5147c8ea66ae1b4e62b2e485fddec2372d37ec88b5f1a3b074068148839a2160001b61068b565b438260600151836080015167ffffffffffffffff16615bb4919061f262565b11615c4657615be57fe18db84dd527a6ec99971b53524da1b2a63fce4c68beb2671b6f3f1919e80f9460001b61068b565b615c117f468fd0fe7625cbe47786c1f59c7c8fe9f56840645ddfa5a1c8f96bb8b0d8f21f60001b61068b565b615c3d7f1fb248eb7f8264591dcdcae61d88dc0f61c023e74af8ed07bcdf028efb71c97e60001b61068b565b60009050615cf6565b615c727f87f09c472d586023d8dbf900f67265fb8da4ce868d4bc9f4992419bd32a6d48060001b61068b565b615c9e7f7f9889fb699781e886f1bdcd9e805689acaaf47203959db4db6eb1f96624cae360001b61068b565b615cca7f4f4c9c78c5c3025b8c192d02e6c2b9991371bb8441ed5860560d4cd0b151c30e60001b61068b565b438260600151836080015167ffffffffffffffff16615ce9919061f262565b615cf3919061f2f6565b90505b615d227ff4e84b3f4b160009cb3bff747bde1233872a7a2319effc882a328a4028fc628060001b61068b565b615d4e7f61c8546dcd2154f43aba4928981159d5688f017396a42b40a7736dda57e66eaa60001b61068b565b826000015160058460000151604051615d67919061ec3c565b90815260200160405180910390206000019080519060200190615d8b92919061d65a565b50615db87fb420003b761aa897b9596e8e77f108235c90271759426d1e854e8bd19a02b8ac60001b61068b565b615de47f285d6fe612dd1a3130294e48db6ca092bb8f0ef11595f6fd1ec34279b4d4df5d60001b61068b565b826060015160058460000151604051615dfd919061ec3c565b908152602001604051809103902060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550615e787ff92f2ac98ebc42dc21436f0d88a1573dc43a35d42362ee9dd866365e088e295f60001b61068b565b615ea47f7abaac172cf351c427fb55f79a2bbe1f2790c75d88de563579ceb8939ce4256d60001b61068b565b600143615eb1919061f262565b60058460000151604051615ec5919061ec3c565b908152602001604051809103902060030181905550615f067f06522fa0cdee23b7aeff6b6407d57fd3ba916ac79e975a4a7eddc02066dc32da60001b61068b565b615f327fcb038e6e455f584af2c18d52ba4767bb2ad7c80019839dc8c21baff72c9064f360001b61068b565b8060058460000151604051615f47919061ec3c565b908152602001604051809103902060040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550615faa7f08706b3aab3f545649bdec04e31c4e5bbf2bfa519828a2e6cbd1e1e6600e743260001b61068b565b615fd67f4aadcff393dc26003508b862358d6c3650f15b6bfab892ad7d1d5f45da43359f60001b61068b565b7f178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca8360400151846060015185600001516040516160159392919061ec9c565b60405180910390a1505050565b61602a61d60b565b6160567f541cea80be568377adb10f687620c51a0db1bad8116408c0a0d8d345b40ee07060001b61068b565b6160827f947ca5fd0afd53d08004a8c3843e4f50e2925324dd878ded5354eb28da3d499860001b61068b565b6160ae7fc4b0edc870958fdd1e16333f0545c502351ca270e5e2e2ec282a5e9569cdf94260001b61068b565b600582600001516040516160c2919061ec3c565b90815260200160405180910390206040518060a00160405290816000820180546160eb9061f431565b80601f01602080910402602001604051908101604052809291908181526020018280546161179061f431565b80156161645780601f1061613957610100808354040283529160200191616164565b820191906000526020600020905b81548152906001019060200180831161614757829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546161d39061f431565b80601f01602080910402602001604051908101604052809291908181526020018280546161ff9061f431565b801561624c5780601f106162215761010080835404028352916020019161624c565b820191906000526020600020905b81548152906001019060200180831161622f57829003601f168201915b50505050508152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250509050919050565b6162c47fb7f1c28a49c59901d523e07e767519b6ce909da29d3e42e2b85e2a30c267c09760001b61068b565b6162f07fc68d9398cff61578867ae8b3f25172e9ea19759cb46eed07ec734db36caccf8060001b61068b565b61631c7ffb465a589f6061e5144a687a57160441fd2e1042ee985de295aa4e5f5e03d5ee60001b61068b565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506163897ff82fb99cfb46320c5d6dc4d96d7508e436eebc8a42f166ef58647c7b6ba859bf60001b61068b565b6163b57f0ed1a6af39051fabca1eec4eb46641fbf5e11c7aac8578c0184343158fbb5f2760001b61068b565b6163bd61d60b565b6163e97f6cffd1e28395c2ae85e7c53c7bc2c576e56e1c5a98d70f766aa0952a2272ed8760001b61068b565b6164157f93f5498865c7127ef4ced4a7f0042cec7401f039c7b2e9714d33a396028c56a260001b61068b565b600280546164229061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461644e9061f431565b801561649b5780601f106164705761010080835404028352916020019161649b565b820191906000526020600020905b81548152906001019060200180831161647e57829003601f168201915b505050505081600001819052506164d47fec9745fdae12b7890e8e2a54f3e6fa61355da206741a02d713474e870e0b4d9f60001b61068b565b6165007f914493cd306e972df2562a2dda2f7921cf9ff68181b56481892074be25c7174d60001b61068b565b81816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506165647f40422c952647a1c77652b23c1e20c7740de83bda8ec417e58ff6e258f67d756860001b61068b565b6165907faf513138e917b457bd58be102a948ba5cee5fcb117a7dff4e7945b25103141c460001b61068b565b6040518060400160405280601581526020017f7265736572766564206473702070726f746f636f6c000000000000000000000081525081604001819052506165fa7f37297149614195205356118f3868d1d730a8e2f69700457970ef9a6fdf79cdad60001b61068b565b6166267fd0aa90b57b8fdbdfe01390053cf7f64213e3a2a5487c927845b6c669769826f660001b61068b565b600081606001818152505061665d7f4dd5585e1a49e683acc98bf8b9e74f97744e61728c9e9027ff9795635f204ecf60001b61068b565b6166897f427cbe2964a6aef40ff920c02ac3ce29840131a410ad5c7f68030efa5497875660001b61068b565b6000816080019067ffffffffffffffff16908167ffffffffffffffff16815250506166d67fc79e59b5e55bf1e5b94b0eaa38012854f719e3785717505bc9886902e596d0df60001b61068b565b6167027f8494157f6564f94ed442a147ad31c6a8bc2c6cfe50dc288be1edaf9d1db1760060001b61068b565b8060058260000151604051616717919061ec3c565b9081526020016040518091039020600082015181600001908051906020019061674192919061d65a565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906167a592919061d65a565b506060820151816003015560808201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050505050565b60606168147f03c89963924b33744cd0f4ac85361a57515bad2a3dd002d8ea6034f56a3ecd6a60001b61068b565b6168407f2b57f15a9a5eb048286ee085198b1102508d78601ec44adcaa5e138a5649766160001b61068b565b61686c7fa8cf860a2cc031e05a84f88ba43a0045a506fcf4b6696dc8fbe1b8d7e1d93ccc60001b61068b565b600060066002015467ffffffffffffffff8111156168b3577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156168ec57816020015b6168d961d55c565b8152602001906001900390816168d15790505b50905061691b7f10a68ccce1d72aa60ec3efce8ac5576dcefd7ab18ef88b7342e59ec2080a613560001b61068b565b6169477f7fbefad15ad5b43005c5cf9dba68132daceb1506a2ebe168101ff7570a489f5760001b61068b565b6000616953600661cc37565b90505b61696a81600661cd3990919063ffffffff16565b15616a995761699b7f675298f0d4d6b00f8810fa65fcde690aaa014435f87aafe8b16b2535b069d60a60001b61068b565b6169c77f7ac14aa2e66d24ab7f44aaad05dd9dd96395decd831eaaf1dea3ee85f261393860001b61068b565b60006169dd82600661cdd190919063ffffffff16565b915050616a0c7fe42c00ed219430da4767c0217ed7349f49a656bf57acaf4b85b113a025b3830e60001b61068b565b616a387ff4fd54f4efe0cf7115ac6a7a6643cf501ddbce0476f3881715b53f78c7a008e160001b61068b565b80838381518110616a72577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525050616a9281600661d2d390919063ffffffff16565b9050616956565b50616ac67fed32139c67660803369a040aea594f0d712eae7fccc3fbe21480784dff3347fe60001b61068b565b616af27f86112798b51decb49a22b11c7c54a934b981e381d868c02bce3964bc08c4d7d960001b61068b565b8091505090565b616b257f88b5019dff649e46e01ebcfb27e8319f4b2daf171bedf28aa66b798dfa5a5bbd60001b61068b565b616b517f7dfb1fa8c2b7155cb111a1089dfabffae0ce0e1f479ce56fc672108ce22e602360001b61068b565b616b7d7f74a9558a209c87a0df1c4745e71ed2896f4ff326b51b8a7da2f8dbaff37e726a60001b61068b565b3373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614616c7357616be07f781c3e0bdc6ae6dac85724fc4b4cf70ac59bc1b8632ff94d01b4029473474a4b60001b61068b565b616c0c7f74a8863779a31b1df8a17f4b2ef354a70c0a2bb21c1fa051ed4c15cf1cdb112560001b61068b565b616c387f793f06980ad6fb442ab48369e954a523da3a892041199a00a910e295a74c3c6460001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616c6a9061efc3565b60405180910390fd5b616c9f7f614f52b035f59a9488011ad5aee40b1519544eaca2eddde87fa837a869dae60760001b61068b565b616ccb7f3702d731fd77ca680eee547f7c4ff6c790ae3e0e3a2732fe0559a466d1f5d47660001b61068b565b616cf77fd0866313cdd4b119775dbd336548fb6cada2642b37903fce298f16b5a43d335460001b61068b565b6000616d118260000151600a61a87590919063ffffffff16565b9050616d3f7fe3897d6867332d1445b8ec3570bf2780462e8c952ae8902b77b84229e5ec49d660001b61068b565b616d6b7ff2d7d8b61dff84adfb30227047300df57a3dc53d67ddbc87589ca7b3806539f160001b61068b565b816020015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614616e6557616dd27f5a3736baca585105c97aacab8d4d0cdd1f7031cd9f23f7b66c7eaf6ceb5693cb60001b61068b565b616dfe7f36fa33916994a51cca195f5a9d4c382d673cc89d09e3cf002162d07b24208b6760001b61068b565b616e2a7fc6d3cd67ae63c2fd3aa5c851b457446cb1893144c80620beab81eef92f94de5a60001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616e5c9061efc3565b60405180910390fd5b616e917fad33324c6bb87a91b720cf207de78018353bdc6c709ed0ba52473f4dda6bde4260001b61068b565b616ebd7f8334d166f57c5cda8feedd9f886437044c2dce5ed26a7623a8263fd18700d04f60001b61068b565b616ee97f916ac7e703f8b96b8817a6887a8fdcba94aac4b2cf17907f0c85a0b996f96e1960001b61068b565b60026005811115616f23577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff1614158015616f7f575060006005811115616f71577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff1614155b1561704357616fb07f49027258703605a998f18743de37d06b37e8baf2594b2bc89f54426544f19bd760001b61068b565b616fdc7f5f253d219167f50f05911e4b91c193615c6bd723b62d4375e349ca841ec8d08e60001b61068b565b6170087f5cac3b3c18954ce67384973978d3b142ea27934cebbec72ba40bf45d52313e6360001b61068b565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161703a9061ee8a565b60405180910390fd5b61706f7fcf75d6afd9f9adfd71e77e946c20dc4a6b3a80267ec1c81163ca0da70e56d4c660001b61068b565b61709b7fa57d8bfef2a0da05ca9b3a282aca44db1f38855c4c4500902ac94198ceb1610760001b61068b565b6170c77f049041322cad138925a6c2cc14599d3300cf3600603dd99da712d55cf67cce9560001b61068b565b60026005811115617101577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff1614156171e65761713e7fe872743dc043c7ae0dc1c048f4fb35ba34dc615e729ce1fdd3d9a9e339c6a56d60001b61068b565b61716a7f1fe1748fe68a3a5e9ff63f090d727346a270f8f3e889c968c8445bfb466b943a60001b61068b565b6171967f5a160bf772d4baa265b3238a5d693d4f243d33d62b00b3865f1f728632f1f78a60001b61068b565b600360058111156171d0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816040019060ff16908160ff16815250506172b6565b6172127fefe060c062cb29f3e115dd835deff3f4d05c33b03af16260ac99b83fe5424eab60001b61068b565b61723e7fee526e61ac28c6ba27bae8fdae0adac948e7a15a6f854a4424e440a792e3422360001b61068b565b61726a7fb6a05e11dec7d481680188e88d9385e696cea80cb7d9c887226c7f228357e55960001b61068b565b600460058111156172a4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816040019060ff16908160ff16815250505b6172e27f2961bf55d385f4dee483653def7871e85d74f2406683f86c2dd025fb47aa800860001b61068b565b61730e7fecbf578fbb57075db22064ad4974e09fb94cb067251a6821d578e985dd98dc6c60001b61068b565b617328826000015182600a61aa6e9092919063ffffffff16565b506173557f5891da1fa0c234e697ad96c0361854118031335704069858c7037bcaad71ab2460001b61068b565b6173817fff924dd835e971ca87e2482ef5eb1f7b9a9f72bad7cd72cb6156afea683c56e160001b61068b565b6173998260200151600d61b38190919063ffffffff16565b505050565b60606173cc7fce6742aacf387566c418908864311d83de4ac8b4afe21c197043ef1564633fc760001b61068b565b6173f87f73d74b6650ded5639c1b66eb91586ea44c13250e3edc11aa29e7a6bcb72bc00960001b61068b565b6174247f466d2d49c18964a4da3cf3d0bb7a9adaaacb8736959059976b5903b81771b4f760001b61068b565b61747c600283604051617437919061ec3c565b602060405180830381855afa158015617454573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190617477919061df47565b611da5565b9050919050565b6174af7fdb7a5c48d4482554fa0d7d54e9b7eb1547e890a9451053b3aece10f8f41b7a0260001b61068b565b6174db7ffda111a7a6ebd6cc9aa095797e8b097ebc3bd963af174088dc44a0bf36ceaede60001b61068b565b6175077f9dc5540cfac1f62fdb8cab176a50cacb22136a752ff0f2cb64ae6155a97bc4eb60001b61068b565b600061751d82600a61a87590919063ffffffff16565b905061754b7fbe01a1403e5448d92ff535fe62efc932fc1b20927419303cd83998cc35ed1ab560001b61068b565b6175777f4c7c6212af7f23bfab8e47a1d4f0fb590924a793b77087702f855b64e4b40a8360001b61068b565b600060058111156175b1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16816040015160ff1614617680576175ed7f2d6f4146f7a91b50fa93fc9678fdb8cd611599b933b32e620e80cbf84903399b60001b61068b565b6176197faada2cf32cabf736fc8cb94f65b64e4f34707d328b0b8107cde50b5ec61c6bd560001b61068b565b6176457f7990726fca46c365be43be3549e5737a5d737447973e022a77f916bd266a272860001b61068b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016176779061eeaa565b60405180910390fd5b6176ac7fedf0c5606eece22933a1c3e487c9ef131968abdd76e14b346d8a621e8fca722860001b61068b565b6176d87ff667845b88d4201638af819b39157245b9a268377f553c6180826b4aaddc1c7c60001b61068b565b6177047f446f21e045d986adfa802f51d7294daa9575d6667ec3ddb485962abb1aeba0ef60001b61068b565b806020015173ffffffffffffffffffffffffffffffffffffffff166108fc826060015167ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801561775c573d6000803e3d6000fd5b506177897f3a0525141cfb0fa9c8cfcfe74fd13a95c4fab0f02a79b1bfcd2146edfaaaa87360001b61068b565b6177b57fd2a9a32a9e6b75ee336b2f2b383085d5faece3b69f0d806fefa467a99a76481060001b61068b565b6177c982600a61b00890919063ffffffff16565b506177f67ff5530b2550eb2f3b3aa72a7d15fcd8965fcdcc041e1409503a8c1d3d3a6d8e3e60001b61068b565b6178227f6211200f206ecdd7311286194e46bc82352998ebad4b3e9318e7eca1cbd60a8260001b61068b565b61783a8160200151600d61b38190919063ffffffff16565b505050565b61786b7fc7ef7b4aac51b10d9109e81b24831812ed816ddb376a955f594d23864b5437cd60001b61068b565b6178977fb57012e94c719b5541bc0e9465eed56b5f835900b1629da2c692db742ac8d12a60001b61068b565b6178c37f32bfe10924aeeed94d8a4e8743c38aa732dac8309f74fb041b8ceb0f20852e2960001b61068b565b6178ef7f56799eda85f89e20e656fa5d0cfe3720635202b909774ecae5ced3ac7ccb92ac60001b61068b565b6000816040015167ffffffffffffffff1611617940576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016179379061f003565b60405180910390fd5b61796c7f50e3b4a94cc520304800225af40e9689f17c0f1854aaf5f36dbfac7daa667c9c60001b61068b565b6179987fcbbd12096ef957bfc96398b8e82c7100d2211a90327fc3f6f434ce9c68986bae60001b61068b565b6179c47f1089af3660613b10df65d1a87ae9c369e9e405b5fc745f67e8de9d1ff8f1953960001b61068b565b3373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614617aba57617a277f8f68bf1083e9f0da8cf00168160e4ac55887d2f52650f352527d16ec96c94a2260001b61068b565b617a537f5182971ba91c2643b4c5968c3126437cd46a85c2a063d3d8c6a88aba7bd6b66b60001b61068b565b617a7f7ffa6e0ede8a3151b73dba436f2bef4a598880dcdda3a5232e087a7a85b38c844a60001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617ab19061efc3565b60405180910390fd5b617ae67f5a3be9a6576eacf3074fdbccf836aef024f5b05216021d8371d7d5c33e12502b60001b61068b565b617b127f8fedd76d6f59d2de317be0a66b358c4a61ece80f59824cdcf27439ed2e11491660001b61068b565b617b3e7f364797623d1a96d4444619e9e4a54acbf3a9b41bdaa8cac062dc39175f84374d60001b61068b565b6000617b588260000151600a61a87590919063ffffffff16565b9050617b867f7a935006291d21aa1cfe75aeee60590b3bc94f4a65e9a06572f8df522d88eb8b60001b61068b565b617bb27fec2fd633334baf9a05835909a1c76f06aef72aa5622affe99221ae879be76a1a60001b61068b565b816020015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614617cac57617c197f8a17a274550ea507a7870197a6b19d262e1e355c93f7d66630e5d0edd8fe305560001b61068b565b617c457f0d6126b695fe0085c7e9b76564dbcae294b3c27f597c7c65d985f3ecd98bb7eb60001b61068b565b617c717f87657bed2f24d4ab7ea01b70e156c5a5cbaa667684ae7cb5a3d564b6ad6489b560001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617ca39061efc3565b60405180910390fd5b617cd87fe2ad006a6d630fc69ec2371e7d49dc71f8bd61bc57d0db34f5d28625c0ac0c0360001b61068b565b617d047fb3c82530b33b4757e1043cbe906603853b254ecf21a524339eace9947297d2d860001b61068b565b617d307fcf352958e951cbd5afb0100a6cadd3536f9e7c4e0067e9896eaaadcd27be7e5460001b61068b565b816040015167ffffffffffffffff16816060015167ffffffffffffffff161015617e1357617d807fd42ee934f0b8a53fd30553094cdf475cf89c217abe3ed81586ac6d793bb7434f60001b61068b565b617dac7fc0b8493c640e2c5d9bf7a031e4b7ba6a4ffb5c3c3c593a6fff2417184e2f075060001b61068b565b617dd87fcf48d1e922594d96398e79c383fd43b633646a93b81880316fa9306f5903e26460001b61068b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617e0a9061efe3565b60405180910390fd5b617e3f7fe4200e3e423a4778c8ab2cedabed40f93ec0b0d6f892211d8e7948310f91e73160001b61068b565b617e6b7f9841e2887a2269b5864e216297fe39a3a20b2e20d508570f1d739607eaec6b1260001b61068b565b617e977fe2f66ce7614bc91cddb4d6cfdfd602917091d88d585c44db7113d88730c4adcc60001b61068b565b816040015181606001818151617ead919061f32a565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050617ef67fec27deead839553582302497f5c1a63807224c0ec96a42740b8ab8cee7748c4b60001b61068b565b617f227f6e7f9bbc99ec5a5485616d34dbff264ae9681b7bce0b9a02e0bfc5eb32e5d00460001b61068b565b617f3c826000015182600a61aa6e9092919063ffffffff16565b505050565b617f6d7fe1d1a368b6db8874f128e6bc8f18faff60de588044091242d7e5712f5a33228960001b61068b565b617f997fc2faa3bf245ae05f17bc59aadcbd0b343e4ee539ca0b5f3f8f5a59a7619668a660001b61068b565b617fc57fbf63b2e464c46df4654d0ad2a38c46e26190494934fa55bed3b6211749ef1ac660001b61068b565b600481606001515110156180e957617fff7fdc721fc858c7b465429dd3a54b98151d6cd3c70517d70dd8aa8b7cd9f4d39e9760001b61068b565b61802b7ff9c15e43580ded4b74b30a96e9bd5a2f1b931e8452fbe5e3d41f89837e41efc460001b61068b565b6180577fd1f698de40d380f318fa2d75c06877365ec84cd6e6f62d15291eda905a7ccc1860001b61068b565b7fd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e356040516180849061eeea565b60405180910390a16180b87fd7beff24f77419795791eb1a8dc4a2f6d81c750b4c5b60c0bc35dcfbd382db6860001b61068b565b6180e47ff11060df64ccc4a52fb4216a02ca2f0431b7d32ae1ec75b4a7e3d752953421fb60001b61068b565b61987c565b6181157ffd3b0b8a9f5a9697ed03177e3981275ea7b5cfc6228d4f86be5c9d3e9817259860001b61068b565b6181417f324f74c9760b33354a8cd17302b62dc7c4ab2e022c20517641fd290bdc51fb6d60001b61068b565b61816d7f6a214738e5bb3f03b5adf14924b5302ff1ab1b2244029415402d0956308766d260001b61068b565b61817561d55c565b6181a17f0ce8f81b77f17e24ad8df7c6e4a9dcc93df83b5d2d7de3b95f9747462dce408860001b61068b565b6181cd7f4b42fdfcfc935b17429b7e572d4bbe0635ae450cfc25880369fc61760f2ec95a60001b61068b565b600060029054906101000a900467ffffffffffffffff1667ffffffffffffffff16826000015167ffffffffffffffff1614156186855761822f7fc952a2d20028a99f47f6a8d35bb4c920509b0091d2a70547582944c0b282f33f60001b61068b565b61825b7f17430870cd3d317248ad70b4f81a789f3c40d938c7f9cdb3bf3bd0d7e7d9e62d60001b61068b565b6182877f0adf234fe199f3daad9ed24b4623f1957fa3de7f6e7451940ef6c58fdf2e42bf60001b61068b565b600060018111156182c1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816000019067ffffffffffffffff16908167ffffffffffffffff168152505061830c7f0653e15752b325e9fbce7c486ea7b9d3901720a9f51e4bddd4182a220f559a0b60001b61068b565b6183387fa4f7db2873ed375a4189f3f6048c8220d567b47e64a61e590779b84496fb313b60001b61068b565b600280546183459061f431565b80601f01602080910402602001604051908101604052809291908181526020018280546183719061f431565b80156183be5780601f10618393576101008083540402835291602001916183be565b820191906000526020600020905b8154815290600101906020018083116183a157829003601f168201915b505050505081602001819052506183f77fe06f194c9755c4e990e204cd7b46c8ad400b9bd402a8033837f7a067b22df6a560001b61068b565b6184237fada5bf775388ddc59d9a8143e2440be600d6701060f73a788988fbcceec1309360001b61068b565b618430826060015161739e565b81604001819052506184647ff3ae0bfa57e0dcb6d0db13f5ea89ead306f5db9cc798b5e240e2e32ff9f168a660001b61068b565b6184907f7c43eec5e2635e3a7a01a38c58922b36a23b61fbb18f288178604caa8754a52d60001b61068b565b816060015181606001819052506184c97f3dea5976dcdb6eb6a5cf8febcc55de3c27972d05fb50a66e6b1c8e488a35810960001b61068b565b6184f57f12cf19bf9c1afefb95c3fbce9028b5a61b5e8103eedc80676e73d9cecbfefa5860001b61068b565b8160800151816080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061855d7fcc1ee0c4e0cfb45586c1dfe880812a0a9474f66453a9f12a48764a8b3470b38860001b61068b565b6185897f8c9b61e8fa777a977edf78fc292da065327371c18b5a6c7ef67d9270c8db2dc060001b61068b565b8160a001518160a001819052506185c27f8925dd3545ef8afb010d68f29096537c169c8a5972b2631760a76ca6144ef01060001b61068b565b6185ee7fcc21850c645c73ddc9efdf935647742ddacde4492dfc766214db9eb02d0c651060001b61068b565b6001436185fb919061f262565b8160c00181815250506186307fc454d9b5ddb9a441daf20516a8756551b0a05f3593d31d0b09e739256d43a52660001b61068b565b61865c7facf72d13b4e164cd78c5530aac7ee396303c872a6d4bd69ecbcc3a6aebf6272b60001b61068b565b8160c001518160e0019067ffffffffffffffff16908167ffffffffffffffff16815250506186b2565b6186b17fff37e416c9b07ef1a85936c9a99ead00066bedd910dbc9f328ac5495c9039be860001b61068b565b5b6186de7ff05cdd7917ace9c37a9e88d6f7e13185271795cd2e1d358bf59be5fc699d6ef760001b61068b565b61870a7f6be2d323f1021a6c35c019a828fb71e0bdb44bf7d8d6f3dd5e1dd6c71c05784b60001b61068b565b6000600a9054906101000a900467ffffffffffffffff1667ffffffffffffffff16826000015167ffffffffffffffff161415618b3c5761876c7f364f1c12e04b19ff8a76600eeb2f25e192ba15f10bd23315617f6e4515d21f9260001b61068b565b6187987f255d7bb50d19ce2aca219b2c768021a9117cc933363c6c86db16d369218a5e9760001b61068b565b6187c47f6122922d9ad2b1c00e757e2995a74776ab3a406838bd85c007a41c1e25f254e160001b61068b565b600060018111156187fe577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816000019067ffffffffffffffff16908167ffffffffffffffff16815250506188497f69efadfe0c48314909ce552330c374201d82b4cf122fcbd00f2c5d7b00ebd51360001b61068b565b6188757f39ba41330bf2f881b7a6a93a642152a4ca3f247b1469b337498f5b74dfce8f6d60001b61068b565b816020015181602001819052506188ae7f497f6c91888f60edf923a11522eafe93d52726c187ec036b29d63d97acb77e6260001b61068b565b6188da7ff9313b2216385cfe0eb2b637e3e8aa0c9c15a2e131abcc1107337b2886f8769160001b61068b565b6188e7826060015161739e565b816040018190525061891b7fc6c34f042bf6aafa9258eb6f3a5bfb0928bedf074763fa55e850c47bbfa0b43460001b61068b565b6189477f9e10944566c170a626a03a44f557d09c4b23cbf43aad1617265aad35776692c360001b61068b565b816060015181606001819052506189807fb4f274cd7a958a9451085d2a1558bb4139be00b05a4a41bef74a25d5cfee3b5760001b61068b565b6189ac7f831b31a8eaa8d310fd48bf1c74d87659891abb128594cfb6b1dd17c014c6ab7f60001b61068b565b8160800151816080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050618a147f4496da76e9c7cb69511ac82f6175feae7d4309c02f1ed7c3b62b04386570725d60001b61068b565b618a407f9b7ee34b4ab6a5452f115b257ef1c662a3b22e9f059978e33bb1be6de9ee612560001b61068b565b8160a001518160a00181905250618a797fec3bbf2d801faca3ecd1226ac9dcb130a287e3db4ce56d412e13eb9160dc6ee960001b61068b565b618aa57f473d94544d81c4a001ca355b033c79df16824f29af603ffd3de7d8974168e12b60001b61068b565b600143618ab2919061f262565b8160c0018181525050618ae77fac71c38fc9d15be6efb7663ccd7f589752409072b02c2294e80a91ed29ba78aa60001b61068b565b618b137fe6e2ed8c43030590ec99c4c6e75d4304778e42d5a351a7ab277a551abe7475da60001b61068b565b8160c001518160e0019067ffffffffffffffff16908167ffffffffffffffff1681525050618b69565b618b687f32d954e05813e654404ac82ed270631b2581acfc0b8adfd8f22e28241405961160001b61068b565b5b618b957fc93f300062b40344246fddf40b30df2c2da800d0751096ca15fafe45ff3beb6260001b61068b565b618bc17fb0a5f122cc1f707d9726a6730a605369ebd075417946a9e525cdd2876018e30360001b61068b565b600060129054906101000a900467ffffffffffffffff1667ffffffffffffffff16826000015167ffffffffffffffff16141561907157618c237f7585a6ad31acaaca38cb4de095b982956cc4f03ffcd1d6fcdeba1bd4fe05932c60001b61068b565b618c4f7f48c980dfb1dc8c4bcd7efbda2bcf4e329b5c078b4209adb489a5094cbdfdc0ff60001b61068b565b618c7b7f7e9c062c66fb0ddc2d31b1e74c54e7c3f362de3681e8799b9dbbab83c6eeabc960001b61068b565b60006001811115618cb5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816000019067ffffffffffffffff16908167ffffffffffffffff1681525050618d007f86ef11d452ebeb80a7830c830fc039427b7fc86559a06aa934d04e7833107e9660001b61068b565b618d2c7f1121959d7481e40df16035b8e83af3c91aa3097656a5ffefb895af597c6cb91e60001b61068b565b60028054618d399061f431565b80601f0160208091040260200160405190810160405280929190818152602001828054618d659061f431565b8015618db25780601f10618d8757610100808354040283529160200191618db2565b820191906000526020600020905b815481529060010190602001808311618d9557829003601f168201915b50505050508160200181905250618deb7f693e9d91ee8ce7b8bae75e72653c3e403f0c6b59980d569a3d06aa8c3f53f3ee60001b61068b565b618e177fd0f8f6bd4d2e27d45a1a92c8ccccc2f30df3a89816bfe7135a1f908aebbc9b3460001b61068b565b81604001518160400181905250618e507fd70368cc9919112396bafb107b28315e1f66e7662589defd3f4dc0fb867b444760001b61068b565b618e7c7ff856113f5d39a5fcb974cfaf6996f6f2fe568df75f1766a3337f51b3cfeb064360001b61068b565b81606001518160600181905250618eb57f91a0a0961b39a04d177e2ba66f6076caccd86fdfefde8fdc8398c2568228893c60001b61068b565b618ee17fbc3e69f518c0dcf8a1c94c30d4b96e2e7d6451fb171562d74119a2b9546c219060001b61068b565b8160800151816080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050618f497f3d30f37766838a2921c1dc5a53f8747f1bc7d7cc24be8569a8908f990fce215660001b61068b565b618f757fd2e10045f69aa50038c28d49884a5fa47e0d527f6867c66a0ebd9aa552afdb2160001b61068b565b8160a001518160a00181905250618fae7f02cdbfcc28533634ab127cfce2a1cb8abc252cd42d8046f169bed22b0b4a5f7460001b61068b565b618fda7f265af2264d60c583ecefd7a6b228e324b0f67834362c08573f41aeecf487af6060001b61068b565b600143618fe7919061f262565b8160c001818152505061901c7f184c70aefc324e6bcda53d60d1afa69256fb4b4250b2c621668cd4004ef729ee60001b61068b565b6190487f066d3c117fff81f09357094b233f4638f611a51eee0e852fe67154845eddbd4260001b61068b565b8160c001518160e0019067ffffffffffffffff16908167ffffffffffffffff168152505061909e565b61909d7f231999c0b1d03a328cd980758ed154968a8041bdd199e3ee60c1823d32dee87f60001b61068b565b5b6190ca7fa42e50a33237ed776527b29264c8f33321dde016c95015740916c6fcd3adf89660001b61068b565b6190f67ff276596cf4c7a0e803739556c12a954b528a43b2f667a99f22607d816f63276460001b61068b565b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16826000015167ffffffffffffffff161415619520576191587f7d0cd9e41afd2ed9140f803430edaf7e5d8a34c042af05d1395c61c40008c14060001b61068b565b6191847f29c88494b7f24c11f8a79e04dbc15e69fd1c80db6a9021544f2b37e1e6a3b7ee60001b61068b565b6191b07f798cf62b789bdae3c1f5dee66b3bfa0413a57c8321a5a29e43c3c44a2d1891f360001b61068b565b600060018111156191ea577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816000019067ffffffffffffffff16908167ffffffffffffffff16815250506192357ffa858603a3bfc74d8c449c0759947697df85191304df804c0021e1b4ed7be0a260001b61068b565b6192617fc815267cbee789787c79f68141e6dd3af2d82af55e07521dc430aaed7088dff060001b61068b565b8160200151816020018190525061929a7f04495bc63379a793da0e8b46243ed5d30093479ce5760f3345fa2004577bf34d60001b61068b565b6192c67fa124ab25d5c0d1de160551c87a549b2da8d6ec553ed9f7e3d92bb1cf6a5538ef60001b61068b565b816040015181604001819052506192ff7fc3a8fcc4c107e9844e552fcb36082ffb99ddd4c9184de216c8c58fc6af4c7eaf60001b61068b565b61932b7f8856fbafaf0cfe7c879f921e8f86c80d19e98d8200486681e713ab7fb5bad1cd60001b61068b565b816060015181606001819052506193647fdfb9d7407a041efe085fe0b01925350829d8e0eb1d9bf72f2a09dc01628eaacc60001b61068b565b6193907f7d6182245ef30de9c1ef8842fd1d2b10257cb93732a3f14dab4380b84367040760001b61068b565b8160800151816080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506193f87ffc56b049adea2db240651e7936d545e9bd373ca01bb47974c54cc5509cfaa33560001b61068b565b6194247fbaa46d3b18ad1b4065b340028bdc4c867929519c7c032311983eddc3a3a772b060001b61068b565b8160a001518160a0018190525061945d7f1da0a9f563d07f0a6246f638a6df6f6c2c90b5e948fe4899ee550c2d4ff62ec460001b61068b565b6194897f19ddd5b226bf92130a61a13eafa8f8fabcc940f90297bf4fd38fd99cd576c07560001b61068b565b600143619496919061f262565b8160c00181815250506194cb7fd7a7f6c0992849e54bbc7b2a36ea96087f3543385e0bfed19a46b54f442cd8d760001b61068b565b6194f77f1f67041c0752cf20d0cadbb715d80e42545bd9aa14d7dd2d4053a5c14bae739960001b61068b565b8160c001518160e0019067ffffffffffffffff16908167ffffffffffffffff168152505061954d565b61954c7f647ad2cc0d2125c53c9ac0178f0c02cc95000b509b700dc7a055ed90d124387160001b61068b565b5b6195797f36b5d8c0ac6f960b949e7e57ee46d86e7c043bdab36a3d009f192d1f5ecae44b60001b61068b565b6195a57f55db5e9d8cbcd841da8a4f7314ff8e757e3b3a6e3a15eadd14d2c8fe6029fecb60001b61068b565b60006195b983602001518460400151613eb2565b90506195e77ff25ae2b525a41209c602b40d7eba472be1f825b2043bbe8e2482dfcd3c69f5a160001b61068b565b6196137fda11ef5141c955c9d3c9143a33ed97b9ed834033bd0e41f7f39eb8e90110a4e960001b61068b565b6196298183600661bb139092919063ffffffff16565b506196567fb31f1a787a9aed387120d0323a1f989e45a1beec42222620d273e127a45c443960001b61068b565b6196827f6fd6c3e4245c61ed121ebd6fc5d7b13b961b205099f116a0e09d1483cc0ee07260001b61068b565b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16836000015167ffffffffffffffff161480156196df575060036040516196c9919061ec53565b6040518091039020836020015180519060200120145b156197a4576197107f71e89fb50fee18bd9ed898147657e010ea3186dacdd9cc50a5e47048b15d350960001b61068b565b61973c7fa277feac6cd706f064b86e04bf198d591973496488e030e1fff20b2585db3f8e60001b61068b565b6197687f4e49085eda8566486990eb052f775468ffa468433401b51dc8a33aaa96cea5db60001b61068b565b600160098260405161977a919061ec3c565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506197d1565b6197d07fbcb8c64aa5707a9809f69a37c31cdbfb2f6ef6744708e8a176e9030f4e000cd260001b61068b565b5b6197fd7fd5cab4e7dbc2d5fc4c3dda4a7830c423a1952c50ff2322eeb4e2aae6680315b360001b61068b565b6198297fa209369b81c06a4814124130f9669e8ff136cc4d318923303ce1616a3c33ada060001b61068b565b7f517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e418360800151619861846020015185604001516114c4565b846040516198719392919061ed0a565b60405180910390a150505b50565b60006198ad7fb383a0efcde5f4d3dbf4bb165b6eecf3eb182d11a16103d6a6d1284071393a0460001b61d4a6565b6198d97f55e7c67e70c715207fa8c6ff366d74d5e117623fbf989e11cffd2d2969a677a760001b61d4a6565b6199057f7a56769147c94e50cde86b25f58bd7549fc47921928b344870aa0f3f9ea2089e60001b61d4a6565b6000619912836000619d45565b90506199407f56c39a07f04ffe56ed01bccacf90de9517a2841f869f0753f1edc21319cba07460001b61d4a6565b61996c7fef3e55dac51e8b93f31e1c8e2c4f48a4c441c794b2cdf2ace8ef50d4f2602fd860001b61d4a6565b600181619979919061f2f6565b915050919050565b60006199af7feb16669fb06a24188e89da4622cbf8e9dbe42c9cdd8e9a6099ed137ce67a577360001b61d4a6565b6199db7f65b94642af9997693ba51954aca120af910022f6dd50ea56ad15666f1fec602060001b61d4a6565b619a077f5b47e0fd512ad074e661b5bbf2e0a37cbeec3d77824254f9f55a94f1ef926bf660001b61d4a6565b82600101805490508210905092915050565b6060619a2361d4c2565b619a4f7f5638afbc8b0495f6e017d9238356152b94d97dfc149b9efa97c74e08ff053b1160001b61d4a6565b619a7b7fa07e80f8612067e53798aaee6a33aa995d54b54a42e8bb57d5ea44aa646220d960001b61d4a6565b619aa77fedf2313c236a1135eb01b238728cb3dd500a59d4483c1ade826f09b5af1fbcc460001b61d4a6565b836001018381548110619ae3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600202016000018054619aff9061f431565b80601f0160208091040260200160405190810160405280929190818152602001828054619b2b9061f431565b8015619b785780601f10619b4d57610100808354040283529160200191619b78565b820191906000526020600020905b815481529060010190602001808311619b5b57829003601f168201915b50505050509150619bab7f848f08bda43d99e0006a70fca03655de0b08f5155bff17b182385f65875a4de060001b61d4a6565b619bd77fe5722f9d0f0770d4fbab10d961cc9e1bf8802d591542f9f80b4f04f31ae0a99960001b61d4a6565b8360000182604051619be9919061ec6a565b9081526020016040518091039020600101604051806080016040529081600082018054619c159061f431565b80601f0160208091040260200160405190810160405280929190818152602001828054619c419061f431565b8015619c8e5780601f10619c6357610100808354040283529160200191619c8e565b820191906000526020600020905b815481529060010190602001808311619c7157829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff1660ff1660ff1681526020016001820160159054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090509250929050565b6000619d737fd1f406cf1ae1d7c484c5ea44d3b281f81bf5c9afbec4aaa9605d4af488db612060001b61d4a6565b619d9f7f0a226c341652e67c4f7db7d4732b3a14022e2324b2a20ca6c36b8c4ee1e1815760001b61d4a6565b8180619daa9061f494565b925050619dd97f197b73813bdbd56ba0e92170ca3f43a85d89b377658b04ba51703dcb51ce6fdd60001b61d4a6565b619e057f60a342490a285effcaf427aa0e2e0856b315d0425dccd79572a5f771c6d125e660001b61d4a6565b5b826001018054905082108015619e735750826001018281548110619e53577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160009054906101000a900460ff165b15619eb757619ea47fd1254c75f0f51bcf92a5edad0ae46cace1b3be8c878ef184401510443f48e69460001b61d4a6565b8180619eaf9061f494565b925050619e06565b619ee37fa7e5c15b8be53edab731a7cdf19e4bf61daca3afd7c25580cfd2c6b8906292c260001b61d4a6565b619f0f7f65e4772fadb58f7bc37c9a168fa2a669cef625ca650eb90b589e83c0dd5af44b60001b61d4a6565b81905092915050565b619f2061d50d565b619f4c7f0e7d021362adb47112dd46250ce673fb071b7ecc535945fddda03a7d6c82627360001b61d4a9565b619f787f9cdc5e80a60cd683d8b4729b6dcfdde4a4bed919c3a38769835b96c784b853ce60001b61d4a9565b619fa47f5b8d2a68e041c2f332ccfe8a70c8b6703a876bb97767de62d81e3700faa1e46360001b61d4a9565b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461a0599061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461a0859061f431565b801561a0d25780601f1061a0a75761010080835404028352916020019161a0d2565b820191906000526020600020905b81548152906001019060200180831161a0b557829003601f168201915b5050505050815260200160028201805461a0eb9061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461a1179061f431565b801561a1645780601f1061a1395761010080835404028352916020019161a164565b820191906000526020600020905b81548152906001019060200180831161a14757829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160048201805461a1af9061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461a1db9061f431565b801561a2285780601f1061a1fd5761010080835404028352916020019161a228565b820191906000526020600020905b81548152906001019060200180831161a20b57829003601f168201915b505050505081525050905092915050565b600061a2677f8087b515667a1beeedba3b06ef6b51286980813d53d14d90d3ad5d39d56377fe60001b61d4a9565b61a2937fae15b9e7eb535c0af30bcf7211f6b6058f2c2b5cdf34e1024e4c16c527c574dd60001b61d4a9565b61a2bf7facc1c4a12eaf689e8c5a9302204d3467e8d4dabdf6a961802f8ff6c5d9e8741a60001b61d4a9565b60008460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154905061a3347f3af8387c7c120b6a19adf8aebbf75b25522e6ec88179e98816d0a96ef722e84760001b61d4a9565b61a3607f61814946f955635f409ca439d3507c008e81fd5e011883df4b32ad432e1398ff60001b61d4a9565b828560000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908051906020019061a40892919061d65a565b50604082015181600201908051906020019061a42592919061d65a565b5060608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550608082015181600401908051906020019061a47192919061d766565b5090505061a4a17f392d11498afcac3fca8145669cc7eff2dd940e7a05cc1af4d9990a3544a9264260001b61d4a9565b61a4cd7f965b2931e9311d83b0229a30ddfe0781f48a898e818605a6abaed2d38d51396c60001b61d4a9565b600081111561a5385761a5027f8e633912915f86c57c82dcec7f08d6618de0831bd37c2e52593c02c836f8272260001b61d4a9565b61a52e7ff3842316fc7acf644552100dd393529107c9c68dfcdd98434be6675ccf50c81060001b61d4a9565b600191505061a86e565b61a5647f13662e72bb05aa66b755286731a27bab72b4895b66f5ab1e8e20b6d5b3abc00260001b61d4a9565b61a5907fc88c790825eb177d327dca6aae3b7735648231c0abc90ecaa7452fafb641581c60001b61d4a9565b61a5bc7f6972424816cdaea3a2e0252aee5d2d45e0bba8b02fde1bf61bc067636e077a8b60001b61d4a9565b8460010180549050905061a5f27f51ff843092b149dc37025bbc7cb1d6629a0bfe4729088f248072580b17d2c8c860001b61d4a9565b61a61e7ff1766afda0d0b9d7233ffe42abd54a0a4b6fcdd331bf3c325699765ac004eaac60001b61d4a9565b8460010160018160018154018082558091505003906000526020600020505061a6697ff40e7ce1328c5b3016130465f437085766623171a037c7ff3477de9c33af75bf60001b61d4a9565b61a6957f44bbeb566f70edf2ece620030eab3faeea1d2a62dfa7bfd8d943d79c8c1f2f3260001b61d4a9565b60018161a6a2919061f262565b8560000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555061a7167f88b6c30d0ef71d4b26fb0ef1b9bc255ff75859d309b18d2a8859f255a5bec26060001b61d4a9565b61a7427fd1b8ece9167f8e93911e37d43dc67dca0bff02c448ec1e2e5c736c5c7d291a2760001b61d4a9565b8385600101828154811061a77f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061a7f67f21598e5dd37fa4a7d873e9d84c014360fbbe5873daa8bdd300b31663dfc92a0860001b61d4a9565b84600201600081548092919061a80b9061f494565b919050555061a83c7f44bc82bc0c1288048b0a9dadf9f305c99c18093051febe567b02dcc91573466460001b61d4a9565b61a8687f50d6fe0907b8ab06eaeda1050b5eb5ce91be851755b04ec4e3e201783f88201860001b61d4a9565b60009150505b9392505050565b61a87d61d4c2565b61a8a97f738bdabc0471eb1163f6e2f990696b93af705b041aa311ea65ca382dd4c4167960001b61d4a6565b61a8d57f58bfcef092feb61ed9a4865df869551f4ce5d721fd99fa48650b6fe768184bb460001b61d4a6565b61a9017f3a7aa24af5adfb0fa96b7a5133be1ff29eff0610211a6d273364b75d6a5fc1fd60001b61d4a6565b826000018260405161a913919061ec6a565b908152602001604051809103902060010160405180608001604052908160008201805461a93f9061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461a96b9061f431565b801561a9b85780601f1061a98d5761010080835404028352916020019161a9b8565b820191906000526020600020905b81548152906001019060200180831161a99b57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff1660ff1660ff1681526020016001820160159054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050905092915050565b600061aa9c7fa43a753ff5eca1a01850e23643e4cda2e8f4d646fcf1d1158f337fb86542a42060001b61d4a6565b61aac87f2a0c1256b5fafbf51b029791fa14038f9512d7ddf850f7583d6741d1b68daa1260001b61d4a6565b61aaf47f38a27fc4df42a59c2fcf021f2da44e559a2c27c3fbe82ac441ce9ac54faacc5160001b61d4a6565b6000846000018460405161ab08919061ec6a565b908152602001604051809103902060000154905061ab487fba9a895b966129cc955ef0de0b74ec45d2de29dd1acfae347b8d72a37106a9e960001b61d4a6565b61ab747f289fd3da4ae9fe7ea2f39fdfc788f510e40a2240e2d56b30441f2bced1a7291560001b61d4a6565b82856000018560405161ab87919061ec6a565b9081526020016040518091039020600101600082015181600001908051906020019061abb492919061d766565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff021916908360ff16021790555060608201518160010160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555090505061ac7a7e5c975b12caca9640446685a054860addd8bd69f6ce2a81cbc95c2fe5f62a0f60001b61d4a6565b61aca67ffce94ada74d5f5600cf281af61c6bc8d5624409282c28b87b7ebc8e13b8c25c960001b61d4a6565b600081111561ad115761acdb7f36d8aac8547e1c916da1b7582cae336d0f1c1b311ecbf667f11f00cec905c4c060001b61d4a6565b61ad077f62408aca115f945689610d20bd7ff068e7450c455921987ffafe02c264b8f25560001b61d4a6565b600191505061b001565b61ad3d7f3dd6f3ae35ad505fc593a08ba265376de2c5254cdbf353833fc85e5f0da4330560001b61d4a6565b61ad697fb96328241b4e9c993674126748ecd3385e9d83991354297af33940666622f86c60001b61d4a6565b61ad957fb6a73cd5456b2cfcc3d41baea7ddbdfc4e6b868269d8778624af396167967dbc60001b61d4a6565b8460010180549050905061adcb7f50e19cf330806d55f2a3bc5844d5955438ffefd378c11233ae515e2e8c7948c360001b61d4a6565b61adf77ff7377904345a1e7ae9e6bc7e88c0f8b12eea7c4c4a50d83779e291d4f748bfac60001b61d4a6565b846001016001816001815401808255809150500390600052602060002090505061ae437f1478fa80c63c46f4452b3d500e6cd260c713c51a8b79e0427451ec0b82a093b260001b61d4a6565b61ae6f7fbbdf7f75b3337fbf1068b2c60dca365779acbec4050629ef18b0b25c4eced62b60001b61d4a6565b60018161ae7c919061f262565b856000018560405161ae8e919061ec6a565b90815260200160405180910390206000018190555061aecf7f9dbdc7c0b43cd2538b0e03bdb1787911b55bff0784fc870e5ff0a22f6e2f176960001b61d4a6565b61aefb7f7eab8c7754ee6be0609ee044a384b47019c5487c939e43642df6118f9cf5b00d60001b61d4a6565b8385600101828154811061af38577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600001908051906020019061af5c92919061d766565b5061af897fff13160362183a01f1bc2e5e0cb2d1f632fcca6d9274a0d35fce193e8c54a93760001b61d4a6565b84600201600081548092919061af9e9061f494565b919050555061afcf7f1c7cc2791f6e39276f2fea114fa156ad029e252b07edc61a8ee7680950e7b99460001b61d4a6565b61affb7f180c21ec81d31cbb0cf00250b9566a700828b07baf6ad30e71c372b83069f15460001b61d4a6565b60009150505b9392505050565b600061b0367f94f39fba1c7e9e20d5296b487586761d60fa050dce963e91162deb5a1c3e86b360001b61d4a6565b61b0627f28f790d721aba6b5aad7e72a6df51f7cafdc6906af1ea183e85e987e281d8c2e60001b61d4a6565b61b08e7fd8bd77790dde30e31b3816ced48cccf9e3b77eb7c57a782bdca240cff1e38c7460001b61d4a6565b6000836000018360405161b0a2919061ec6a565b908152602001604051809103902060000154905061b0e27f9c63fcd68d8c5014241740dba06ad129f5664c8ddbc39a5bc6ba88d76e46ea7960001b61d4a6565b61b10e7f615c6ddccc936c74c925471adc60b0fde8708cb11ac83ba3b2561aad8dd4d96260001b61d4a6565b600081141561b1795761b1437fed3d88ba371925fe033b2dcba2c2cceac2a7bb602cbcfae50d9e70095505b53260001b61d4a6565b61b16f7f9679d9b774ad7482477a46fe33e4eb7bf22e5bddda7ec09c09e2abf60e6958d260001b61d4a6565b600091505061b37b565b61b1a57fbf802e2b70c77fbcb2bd6726463d962e1754278ea153dc85bba6485092de508560001b61d4a6565b61b1d17f04d124d12d123eca6f4cc2ed2a3e290c2d9ba831a77c6020275a3ef22264d9e860001b61d4a6565b836000018360405161b1e3919061ec6a565b9081526020016040518091039020600080820160009055600182016000808201600061b20f919061d7ec565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160146101000a81549060ff02191690556001820160156101000a81549067ffffffffffffffff02191690555050505061b2957f15ea82bcab93bc5b74d9e0a3529c1afdb2531a6a69333e734b5ca751fa682cb160001b61d4a6565b61b2c17fe10c4cbad61920c82ad756a4da44c46f75972993003b090a3af7573c374c252c60001b61d4a6565b60018460010160018361b2d4919061f2f6565b8154811061b30b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555061b35f7f07d39a1bcfb195069832049529b3f1504afb8e21d216f094789c3792e85f841560001b61d4a6565b83600201600081548092919061b3749061f407565b9190505550505b92915050565b600061b3af7fc65aba6e61cf4328d3711b7e14a1c6439e342a595c2e976557e70aae0b6a2ca260001b61d4a9565b61b3db7f09268a469ac7ac147b84d764e05b531857f5ae966a49320bdaf76a013781107560001b61d4a9565b61b4077f5ea5efc51e38941d4750c4083671ddfa990999332619bfb54b49847b41c7496260001b61d4a9565b60008360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154905061b47c7f7be6b9e4c9556ac2cbdbad39ce4018376fa5bf1d9570f08dd0afd47c959d406460001b61d4a9565b61b4a87f1f8bd4a3c573dd143cdf29f6bcdbaed42ad69be7e30d1eb8658ee67b597c4ed960001b61d4a9565b600081141561b5135761b4dd7f127355c97c81f2b880f11aa2370b4f9b4b21b626b1e1d4c83d587a4705fa9af260001b61d4a9565b61b5097f83e49c7336afb9b2db44167d587298c5bd65b5d2740952411297f394d3c35f2e60001b61d4a9565b600091505061b73e565b61b53f7fb69f8a6afd39d040e265a2ea329c1a6ec772ae1cd86b3fff1f7cdff85c1a84b860001b61d4a9565b61b56b7f752ca27bf6af1d7ecfa4de60b0e44a407ee81350769191d26ce13d1c292bb20960001b61d4a9565b8360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600061b5f1919061d5cb565b60028201600061b601919061d5cb565b6003820160006101000a81549067ffffffffffffffff021916905560048201600061b62c919061d7ec565b5050505061b65c7fc22186f5eb6293580915d3ba007c5ecabffb6a6e6e374c3aa95f87b78f4ffed760001b61d4a9565b61b6887f03bacbe991c4b7d0c6a1fe77ec03d623fa2edc0358e7a808669b772074c48cf660001b61d4a9565b60018460010160018361b69b919061f2f6565b8154811061b6d2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160146101000a81548160ff02191690831515021790555061b7227fd16e51e22de8e6b42b792699c61ae72e78b84c4f659af6072ea80af81b49511060001b61d4a9565b83600201600081548092919061b7379061f407565b9190505550505b92915050565b61b74c61d55c565b61b7787f221afdfd06ff554cabb9923c56bd59eab0f54ec6ade96fe9e2afc75c5e74356c60001b61d4ac565b61b7a47fd220eac2469541adbf15f11825b6cb64dfddde4d8b78ba4db4136c2d132bc64b60001b61d4ac565b61b7d07fd07aeb08788c154234295ea66dd7c17c21b5de146bf43393bf37ee9646b0e97c60001b61d4ac565b826000018260405161b7e2919061ec3c565b9081526020016040518091039020600101604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461b8419061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461b86d9061f431565b801561b8ba5780601f1061b88f5761010080835404028352916020019161b8ba565b820191906000526020600020905b81548152906001019060200180831161b89d57829003601f168201915b5050505050815260200160028201805461b8d39061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461b8ff9061f431565b801561b94c5780601f1061b9215761010080835404028352916020019161b94c565b820191906000526020600020905b81548152906001019060200180831161b92f57829003601f168201915b5050505050815260200160038201805461b9659061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461b9919061f431565b801561b9de5780601f1061b9b35761010080835404028352916020019161b9de565b820191906000526020600020905b81548152906001019060200180831161b9c157829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201805461ba4d9061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461ba799061f431565b801561bac65780601f1061ba9b5761010080835404028352916020019161bac6565b820191906000526020600020905b81548152906001019060200180831161baa957829003601f168201915b50505050508152602001600682015481526020016007820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050905092915050565b600061bb417fc2868a4059bbae7e09f80f29e3e73f92d7da27b2a8145931dcdf7237f3bc247460001b61d4ac565b61bb6d7f505cb07d155a99988d1601beda3a2c4fe1616a8e57cb21f427c6385edd0bf3c160001b61d4ac565b61bb997fe594d96b6f60c626d5637f18c9f0300612ff4cfc39be1bf0b1a1f57c8306c6d160001b61d4ac565b6000846000018460405161bbad919061ec3c565b908152602001604051809103902060000154905061bbed7f9a0adebb92b2dbe6c5c96230999eed4415130b6e7b322c1de24754068ad817b160001b61d4ac565b61bc197f5ade4149ad458a74efdbddbf17d495280a5755fe6e8316e67e9a75c0714d61e160001b61d4ac565b82856000018560405161bc2c919061ec3c565b908152602001604051809103902060010160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550602082015181600101908051906020019061bc8892919061d65a565b50604082015181600201908051906020019061bca592919061d65a565b50606082015181600301908051906020019061bcc292919061d65a565b5060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501908051906020019061bd2692919061d65a565b5060c0820151816006015560e08201518160070160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555090505061bd8f7f506d12a376a51c5a19eabaf59349a20d60ceba6c57c7b763bb2ca47e0561621d60001b61d4ac565b61bdbb7fe546ea4ec05918c177b7cefe82ddac1614f98de879792fee775307fcf6d9be4060001b61d4ac565b600081111561be265761bdf07f968501f5b96e5925d47892788ce7672d10144f8ed662072c73a214492fecf81860001b61d4ac565b61be1c7fa67780f1f943328605236ffa52b6b07e8e7976257b01ebd15c1995ea1965401a60001b61d4ac565b600191505061c116565b61be527f844eec9983b8a791be2b4adf314591700f298d1216837902df3340606e10ed9360001b61d4ac565b61be7e7f04af785957bb3a19a7aa6ee6c50328aed4967e27de9c8af35816c6f9612f2a8960001b61d4ac565b61beaa7f75f0131d5aa828bcd6827ce0d2591cb0f92a49d2f600c16f89d5639c00b8661460001b61d4ac565b8460010180549050905061bee07fa074913b772ff1282674c79431da46b993efe37838dee97b7eff5fd24d5031b560001b61d4ac565b61bf0c7fd965cc23f27e682528d62828fb49aa031df7b3774c720376f95b56ecae2a97f160001b61d4ac565b846001016001816001815401808255809150500390600052602060002090505061bf587f69c78817124639b8120850776c2636a10c1c954ed67e11e20ce975996947683f60001b61d4ac565b61bf847fd0c25fb18c711d113702c1510daceaa056a86104499afe104f38241b6fe43a6760001b61d4ac565b60018161bf91919061f262565b856000018560405161bfa3919061ec3c565b90815260200160405180910390206000018190555061bfe47f5df09431c77236c72443f6c6b2859f61eebb4bf0601063007a32c08d2d90daab60001b61d4ac565b61c0107f6858bf2bdc712f32fb8af171a9a2c8145ab59e7a443cdbb0f4b11bd71640f5cf60001b61d4ac565b8385600101828154811061c04d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600001908051906020019061c07192919061d65a565b5061c09e7f0d520bcfb9fef83c24152bef20c51038142ed1ed1875b69db1065b8aec2d9c8d60001b61d4ac565b84600201600081548092919061c0b39061f494565b919050555061c0e47fbe3dca0f7562c54744c4e322c237017b22dd7a20f56915f3f13316a8011dbef760001b61d4ac565b61c1107fc7bbac8170187e5f7aeff3d32aeca24baa8ac1178648bb00f7c3443c33b7186860001b61d4ac565b60009150505b9392505050565b600061c14b7fd2ea2c5f4e69985101587c43fdb99bc958197245f4843e72be988ab82ca83a0160001b61d4ac565b61c1777f5e3fc0190250497ec3b6335bb59892d82025ed1f2e1984d2a8f2bd3dbf517c9f60001b61d4ac565b61c1a37f8f91ae81cba2a1878ead79362512e27dfa8787cdb05d57fcc326dd58da6cfd0960001b61d4ac565b6000836000018360405161c1b7919061ec3c565b908152602001604051809103902060000154905061c1f77f2d70e10b947d87040893d164d298e81757f7da06b1206daa551649e6296aa2d260001b61d4ac565b61c2237f894782006ec3feb183423dcab429dd25a115709ac530702c8c7892d9da3b5fe560001b61d4ac565b600081141561c28e5761c2587f0596b2ca8b1a12a9de2429f84248577d917ff45a6037b46bd6cd494c08f9e9fc60001b61d4ac565b61c2847fb16ade825985893db3ed7bb78052bba8f1542e7b5250d32370faa3bd1ba4d84160001b61d4ac565b600091505061c4cf565b61c2ba7f88fa0d00cf81808908c33583d766888696c1d7caa0b4cc1a22f6d9fc4e05b5b660001b61d4ac565b61c2e67f8c1d9e1857456cf81d3e53ff7da8aded2825de28e861514fd3da469c8c83403460001b61d4ac565b836000018360405161c2f8919061ec3c565b908152602001604051809103902060008082016000905560018201600080820160006101000a81549067ffffffffffffffff021916905560018201600061c33f919061d5cb565b60028201600061c34f919061d5cb565b60038201600061c35f919061d5cb565b6004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560058201600061c396919061d5cb565b60068201600090556007820160006101000a81549067ffffffffffffffff02191690555050505061c3e97fb098cefb801dc6a9accb37c6cb39af21240603ccb144e5553f4b6ac50bb95e5260001b61d4ac565b61c4157f9f00185e1e47e50b7d1bc2eccc0ab2c104a1f568c532d97c7fb32364b53a7b3c60001b61d4ac565b60018460010160018361c428919061f2f6565b8154811061c45f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555061c4b37faf1316efbbd405bd68ab03adf06718d2ff5f57bc48287e775c024f962f58937d60001b61d4ac565b83600201600081548092919061c4c89061f407565b9190505550505b92915050565b600061c5037f273bee86e7e5f8a0e05bc2c7f8353c69d02b3b3bd88788c26d7f64155561dbfb60001b61d4a9565b61c52f7ffac3b926e8c1d84df5fd7a4dee26bbe6053d5295f775c407c5bcc7ed5fd0508f60001b61d4a9565b61c55b7fdc799cf0573305129a038008235fdb6ea8854adbb50baaecfcc18ba93cec90a160001b61d4a9565b600061c56883600061ca57565b905061c5967fb515ed25f7e00fc1397cc1e359b7d91b1203292b4b66daf373b87f6dff7090c660001b61d4a9565b61c5c27f07f0bf1c284fc3e059b4d4efdb6a3557db24aced3b1ebf082f6ba1dc6fda8b9b60001b61d4a9565b60018161c5cf919061f2f6565b915050919050565b600061c6057fc08c0ef1764714aa08ac6f31dfed3dd63b17f41b2a87594b2569488c8769ac0760001b61d4a9565b61c6317f92e1511f16d42703e360d1b671b2b02572a06b27f7880014961ab86d01e8d55560001b61d4a9565b61c65d7f7d30f84e3417d2d46de70aefc3827aac0bfe823b8cebbfc5f8e261b16abfc75960001b61d4a9565b82600101805490508210905092915050565b600061c67961d50d565b61c6a57f37f32662849c88272abbe0c8e0bab07a3f63fbb4faf6d373d25cd8f103936e8960001b61d4a9565b61c6d17f1a37f2b17e5bbbcc71603879ad461b6169c94b1d75cfdb10bfe55dcca269b80560001b61d4a9565b61c6fd7f31dd8dd00b447ed1527675410addfb6adbc77d6ac6de05d4e765dca993cd7b2160001b61d4a9565b83600101838154811061c739577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915061c7957fc0d2353442e772344d44f2828a3c55b6ebd4a1753231508a1e375f52c813cd9c60001b61d4a9565b61c7c17f1cc5dcdacc2155d3bb0c1c3c0af10810886fcf0a2d7f5554d988c6d505d238d560001b61d4a9565b8360000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461c8769061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461c8a29061f431565b801561c8ef5780601f1061c8c45761010080835404028352916020019161c8ef565b820191906000526020600020905b81548152906001019060200180831161c8d257829003601f168201915b5050505050815260200160028201805461c9089061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461c9349061f431565b801561c9815780601f1061c9565761010080835404028352916020019161c981565b820191906000526020600020905b81548152906001019060200180831161c96457829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160048201805461c9cc9061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461c9f89061f431565b801561ca455780601f1061ca1a5761010080835404028352916020019161ca45565b820191906000526020600020905b81548152906001019060200180831161ca2857829003601f168201915b50505050508152505090509250929050565b600061ca857fd2ce73549559727d8279feba875c82ce8cb4317db2981934ff5c8463134b584a60001b61d4a9565b61cab17f51a21018693c9ddb8774de7a67c97f3c17d7c1e415bb77e7eb124d5eda08223660001b61d4a9565b818061cabc9061f494565b92505061caeb7f3a917ee860ea1019dd7a06da933188138b40d5a98388678e5a83290545d4f6e960001b61d4a9565b61cb177f27880a108071f2784bd0cb600bdbf23003cfab95540ea458a0e2da41bf1bdc3b60001b61d4a9565b5b82600101805490508210801561cb81575082600101828154811061cb65577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160149054906101000a900460ff165b1561cbc55761cbb27fa8b7d268e95962401d3498b1c62d1f436b5332b6fcd4bacb9352a3136a89a58f60001b61d4a9565b818061cbbd9061f494565b92505061cb18565b61cbf17faadfa02caea12f96559d95a6b309adbf58a2017da67eb0d7a9b3a202b350e49660001b61d4a9565b61cc1d7fd512b343aecb423395225b6c52cb659dd37b681ad334c03059b297776c20f71360001b61d4a9565b81905092915050565b600061cc313061d4af565b15905090565b600061cc657fc8742d9e3bdb36a600283808962d935f6b07a40a97affa1c84f7a1c729c2aab260001b61d4ac565b61cc917f6123b657ee2e73c8d5351b525c37a46fc362f161c51e14dff32581ce99393d8f60001b61d4ac565b61ccbd7f6fd7f9ef94d043be6a84eabe2f5b1ac942756d8d15a04ac6385a87014503f55460001b61d4ac565b600061ccca83600061d2d3565b905061ccf87fcb97919e40e1d52e53036f30a25e3f155bdeda0c5629792ca43d3af81797026c60001b61d4ac565b61cd247fa23faabcd5682f59259c1802cb36b4b321dbd97ec041b27b15667eb0507adac260001b61d4ac565b60018161cd31919061f2f6565b915050919050565b600061cd677f978f65843eb2085eeaf6fa050c32a2e52759a2a79e03835081465233f2d162f560001b61d4ac565b61cd937f75eaa14c9b860a70f1db9ac285c4ffcec53f966f72ac9ad05bb468ecd2d11d1f60001b61d4ac565b61cdbf7fd81f7be7b6db562a291f9644e8442d15baac103608f5ebfac7b19ceff3be184860001b61d4ac565b82600101805490508210905092915050565b606061cddb61d55c565b61ce077f4677113a72b688547033889e54972268d5cc448b1da89f8cbcc018570c2ae41f60001b61d4ac565b61ce337f608028e45364cb876409203b751b4008088292bbf7a567cd486b3396a616fd4f60001b61d4ac565b61ce5f7f4686bbc730c8a8d3cd45162b55d50d4756cc72d4d65f96d43c379898fa77866960001b61d4ac565b83600101838154811061ce9b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600001805461ceb79061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461cee39061f431565b801561cf305780601f1061cf055761010080835404028352916020019161cf30565b820191906000526020600020905b81548152906001019060200180831161cf1357829003601f168201915b5050505050915061cf637f66d33bd97f03167ae4be3c2c7665469823a72f784a0200ad167d61b21f71919d60001b61d4ac565b61cf8f7fae2e212b887edb183ed7379ce8eb6626f1503b2ea8312773ff67ea264d443d1360001b61d4ac565b836000018260405161cfa1919061ec3c565b9081526020016040518091039020600101604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461d0009061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461d02c9061f431565b801561d0795780601f1061d04e5761010080835404028352916020019161d079565b820191906000526020600020905b81548152906001019060200180831161d05c57829003601f168201915b5050505050815260200160028201805461d0929061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461d0be9061f431565b801561d10b5780601f1061d0e05761010080835404028352916020019161d10b565b820191906000526020600020905b81548152906001019060200180831161d0ee57829003601f168201915b5050505050815260200160038201805461d1249061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461d1509061f431565b801561d19d5780601f1061d1725761010080835404028352916020019161d19d565b820191906000526020600020905b81548152906001019060200180831161d18057829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201805461d20c9061f431565b80601f016020809104026020016040519081016040528092919081815260200182805461d2389061f431565b801561d2855780601f1061d25a5761010080835404028352916020019161d285565b820191906000526020600020905b81548152906001019060200180831161d26857829003601f168201915b50505050508152602001600682015481526020016007820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090509250929050565b600061d3017f5ec4d151e808a890ad4c1d262eec962101feb35fbc7ac2666c3876ac286999ad60001b61d4ac565b61d32d7ffed3c4d11cf000ccac79a5094ed83e7a4c52aec8484e6d3f41a61b64d836019a60001b61d4ac565b818061d3389061f494565b92505061d3677f5a9d6be27951e631f16558475fd5c0ce4f30a93f4bda20ea9277f55c2cf54b8f60001b61d4ac565b61d3937f83ee8bc1d2850292e6730fdd3ab19d4ed228eab8706e3a63fec2ef206d120c6460001b61d4ac565b5b82600101805490508210801561d401575082600101828154811061d3e1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160009054906101000a900460ff165b1561d4455761d4327f53ca2cac28cf17a5dcc954eac340d63989411ea38ddf440412ea605653f56d1e60001b61d4ac565b818061d43d9061f494565b92505061d394565b61d4717fbc3533477ebffdeeff58335e439caefcb7f4a10837840e02ad6a41b302a3850060001b61d4ac565b61d49d7f9de6e0a748ec21a32e938f692677a782ab21f23afc54a7bb7893e1b2d642486260001b61d4ac565b81905092915050565b50565b50565b50565b600080823b905060008111915050919050565b604051806080016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600067ffffffffffffffff168152602001606081525090565b604051806101000160405280600067ffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160008152602001600067ffffffffffffffff1681525090565b50805461d5d79061f431565b6000825580601f1061d5e9575061d608565b601f01602090049060005260206000209081019061d607919061d82c565b5b50565b6040518060a0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160008152602001600067ffffffffffffffff1681525090565b82805461d6669061f431565b90600052602060002090601f01602090048101928261d688576000855561d6cf565b82601f1061d6a157805160ff191683800117855561d6cf565b8280016001018555821561d6cf579182015b8281111561d6ce57825182559160200191906001019061d6b3565b5b50905061d6dc919061d82c565b5090565b82805461d6ec9061f431565b90600052602060002090601f01602090048101928261d70e576000855561d755565b82601f1061d72757805160ff191683800117855561d755565b8280016001018555821561d755579182015b8281111561d75457825182559160200191906001019061d739565b5b50905061d762919061d82c565b5090565b82805461d7729061f431565b90600052602060002090601f01602090048101928261d794576000855561d7db565b82601f1061d7ad57805160ff191683800117855561d7db565b8280016001018555821561d7db579182015b8281111561d7da57825182559160200191906001019061d7bf565b5b50905061d7e8919061d82c565b5090565b50805461d7f89061f431565b6000825580601f1061d80a575061d829565b601f01602090049060005260206000209081019061d828919061d82c565b5b50565b5b8082111561d84557600081600090555060010161d82d565b5090565b600061d85c61d8578461f0d0565b61f0ab565b90508281526020810184848401111561d87457600080fd5b61d87f84828561f3c5565b509392505050565b600061d89a61d8958461f101565b61f0ab565b90508281526020810184848401111561d8b257600080fd5b61d8bd84828561f3c5565b509392505050565b60008135905061d8d48161f812565b92915050565b60008135905061d8e98161f829565b92915050565b60008151905061d8fe8161f829565b92915050565b600082601f83011261d91557600080fd5b813561d92584826020860161d849565b91505092915050565b600082601f83011261d93f57600080fd5b813561d94f84826020860161d887565b91505092915050565b60006060828403121561d96a57600080fd5b61d974606061f0ab565b9050600082013567ffffffffffffffff81111561d99057600080fd5b61d99c8482850161d92e565b600083015250602061d9b08482850161d8c5565b602083015250604061d9c48482850161dee0565b60408301525092915050565b600060a0828403121561d9e257600080fd5b61d9ec60a061f0ab565b9050600061d9fc8482850161d8c5565b600083015250602082013567ffffffffffffffff81111561da1c57600080fd5b61da288482850161d904565b602083015250604082013567ffffffffffffffff81111561da4857600080fd5b61da548482850161d904565b604083015250606061da688482850161dee0565b606083015250608082013567ffffffffffffffff81111561da8857600080fd5b61da948482850161d92e565b60808301525092915050565b60006040828403121561dab257600080fd5b61dabc604061f0ab565b9050600082013567ffffffffffffffff81111561dad857600080fd5b61dae48482850161d92e565b600083015250602061daf88482850161d8c5565b60208301525092915050565b60006060828403121561db1657600080fd5b61db20606061f0ab565b9050600082013567ffffffffffffffff81111561db3c57600080fd5b61db488482850161d904565b600083015250602082013567ffffffffffffffff81111561db6857600080fd5b61db748482850161d904565b602083015250604061db888482850161d8c5565b60408301525092915050565b60006080828403121561dba657600080fd5b61dbb0608061f0ab565b9050600082013567ffffffffffffffff81111561dbcc57600080fd5b61dbd88482850161d904565b600083015250602061dbec8482850161d8c5565b602083015250604082013567ffffffffffffffff81111561dc0c57600080fd5b61dc188482850161d904565b604083015250606061dc2c8482850161dee0565b60608301525092915050565b600060e0828403121561dc4a57600080fd5b61dc5460e061f0ab565b9050600061dc648482850161dee0565b600083015250602082013567ffffffffffffffff81111561dc8457600080fd5b61dc908482850161d904565b602083015250604082013567ffffffffffffffff81111561dcb057600080fd5b61dcbc8482850161d904565b604083015250606082013567ffffffffffffffff81111561dcdc57600080fd5b61dce88482850161d904565b606083015250608061dcfc8482850161d8c5565b60808301525060a082013567ffffffffffffffff81111561dd1c57600080fd5b61dd288482850161d904565b60a08301525060c061dd3c8482850161dee0565b60c08301525092915050565b60006080828403121561dd5a57600080fd5b61dd64608061f0ab565b9050600082013567ffffffffffffffff81111561dd8057600080fd5b61dd8c8482850161d904565b600083015250602082013567ffffffffffffffff81111561ddac57600080fd5b61ddb88482850161d904565b602083015250604061ddcc8482850161d8c5565b604083015250606061dde08482850161d8c5565b60608301525092915050565b60006040828403121561ddfe57600080fd5b61de08604061f0ab565b9050600082013567ffffffffffffffff81111561de2457600080fd5b61de308482850161d92e565b600083015250602061de448482850161d8c5565b60208301525092915050565b60006060828403121561de6257600080fd5b61de6c606061f0ab565b9050600061de7c8482850161d8c5565b600083015250602082013567ffffffffffffffff81111561de9c57600080fd5b61dea88482850161d904565b602083015250604082013567ffffffffffffffff81111561dec857600080fd5b61ded48482850161d904565b60408301525092915050565b60008135905061deef8161f840565b92915050565b60006020828403121561df0757600080fd5b600061df158482850161d8c5565b91505092915050565b60006020828403121561df3057600080fd5b600061df3e8482850161d8da565b91505092915050565b60006020828403121561df5957600080fd5b600061df678482850161d8ef565b91505092915050565b60006020828403121561df8257600080fd5b600082013567ffffffffffffffff81111561df9c57600080fd5b61dfa88482850161d904565b91505092915050565b6000806040838503121561dfc457600080fd5b600083013567ffffffffffffffff81111561dfde57600080fd5b61dfea8582860161d904565b925050602083013567ffffffffffffffff81111561e00757600080fd5b61e0138582860161d904565b9150509250929050565b60006020828403121561e02f57600080fd5b600082013567ffffffffffffffff81111561e04957600080fd5b61e0558482850161d92e565b91505092915050565b60006020828403121561e07057600080fd5b600082013567ffffffffffffffff81111561e08a57600080fd5b61e0968482850161d958565b91505092915050565b60006020828403121561e0b157600080fd5b600082013567ffffffffffffffff81111561e0cb57600080fd5b61e0d78482850161d9d0565b91505092915050565b60006020828403121561e0f257600080fd5b600082013567ffffffffffffffff81111561e10c57600080fd5b61e1188482850161daa0565b91505092915050565b60006020828403121561e13357600080fd5b600082013567ffffffffffffffff81111561e14d57600080fd5b61e1598482850161db04565b91505092915050565b60006020828403121561e17457600080fd5b600082013567ffffffffffffffff81111561e18e57600080fd5b61e19a8482850161db94565b91505092915050565b60006020828403121561e1b557600080fd5b600082013567ffffffffffffffff81111561e1cf57600080fd5b61e1db8482850161dc38565b91505092915050565b60006020828403121561e1f657600080fd5b600082013567ffffffffffffffff81111561e21057600080fd5b61e21c8482850161dd48565b91505092915050565b60006020828403121561e23757600080fd5b600082013567ffffffffffffffff81111561e25157600080fd5b61e25d8482850161ddec565b91505092915050565b60006020828403121561e27857600080fd5b600082013567ffffffffffffffff81111561e29257600080fd5b61e29e8482850161de50565b91505092915050565b600061e2b3838361e810565b905092915050565b600061e2c7838361e995565b905092915050565b600061e2db838361eb1f565b905092915050565b61e2ec8161f35e565b82525050565b61e2fb8161f35e565b82525050565b600061e30c8261f177565b61e316818561f1d5565b93508360208202850161e3288561f132565b8060005b8581101561e364578484038952815161e345858261e2a7565b945061e3508361f1ae565b925060208a0199505060018101905061e32c565b50829750879550505050505092915050565b600061e3818261f182565b61e38b818561f1e6565b93508360208202850161e39d8561f142565b8060005b8581101561e3d9578484038952815161e3ba858261e2bb565b945061e3c58361f1bb565b925060208a0199505060018101905061e3a1565b50829750879550505050505092915050565b600061e3f68261f18d565b61e400818561f1f7565b93508360208202850161e4128561f152565b8060005b8581101561e44e578484038952815161e42f858261e2cf565b945061e43a8361f1c8565b925060208a0199505060018101905061e416565b50829750879550505050505092915050565b61e47161e46c8261f370565b61f4dd565b82525050565b600061e4828261f198565b61e48c818561f208565b935061e49c81856020860161f3d4565b61e4a58161f574565b840191505092915050565b600061e4bb8261f198565b61e4c5818561f219565b935061e4d581856020860161f3d4565b61e4de8161f574565b840191505092915050565b600061e4f48261f198565b61e4fe818561f22a565b935061e50e81856020860161f3d4565b80840191505092915050565b6000815461e5278161f431565b61e531818661f22a565b9450600182166000811461e54c576001811461e55d5761e590565b60ff1983168652818601935061e590565b61e5668561f162565b60005b8381101561e5885781548189015260018201915060208101905061e569565b838801955050505b50505092915050565b600061e5a48261f1a3565b61e5ae818561f235565b935061e5be81856020860161f3d4565b61e5c78161f574565b840191505092915050565b600061e5dd8261f1a3565b61e5e7818561f257565b935061e5f781856020860161f3d4565b80840191505092915050565b600061e61060108361f246565b915061e61b8261f585565b602082019050919050565b600061e63360118361f246565b915061e63e8261f5ae565b602082019050919050565b600061e656600d8361f246565b915061e6618261f5d7565b602082019050919050565b600061e679600d8361f246565b915061e6848261f600565b602082019050919050565b600061e69c600c8361f246565b915061e6a78261f629565b602082019050919050565b600061e6bf602e8361f246565b915061e6ca8261f652565b604082019050919050565b600061e6e2600e8361f246565b915061e6ed8261f6a1565b602082019050919050565b600061e705600c8361f246565b915061e7108261f6ca565b602082019050919050565b600061e72860158361f246565b915061e7338261f6f3565b602082019050919050565b600061e74b600a8361f246565b915061e7568261f71c565b602082019050919050565b600061e76e60098361f246565b915061e7798261f745565b602082019050919050565b600061e791600e8361f246565b915061e79c8261f76e565b602082019050919050565b600061e7b460098361f246565b915061e7bf8261f797565b602082019050919050565b600061e7d760178361f246565b915061e7e28261f7c0565b602082019050919050565b600061e7fa600c8361f246565b915061e8058261f7e9565b602082019050919050565b600060a08301600083015161e828600086018261e2e3565b506020830151848203602086015261e840828261e477565b9150506040830151848203604086015261e85a828261e477565b915050606083015161e86f606086018261ebf4565b506080830151848203608086015261e887828261e599565b9150508091505092915050565b600060a08301600083015161e8ac600086018261e2e3565b506020830151848203602086015261e8c4828261e477565b9150506040830151848203604086015261e8de828261e477565b915050606083015161e8f3606086018261ebf4565b506080830151848203608086015261e90b828261e599565b9150508091505092915050565b600060a083016000830151848203600086015261e935828261e477565b915050602083015161e94a602086018261e2e3565b506040830151848203604086015261e962828261e477565b915050606083015161e977606086018261ebe5565b50608083015161e98a608086018261ebf4565b508091505092915050565b60006101008301600083015161e9ae600086018261ebf4565b506020830151848203602086015261e9c6828261e477565b9150506040830151848203604086015261e9e0828261e477565b9150506060830151848203606086015261e9fa828261e477565b915050608083015161ea0f608086018261e2e3565b5060a083015184820360a086015261ea27828261e477565b91505060c083015161ea3c60c086018261ebe5565b5060e083015161ea4f60e086018261ebf4565b508091505092915050565b60006101008301600083015161ea73600086018261ebf4565b506020830151848203602086015261ea8b828261e477565b9150506040830151848203604086015261eaa5828261e477565b9150506060830151848203606086015261eabf828261e477565b915050608083015161ead4608086018261e2e3565b5060a083015184820360a086015261eaec828261e477565b91505060c083015161eb0160c086018261ebe5565b5060e083015161eb1460e086018261ebf4565b508091505092915050565b6000608083016000830151848203600086015261eb3c828261e599565b915050602083015161eb51602086018261e2e3565b50604083015161eb64604086018261ec12565b50606083015161eb77606086018261ebf4565b508091505092915050565b6000608083016000830151848203600086015261eb9f828261e599565b915050602083015161ebb4602086018261e2e3565b50604083015161ebc7604086018261ec12565b50606083015161ebda606086018261ebf4565b508091505092915050565b61ebee8161f39a565b82525050565b61ebfd8161f3a4565b82525050565b61ec0c8161f3a4565b82525050565b61ec1b8161f3b8565b82525050565b600061ec2d828461e460565b60208201915081905092915050565b600061ec48828461e4e9565b915081905092915050565b600061ec5f828461e51a565b915081905092915050565b600061ec76828461e5d2565b915081905092915050565b600060208201905061ec96600083018461e2f2565b92915050565b600060608201905061ecb1600083018661e2f2565b61ecbe602083018561e2f2565b818103604083015261ecd0818461e4b0565b9050949350505050565b600060408201905061ecef600083018561e2f2565b818103602083015261ed01818461e4b0565b90509392505050565b600060608201905061ed1f600083018661e2f2565b818103602083015261ed31818561e4b0565b9050818103604083015261ed45818461ea5a565b9050949350505050565b6000602082019050818103600083015261ed69818461e301565b905092915050565b6000602082019050818103600083015261ed8b818461e376565b905092915050565b6000602082019050818103600083015261edad818461e3eb565b905092915050565b6000602082019050818103600083015261edcf818461e4b0565b905092915050565b6000608082019050818103600083015261edf1818761e4b0565b9050818103602083015261ee05818661e4b0565b905061ee14604083018561e2f2565b61ee21606083018461ec03565b95945050505050565b6000602082019050818103600083015261ee438161e603565b9050919050565b6000602082019050818103600083015261ee638161e626565b9050919050565b6000602082019050818103600083015261ee838161e649565b9050919050565b6000602082019050818103600083015261eea38161e66c565b9050919050565b6000602082019050818103600083015261eec38161e68f565b9050919050565b6000602082019050818103600083015261eee38161e6b2565b9050919050565b6000604082019050818103600083015261ef038161e6f8565b9050818103602083015261ef168161e71b565b9050919050565b6000604082019050818103600083015261ef368161e73e565b9050818103602083015261ef498161e6d5565b9050919050565b6000604082019050818103600083015261ef698161e73e565b9050818103602083015261ef7c8161e7a7565b9050919050565b6000602082019050818103600083015261ef9c8161e761565b9050919050565b6000602082019050818103600083015261efbc8161e784565b9050919050565b6000602082019050818103600083015261efdc8161e7a7565b9050919050565b6000602082019050818103600083015261effc8161e7ca565b9050919050565b6000602082019050818103600083015261f01c8161e7ed565b9050919050565b6000602082019050818103600083015261f03d818461e894565b905092915050565b6000602082019050818103600083015261f05f818461e918565b905092915050565b6000602082019050818103600083015261f081818461ea5a565b905092915050565b6000602082019050818103600083015261f0a3818461eb82565b905092915050565b600061f0b561f0c6565b905061f0c1828261f463565b919050565b6000604051905090565b600067ffffffffffffffff82111561f0eb5761f0ea61f545565b5b61f0f48261f574565b9050602081019050919050565b600067ffffffffffffffff82111561f11c5761f11b61f545565b5b61f1258261f574565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061f26d8261f39a565b915061f2788361f39a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561f2ad5761f2ac61f4e7565b5b828201905092915050565b600061f2c38261f3a4565b915061f2ce8361f3a4565b92508267ffffffffffffffff0382111561f2eb5761f2ea61f4e7565b5b828201905092915050565b600061f3018261f39a565b915061f30c8361f39a565b92508282101561f31f5761f31e61f4e7565b5b828203905092915050565b600061f3358261f3a4565b915061f3408361f3a4565b92508282101561f3535761f35261f4e7565b5b828203905092915050565b600061f3698261f37a565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561f3f257808201518184015260208101905061f3d7565b8381111561f401576000848401525b50505050565b600061f4128261f39a565b9150600082141561f4265761f42561f4e7565b5b600182039050919050565b6000600282049050600182168061f44957607f821691505b6020821081141561f45d5761f45c61f516565b5b50919050565b61f46c8261f574565b810181811067ffffffffffffffff8211171561f48b5761f48a61f545565b5b80604052505050565b600061f49f8261f39a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561f4d25761f4d161f4e7565b5b600182019050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f6465706f736974206d757374203e203000000000000000000000000000000000600082015250565b7f6970206d757374206e6f7420656d707479000000000000000000000000000000600082015250565b7f706f7274206d757374203e203000000000000000000000000000000000000000600082015250565b7f6e6f7420636f6e73656e73757300000000000000000000000000000000000000600082015250565b7f6e6f742072656769737465720000000000000000000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f6e616d65206e6f74206578697374000000000000000000000000000000000000600082015250565b7f52656769737465724e616d650000000000000000000000000000000000000000600082015250565b7f6e616d65206c656e677468206d757374203e3d20340000000000000000000000600082015250565b7f5570646174654e616d6500000000000000000000000000000000000000000000600082015250565b7f6e6f742061646d696e0000000000000000000000000000000000000000000000600082015250565b7f696e646578206d757374203e2030000000000000000000000000000000000000600082015250565b7f6e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b7f6e6f7420656e6f75676820696e6974206465706f736974000000000000000000600082015250565b7f706f73206d757374203e20300000000000000000000000000000000000000000600082015250565b61f81b8161f35e565b811461f82657600080fd5b50565b61f8328161f370565b811461f83d57600080fd5b50565b61f8498161f3a4565b811461f85457600080fd5b5056fea2646970667358221220b0dc7eafd341a8622132d0ffa08bf62944b95a599946e0c99cc7ae3d6937a94264736f6c63430008040033",
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

// C0x95c39a6f is a free data retrieval call binding the contract method 0x084cbe86.
//
// Solidity: function c_0x95c39a6f(bytes32 c__0x95c39a6f) pure returns()
func (_Store *StoreCaller) C0x95c39a6f(opts *bind.CallOpts, c__0x95c39a6f [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x95c39a6f", c__0x95c39a6f)

	if err != nil {
		return err
	}

	return err

}

// C0x95c39a6f is a free data retrieval call binding the contract method 0x084cbe86.
//
// Solidity: function c_0x95c39a6f(bytes32 c__0x95c39a6f) pure returns()
func (_Store *StoreSession) C0x95c39a6f(c__0x95c39a6f [32]byte) error {
	return _Store.Contract.C0x95c39a6f(&_Store.CallOpts, c__0x95c39a6f)
}

// C0x95c39a6f is a free data retrieval call binding the contract method 0x084cbe86.
//
// Solidity: function c_0x95c39a6f(bytes32 c__0x95c39a6f) pure returns()
func (_Store *StoreCallerSession) C0x95c39a6f(c__0x95c39a6f [32]byte) error {
	return _Store.Contract.C0x95c39a6f(&_Store.CallOpts, c__0x95c39a6f)
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
