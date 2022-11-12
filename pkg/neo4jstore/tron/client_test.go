package tron

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"os"
	"testing"
	"time"
)

const testNeoFirst = "neo4j://46.4.253.163"
const testNeoUsername = "neo4j"
const testNeoPassword = "RubinNeo#"

var ctx = context.TODO()

func createTestSession() (neo4j.SessionWithContext, error) {
	driver, err := createTestDriver()
	if err != nil {
		return nil, err
	}
	return neo4jstore.CreateSession(driver), nil
}

func createTestDriver() (neo4j.DriverWithContext, error) {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.ErrorLevel)
	ctx := logger.WithContext(context.Background())
	return neo4jstore.CreateDriver(ctx, testNeoFirst, testNeoUsername, testNeoPassword)
}

func TestCreateDriver(t *testing.T) {
	driver, err := createTestDriver()
	require.NoError(t, err)
	require.NotNil(t, driver)
	timed, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	require.NoError(t, driver.VerifyConnectivity(timed))
}
