package p2p

// Self-note : Any type that implements an interface can then 
// be used as the interface type itself and their behaviour
// is defined by their method implementations

// Peer is an interface that represents the remote node
type Peer interface {}


// Transport is anything that handles the communication
// betweeen the nodes in the network. This can be of the 
// form (TCP , UDP , websockets , ...)
type Transport interface {}