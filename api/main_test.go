package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/transparentideas/gobank/db/sqlc"
	"github.com/transparentideas/gobank/token"
	"github.com/transparentideas/gobank/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {

	config := util.Config{
		TokenSymetricKey: util.RandomString(32),
		AccessTokenTTL:   time.Minute,
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		t.Fatal(err)
	}

	return &Server{
		store:      store,
		config:     &config,
		router:     gin.Default(),
		tokenMaker: tokenMaker,
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
