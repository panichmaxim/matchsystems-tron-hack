package btcstore

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
	"testing"
	"time"
)

const testNeoFirst = "neo4j://51.250.16.46"
const testNeoUsername = "neo4j"
const testNeoPassword = "RubinNeo#"

var ctx = context.TODO()

func createTestSession() (neo4j.SessionWithContext, error) {
	driver, err := createTestDriver()
	if err != nil {
		return nil, err
	}
	return neoutils.CreateSession(driver), nil
}

func createTestDriver() (neo4j.DriverWithContext, error) {
	return neoutils.CreateDriver(testNeoFirst, testNeoUsername, testNeoPassword)
}

func TestCreateDriver(t *testing.T) {
	driver, err := createTestDriver()
	require.NoError(t, err)
	require.NotNil(t, driver)
	timed, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	require.NoError(t, driver.VerifyConnectivity(timed))
}
