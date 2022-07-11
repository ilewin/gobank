package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)

	var err error

	errChan := make(chan error)
	resultChan := make(chan TransferTxResult)

	n := 5

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: account1.ID,
				ToAccountId:   account2.ID,
				Amount:        amount,
			})
			errChan <- err
			resultChan <- result
		}()
	}
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errChan
		require.NoError(t, err)
		txRes := <-resultChan
		require.NotEmpty(t, txRes)
		require.NotEmpty(t, txRes.Transfer)
		require.Equal(t, amount, txRes.Transfer.Amount)
		require.NotZero(t, txRes.Transfer.ID)
		require.NotZero(t, txRes.Transfer.CreatedAt)
		_, err = store.GetTransfer(context.Background(), txRes.Transfer.ID)
		require.NoError(t, err)
		require.NotEmpty(t, txRes.FromEntry)
		require.NotEmpty(t, txRes.ToEntry)
		require.Equal(t, txRes.FromEntry.AccountID, account1.ID)
		require.Equal(t, txRes.ToEntry.AccountID, account2.ID)
		require.Equal(t, txRes.FromEntry.Amount, -amount)
		require.Equal(t, txRes.ToEntry.Amount, amount)
		require.NotZero(t, txRes.FromEntry.ID)
		require.NotZero(t, txRes.FromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), txRes.FromEntry.ID)
		require.NoError(t, err)
		_, err = store.GetEntry(context.Background(), txRes.ToEntry.ID)
		require.NoError(t, err)

		fromAccount := txRes.FromAccount
		fmt.Println(fromAccount)
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, account1.ID)

		toAccount := txRes.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, account2.ID)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff2%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updateAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updateAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-(int64(n)*amount), updateAccount1.Balance)
	require.Equal(t, account2.Balance+(int64(n)*amount), updateAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)

	var err error

	errChan := make(chan error)

	n := 10

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccountID,
				ToAccountId:   toAccountID,
				Amount:        amount,
			})
			errChan <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan
		require.NoError(t, err)
	}

	updateAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updateAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)
}
