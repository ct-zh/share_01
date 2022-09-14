package btree

func mallocNewNode(isLeaf bool) *BPFullNode {
	var newNode *BPFullNode
	if isLeaf == true {
		newLeaf := mallocNewLeaf()
		newNode = &BPFullNode{
			nodeCatalogue: nodeCatalogue{
				keyNum:   0,
				key:      make([]int, M+1), // M+1是因为可能暂时出现节点key数量大于M的情况 此时触发页分裂
				children: nil,
			},
			isLeaf:   true,
			leafNode: newLeaf,
			point: point{
				prev: nil,
				next: nil,
			},
		}
	} else {
		newNode = &BPFullNode{
			nodeCatalogue: nodeCatalogue{
				keyNum:   0,
				key:      make([]int, M+1), // M+1是因为可能暂时出现节点key数量大于M的情况 此时触发页分裂
				children: make([]*BPFullNode, M+1),
			},
			isLeaf:   false,
			leafNode: nil,
			point: point{
				prev: nil,
				next: nil,
			},
		}
	}
	for i := range newNode.key { // 切片填充值为-1
		newNode.key[i] = -1
	}

	return newNode
}

func mallocNewLeaf() *bPLeafNode {
	NewLeaf := bPLeafNode{
		Next: nil,
		data: make(map[int]interface{}, M+1),
	}
	return &NewLeaf
}
