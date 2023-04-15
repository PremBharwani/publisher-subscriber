package main
/*
var accounts;
web3.eth.getAccounts(function(err,res){accounts=res;});

Events.deployed().then(function (i){eve=i})
Pub.deployed().then(function (i){pub=i})
Sub.deployed().then(function (i){sub=i})

eve.add_topic()

pub.create_publisher('0xBb8E23BFCde7A2911efcBC90990aD52f4971A51E')
pub.add_publisher(1, '0xBb8E23BFCde7A2911efcBC90990aD52f4971A51E')
pub.publish_to_eventstream("test", 1, '0xBb8E23BFCde7A2911efcBC90990aD52f4971A51E')

sub.create_subscriber('0x03A63aBaFA97efA7FF5e6B224d9B4f9F40755FA6')
sub.subscribe_to_event(1, '0x03A63aBaFA97efA7FF5e6B224d9B4f9F40755FA6')
sub.call_for_events(1, '0x03A63aBaFA97efA7FF5e6B224d9B4f9F40755FA6')

truffle(ganache)> [
	'0xBb8E23BFCde7A2911efcBC90990aD52f4971A51E',
  '0x03A63aBaFA97efA7FF5e6B224d9B4f9F40755FA6',
  '0x2F0f852e20746d320FA5565e4671BA14B92fC5D3',
  '0xC3a83B8209b412092ac6bd692fF073d3420D5C0c',
  '0x30267a1AA03B582ce13987b437850ba8fb1D5d1a',
  '0x9e5613f6d06489Ab2De0F032684CaFd5f7322454',
  '0xeD68A593E74c1eC0cb5f7D69407f70a22168e220',
  '0x5796b22e321F9cB70593B2039D9e69297c2a19e2',
  '0x3dd9C729C46C264919cE05ceaCeB16D76383EeC1',
  '0xEE6950a87dA456D9E86dcc2aFa674F1dce09C133'

]


*/
import (
	"context"
	"log"
	"fmt"

	// "math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


func main() {
	
	pubAddress := "0x3E5005D7A64877c6bA492b69B93e8adA27A32FCA"
	subAddress := "0x5442d9D53fD76D11199cF4e16F8c5Ccc9F63357D"
	eventAddress := "0x4FB614DCC30509346185C96Db3ee9e1d1e5f93a9"
	
	// LastBlockId := big.NewInt(0)

	// Set the LastBlockId to the latest block
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	latestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	LastBlockId := latestBlock.Number()

	for {

		client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
		if err != nil {
			log.Fatal(err)
		}
		query := ethereum.FilterQuery{
			Addresses: []common.Address{
				common.HexToAddress(pubAddress),
				common.HexToAddress(subAddress),
				common.HexToAddress(eventAddress),
			},
			FromBlock: LastBlockId,
		}
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Printf("Error In Filtering Logs \n")
		}
		


		LastBlockId = listen_pub_logs(logs, LastBlockId) 
		LastBlockId = listen_event_logs(logs, LastBlockId) 
		LastBlockId = listen_sub_logs(logs, LastBlockId)


	}
	


}