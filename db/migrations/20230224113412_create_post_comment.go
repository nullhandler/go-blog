package migrations

import (
	"context"
	"fmt"
	"go_crud/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.Post)(nil)).
			Exec(ctx)

		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().
			IfNotExists().
			Model((*models.Comment)(nil)).
			Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")

		_, err := db.NewDropTable().
			Model((*models.Post)(nil)).
			IfExists().
			Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().
			Model((*models.Comment)(nil)).
			IfExists().
			Exec(ctx)
		return err
	})
}
