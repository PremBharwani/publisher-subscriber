## Callable Functions supported in `Pub.sol`

`create_publisher (address pub_id)` 

`publish_to_eventstream (string memory data, uint stream_id, address pub_id)`

`add_publisher (uint stream_id, address pub_id)` 

`remove_publisher(uint stream_id,address pub_id)`

`delete_publisher(address pub_id)`

## Callable Functions supported in `Sub.sol`

`create_subscriber(address sub_id)`

`subscribe_to_event(uint stream_id, address sub_id)`

`unsubscribe_to_event(uint stream_id, address sub_id)`

`delete_subscriber(address sub_id)`

`call_for_events(uint stream_id, address sub_id)` 

`get_back_events(address sub_id) returns (events_data memory)`

The functions `call_for_events()` and `get_back_events()` needs to be called one after other to get events from a queue. Here is how you can get events from a stream:

```
call_for_events(stream_id, sub_address)
return_object = get_back_events(sub_address)
```

Here, `return_object` is an object of a struct defined in `Sub.sol`:

```
struct events_data {
        string[50] events;
        uint8 last_index;
    }
```

The events are filled from `index=0` to `index=last_index` in that order. Currently, we have hardcoded 50 events at a time, this limit can be changed easily. 

## Callable Functions supported in `Events.sol`

`add_topic() returns (uint)`

`delete_topic(uint _topic_id)`

Get the `topic_id` / `stream_id` by calling `add_topic()` and use that to call the publisher and subscriber functions given above that require `stream_id`. 



