package main

import "Kademlia"

/* In this file, you should implement function "NewNode" and
 * a struct which implements the interface "dhtNode".
 */

func NewNode(port int) dhtNode {
	// Todo: create a node and then return it.
	var tempNode Kademlia.KademliaServer
	tempNode.SetPort(port)
	return &tempNode
}

// Todo: implement a struct which implements the interface "dhtNode".
