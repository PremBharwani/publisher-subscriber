// SPDX-License-Identifier: GPL-3.0

pragma solidity >= 0.8.2 <0.9.0;


contract Sub {
    
    address owner;
    mapping(address => bool) relay_eventsCalled;
    mapping(address => string[50]) private ret_events;
    mapping(address => uint8) private filled_till; 
    
    constructor() {
        owner = msg.sender;
        address test_addr = 0xc0ffee254729296a45a3885639AC7E10F9d54979;
        relay_eventsCalled[test_addr]=false;
        filled_till[test_addr]=0;
        string[50] memory m; 
        ret_events[test_addr]=m;

    }
    
    modifier OwnerOnly() {
        require(msg.sender == owner, "Only Owner Can Call This Function.");
        _;
    }

    struct subscriber{
        bool exist;
        uint256[] access;
    }

    mapping (address => subscriber) public subscriber_list;


    uint256 public event_subscribe_limit = 50 ; // default 50

    event subscriber_created(address subscriber_id) ;
    event subscriber_removed(address subscriber_id) ;
    event subscribed_to_event(address subscriber_id , uint256 event_stream_id) ;
    event unsubscribed_to_event(address subscriber_id , uint256 event_stream_id) ;
    event requested_for_events(address subscriber_id, uint256 event_stream_id);

    function set_limit(uint256 limit) public OwnerOnly {
        event_subscribe_limit = limit ; 
        // emit subscriber_limit_set(limit) ;
    }
    
    function create_subscriber(address _address_subscriber) public {
        
        require(subscriber_list[_address_subscriber].exist == false, "subscriber already exist");
        subscriber_list[_address_subscriber] = subscriber(true, new uint256[](0));

        emit subscriber_created(_address_subscriber) ;

    }

    function delete_subscriber(address _id) public {
        require(subscriber_list[_id].exist == true, "subscriber does not exist");
        // require(subscriber[_id].address_publisher == msg.sender, "you are not allowed to delete this publisher");
        subscriber_list[_id].exist = false;
        subscriber_list[_id].access = new uint256[](0);
        emit subscriber_removed(_id) ;
    }


    function subscribe_to_event(uint256 stream_id, address sub_id) public {
        require(subscriber_list[sub_id].exist == true, "subscriber does not exist");
        // require(subscriber_list[sub_id].address_publisher == msg.sender, "you are not allowed to add this publisher");
        bool check= false;
        for(uint256 i = 0; i < subscriber_list[sub_id].access.length; i++){
            if(subscriber_list[sub_id].access[i] == stream_id){
                check = true;
            }
        }
        require(check == false, "subscriber already has access to this event");
        require(subscriber_list[sub_id].access.length<event_subscribe_limit,"Subscription limit reached");
        subscriber_list[sub_id].access.push(stream_id);
        emit subscribed_to_event(sub_id, stream_id);
    }

    function unsubscribe_to_event(uint256 stream_id, address sub_id) public OwnerOnly {
        require(subscriber_list[sub_id].exist == true, "subscriber does not exist");
        // require(subscriber_list[sub_id].address_publisher == msg.sender, "you are not allowed to remove this publisher");
        bool check= false;
        for(uint256 i = 0; i < subscriber_list[sub_id].access.length; i++){
            if(subscriber_list[sub_id].access[i] == stream_id){
                check = true;
            }
        }
        require(check == true, "subscriber does not have access to this event");
        for(uint256 i = 0; i < subscriber_list[sub_id].access.length; i++){
            if(subscriber_list[sub_id].access[i] == stream_id){
                subscriber_list[sub_id].access[i] = 0;
            }
        }
    }

    function get_subscriber(address sub_id) public view returns( uint[] memory){
        require(subscriber_list[sub_id].exist == true, "subscriber does not exist" );
        return (subscriber_list[sub_id].access);
    }


    struct events_data {
        string[50] events;
        uint8 last_index;
    }

    bool public relay_events_called = false ;
    string public checkEvent;

    function call_for_events(uint stream_id, address sub_id) public {
        
        relay_eventsCalled[sub_id]=false;
        filled_till[sub_id]=0;
        string[50] memory m; 
        ret_events[sub_id]=m;

        // emit saying the sub needs the events in the particular stream
        emit requested_for_events(sub_id, stream_id);

        // wait until relay_events() is called

        // while(!relay_eventsCalled[sub_id]){
        //     //wait
        // }

        // // relay_eventsCalled[sub_id]=false;
        // events_data memory ev;
        // ev.events=ret_events[sub_id];
        // ev.last_index=filled_till[sub_id];

        // //delete the flag

        // delete relay_eventsCalled[sub_id];
        // delete filled_till[sub_id];
        // delete ret_events[sub_id];

        // return ev;

    }

    function get_back_events(address sub_id) public returns (events_data memory){

         while(!relay_eventsCalled[sub_id]){
            //wait
        }

        // relay_eventsCalled[sub_id]=false;
        events_data memory ev;
        ev.events=ret_events[sub_id];
        ev.last_index=filled_till[sub_id];

        //delete the flag

        delete relay_eventsCalled[sub_id];
        delete filled_till[sub_id];
        delete ret_events[sub_id];

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
        relay_events_called=true ;
        checkEvent = events[0];
        relay_eventsCalled[sub_id]=true;
        return;

    }
}