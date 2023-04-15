// SPDX-License-Identifier: GPL-3.0
pragma solidity >= 0.8.2 <0.9.0;


contract subscriber_Functions {
    
    address owner;
    
    constructor() {
        owner = msg.sender;
    }
    
    modifier OwnerOnly() {
        require(msg.sender == owner, "Only Owner Can Call This Function.");
        _;
    }

    uint public event_subscribe_limit = 0 ;

    event subscriber_limit_set(uint limit) ;
    event subscriber_created(string name, address subscriber_id) ;
    event subscriber_removed(address subscriber_id) ;
    event subscribed_to_event(address subscriber_id, uint event_stream_id) ;
    event unsubscribed_to_event(address subscriber_id, uint event_stream_id) ;

// setting up limit of event stream subscription
    function set_limit(uint limit) public OwnerOnly {
        event_subscribe_limit = limit ; 
        emit subscriber_limit_set(limit) ;
    }

    struct subscriber {
        string name;
        address subscriber_id ;
        bool exist;
        uint[] event_streams_subscribed ; // stores address of events streams subscribed
    }

    struct subscriber_stack{
        int last_seen_pos;
        bool subscription;
    }
    
    struct stream_collection{
        mapping(uint => subscriber_stack)  lastseen_stack;
        bool _isDeleted;
    }
    mapping (address => subscriber) public subscriber_list;
    mapping(address => mapping(uint => subscriber_stack)) subscription_map;
    // mapping(address =>  stream_collection) subscription_map ; // maps subscriber -> List[ his subscriptions ]
    //List[ subscriptions ] are stored as a mapping from event_stream_id -> lastseen_stack for each eventstream 

    
    function create_subscriber(string memory _name, address _subscriber_id) public OwnerOnly returns(subscriber memory) {
    	subscriber memory s = subscriber(_name, _subscriber_id, true, new uint[](0));
        subscriber_list[_subscriber_id] = s;
        emit subscriber_created(s.name, s.subscriber_id) ;
	    return s;
    }


    function delete_subscriber(address sub_id) public OwnerOnly {
        require(subscriber_list[sub_id].exist == true, "publisher does not exist");
        require(subscriber_list[sub_id].subscriber_id == msg.sender, "you are not allowed to delete this publisher");
        subscriber_list[sub_id].exist = false;
        subscriber_list[sub_id].name = "";
        subscriber_list[sub_id].subscriber_id = address(0);    // removing the address of subscriber
        subscriber_list[sub_id].event_streams_subscribed = new uint[](0); // clearing information of events_streams_subscribed
        emit subscriber_removed(subscriber_list[sub_id].subscriber_id) ;
    }

    function subscribe_to_eventstream(uint stream_id, address sub_id) public {
        require(subscriber_list[sub_id].exist == true, "Subscriber does not exist");
        require(subscriber_list[sub_id].subscriber_id == msg.sender, "you are not allowed to delete this publisher");
        
        // Issue left to resolve- Check whetehr the stream_id is valid

        uint count_subscribed = subscriber_list[sub_id].event_streams_subscribed.length;
        require(count_subscribed == event_subscribe_limit, "Subscription limit reached");
        subscriber_list[sub_id].event_streams_subscribed[count_subscribed + 1] = stream_id;

        subscription_map[sub_id][stream_id].last_seen_pos = 0;
        subscription_map[sub_id][stream_id].subscription = true;

        emit subscribed_to_event(sub_id, stream_id) ;

    }

    function unsubscribe_to_eventstream(uint stream_id, address sub_id) public OwnerOnly {
        require(subscriber_list[sub_id].exist == true, "Subscriber does not exist");
        require(subscriber_list[sub_id].subscriber_id == msg.sender, "you are not allowed to delete this publisher");

        subscription_map[sub_id][stream_id].last_seen_pos = -1;
        subscription_map[sub_id][stream_id].subscription = false;

        uint i =0 ;
        while(subscriber_list[sub_id].event_streams_subscribed[i] != stream_id){
            i++;
        }
        delete subscriber_list[sub_id].event_streams_subscribed[i] ;
        emit unsubscribed_to_event(subscriber_list[sub_id].subscriber_id, stream_id) ;
    }

    function get_subscriber(address sub_id) public view returns(string memory, address , uint[] memory){
        require(subscriber_list[sub_id].exist == true, "subscriber does not exist");
        return (subscriber_list[sub_id].name, subscriber_list[sub_id].subscriber_id, subscriber_list[sub_id].event_streams_subscribed);
    }


}
