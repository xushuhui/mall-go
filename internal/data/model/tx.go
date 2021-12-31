// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Activity is the client for interacting with the Activity builders.
	Activity *ActivityClient
	// Banner is the client for interacting with the Banner builders.
	Banner *BannerClient
	// BannerItem is the client for interacting with the BannerItem builders.
	BannerItem *BannerItemClient
	// Brand is the client for interacting with the Brand builders.
	Brand *BrandClient
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Charge is the client for interacting with the Charge builders.
	Charge *ChargeClient
	// Coupon is the client for interacting with the Coupon builders.
	Coupon *CouponClient
	// CouponTemplate is the client for interacting with the CouponTemplate builders.
	CouponTemplate *CouponTemplateClient
	// CouponType is the client for interacting with the CouponType builders.
	CouponType *CouponTypeClient
	// GridCategory is the client for interacting with the GridCategory builders.
	GridCategory *GridCategoryClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// OrderDetail is the client for interacting with the OrderDetail builders.
	OrderDetail *OrderDetailClient
	// OrderSnap is the client for interacting with the OrderSnap builders.
	OrderSnap *OrderSnapClient
	// OrderSub is the client for interacting with the OrderSub builders.
	OrderSub *OrderSubClient
	// Refund is the client for interacting with the Refund builders.
	Refund *RefundClient
	// SaleExplain is the client for interacting with the SaleExplain builders.
	SaleExplain *SaleExplainClient
	// Sku is the client for interacting with the Sku builders.
	Sku *SkuClient
	// SkuSpec is the client for interacting with the SkuSpec builders.
	SkuSpec *SkuSpecClient
	// SpecKey is the client for interacting with the SpecKey builders.
	SpecKey *SpecKeyClient
	// SpecValue is the client for interacting with the SpecValue builders.
	SpecValue *SpecValueClient
	// Spu is the client for interacting with the Spu builders.
	Spu *SpuClient
	// SpuDetailImg is the client for interacting with the SpuDetailImg builders.
	SpuDetailImg *SpuDetailImgClient
	// SpuImg is the client for interacting with the SpuImg builders.
	SpuImg *SpuImgClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Theme is the client for interacting with the Theme builders.
	Theme *ThemeClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserCoupon is the client for interacting with the UserCoupon builders.
	UserCoupon *UserCouponClient
	// UserFavor is the client for interacting with the UserFavor builders.
	UserFavor *UserFavorClient
	// UserInfo is the client for interacting with the UserInfo builders.
	UserInfo *UserInfoClient
	// UserPoint is the client for interacting with the UserPoint builders.
	UserPoint *UserPointClient
	// UserPointDetail is the client for interacting with the UserPointDetail builders.
	UserPointDetail *UserPointDetailClient
	// UserWallet is the client for interacting with the UserWallet builders.
	UserWallet *UserWalletClient
	// UserWalletDetail is the client for interacting with the UserWalletDetail builders.
	UserWalletDetail *UserWalletDetailClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Committer method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollbacker method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Activity = NewActivityClient(tx.config)
	tx.Banner = NewBannerClient(tx.config)
	tx.BannerItem = NewBannerItemClient(tx.config)
	tx.Brand = NewBrandClient(tx.config)
	tx.Category = NewCategoryClient(tx.config)
	tx.Charge = NewChargeClient(tx.config)
	tx.Coupon = NewCouponClient(tx.config)
	tx.CouponTemplate = NewCouponTemplateClient(tx.config)
	tx.CouponType = NewCouponTypeClient(tx.config)
	tx.GridCategory = NewGridCategoryClient(tx.config)
	tx.Order = NewOrderClient(tx.config)
	tx.OrderDetail = NewOrderDetailClient(tx.config)
	tx.OrderSnap = NewOrderSnapClient(tx.config)
	tx.OrderSub = NewOrderSubClient(tx.config)
	tx.Refund = NewRefundClient(tx.config)
	tx.SaleExplain = NewSaleExplainClient(tx.config)
	tx.Sku = NewSkuClient(tx.config)
	tx.SkuSpec = NewSkuSpecClient(tx.config)
	tx.SpecKey = NewSpecKeyClient(tx.config)
	tx.SpecValue = NewSpecValueClient(tx.config)
	tx.Spu = NewSpuClient(tx.config)
	tx.SpuDetailImg = NewSpuDetailImgClient(tx.config)
	tx.SpuImg = NewSpuImgClient(tx.config)
	tx.Tag = NewTagClient(tx.config)
	tx.Theme = NewThemeClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.UserCoupon = NewUserCouponClient(tx.config)
	tx.UserFavor = NewUserFavorClient(tx.config)
	tx.UserInfo = NewUserInfoClient(tx.config)
	tx.UserPoint = NewUserPointClient(tx.config)
	tx.UserPointDetail = NewUserPointDetailClient(tx.config)
	tx.UserWallet = NewUserWalletClient(tx.config)
	tx.UserWalletDetail = NewUserWalletDetailClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Activity.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
