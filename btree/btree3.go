package btree

// recursiveInsert 递归插入
// beInsertedElement 被插入的页
// id 插入id
// posAtParent
// Parent
// data 插入数据
func (b *BPTree) recursiveInsert(beInsertedElement Position, id int, posAtParent int, parent Position, data interface{}) (Position, error) {
	var insertIndex int
	var sibling Position

	// 查找分支
	insertIndex = 0
	for insertIndex < beInsertedElement.keyNum && id >= beInsertedElement.key[insertIndex] {
		if id == beInsertedElement.key[insertIndex] { // 重复值不插入
			return beInsertedElement, AlreadyExistsErr
		}
		insertIndex++
	}

	//key必须大于被插入节点的最小元素，才能插入到此节点，故需回退一步
	if insertIndex != 0 && beInsertedElement.isLeaf == false {
		insertIndex--
	}

	if beInsertedElement.isLeaf == true {
		beInsertedElement = b.insertData(parent, beInsertedElement, id, posAtParent, insertIndex, data)
	} else {
		p, err := b.recursiveInsert(beInsertedElement.children[insertIndex], id, insertIndex, beInsertedElement, data)
		if err != nil {
			return nil, err
		}
		beInsertedElement.children[insertIndex] = p
	}

	// 触发了页分裂
	if beInsertedElement.keyNum > M {
		if parent == nil {
			// 根节点直接 分裂
			beInsertedElement = b.splitNode(parent, beInsertedElement, posAtParent)
		} else {
			sibling = findSibling(parent, posAtParent)
			if sibling != nil {
				// 将T的一个元素（Key或者Child）移动的Sibing中
				b.moveElement(beInsertedElement, sibling, parent, posAtParent, 1)
			} else {
				// 分裂节点
				beInsertedElement = b.splitNode(parent, beInsertedElement, posAtParent)
			}
		}
	}

	if parent != nil {
		parent.key[posAtParent] = beInsertedElement.key[0]
	}

	return beInsertedElement, nil
}

func (b *BPTree) recursiveDelete(beRemovedElement Position, key int, posAtParent int, Parent Position) (Position, error) {
	var deleteIndex int
	var Sibling Position
	var NeedAdjust bool
	Sibling = nil

	// 查找分支   TODO查找函数可以在参考这里的代码 或者实现一个递归遍历
	deleteIndex = 0
	for deleteIndex < beRemovedElement.keyNum && key >= beRemovedElement.key[deleteIndex] {
		if key == beRemovedElement.key[deleteIndex] {
			break
		}
		deleteIndex++
	}

	if beRemovedElement.isLeaf == true {
		// 没找到
		if key != beRemovedElement.key[deleteIndex] || deleteIndex == beRemovedElement.keyNum {
			return beRemovedElement, NotExistErr
		}
	} else {
		if deleteIndex == beRemovedElement.keyNum || key < beRemovedElement.key[deleteIndex] {
			deleteIndex-- //准备到下层节点查找
		}
	}

	// 树叶
	if beRemovedElement.isLeaf == true {
		beRemovedElement = b.removeElement(true, Parent, beRemovedElement, posAtParent, deleteIndex)
	} else {
		p, err := b.recursiveDelete(beRemovedElement.children[deleteIndex], key, deleteIndex, beRemovedElement)
		if err != nil {
			return nil, err
		}
		beRemovedElement.children[deleteIndex] = p
	}

	NeedAdjust = false
	//有子节点的root节点，当keyNum小于2时
	if Parent == nil && beRemovedElement.isLeaf == false && beRemovedElement.keyNum < 2 {
		NeedAdjust = true
	} else if Parent != nil && beRemovedElement.isLeaf == false && beRemovedElement.keyNum < LimitM2 {
		// 除根外，所有中间节点的儿子数不在[M/2]到M之间时。(符号[]表示向上取整)
		NeedAdjust = true
	} else if Parent != nil && beRemovedElement.isLeaf == true && beRemovedElement.keyNum < LimitM2 {
		// （非根）树叶中关键字的个数不在[M/2]到M之间时
		NeedAdjust = true
	}

	// 调整节点
	if NeedAdjust {
		// 根
		if Parent == nil {
			if beRemovedElement.isLeaf == false && beRemovedElement.keyNum < 2 {
				//树根的更新操作 树高度减一
				beRemovedElement = beRemovedElement.children[0]
				b.root = beRemovedElement.children[0]
				return beRemovedElement, nil
			}

		} else {
			// 查找兄弟节点，其关键字数目大于M/2
			Sibling = findSiblingKeyNumM2(Parent, posAtParent, &deleteIndex)
			if Sibling != nil {
				b.moveElement(Sibling, beRemovedElement, Parent, deleteIndex, 1)
			} else {
				if posAtParent == 0 {
					Sibling = Parent.children[1]
				} else {
					Sibling = Parent.children[posAtParent-1]
				}

				Parent = b.mergeNode(Parent, beRemovedElement, Sibling, posAtParent)
				//Merge中已考虑空节点的删除
				beRemovedElement = Parent.children[posAtParent]
			}
		}

	}

	return beRemovedElement, nil
}

