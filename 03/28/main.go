package _8

import (
	"math/rand"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

// RandomScheduler chooses machines randomly.
type RandomScheduler struct {
	rand *rand.Rand
}

// NewRandomScheduler creates a new RandomScheduler instance
func NewRandomScheduler() *RandomScheduler {
	source := rand.NewSource(time.Now().UnixNano())
	return &RandomScheduler{
		rand: rand.New(source),
	}
}

// Schedule selects a machine randomly from the given list
func (s *RandomScheduler) Schedule(pod api.Pod, minionLister MinionLister) (string, error) {
	machines, err := minionLister.List()
	if err != nil {
		return "", err
	}
	randomIndex := s.rand.Intn(len(machines))
	result := machines[randomIndex]
	return result, nil
}
