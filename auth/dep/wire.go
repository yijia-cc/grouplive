//+build wireinject

package dep

import (
	"database/sql"
	"net/http"

	"github.com/google/wire"
	"github.com/yijia-cc/grouplive/auth/db/dao"
	"github.com/yijia-cc/grouplive/auth/routing"
	"github.com/yijia-cc/grouplive/auth/tm"
	"github.com/yijia-cc/grouplive/auth/tx"
)

type JWTSigningKey string
type CaesarCipherOffset int

func InitRoutingServer(jwtSigningKey JWTSigningKey, caesarCipherOffset CaesarCipherOffset, sqlDB *sql.DB) *http.ServeMux {
	wire.Build(
		wire.Bind(new(tm.Timer), new(tm.LocalTimer)),
		wire.Bind(new(tx.TransactionFactory), new(tx.SafeTransactionFactory)),
		wire.Bind(new(dao.User), new(dao.UserSQL)),

		tm.NewLocalTimer,
		tx.NewSafeTransactionFactory,
		dao.NewUserSQL,
		newRoutingServer,
	)
	return &http.ServeMux{}
}

func newRoutingServer(timer tm.Timer, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey JWTSigningKey, caesarCipherOffset CaesarCipherOffset) *http.ServeMux {
	return routing.NewServer(timer, txFactory, userDao, string(jwtSigningKey), int(caesarCipherOffset))
}
