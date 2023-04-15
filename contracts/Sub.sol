// SPDX-License-Identifier: GPL-3.0

pragma solidity >= 0.8.2 <0.9.0;


contract Sub {
    
    address owner;
    mapping(address => bool) relay_eventsCalled;
    mapping(address => string[50]) private ret_events;
    mapping(address => uint8) private filled_till; 
    
    constructor() {
        owner = msg.sender;
        // address test_addr = 0xc0ffee254729296a45a3885639AC7E10F9d54979;
        // relay_eventsCalled[test_addr]=false;
        // filled_till[test_addr]=0;
        // string[50] memory m; 
        // ret_events[test_addr]=m;

    }
    
    modifier OwnerOnly() {
        require(msg.sender == owner, "Only Owner Can Call This Function.");
        _;
    }
    uint public event_subscribe_limit = 0 ;
   
    event subscriber_created(address subscriber_id) ;
    event subscriber_removed(address subscriber_id) ;
    event subscribed_to_event(address subscriber_id , address event_stream_id) ;
    event unsubscribed_to_event(address subscriber_id , address event_stream_id) ;
    event requested_for_events(address subscriber_id, address event_stream_id);


    function set_limit(uint256 limit) public OwnerOnly {
        require(limit<=100, "Limit must be less than or equal to 100");
        event_subscribe_limit = limit ; 
        // emit subscriber_limit_set(limit) ;
    }

    struct Subscriber {
        address subscriber_id ;
        address[] event_streams_subscribed ; // stores address of events streams subscribed
    }

    struct Event_Stream_Iterator{
        uint position;
        bool event_exists;
    }
    
    mapping(address => mapping(address => Event_Stream_Iterator)) event_stream_iterator_map ; // maps subscriber IDs to map of IDs of events 
    // subscribed to the iterator of last seen message
    mapping(address => Subscriber) addr_to_sub;

    function create_subscriber(address subscriber_id) public OwnerOnly returns(Subscriber memory) {
    	Subscriber memory s = Subscriber(subscriber_id, new address[](0));
        addr_to_sub[subscriber_id]=s;
        relay_eventsCalled[subscriber_id]=false;
        filled_till[subscriber_id]=0;
        string[50] memory m; 
        ret_events[subscriber_id]=m;
        emit subscriber_created(s.subscriber_id) ;
	    return s;
    }

    function delete_subscriber(address add) public OwnerOnly {
        uint len = addr_to_sub[add].event_streams_subscribed.length;
        for (uint i = 0; i < len; i++) {
            delete event_stream_iterator_map[addr_to_sub[add].subscriber_id][addr_to_sub[add].event_streams_subscribed[i]]; // clearing associated mappings first
        }
        delete addr_to_sub[add].event_streams_subscribed; // clearing information of events_streams_subscribed
        addr_to_sub[add].subscriber_id = address(0); // removing the address of subscriber
        emit subscriber_removed(addr_to_sub[add].subscriber_id) ;
    }


    function subscribe_to_event(address event_stream_id, address add) public OwnerOnly{
        require( event_stream_iterator_map[addr_to_sub[add].subscriber_id][event_stream_id].event_exists==false , "Subscriber Has already Subscribed To This Event"); // chec
        require(addr_to_sub[add].event_streams_subscribed.length<event_subscribe_limit,"Subscription limit reached");
        event_stream_iterator_map[addr_to_sub[add].subscriber_id][event_stream_id] = Event_Stream_Iterator(0, true);
        addr_to_sub[add].event_streams_subscribed.push(event_stream_id);
        emit subscribed_to_event(addr_to_sub[add].subscriber_id, event_stream_id);
    }

    function unsubscribe_to_event(address event_stream_id, address add) public OwnerOnly {
        require( event_stream_iterator_map[addr_to_sub[add].subscriber_id][event_stream_id].event_exists , "Subscriber Has Not Subscribed To This Event"); // checking if the event has been subscribed        uint i = 0;
        uint i =0 ;
        while(addr_to_sub[add].event_streams_subscribed[i] != event_stream_id){
            i++;
        }
        delete event_stream_iterator_map[addr_to_sub[add].subscriber_id][event_stream_id] ; // clearing the map associated with the subscriber_id and event_stream_id pair
        addr_to_sub[add].event_streams_subscribed[i] = address(0); // clearing the associated address
        emit unsubscribed_to_event(addr_to_sub[add].subscriber_id, event_stream_id) ;
    }


    struct events_data {
        string[50] events;
        uint8 last_index;
    }

    // bool public relay_events_called = false ;
    // string public checkEvent;

    function get_events(address stream_id, address sub_id) public returns (events_data memory){
        
        // emit saying the sub needs the events in the particular stream
        emit requested_for_events(sub_id, stream_id);

        // wait until relay_events() is called

        while(!relay_eventsCalled[sub_id]){
            //wait
        }

        relay_eventsCalled[sub_id]=false;
        events_data memory ev;
        ev.events=ret_events[sub_id];
        ev.last_index=filled_till[sub_id];

        return ev;

    }



    function relay_events(string[] memory events, address sub_id) public {
        // this function will be called from the js script (which gets the  data from the go script)

        // need to return this data to the call that called get_events()

        uint8 i=0;
        for(; i<events.length; i++){
            ret_events[sub_id][i]=events[i];
        }

        filled_till[sub_id]=i-1;
        // relay_events_called=true ;
        // checkEvent = events[0];
        relay_eventsCalled[sub_id]=true;
        return;

    }
}