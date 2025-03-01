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

func (c *gormRepository[E]) FindByID(ctx context.Context, id uint) (E, error) {
	return c.FindByField(ctx, "id", id)
}

func (c *gormRepository[E]) FindByField(ctx context.Context, field string, value interface{}) (E, error) {
	var e E
	err := c.Model(ctx).Where(field+"=?", value).First(&e).Error
	return e, err
}

func (c *gormRepository[E]) Remove(ctx context.Context, model E) error {
	return c.db.WithContext(ctx).Delete(model).Error
}
func (c *gormRepository[E]) Save(ctx context.Context, model E) error {
	return c.db.WithContext(ctx).Save(model).Error
}

func (c *gormRepository[E]) Model(ctx context.Context) *gorm.DB {
	var e E
	return c.db.WithContext(ctx).Model(e)
}
