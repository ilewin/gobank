package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/transparentideas/gobank/util"
)

func createRandomEntry(t *testing.T) Entry {
	a1 := createRandomAccount(t)

	ep := AddEntryParams{
		AccountID: a1.ID,
		Amount:    util.RandomMoney()}
	e, err := testQueries.AddEntry(context.Background(), ep)

	require.NoError(t, err)
	require.NotEmpty(t, e)
	require.Equal(t, a1.ID, e.AccountID)
	require.Equal(t, ep.Amount, e.Amount)
	return e
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	e := createRandomEntry(t)
	e2, err := testQueries.GetEntry(context.Background(), e.ID)
	require.NoError(t, err)
	require.NotEmpty(t, e2)
	require.Equal(t, e.ID, e2.ID)
	require.Equal(t, e.AccountID, e2.AccountID)
	require.Equal(t, e.Amount, e2.Amount)
	require.WithinDuration(t, e.CreatedAt, e2.CreatedAt, 1*time.Second)
}

func TestUpdateEntry(t *testing.T) {
	e1 := createRandomEntry(t)
	ep := UpdateEntryParams{
		ID:     e1.ID,
		Amount: util.RandomMoney(),
	}

	e2, err := testQueries.UpdateEntry(context.Background(), ep)
	require.NoError(t, err)
	require.NotEmpty(t, e2)
	require.Equal(t, e1.ID, e2.ID)
	require.Equal(t, ep.Amount, e2.Amount)
}

func TestListEntries(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {

		ep1 := AddEntryParams{
			AccountID: a1.ID,
			Amount:    util.RandomMoney(),
		}

		testQueries.AddEntry(context.Background(), ep1)

		ep2 := AddEntryParams{
			AccountID: a2.ID,
			Amount:    util.RandomMoney(),
		}

		testQueries.AddEntry(context.Background(), ep2)

	}

	lp := ListEntriesParams{
		AccountID: a1.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), lp)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, e := range entries {
		require.Equal(t, a1.ID, e.AccountID)
	}

}
