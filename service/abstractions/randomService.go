package service

type RandomService interface {
	UpdateSeededRand()
	RandString(length int) string
	RandInt() int
	RandIntRange(min int, max int) int
}
