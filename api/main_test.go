package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "github.com/transparentideas/gobank/db/sqlc"
	"github.com/transparentideas/gobank/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {

	config := util.Config{
		TokenSymetricKey: util.RandomString(32),
		AccessTokenTTL:   time.Minute,
	}

	server, err := NewServer(&config, store)
	require.NoError(t, err)

	return server

}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
