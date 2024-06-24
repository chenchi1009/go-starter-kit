package repository

import (
	"context"
	"fmt"

	"time"

	"github.com/chenchi1009/go-starter-kit/internal/model"
	"github.com/chenchi1009/go-starter-kit/pkg/log"
	"github.com/chenchi1009/go-starter-kit/pkg/zapgorm"

	"github.com/redis/go-redis/v9"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	//rdb    *redis.Client
	logger *log.Logger
}

// Create implements UserReposisotry.
func (r *Repository) Create(ctx context.Context, user *model.User) error {
	panic("unimplemented")
}

// Delete implements UserReposisotry.
func (r *Repository) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// GetByID implements UserReposisotry.
func (r *Repository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	panic("unimplemented")
}

// GetByUsername implements UserReposisotry.
func (r *Repository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	panic("unimplemented")
}

// List implements UserReposisotry.
func (r *Repository) List(ctx context.Context, page int, pageSize int) ([]*model.User, error) {
	panic("unimplemented")
}

// Update implements UserReposisotry.
func (r *Repository) Update(ctx context.Context, user *model.User) error {
	panic("unimplemented")
}

func NewRepository(
	logger *log.Logger,
	db *gorm.DB,
	// rdb *redis.Client,
) *Repository {
	return &Repository{
		db: db,
		//rdb:    rdb,
		logger: logger,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx)
	})
}

func NewDB(driver, dsn string, l *log.Logger) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	logger := zapgorm.New(l.Logger)

	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html
	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger,
		})
	case "postgres":
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		panic("unknown db driver")
	}
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func NewRedis(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}
