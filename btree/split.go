package btree

import "fmt"

// splitNode
// parent: 父节点
// beSplitNode: 当前分裂的节点
// i: 当前节点在父节点的位置
func (b *BPTree) splitNode(parent Position, beSplitNode Position, i int) Position {
	fmt.Printf("start split: parent: %v currentNode: %v key: %d\n", parent, beSplitNode, i)

	newNode := mallocNewNode(beSplitNode.isLeaf) // 新分裂出来的叶子节点

	index1 := 0
	index2 := beSplitNode.keyNum / 2
	for index2 < beSplitNode.keyNum {
		if beSplitNode.isLeaf == false {
			newNode.children[index1] = beSplitNode.children[index2]
			beSplitNode.children[index2] = nil
		} else {
			newNode.leafNode.data[index1] = beSplitNode.leafNode.data[index2]
			beSplitNode.leafNode.data[index2] = nil
		}
		newNode.key[index1] = beSplitNode.key[index2]
		beSplitNode.key[index2] = -1
		newNode.keyNum++
		index2++
		index1++
	}
	beSplitNode.keyNum = beSplitNode.keyNum - newNode.keyNum

	if parent != nil {
		b.insertNode(parent, newNode, i+1)
		// parent > limit 时的递归split recurvie中实现
	} else {
		// 如果是X是根，那么创建新的根并返回
		parent = mallocNewNode(false)
		b.insertNode(parent, beSplitNode, 0)
		b.insertNode(parent, newNode, 1)
		b.root = parent
		return parent
	}

	return beSplitNode
}

// MergeNode 合并节点,X少于M/2关键字，S有大于或等于M/2个关键字
func (b *BPTree) mergeNode(Parent Position, X Position, S Position, i int) Position {
	var Limit int

	// S的关键字数目大于M/2
	if S.keyNum > LimitM2 {
		// 从S中移动一个元素到X中
		b.moveElement(S, X, Parent, i, 1)
	} else {
		// 将X全部元素移动到S中，并把X删除
		Limit = X.keyNum
		b.moveElement(X, S, Parent, i, Limit) //最多时S恰好MAX MoveElement已考虑了parent.key的索引更新
		b.removeElement(false, Parent, X, i, IntMin)
	}
	return Parent
}

// findSibling  寻找一个兄弟节点，其存储的关键字未满，若左右都满返回nil
func findSibling(parent Position, i int) (sibling Position) {
	upperLimit := M
	sibling = nil
	if i == 0 {
		if parent.children[1].keyNum < upperLimit {
			sibling = parent.children[1]
		}
	} else if parent.children[i-1].keyNum < upperLimit {
		sibling = parent.children[i-1]
	} else if i+1 < parent.keyNum && parent.children[i+1].keyNum < upperLimit {
		sibling = parent.children[i+1]
	}
	return sibling
}

// FindSiblingKeyNumM2 查找兄弟节点，其关键字数大于M/2 ;没有返回nil j用来标识是左兄还是右兄
func findSiblingKeyNumM2(Parent Position, i int, j *int) Position {
	var lowerLimit int
	var Sibling Position
	Sibling = nil

	lowerLimit = LimitM2

	if i == 0 {
		if Parent.children[1].keyNum > lowerLimit {
			Sibling = Parent.children[1]
			*j = 1
		}
	} else {
		if Parent.children[i-1].keyNum > lowerLimit {
			Sibling = Parent.children[i-1]
			*j = i - 1
		} else if i+1 < Parent.keyNum && Parent.children[i+1].keyNum > lowerLimit {
			Sibling = Parent.children[i+1]
			*j = i + 1
		}

	}
	return Sibling
}

func findMostLeft(P Position) Position {
	var Tmp Position
	Tmp = P
	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.children[0].isLeaf == true {
		return Tmp.children[0]
	} else {
		for Tmp != nil && Tmp.children[0].isLeaf != true {
			Tmp = Tmp.children[0]
		}
	}
	return Tmp.children[0]
}

func findMostRight(P Position) Position {
	var Tmp Position
	Tmp = P

	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.children[Tmp.keyNum-1].isLeaf == true {
		return Tmp.children[Tmp.keyNum-1]
	} else {
		for Tmp != nil && Tmp.children[Tmp.keyNum-1].isLeaf != true {
			Tmp = Tmp.children[Tmp.keyNum-1]
		}
	}

	return Tmp.children[Tmp.keyNum-1]
}
