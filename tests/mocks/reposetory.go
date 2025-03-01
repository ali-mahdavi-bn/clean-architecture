package mocks

import (
	"clean-hex/pkg/framwork/adapter"
	"context"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type FakRepository[E adapter.Entity] struct {
	mock.Mock
}

func NewFakeRepository[E adapter.Entity]() *FakRepository[E] {
	return &FakRepository[E]{}
}

func (c *FakRepository[E]) FindByID(ctx context.Context, id uint) (E, error) {
	args := c.Called(ctx, id)
	var e E
	if args.Get(0) != nil {
		e = args.Get(0).(E)
	}
	return e, args.Error(1)
}

func (c *FakRepository[E]) FindByField(ctx context.Context, field string, value interface{}) (E, error) {
	args := c.Called(ctx, field, value)
	var e E
	if args.Get(0) != nil {
		e = args.Get(0).(E)
	}
	return e, args.Error(1)
}

func (c *FakRepository[E]) Remove(ctx context.Context, model E) error {
	args := c.Called(ctx, model)
	return args.Error(0)
}

func (c *FakRepository[E]) Save(ctx context.Context, model E) error {
	args := c.Called(ctx, model)
	return args.Error(0)
}

func (c *FakRepository[E]) Model(ctx context.Context) *gorm.DB {
	args := c.Called(ctx)
	return args.Get(0).(*gorm.DB)
}
