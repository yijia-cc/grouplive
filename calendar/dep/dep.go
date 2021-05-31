//+build wireinject

package dep

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/gql/resolver"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

func InitGraphQLResolver(db *sql.DB) *resolver.Resolver {
	wire.Build(
		wire.Bind(new(tx.TransactionFactory), new(tx.SafeTransactionFactory)),
		wire.Bind(new(dao.Amenity), new(dao.AmenitySQL)),
		wire.Bind(new(dao.AmenityType), new(dao.AmenityTypeSQL)),

		tx.NewSafeTransactionFactory,
		dao.NewAmenitySQL,
		dao.NewAmenityTypeSQL,
		resolver.NewResolver,
	)
	return &resolver.Resolver{}
}
