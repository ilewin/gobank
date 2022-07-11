package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/themeszone/gobank/util"
)

func createRandomTransfer(t *testing.T, a1 Account, a2 Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: a1.ID,
		ToAccountID:   a2.ID,
		Amount:        util.RandomMoney(),
	}

	tr, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tr)
	require.Equal(t, a1.ID, tr.FromAccountID)
	require.Equal(t, a2.ID, tr.ToAccountID)
	require.Equal(t, arg.Amount, tr.Amount)

	return tr

}

func TestCreateTransfer(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	createRandomTransfer(t, a1, a2)
}

func TestGetTransfer(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	tr := createRandomTransfer(t, a1, a2)
	tr2, err := testQueries.GetTransfer(context.Background(), tr.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tr2)
	require.Equal(t, tr.ID, tr2.ID)
	require.Equal(t, tr.FromAccountID, tr2.FromAccountID)
	require.Equal(t, tr.ToAccountID, tr2.ToAccountID)
	require.Equal(t, tr.Amount, tr2.Amount)
	require.WithinDuration(t, tr.CreatedAt, tr2.CreatedAt, 1*time.Second)
}

func TestListAllTransfersFrom(t *testing.T) {

	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	a3 := createRandomAccount(t)
	a4 := createRandomAccount(t)

	var accs []*Account

	accs = append(accs, &a1, &a2, &a3, &a4)

	for i := 0; i < 16; i++ {
		aF := accs[rand.Int()%len(accs)]
		aT := accs[rand.Int()%len(accs)]
		createRandomTransfer(t, *aF, *aT)
	}

	aTest := *accs[rand.Int()%len(accs)]

	lta := ListAllTransfersFromParams{
		FromAccountID: aTest.ID,
		Limit:         3,
		Offset:        0,
	}

	trs, err := testQueries.ListAllTransfersFrom(context.Background(), lta)
	require.NoError(t, err)
	require.NotEmpty(t, trs)
	require.Len(t, trs, 3)

	for _, ts := range trs {
		require.NotEmpty(t, ts)
		require.Equal(t, aTest.ID, ts.FromAccountID)
	}
}

func TestListAllTransfersTo(t *testing.T) {

	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	a3 := createRandomAccount(t)
	a4 := createRandomAccount(t)

	var accs []*Account

	accs = append(accs, &a1, &a2, &a3, &a4)

	for i := 0; i < 16; i++ {
		aF := accs[rand.Int()%len(accs)]
		aT := accs[rand.Int()%len(accs)]
		createRandomTransfer(t, *aF, *aT)
	}

	aTest := *accs[rand.Int()%len(accs)]

	lta := ListAllTransfersToParams{
		ToAccountID: aTest.ID,
		Limit:       3,
		Offset:      0,
	}

	trs, err := testQueries.ListAllTransfersTo(context.Background(), lta)
	require.NoError(t, err)
	require.NotEmpty(t, trs)
	require.Len(t, trs, 3)

	for _, ts := range trs {
		require.NotEmpty(t, ts)
		require.Equal(t, aTest.ID, ts.ToAccountID)
	}

}

func TestListAllTransfersFromTo(t *testing.T) {

	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, a1, a2)
	}

	lta := ListAllTransfersFromToParams{
		FromAccountID: a1.ID,
		ToAccountID:   a2.ID,
		Limit:         5,
		Offset:        0,
	}

	trs, err := testQueries.ListAllTransfersFromTo(context.Background(), lta)
	require.NoError(t, err)
	require.NotEmpty(t, trs)
	require.Len(t, trs, 5)

	for _, ts := range trs {
		require.NotEmpty(t, ts)
		require.True(t, ts.ToAccountID == a2.ID && ts.FromAccountID == a1.ID)
	}

}
