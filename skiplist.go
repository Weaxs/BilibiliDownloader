package main

type zskiplistLevel struct {
	// 前进指针
	forward *zskiplistNode
	//跨度
	span int64
}

type zskiplistNode struct {
	// 后退指针
	backward *zskiplistNode
	// 权重 (用于排序)
	score float64

	// 层
	level []zskiplistLevel

	//level[] struct{
	//	// 前进指针
	//	forward *zskiplistNode;
	//	//跨度
	//	span int64;
	//}
}

type zskiplist struct {
	// 链头、链尾
	header, tail *zskiplistNode

	length int64

	level int16
}

func initSkipList() zskiplist {

	tail := &zskiplistNode{
		backward: nil,
		score:    1,
		level:    []zskiplistLevel{},
	}

	header := &zskiplistNode{
		backward: tail,
		score:    -1,
		level:    []zskiplistLevel{},
	}

	return zskiplist{
		header: header,
		tail:   tail,
		length: 0,
		level:  0,
	}
}

func main() {

}
