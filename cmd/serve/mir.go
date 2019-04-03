package serve

import (
	"github.com/okmall/okmall/models"
	"github.com/okmall/okmall/pkg/mina"
)

// mirEntries get all entries that used to register to Mir
// Notice: this func must call after models.InitWith(...)
func mirEntries() []interface{} {
	ctx := models.NewContext()

	return []interface{}{
		&mina.Address{Context: ctx},
		&mina.Brand{Context: ctx},
		&mina.Cart{Context: ctx},
		&mina.Catalog{Context: ctx},
		&mina.Collect{Context: ctx},
		&mina.Comment{Context: ctx},
		&mina.Footprint{Context: ctx},
		&mina.Goods{Context: ctx},
		&mina.Index{Context: ctx},
		&mina.Order{Context: ctx},
		&mina.Search{Context: ctx},
		&mina.Topic{Context: ctx},
		&mina.Wx{Context: ctx},
	}
}
