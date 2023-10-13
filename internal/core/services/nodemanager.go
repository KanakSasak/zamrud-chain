package services

type NodeManager struct {
	peerNodes []string
}

func NewNodeManager() *NodeManager {
	return &NodeManager{
		peerNodes: []string{}, // initialize with known peers or keep it empty
	}
}

func (nm *NodeManager) AddPeer(peer string) {
	nm.peerNodes = append(nm.peerNodes, peer)
}

func (nm *NodeManager) GetPeers() []string {
	return nm.peerNodes
}
