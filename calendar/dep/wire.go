//+build wireinject

package dep

import (
	"database/sql"
	"net/http"

	"github.com/yijia-cc/grouplive/calendar/gql"
	"github.com/yijia-cc/grouplive/calendar/obs"

	"github.com/google/wire"
	"github.com/yijia-cc/grouplive/calendar/auth"
	"github.com/yijia-cc/grouplive/calendar/config"
	"github.com/yijia-cc/grouplive/calendar/db/dao"
	"github.com/yijia-cc/grouplive/calendar/tx"
)

func InitGraphQLServer(cfg config.Config, logger obs.Logger, db *sql.DB) (*http.ServeMux, error) {
	wire.Build(
		wire.Bind(new(auth.Authorizer), new(auth.GroupLiveAuthClient)),
		wire.Bind(new(auth.Authenticator), new(auth.GroupLiveAuthClient)),
		wire.Bind(new(auth.UserProvider), new(auth.GroupLiveAuthClient)),
		wire.Bind(new(tx.TransactionFactory), new(tx.SafeTransactionFactory)),
		wire.Bind(new(dao.Amenity), new(dao.AmenitySQL)),
		wire.Bind(new(dao.AmenityType), new(dao.AmenityTypeSQL)),

		newGroupLiveAuthClient,
		tx.NewSafeTransactionFactory,
		dao.NewAmenitySQL,
		dao.NewAmenityTypeSQL,
		gql.NewServer,
	)
	return &http.ServeMux{}, nil
}

func newGroupLiveAuthClient(cfg config.Config) (auth.GroupLiveAuthClient, error) {
	return auth.NewGroupLiveAuthClient(cfg.AuthGRPCEndpoint)
}
