// `auth_helper.go` helps us to authenticate that the user request to perform any events (Publish, subscribe, create, delete, etc.) is allowable
package main

import "fmt"

var accessListMap map[string]UserAcess //! Note the value of this will be nil initially. This is a map to map userWalletAddress --> Permission list(s) for the user. 
var uniqEventQueueMap map[string]bool //! Note the value of this will be nil initially. This is a map to map eventQueueId --> true/false

type UserAcess struct {
	// XYZAccessList is a set that contains the list of the event queues the user can XYZ into
	//! Implement the REMOVING access methods very carefully after reviewing the verify & create methods!
	PublishAccessList map[string]bool `json:"publish_access_list"`
	SubscribeAccessList map[string]bool `json:"subscribe_access_list"`
	AdminAccessList map[string]bool `json:"admin_access_list"`
}

func initializeLists(){
	fmt.Println("Initializing the access lists")
	accessListMap = make(map[string]UserAcess) // Initializing the list 
	uniqEventQueueMap = make(map[string]bool) // Initializing the list
}

func get_user_access_list() map[string]UserAcess{
	return accessListMap
}

func get_queue_access_list() map[string]bool{
	return uniqEventQueueMap
}


func verify_user(userWalletAddress string, eventQueueId string, action string) bool { // Method to verify if the given userWalletAddress has rights to perform the 'action'(publish/subscribe/admin) to the eventQueueId 
	uAccessObj, ok := accessListMap[userWalletAddress]
	if (ok){
		var mList map[string]bool
		switch action{
			case "publish":
				mList = uAccessObj.PublishAccessList
			case "subscribe":
				mList = uAccessObj.SubscribeAccessList
			case "admin":
				mList = uAccessObj.AdminAccessList
		
		}
		_, ok  = mList[eventQueueId]		//* NOTE: Here I am just checking if any entry exists for the 
		if(ok){
			return true
		}
	}
	return false
}

//* === Methods to add/remove queue to/from the map we store on the server ===

func add_queue_to_map(eventQueueId string) { // Method to add a queue to the access list
	_, already_exists  := uniqEventQueueMap[eventQueueId]
	if(!already_exists) {
		uniqEventQueueMap[eventQueueId] = true
	}
}

func remove_queue_from_map(eventQueueId string) { // Method to remove a queue from the access list
	//! Check if you want to also remove all the subscribers & publishers from the access list since you're deleting the queue itself.
	_, already_exists  := uniqEventQueueMap[eventQueueId]
	if(already_exists) {
		delete(uniqEventQueueMap, eventQueueId)
	}
}

//* === Methods to add/remove user to/from the map we store on the server ===

func add_user_to_map(userWalletAddress string) { // Method to add a user to the access list
	_, already_exists  := accessListMap[userWalletAddress]
	if(!already_exists) {
		var uAccessObj UserAcess
		uAccessObj.PublishAccessList = make(map[string]bool); uAccessObj.SubscribeAccessList = make(map[string]bool); uAccessObj.AdminAccessList = make(map[string]bool)
		accessListMap[userWalletAddress] = uAccessObj
	}
}
func remove_user_from_map(userWalletAddress string) {
	// Method to remove a user from the access list
	delete(accessListMap, userWalletAddress)
}

func add_access_for_user(userWalletAddress string, eventQueueId string, action string) bool {
	// Method to add access for a user to publish/subscribe to an event queue
	// TODO: Add a check to verify who can perform the operations later on based on maybe admin access?
	// checkUserPermissions() // & return false if the user is not authorized.
	
	_, already_exists  := accessListMap[userWalletAddress]
	if(!already_exists) {
		add_user_to_map(userWalletAddress)
	}

	var mList map[string]bool // The list to be mutated

	switch action{ // To choose the list to be mutated according to the action
		case "publish":
			mList = accessListMap[userWalletAddress].PublishAccessList
		case "subscribe":
			mList = accessListMap[userWalletAddress].SubscribeAccessList
		case "admin":
			mList = accessListMap[userWalletAddress].AdminAccessList
	}

	_, eventQueueExists := uniqEventQueueMap[eventQueueId]
	if(!eventQueueExists) {
		fmt.Printf("The event queue [%s] does not exist", eventQueueId)
		return false
	}

	mList[eventQueueId] = true

	return true
}


func remove_access_for_user(userWalletAddress string, eventQueueId string, action string) bool {
	// Method to remove access for a user to publish/subscribe to an event queue
	// TODO: Add a check to verify who can perform the operations later on based on maybe admin access?
	
	// checkUserPermissions() // & return false if the user is not authorized.
	
	_, already_exists  := accessListMap[userWalletAddress]
	if(!already_exists) { // Since this USER doesn't already exists we do not need to do anything.
		return true 
	}

	var mList map[string]bool // The list to be mutated

	switch action{ // To choose the list to be mutated according to the action
		case "publish":
			mList = accessListMap[userWalletAddress].PublishAccessList
		case "subscribe":
			mList = accessListMap[userWalletAddress].SubscribeAccessList
		case "admin":
			mList = accessListMap[userWalletAddress].AdminAccessList
	}

	delete(mList, eventQueueId)
	
	return true
}


