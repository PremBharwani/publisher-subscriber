// ========================================================================
//*Contains the functions that facilitate access rights for users --> Topics
//*and maintains a list of unique event queues/topics
// ========================================================================
package main

import (
	"fmt"
)

// ================================================================
// To store (1)access list for users (2)List of event queues/topics
// ================================================================
var accessListMap map[string]UserAcess //* userWalletAddress --> Permission list(s) for the user. 
var uniqEventQueueMap map[string]bool  //* eventQueueId --> true/false

type UserAcess struct { //*Struct to store info about the topics where user has publish, subscribe & admin rights
	// XYZAccessList is a set that contains the list of the event queues the user can XYZ into
	//! Implement the REMOVING access methods very carefully after reviewing the verify & create methods!
	PublishAccessList map[string]bool `json:"publish_access_list"`
	SubscribeAccessList map[string]bool `json:"subscribe_access_list"`
	AdminAccessList map[string]bool `json:"admin_access_list"`
}

func initializeLists(){ // Method to initialize the lists
	fmt.Println("Initializing the access lists")
	accessListMap = make(map[string]UserAcess) // Initializing the list 
	uniqEventQueueMap = make(map[string]bool) // Initializing the list
}

//* ==========================================================================
//* === Methods to add/remove queue to/from the map we store on the server ===
//* ==========================================================================
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
//* =========================================================================
//* === Methods to add/remove user to/from the map we store on the server ===
//* =========================================================================
func verify_user_access_to_topic(userWalletAddress string, eventQueueId string, action string) bool { //*Method to verify if the given userWalletAddress has rights to perform the 'action'(publish/subscribe/admin) to the eventQueueId 
	uAccessObj, found := accessListMap[userWalletAddress]
	if (found){
		var mList map[string]bool
		switch action{
			case "publish":
				mList = uAccessObj.PublishAccessList
			case "subscribe":
				mList = uAccessObj.SubscribeAccessList
			case "admin":
				mList = uAccessObj.AdminAccessList
		
		}
		_, found  = mList[eventQueueId]		//* NOTE: Here I am just checking if any entry exists for the 
		if(found){
			return true
		}
	}
	return false
}

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
	
	add_user_to_map(userWalletAddress)

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


