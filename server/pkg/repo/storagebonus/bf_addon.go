package storagebonus

import (
	"context"
	"fmt"
	"github.com/ente-io/museum/ente/storagebonus"
)

func (r *Repository) InsertBFBonus(ctx context.Context, userID int64, validTill int64, storage int64) error {
	bonusID := fmt.Sprintf("%s-%d", storagebonus.AddOnBf2023, userID)
	_, err := r.DB.ExecContext(ctx, "INSERT INTO storage_bonus (bonus_id, user_id, storage, type, valid_till) VALUES ($1, $2, $3, $4, $5)", bonusID, userID, storage, storagebonus.AddOnBf2023, validTill)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RemoveBFBonus(ctx context.Context, userID int64) (int64, error) {
	bonusID := fmt.Sprintf("%s-%d", storagebonus.AddOnBf2023, userID)
	res, err := r.DB.ExecContext(ctx, "DELETE FROM storage_bonus WHERE bonus_id = $1", bonusID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *Repository) UpdateBFBonus(ctx context.Context, userID int64, validTill int64, storage int64) error {
	bonusID := fmt.Sprintf("%s-%d", storagebonus.AddOnBf2023, userID)
	_, err := r.DB.ExecContext(ctx, "UPDATE storage_bonus SET storage = $1, valid_till = $2 WHERE bonus_id = $3", storage, validTill, bonusID)
	if err != nil {
		return err
	}
	return nil
}
