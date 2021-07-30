package generator

type Generator interface {
	NextID() (uint64, error)
}
