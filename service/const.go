package service

import (
	"log"
)

const DBC_MAINNET = "https://rpc.dbcwallet.io"
const MAIN_NET_VERIFIER_URL = "https://www.dbcscan.io/api"
const XAAIAO_TOKEN_IN_CONTRACT = "0x16d83F6B17914a4e88436251589194CA5AC0f452"
const XAAIAO_NFT_HOLDER_CONTRACT = "0xc488736c09ab088e5203b48d973dca30581d6118"
const DBC_AI_PROXY = "0xa7B9f404653841227AF204a561455113F36d8EC8"
const SLASH_TO_ADDRESS = "0x0C1a2890e5575Be1a7E0C4a35c866CdF3F6D6300"
var (
	ContractEnvPath = "./contracts/.env"
	ContractPath    = "./contracts"
)

func Init(env string) {
	// if env == "prod" {
	// 	ContractEnvPath = "../contracts/.env"
	// 	ContractPath = "../contracts"
	// }
	log.Println("contract env file path: ", ContractEnvPath)
	log.Println("contract file path: ", ContractPath)

}
