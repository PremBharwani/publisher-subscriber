package main
/*
var accounts;
web3.eth.getAccounts(function(err,res){accounts=res;});

Events.deployed().then(function (i){eve=i})
Pub.deployed().then(function (i){pub=i})
Sub.deployed().then(function (i){sub=i})

eve.add_topic()

pub.create_publisher('0xac8084232d2f5f459F480d903ac9315B011371a8')
pub.add_publisher(1, '0xac8084232d2f5f459F480d903ac9315B011371a8')
pub.publish_to_eventstream("nikhil", 1, '0xac8084232d2f5f459F480d903ac9315B011371a8')

sub.create_subscriber('0x77Aca80c0510685F6C9d6De87aF315c1817e8bC3')
sub.subscribe_to_event(1, '0x77Aca80c0510685F6C9d6De87aF315c1817e8bC3')
sub.call_for_events(1, '0x77Aca80c0510685F6C9d6De87aF315c1817e8bC3')

truffle(ganache)> [
  '0xac8084232d2f5f459F480d903ac9315B011371a8',
  '0x77Aca80c0510685F6C9d6De87aF315c1817e8bC3',
  '0x19FF33D1Bb54C868c4efCC555d69Da07218fB9Dd',
  '0xFcadf6Ce8a3e5461d3D9702574B68273c299251E',
  '0xB6b0Db2ff1e44936EeBB9F33369F8678fF80DD46',
  '0xD81BA89f99Bb36FbeaC0Db40453b53013f93D06A',
  '0xc8dD61E2796A11319C1cC9DDC245004906f8c690',
  '0x550E6aC95B95F57E8c9678E7e3e57C051201243b',
  '0x2F2c1C9DeC3547941c845c75dCF7139C6c74542A',
  '0x19633C11AE598c9EA10024f45321Aa06c7a31517'
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
	
	pubAddress := "0x5751B874ed01344b9b87d03a800298502a41bf35"
	subAddress := "0x01fDC4F2E5583E700f962FeCf8bdf85973993411"
	eventAddress := "0x4222843246A77FaBa6a963B524076849BE78BddA"
	
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