package cmdutils

import (
	"errors"
	"fmt"
	"os"

	"github.com/itera-io/taikun-cli/utils/types"
)

func ArgsToNumericalIDs(args []string) ([]int32, error) {
	ids := make([]int32, len(args))

	for argIndex, arg := range args {
		id, err := types.Atoi32(arg)
		if err != nil {
			return nil, err
		}

		ids[argIndex] = id
	}

	return ids, nil
}

type DeleteFunc func(int32) error

func DeleteMultiple(ids []int32, deleteFunc DeleteFunc) error {
	errorOccured := false

	for _, id := range ids {
		if err := deleteFunc(id); err != nil {
			errorOccured = true

			fmt.Fprintln(os.Stderr, err)
		}
	}

	if errorOccured {
		fmt.Fprintln(os.Stderr)
		return errors.New("Failed to delete one or more resources")
	}

	return nil
}

type DeleteFuncStringID func(string) error

func DeleteMultipleStringID(ids []string, deleteFunc DeleteFuncStringID) error {
	errorOccured := false

	for _, id := range ids {
		if err := deleteFunc(id); err != nil {
			errorOccured = true

			fmt.Fprintln(os.Stderr, err)
		}
	}

	if errorOccured {
		fmt.Fprintln(os.Stderr)
		return errors.New("Failed to delete one or more resources")
	}

	return nil
}

type DeleteFuncChildResource func(int32, int32) error

func DeleteMultipleChildResources(parentID int32, ids []int32, deleteFunc DeleteFuncChildResource) error {
	errorOccured := false

	for _, id := range ids {
		if err := deleteFunc(parentID, id); err != nil {
			errorOccured = true

			fmt.Fprintln(os.Stderr, err)
		}
	}

	if errorOccured {
		fmt.Fprintln(os.Stderr)
		return errors.New("Failed to delete one or more resources")
	}

	return nil
}
