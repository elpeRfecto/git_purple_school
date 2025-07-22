package bins

import (
	"time"
)

type Bin struct {
	CreatedAt time.Time
	Id        string
	Name      string
	private   bool
}

var BinList []Bin

func NewBin(id, name string, private bool, createdAt time.Time) *Bin {
	return &Bin{
		Id:        id,
		Name:      name,
		CreatedAt: time.Now(),
		private:   private,
	}
}
