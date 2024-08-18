package gofer

type GroupOption func(*Group)

func UseErrorGroup() GroupOption {
	return func(g *Group) {
		g.isErrorGroup = true
	}
}

func WithUsableG(num int) GroupOption {
	return func(g *Group) {
		g.numG = num
	}
}

func WithWaitQueue(size int) GroupOption {
	return func(g *Group) {
		g.queueSize = size
	}
}
