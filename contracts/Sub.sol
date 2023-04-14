pragma solidity >= 0.8.2 <0.9.0;


contract Sub {
    
    address owner;
    
    constructor() {
        owner = msg.sender;
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
    
    
    
    bool private relay_eventsCalled = false;
    // uint private max_events_at_a_time=100;
    string[100] private ret_events ; // keeping a dynamic array will increase gas usage (clearing the array). Just overwrite
    uint private filled_till;

    struct events_data {
        string[100] events;
        uint last_index;
    }

    function get_events(address stream_id, address sub_id) public returns (events_data memory){
        
        // emit saying the sub needs the events in the particular stream
        emit requested_for_events(sub_id, stream_id);

        // wait until relay_events() is called

        while(!relay_eventsCalled){
            //wait
        }

        relay_eventsCalled=false;
        events_data memory ev;
        ev.events=ret_events;
        ev.last_index=filled_till;

        return ev;

    }

    function relay_events(string[] memory events) public {
        // this function will be called from the js script (which gets the  data from the go script)

        // need to return this data to the call that called get_events()

        uint i=0;
        for(; i<events.length; i++){
            ret_events[i]=events[i];
        }

        filled_till=i-1;
        relay_eventsCalled=true;
        return;

    }
}