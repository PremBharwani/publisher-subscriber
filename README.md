# Publisher-Subscriber
Implementation of a Publisher Subscriber model using blockchain. This project was done under the course: Blockchain Technology and Applications(CS731A).
# Instruction to Run
Firstly we need the blockchain up, before we can get the other servers running.

Start `ganache`, use the `truffle-config.js` file found in the directory to setup the workspace.

Run the following command in the terminal (base working directory,i.e., publisher-subscriber)
```
truffle migrate
```
This compiles the contracts and prints their metadata, which also contains their deployed contract address.
Copy the contract addresses and paste(replace) them into `listener.go`(present in the `./scripts/` directory) at lines mentioned below:
- Pub address goes in line 51
- Sub address goes in line 52
- Event address goes in line 53

Now execute the following commands to complete the setup:

#### Server setup commands
Our implementation requires 2 servers running simultaneously. To setup run the following commands (make sure `docker` is setup on your system): 

```
cd kakfa-event-queue
docker compose up -d
go run *.go
```
If `docker compose up -d` gives permission error, run it with `sudo`. 

This sets up the `gin-server` which is used by the scripts and event-queue to communicate.

In a separate terminal, navigate to the `scripts` directory and get the scripts running:
```
cd scripts
go run *.go
```


## Instructions to run the relay-server
Navigate to the `relay-server` directory, run the commands:

```
npm install
node relay.js
```

This will start a server listening on port `3000`. After the API is called and the data is sent back to the contract, the first event is also printed on the console (just as a sanity check). 

## Testing Guidelines

Note: If you want to use any other method (say `js` scripts) to call the functions, please refer to `/contracts/README.md` for all the function specifications and calling conventions.

For testing in `truffle`:

```
truffle console
```

This connects to `ganache` chain if it is running in the background, and opens up it's own console. To test and call functions, you can use the following in the truffle console:

**Getting accounts**:
```
var accounts;
> 
web3.eth.getAccounts(function(err,res){accounts=res;});
> 
```
This displays all the account addresses from the chain truffle is connected to. These addresses can be used for testing.

Note: `>` indicates output of the respective command

**Get objects of deployed contracts**:

```
Events.deployed().then(function (i){eve=i})
>
Pub.deployed().then(function (i){pub=i})
>
Sub.deployed().then(function (i){sub=i})
>
```
These objects can be used to call the respective functions present in the smart contract.

**Create Publisher**:

```
pub.create_publisher('publisher_address')
>
```
**Create Subscriber**:

```
sub.create_subscriber('subscriber_address')
>
```

**Create new event stream**:

```
eve.add_topic()
>
```

**Add Publisher to event stream**:

```
pub.add_publisher(stream_id, 'publisher_address')
>
```
`stream_id` is the unique stream identifier that is returned when you call `eve.add_topic()`. 

**Add Subscriber to event stream**:

```
sub.subscribe_to_event(stream_id, 'subscriber_address')
>
```

**Publish a message to event stream**:

```
pub.publish_to_eventstream("msg", stream_id, 'publisher_address')
>
```

**Get events from event stream**

You need to make two sequential calls to get the events. The second call's return object contains the relevant events. You might want to refer to `/contracts/README.md` on how to unpack the events.

```
sub.call_for_events(stream_id, 'subscriber_address')
>
sub.get_back_events('subscriber_address')
>
```