// 插入节点
// insertNode 当要对X插入data的时候，i是X在Parent的位置，
// 当要对Parent插入X节点的时候，posAtParent是要插入的位置
func (b *BPTree) insertNode(Parent Position, X Position, posAtParent int) Position {
	if X.isLeaf == true {
		if posAtParent > 0 {
			Parent.children[posAtParent-1].leafNode.Next = X
		}
		X.leafNode.Next = Parent.children[posAtParent]
		//更新叶子指针
		if X.key[0] <= b.ptr.key[0] {
			b.ptr = X
		}
	}

	k := Parent.keyNum - 1
	for k >= posAtParent { //插入节点时key也要对应的插入
		Parent.children[k+1] = Parent.children[k]
		Parent.key[k+1] = Parent.key[k]
		k--
	}
	Parent.key[posAtParent] = X.key[0]
	Parent.children[posAtParent] = X
	Parent.keyNum++
	return X
}

// insertData 插入data
// insertIndex data要插入的位置，j可由查找得到
func (b *BPTree) insertData(parent Position, x Position, id int, posAtParent int, insertIndex int, data interface{}) Position {
	//fmt.Printf("parent: %+v current: %+v id: %d posAtParent: %d insertIndex %d data: %+v \n",
	//	parent, x, id, posAtParent, insertIndex, data)

	k := x.keyNum - 1
	for k >= insertIndex {
		x.key[k+1] = x.key[k]
		x.leafNode.data[k+1] = x.leafNode.data[k]
		k--
	}

	x.key[insertIndex] = id
	x.leafNode.data[insertIndex] = data
	if parent != nil {
		parent.key[posAtParent] = x.key[0] //可能min_key 已发生改变
	}

	x.keyNum++

	//fmt.Printf("x: %+v \n", x)
	return x
}

func (b *BPTree) moveElement(src Position, dst Position, parent Position, posAtParent int, eNum int) Position {
	var TmpKey int
	var data interface{}
	var Child Position
	var j int
	var srcInFront bool

	srcInFront = false

	if src.key[0] < dst.key[0] {
		srcInFront = true
	}
	j = 0
	// 节点Src在Dst前面
	if srcInFront {
		if src.isLeaf == false {
			for j < eNum {
				Child = src.children[src.keyNum-1]
				b.removeElement(false, src, Child, src.keyNum-1, IntMin) //每删除一个节点keyNum也自动减少1 队尾删
				b.insertNode(dst, Child, 0)                              //队头加
				j++
			}
		} else {
			for j < eNum {
				TmpKey = src.key[src.keyNum-1]
				data = src.leafNode.data[src.keyNum-1]
				b.removeElement(true, parent, src, posAtParent, src.keyNum-1)
				b.insertData(parent, dst, TmpKey, posAtParent+1, 0, data)
				j++
			}

		}

		parent.key[posAtParent+1] = dst.key[0]
		// 将树叶节点重新连接
		if src.keyNum > 0 {
			findMostRight(src).leafNode.Next = findMostLeft(dst) //似乎不需要重连，src的最右本身就是dst最左的上一元素
		} else {
			if src.isLeaf == true {
				parent.children[posAtParent-1].leafNode.Next = dst
			}
			//  此种情况肯定是merge merge中有实现先移动再删除操作
			//b.removeElement(false ,parent.parent，parent ,parentIndex,Int_Min )
		}
	} else {
		if src.isLeaf == false {
			for j < eNum {
				Child = src.children[0]
				b.removeElement(false, src, Child, 0, IntMin) //从src的队头删
				b.insertNode(dst, Child, dst.keyNum)
				j++
			}

		} else {
			for j < eNum {
				TmpKey = src.key[0]
				data = src.leafNode.data[0]
				b.removeElement(true, parent, src, posAtParent, 0)
				b.insertData(parent, dst, TmpKey, posAtParent-1, dst.keyNum, data)
				j++
			}
		}

		parent.key[posAtParent] = src.key[0]
		if src.keyNum > 0 {
			findMostRight(dst).leafNode.Next = findMostLeft(src)
		} else {
			if src.isLeaf == true {
				dst.leafNode.Next = src.leafNode.Next
			}
			//b.removeElement(false ,parent.parent，parent ,parentIndex,Int_Min )
		}
	}

	return parent
}

// removeElement 两个参数X posAtParent 有些重复 posAtParent可以通过X的最小关键字查找得到
func (b *BPTree) removeElement(isData bool, Parent Position, X Position, posAtParent int, deleteIndex int) Position {

	var k, keyNum int

	if isData {
		keyNum = X.keyNum
		// 删除key
		k = deleteIndex + 1
		for k < keyNum {
			X.key[k-1] = X.key[k]
			X.leafNode.data[k-1] = X.leafNode.data[k]
			k++
		}

		X.key[keyNum-1] = IntMin
		X.leafNode.data[keyNum-1] = IntMin
		Parent.key[posAtParent] = X.key[0]
		X.keyNum--
	} else {
		// 删除节点
		// 修改树叶节点的链接
		if X.isLeaf == true && posAtParent > 0 {
			Parent.children[posAtParent-1].leafNode.Next = Parent.children[posAtParent+1]
		}

		keyNum = Parent.keyNum
		k = posAtParent + 1
		for k < keyNum {
			Parent.children[k-1] = Parent.children[k]
			Parent.key[k-1] = Parent.key[k]
			k++
		}

		if X.key[0] == b.ptr.key[0] { // refresh ptr
			b.ptr = Parent.children[0]
		}
		Parent.children[Parent.keyNum-1] = nil
		Parent.key[Parent.keyNum-1] = IntMin

		Parent.keyNum--

	}
	return X
}
