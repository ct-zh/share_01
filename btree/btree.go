package btree

// todo 单元测试

// B+ 树和 B 树有什么不同：
// 1. B+树非叶子节点上是不存储数据的，仅存储键值，而B树节点中不仅存储键值，也会存储数据。
// 数据库中页的大小是固定的，InnoDB 中页的默认大小是 16kB。
// 非叶子结点不存储数据 = 存储更多的key = 树的阶数变大、层数变小(更矮更胖) = IO次数减少 = 数据查询的效率更快。
// B+树的分叉数等于键值的数量，如果一个节点存储1000个键值，那么3层 B+ 树可以存储 1000×1000×1000=10亿个数据,
// 只需要2次磁盘IO(根结点二分查找到第二层的位置，一次io；第二层结点继续二分查找到第三层叶子结点的位置，一次io，取出数据)
//
// 补充： InnoDB页默认大小是16kB
// - 对于叶子节点，一般行大小假设是1kb，那么一页能存储16kB/1kB = 16 条记录；
// - 不过为了探测页是否满了以便进行页分裂，InnoDB的默认填充因子是页大小的15/16；所以这里算15条记录
// - 对于非叶子节点，因为只需要存主键id+指针，假设主键id类型是bigint，bigint占8字节；指针在InnoDB源码中占6字节；
//   也就是一个主键+指针=14字节，(16kB*1024)/14byte=1170 可以存储1170条数据
// - root存储1170条数据，高度为2的B+Tree可以存储 1170*15=17550条数据
// - 高度为3的B+Tree可以存储 1170*1170*15=20,533,500 可以存2000w条数据
//
// 2. B+树 叶子节点数据是按照顺序排列的,范围查找，排序查找，分组查找以及去重查找变得简单
// innodb中各个页之间通过双向链表连接，叶子节点中的数据通过单向链表连接的,效率更高

// B+树特性
// - 所有的值按顺序存储
// - 每个叶子节点到根节点的距离相同

// M MySQL的填充因子为15/16，这里简化一下逻辑，假设每条数据为1kB大小
// 也就是说达到15条数据时此页已满，达到M+1也就是16条数据时触发页分裂
const M = 6

// LimitM2 说明该页数据不足一半，触发页合并
const LimitM2 = (M + 1) / 2
const IntMin = -1 // 还未写入的slot，id值默认为-1

type bPTree struct {
	keyMax int         // 当前tree的最大关键字；也就是AUTO_INCREMENT的值
	width  int         // 阶，也就是树的高度
	root   *bPFullNode // 根节点
	ptr    *bPFullNode

	cfg *Cfg
}

// bPFullNode 页结构
type bPFullNode struct {
	isLeaf   bool        // 是否是叶子节点？
	leafNode *bPLeafNode // 页存储的具体数据；Innodb的非叶子节点是不存数据的，为nil；叶子节点是存数据的，不为nil；

	nodeCatalogue

	point // 每个页都存储了前一页和后一页的指针
}

type Position *bPFullNode

// BPLeafNode 叶子节点的数据
type bPLeafNode struct {
	Next *bPFullNode
	data map[int]interface{} // 详细数据 key=id value=具体的数据;实际Innodb这里不是简单存储了数据，而是以链表的形式存储了多版本号的数据;
}

// nodeCatalogue 页目录
type nodeCatalogue struct {
	keyNum   int           // 当前页目录的大小,当前页的最大id和最小id通过ta求出来
	key      []int         // 当前页的页目录,int 值为主键id的值
	children []*bPFullNode // 子节点;
}

// point 页的指针，每个层级的页组成链表结构
type point struct {
	prev *bPFullNode
	next *bPFullNode
}

type Cfg struct {
	PageMaxSize int
}

var NotExistErr error

var AlreadyExistsErr error
