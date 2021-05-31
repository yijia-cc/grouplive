//+build wireinject

package dep

import (
	"database/sql"
	"net/http"

	"github.com/google/wire"
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/config"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/gql/server"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

func InitGraphQLServer(cfg config.Config, db *sql.DB) *http.ServeMux {
	wire.Build(
		wire.Bind(new(auth.Authorizer), new(auth.Client)),
		wire.Bind(new(auth.Authenticator), new(auth.Client)),
		wire.Bind(new(tx.TransactionFactory), new(tx.SafeTransactionFactory)),
		wire.Bind(new(dao.Amenity), new(dao.AmenitySQL)),
		wire.Bind(new(dao.AmenityType), new(dao.AmenityTypeSQL)),

		auth.NewClient,
		tx.NewSafeTransactionFactory,
		dao.NewAmenitySQL,
		dao.NewAmenityTypeSQL,
		server.NewServer,
	)
	return &http.ServeMux{}
}
