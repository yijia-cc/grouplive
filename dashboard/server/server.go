package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yijia-cc/grouplive/dashboard/auth"
	"github.com/yijia-cc/grouplive/dashboard/config"
	"github.com/yijia-cc/grouplive/dashboard/service"
	"log"
	"net/http"
)

/*func StartUp(cfg *config.Config) {
	jwtMiddleware := middleware.StartupJWT(cfg)
	httpRouter := router.StartupHttpRouter(cfg, jwtMiddleware)

	log.Printf("Dashboard server started at port %d\n", cfg.App.WebServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.WebServerPort), httpRouter))
}*/

func StartUp(cfg *config.Config) {
	//jwtMiddleware := middleware.StartupJWT(cfg)

	client, err := auth.NewGroupLiveAuthClient("127.0.0.1:8080")
	if err != nil {
		return
	}

	r := mux.NewRouter()

	// 3 valid route variables are accepted: mixed/grouped/dashboard
	r.Handle("/search/{search_type}/", auth.WithMiddleware(client, client, service.SearchHandler)).Methods("GET", "OPTIONS")

	r.Handle("/post", auth.WithMiddleware(client, client, service.PostHandler)).Methods("POST", "OPTIONS")
	r.Handle("/confirm", auth.WithMiddleware(client, client, service.ConfirmHandler)).Methods("POST", "OPTIONS")
	r.Handle("/update", auth.WithMiddleware(client, client, service.UpdateHandler)).Methods("PUT", "OPTIONS")
	r.Handle("/delete/{id}", auth.WithMiddleware(client, client, service.DeleteHandler)).Methods("DELETE", "OPTIONS")

	// Serve any files located at <STAIC_MEDIA_DIR>/<filename> via the URL: http://host:port/<STAIC_MEDIA_DIR>/<filename>
	r.PathPrefix(cfg.App.StaticMediaDir[1:]).Handler(http.StripPrefix(cfg.App.StaticMediaDir[1:], http.FileServer(http.Dir(cfg.App.StaticMediaDir))))

	//log.Printf("Dashboard server started at port %d\n", cfg.App.WebServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.WebServerPort), r))
}
