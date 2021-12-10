package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePerson(t *testing.T) {

	arg := CreatePersonParams{
		Name: "Atiyeh",
		Age:  "23 years",
	}

	person, err := testQueries.CreatePerson(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, arg.Name, person.Name)
	require.Equal(t, arg.Age, person.Age)

	require.NotZero(t, person.ID)

}
