//+build wireinject

package dep

import (
	"database/sql"
	"net/http"

	"github.com/yijia-cc/grouplive/auth/idgen"

	"github.com/yijia-cc/grouplive/auth/rpc"
	"google.golang.org/grpc"

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
		wire.Bind(new(dao.PermissionBinding), new(dao.PermissionBindingSQL)),
		wire.Bind(new(idgen.IDGenerator), new(idgen.UUIDGenerator)),

		tm.NewLocalTimer,
		tx.NewSafeTransactionFactory,
		dao.NewUserSQL,
		dao.NewPermissionBindingSQL,
		idgen.NewUUIDGenerator,
		newRoutingServer,
	)
	return &http.ServeMux{}
}

func InitGRPCServer(jwtSigningKey JWTSigningKey, caesarCipherOffset CaesarCipherOffset, sqlDB *sql.DB) *grpc.Server {
	wire.Build(
		wire.Bind(new(tm.Timer), new(tm.LocalTimer)),
		wire.Bind(new(tx.TransactionFactory), new(tx.SafeTransactionFactory)),
		wire.Bind(new(dao.User), new(dao.UserSQL)),
		wire.Bind(new(dao.PermissionBinding), new(dao.PermissionBindingSQL)),
		wire.Bind(new(idgen.IDGenerator), new(idgen.UUIDGenerator)),

		tm.NewLocalTimer,
		tx.NewSafeTransactionFactory,
		dao.NewUserSQL,
		dao.NewPermissionBindingSQL,
		idgen.NewUUIDGenerator,
		newGRPCServer,
	)
	return nil
}

func newGRPCServer(timer tm.Timer, idGenerator idgen.IDGenerator, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey JWTSigningKey, caesarCipherOffset CaesarCipherOffset, permissionBinding dao.PermissionBinding) *grpc.Server {
	return rpc.NewServer(timer, idGenerator, txFactory, userDao, string(jwtSigningKey), int(caesarCipherOffset), permissionBinding)
}

func newRoutingServer(timer tm.Timer, idGenerator idgen.IDGenerator, txFactory tx.TransactionFactory, userDao dao.User, jwtSigningKey JWTSigningKey, caesarCipherOffset CaesarCipherOffset, permissionBinding dao.PermissionBinding) *http.ServeMux {
	return routing.NewServer(timer, idGenerator, txFactory, userDao, string(jwtSigningKey), int(caesarCipherOffset), permissionBinding)
}
