pragma solidity >= 0.8.2 <0.9.0;


contract Subscriber_Functions {
    
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
    event subscriber_created(address subscriber_id) ;
    event subscriber_removed(address subscriber_id) ;
    event subscribed_to_event(address subscriber_id , address event_stream_id) ;
    event unsubscribed_to_event(address subscriber_id , address event_stream_id) ;

// setting up limit of event stream subscription
    function set_limit(uint limit) public OwnerOnly {
        event_subscribe_limit = limit ; 
        emit subscriber_limit_set(limit) ;
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

    
    function create_subscriber(address subscriber_id) public OwnerOnly returns(Subscriber memory) {
    	Subscriber memory s = Subscriber(subscriber_id, new address[](0));
        emit subscriber_created(s.subscriber_id) ;
	    return s;
    }


    function delete_subscriber(Subscriber memory s) public OwnerOnly {
        uint len = s.event_streams_subscribed.length;
        for (uint i = 0; i < len; i++) {
            delete event_stream_iterator_map[s.subscriber_id][s.event_streams_subscribed[i]]; // clearing associated mappings first
        }
        delete s.event_streams_subscribed; // clearing information of events_streams_subscribed
        s.subscriber_id = address(0); // removing the address of subscriber
        emit subscriber_removed(s.subscriber_id) ;
    }


    function unsubscribe_to_event(address event_stream_id, Subscriber memory s) public OwnerOnly {
        require( event_stream_iterator_map[s.subscriber_id][event_stream_id].event_exists , "Subscriber Has Not Subscribed To This Event"); // checking if the event has been subscribed        uint i = 0;
        uint i =0 ;
        while(s.event_streams_subscribed[i] != event_stream_id){
            i++;
        }
        delete event_stream_iterator_map[s.subscriber_id][event_stream_id] ; // clearing the map associated with the subscriber_id and event_stream_id pair
        s.event_streams_subscribed[i] = address(0); // clearing the associated address
        emit unsubscribed_to_event(s.subscriber_id, event_stream_id) ;
    }


}
