package user

import (
	"context"
)

func (s *service) Delete(ctx context.Context, IDs []int64) error {
	for _, id := range IDs {
		if err := s.hub.User().Delete(ctx, id); err != nil {
			return err
		}
	}

	return nil
}
