package pkg

import (
	"context"
	"errors"
)

func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return "", errors.New("User ID not found in context")
	}
	if userID == "" {
		return "", errors.New("User ID is empty")
	}
	return userID, nil
}
