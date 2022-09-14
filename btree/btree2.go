package btree

func New() *BPTree {
	root := mallocNewNode(true)
	return &BPTree{
		keyMax: 0,
		width:  0,
		root:   root,
		ptr:    root,
	}
}

// Insert 往表里面写数据  key=主键id  data=表数据
func (b *BPTree) Insert(id int, data interface{}) (Position, error) {
	return b.recursiveInsert(b.root, id, 0, nil, data)
}

// Delete 删除数据，目前支持主键删除
func (b *BPTree) Delete(id int) (Position, error) {
	return b.recursiveDelete(b.root, id, 0, nil)
}

// Find 查找数据，目前支持主键查找
func (b *BPTree) Find(id int) (interface{}, error) {
	var currentNode *BPFullNode // 当前页
	var index int               // 当前页目录的索引
	currentNode = b.root

	if b.keyMax < id {
		return nil, NotExistErr
	}

	// 这里是直接遍历查找，实际上MySQL这里进行的是二分查找
	for index < currentNode.keyNum {
		index = 0 // 每次翻页都将索引重制为0
		// 遍历页目录，直到找到大于等于id的目录值
		for index < currentNode.keyNum && id >= currentNode.key[index] {
			index++
		}
		if index == 0 {
			return nil, NotExistErr
		}
		index--
		if currentNode.isLeaf == false { // 不是叶子节点,继续往下一层级找
			currentNode = currentNode.children[index]
		} else {
			if id == currentNode.key[index] {
				return currentNode.leafNode.data[index], nil
			} else {
				return nil, NotExistErr
			}
		}
	}
	return nil, NotExistErr
}
