package adapter

import (
	"context"
	"gorm.io/gorm"
)

type gormRepository[E Entity] struct {
	db *gorm.DB
}

func NewGormRepository[E Entity](db *gorm.DB) BaseRepository[E] {
	return &gormRepository[E]{db: db}
}

func (c *gormRepository[E]) ByID(ctx context.Context, id uint) (E, error) {
	return c.ByField(ctx, "id", id)
}

func (c *gormRepository[E]) ByField(ctx context.Context, field string, value interface{}) (E, error) {
	var e E
	err := c.db.WithContext(ctx).Where(field+"=?", value).First(&e).Error
	return e, err
}

func (c *gormRepository[E]) Add(ctx context.Context, model E) error {
	return c.db.WithContext(ctx).Create(model).Error
}
