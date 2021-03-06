package inmem

import (
	"sync"

	"github.com/friendsofgo/go-api-structure-example/pkg"
)

type gameRepository struct {
	mtx   sync.RWMutex
	games map[string]*pkg.Game
}

// NewGameRepository returns a new instance of a in-memory game repository
func NewGameRepository() pkg.GameRepository {
	return &gameRepository{
		games: make(map[string]*pkg.Game),
	}
}

func (r *gameRepository) Find(ID string) ([]*pkg.Game, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var values []*pkg.Game
	for _, value := range r.games {
		if value.ID == ID {
			values = append(values, value)
		}
	}
	return values, nil
}
