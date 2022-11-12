package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"gitlab.com/rubin-dev/api/internal/tools"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"time"
)

var ErrInsufficientBalance = errors.New("insufficient balance")

var _ BillingStore = (*storeImpl)(nil)

type BillingStore interface {
	BillingRisks(ctx context.Context, id int64) ([]*models.BillingRisk, error)
	BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error)
	BillingRequestFindByID(ctx context.Context, id int64) (*models.BillingRequest, error)
	BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error)
	BillingStatisticsSummary(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		last *bool,
	) ([]*models.StatisticsSummary, error)
	BillingRegisterRequest(
		ctx context.Context,
		userID int64,
		query string,
		risk *neo4jstore.Risk,
		network string,
	) (*models.Billing, error)
	BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error)
	BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error)
	BillingHistoryList(
		ctx context.Context,
		userID int64,
		limit, offset int,
		from, to *time.Time,
		last *bool,
	) ([]*models.BillingRequest, int, error)
	BillingKeyList(ctx context.Context, userID int64, limit, offset int) ([]*models.BillingKey, int, error)
	BillingKeyRemove(ctx context.Context, userID, id int64) error
	BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error)
	BillingRiskFindDirectories(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) ([]*BillingDirectoryOrCategoryRisk, error)
	BillingRiskFindCategories(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) ([]*BillingDirectoryOrCategoryRisk, error)
	BillingStatisticsRisks(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) ([]*models.BillingRisk, error)
	BillingRiskFindCategoriesAndDirectories(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) (map[string]int, error)
	BillingStatisticsRiskRange(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		riskFrom, riskTo *int,
		last *bool,
	) (int, error)
	BillingStatisticsRiskRanges(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) ([]int, error)
	BillingStatisticsNetwork(
		ctx context.Context,
		userID int64,
		from, to *time.Time,
		network string,
		last *bool,
	) (*models.BillingStatisticsBlockchain, error)
}

func (s *storeImpl) BillingHistoryRequestsCount(ctx context.Context, userID int64) (int, error) {
	var count int
	err := s.db.
		NewSelect().
		Model((*models.BillingRequest)(nil)).
		ColumnExpr("count(id)").
		Where("user_id = ?", userID).
		Scan(ctx, &count)
	return count, err
}

func (s *storeImpl) BillingKeyRemove(ctx context.Context, userID, id int64) error {
	_, err := s.db.
		NewDelete().
		Model((*models.BillingKey)(nil)).
		Where("user_id = ? AND id = ?", userID, id).
		Exec(ctx)
	return err
}

func (s *storeImpl) BillingKeyList(ctx context.Context, userID int64, limit, offset int) ([]*models.BillingKey, int, error) {
	var u []*models.BillingKey
	count, err := s.db.
		NewSelect().
		Model((*models.BillingKey)(nil)).
		Where("user_id = ?", userID).
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx, &u)
	return u, count, err
}

func (s *storeImpl) BillingHistoryList(
	ctx context.Context,
	userID int64,
	limit, offset int,
	from, to *time.Time,
	last *bool,
) ([]*models.BillingRequest, int, error) {
	var u []*models.BillingRequest
	q := s.db.
		NewSelect().
		Model((*models.BillingRequest)(nil)).
		Where("user_id = ?", userID)

	if last != nil && *last {
		q = q.Where("last = ?", true)
	}

	if from != nil {
		q = q.Where("created_at::date >= ?", *from)
	}
	if to != nil {
		q = q.Where("created_at::date <= ?", *to)
	}
	count, err := q.
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx, &u)
	return u, count, err
}

func (s *storeImpl) BillingStatisticsNetwork(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) (*models.BillingStatisticsBlockchain, error) {
	total, err := s.billingHistoryStatisticsTotal(
		ctx,
		userID,
		from, to,
		network,
		last,
	)
	if err != nil {
		return nil, err
	}

	riskMap, err := s.BillingRiskFindCategoriesAndDirectories(
		ctx,
		userID,
		from, to,
		network,
		last,
	)
	if err != nil {
		return nil, err
	}

	return &models.BillingStatisticsBlockchain{
		Total:      total,
		Categories: riskMap,
	}, nil
}

func (s *storeImpl) BillingStatisticsRisks(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) ([]*models.BillingRisk, error) {
	var risks []*models.BillingRisk

	q := s.db.NewSelect().
		Model((*models.BillingRisk)(nil)).
		Join("LEFT JOIN billing_request").
		JoinOn("billing_risk.billing_request_id = billing_request.id").
		JoinOn("billing_risk.is_calculated = billing_request.is_calculated").
		JoinOn("billing_risk.is_reported = billing_request.is_reported").
		JoinOn("billing_risk.is_wallet = billing_request.is_wallet").
		Where("billing_request.user_id = ? AND billing_request.network = ?", userID, network)

	if last != nil && *last {
		q = q.Where("billing_request.last = ?", true)
	}

	if from != nil {
		q = q.Where("billing_request.created_at::date >= ?", *from)
	}
	if to != nil {
		q = q.Where("billing_request.created_at::date <= ?", *to)
	}

	if err := q.Scan(ctx, &risks); err != nil {
		return nil, err
	}

	return risks, nil
}

func (s *storeImpl) billingHistoryStatisticsQuery(
	userID int64,
	from, to *time.Time,
	network string,
) *bun.SelectQuery {
	q := s.db.
		NewSelect().
		Model((*models.BillingRequest)(nil)).
		Where("user_id = ? AND network = ?", userID, network)
	if from != nil {
		q = q.Where("created_at::date >= ?", *from)
	}
	if to != nil {
		q = q.Where("created_at::date <= ?", *to)
	}

	return q
}

func (s *storeImpl) billingHistoryStatisticsTotal(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) (int64, error) {
	var total int64
	q := s.billingHistoryStatisticsQuery(userID, from, to, network).
		ColumnExpr("count(id) AS total")

	if last != nil && *last {
		q = q.Where("last = ?", true)
	}

	if err := q.Scan(ctx, &total); err != nil {
		return -1, err
	}

	return total, nil
}

func (s *storeImpl) BillingStatisticsRiskRanges(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) ([]int, error) {
	low, err := s.BillingStatisticsRiskRange(ctx, userID, from, to, network, nil, tools.Ptr[int](25), last)
	if err != nil {
		return nil, err
	}
	mid, err := s.BillingStatisticsRiskRange(ctx, userID, from, to, network, tools.Ptr[int](25), tools.Ptr[int](50), last)
	if err != nil {
		return nil, err
	}
	high, err := s.BillingStatisticsRiskRange(ctx, userID, from, to, network, tools.Ptr[int](50), tools.Ptr[int](75), last)
	if err != nil {
		return nil, err
	}
	crit, err := s.BillingStatisticsRiskRange(ctx, userID, from, to, network, tools.Ptr[int](75), nil, last)
	if err != nil {
		return nil, err
	}

	return []int{low, mid, high, crit}, nil
}

// BillingStatisticsRiskRange
// green: > 0 & < 25
// yellow: > 25 & < 50
// orange: > 50 & < 75
// red: > 75
func (s *storeImpl) BillingStatisticsRiskRange(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	riskFrom, riskTo *int,
	last *bool,
) (int, error) {
	q := s.db.NewSelect().
		Model((*models.BillingRequest)(nil)).
		ColumnExpr("count(id) AS count").
		Where("user_id = ?", userID).
		Where("network = ?", network)

	if last != nil && *last {
		q = q.Where("last = ?", true)
	}
	if riskFrom != nil {
		q = q.Where("risk > ?", *riskFrom)
	}
	if riskTo != nil {
		q = q.Where("risk <= ?", *riskTo)
	}

	if from != nil {
		q = q.Where("created_at::date >= ?", *from)
	}
	if to != nil {
		q = q.Where("created_at::date <= ?", *to)
	}

	var result int
	if err := q.Scan(ctx, &result); err != nil {
		return -1, err
	}

	return result, nil
}

func (s *storeImpl) BillingFindApiKey(ctx context.Context, apiKey string) (*models.BillingKey, error) {
	billingApiKey := new(models.BillingKey)
	err := s.db.
		NewSelect().
		Model((*models.BillingKey)(nil)).
		Where("key = ?", apiKey).
		Scan(ctx, billingApiKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return billingApiKey, nil
}

type BillingDirectoryOrCategoryRisk struct {
	ID    int
	Name  string
	Count int
}

func (s *storeImpl) BillingRiskFindDirectories(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) ([]*BillingDirectoryOrCategoryRisk, error) {
	var risks []*BillingDirectoryOrCategoryRisk

	q := s.db.NewSelect().
		Model((*models.BillingRisk)(nil)).
		ColumnExpr("directory_id AS id").
		ColumnExpr("count(directory_id) AS count").
		Join("LEFT JOIN billing_request").
		JoinOn("billing_risk.billing_request_id = billing_request.id").
		JoinOn("billing_risk.is_calculated = billing_request.is_calculated").
		JoinOn("billing_risk.is_reported = billing_request.is_reported").
		JoinOn("billing_risk.is_wallet = billing_request.is_wallet").
		Where("billing_request.user_id = ?", userID).
		Where("billing_request.network = ?", network).
		Where("billing_risk.category_id IS NULL").
		Group("billing_risk.directory_id").
		Having("count(billing_risk.directory_id) > 0").
		Order("billing_risk.directory_id")

	if last != nil && *last {
		q = q.Where("billing_request.last = ?", true)
	}

	if from != nil {
		q = q.Where("billing_request.created_at::date >= ?", *from)
	}
	if to != nil {
		q = q.Where("billing_request.created_at::date <= ?", *to)
	}

	if err := q.Scan(ctx, &risks); err != nil {
		return nil, err
	}

	for _, r := range risks {
		d := models.FindDirectory(r.ID)
		if d == nil {
			return nil, errors.New(fmt.Sprintf("unknown directory: %d", r.ID))
		}
		r.Name = d.NameEn
	}

	return risks, nil
}

func (s *storeImpl) BillingRiskFindCategoriesAndDirectories(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) (map[string]int, error) {
	categories, err := s.BillingRiskFindCategories(ctx, userID, from, to, network, last)
	if err != nil {
		return nil, err
	}

	directories, err := s.BillingRiskFindDirectories(ctx, userID, from, to, network, last)
	if err != nil {
		return nil, err
	}

	result := map[string]int{}

	for _, c := range categories {
		if val, ok := result[c.Name]; ok {
			result[c.Name] = val + c.Count
		} else {
			result[c.Name] = c.Count
		}
	}

	for _, c := range directories {
		if val, ok := result[c.Name]; ok {
			result[c.Name] = val + c.Count
		} else {
			result[c.Name] = c.Count
		}
	}

	return result, nil
}

func (s *storeImpl) BillingRiskFindCategories(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	network string,
	last *bool,
) ([]*BillingDirectoryOrCategoryRisk, error) {
	var risks []*BillingDirectoryOrCategoryRisk

	q := s.db.NewSelect().
		Model((*models.BillingRisk)(nil)).
		ColumnExpr("category_id AS id").
		ColumnExpr("category.name AS name").
		ColumnExpr("count(category_id) AS count").
		Join("LEFT JOIN billing_request").
		JoinOn("billing_risk.billing_request_id = billing_request.id").
		JoinOn("billing_risk.is_calculated = billing_request.is_calculated").
		JoinOn("billing_risk.is_reported = billing_request.is_reported").
		JoinOn("billing_risk.is_wallet = billing_request.is_wallet").
		Join("LEFT JOIN category").
		JoinOn("category.number = billing_risk.category_id").
		Where("billing_request.user_id = ? AND billing_request.network = ?", userID, network).
		Group("billing_risk.category_id").
		Group("category.name").
		Having("count(billing_risk.category_id) > 0").
		Order("billing_risk.category_id")

	if last != nil && *last {
		q = q.Where("billing_request.last = ?", true)
	}

	if from != nil {
		q = q.Where("billing_request.created_at::date >= ?", *from)
	}
	if to != nil {
		q = q.Where("billing_request.created_at::date <= ?", *to)
	}

	if err := q.Scan(ctx, &risks); err != nil {
		return nil, err
	}

	return risks, nil
}

func (s *storeImpl) BillingCreateApiKey(ctx context.Context, userID int64) (*models.BillingKey, error) {
	billing, err := s.BillingGetOrCreate(ctx, userID)
	if err != nil {
		return nil, err
	}
	billingApiKey := &models.BillingKey{
		UserID:    userID,
		BillingID: billing.ID,
	}
	_, err = s.db.
		NewInsert().
		Model(billingApiKey).
		Exec(ctx)
	return billingApiKey, nil
}

func (s *storeImpl) BillingGetOrCreate(ctx context.Context, userID int64) (*models.Billing, error) {
	billing := new(models.Billing)
	err := s.db.NewSelect().Model((*models.Billing)(nil)).Where("user_id = ?", userID).Scan(ctx, billing)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		billing = &models.Billing{UserID: userID}
		_, err := s.db.
			NewInsert().
			Model(billing).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

		return billing, nil
	}

	return billing, nil
}

func (s *storeImpl) BillingAddRequests(ctx context.Context, userID int64, requests int) (*models.BillingPacket, error) {
	billing, err := s.BillingGetOrCreate(ctx, userID)
	if err != nil {
		return nil, err
	}

	packet := &models.BillingPacket{UserID: userID, Requests: requests}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	_, err = tx.
		NewInsert().
		Model(packet).
		Exec(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	billing.Requests += requests
	_, err = tx.
		NewUpdate().
		Model(billing).
		WherePK().
		Column("requests").
		Returning("*").
		Exec(ctx)

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return packet, nil
}

func (s *storeImpl) createBillingRequestFromRisk(userID int64, query, network string, risk *neo4jstore.Risk) *models.BillingRequest {
	req := &models.BillingRequest{
		UserID:  userID,
		Query:   query,
		Network: network,
		Last:    true,
	}

	if risk.Calculated != nil {
		req.IsCalculated = true
		req.IsReported = false
		req.IsWallet = false
		req.Risk = risk.Calculated.Risk
	}

	if risk.Wallet != nil {
		req.IsCalculated = false
		req.IsReported = false
		req.IsWallet = true
		req.Risk = risk.Wallet.Risk
	}

	if risk.Reported != nil {
		req.IsCalculated = false
		req.IsReported = true
		req.IsWallet = false
		req.Risk = risk.Reported.Risk
	}

	return req
}

func (s *storeImpl) billingFindCategory(
	ctx context.Context,
	id int64,
	data *neo4jstore.RiskData,
	isReported bool,
	isWallet bool,
) (*models.BillingRisk, error) {
	cat, err := s.CategoryFindByNumber(ctx, data.Category)
	if err != nil {
		return nil, err
	}

	r := &models.BillingRisk{
		BillingRequestID: id,
		IsReported:       isReported,
		IsWallet:         isWallet,
		Risk:             data.Risk,
		RiskRaw:          data.Risk,
		CategoryID:       &cat.Number,
	}
	if cat.CategoryGroupID != nil {
		r.DirectoryID = *cat.CategoryGroupID
	}

	return r, nil
}

func (s *storeImpl) createBillingRisks(ctx context.Context, billingRequest *models.BillingRequest, risk *neo4jstore.Risk) ([]*models.BillingRisk, error) {
	var risks []*models.BillingRisk

	if risk.Reported != nil {
		r, err := s.billingFindCategory(ctx, billingRequest.ID, risk.Reported, true, false)
		if err != nil {
			return nil, err
		}
		risks = append(risks, r)
	} else if risk.Wallet != nil {
		r, err := s.billingFindCategory(ctx, billingRequest.ID, risk.Wallet, false, true)
		if err != nil {
			return nil, err
		}
		risks = append(risks, r)
	} else if risk.Calculated != nil {
		for _, c := range risk.Calculated.Items {
			if c.PercentRaw == 0 {
				continue
			}

			risks = append(risks, &models.BillingRisk{
				BillingRequestID: billingRequest.ID,
				IsCalculated:     true,
				Risk:             c.Risk,
				RiskRaw:          c.RiskRaw,
				Percent:          c.Percent,
				PercentRaw:       c.PercentRaw,
				DirectoryID:      c.ID,
				Total:            c.Total,
			})
		}
	}

	return risks, nil
}

func (s *storeImpl) BillingStatisticsSummary(
	ctx context.Context,
	userID int64,
	from, to *time.Time,
	last *bool,
) ([]*models.StatisticsSummary, error) {
	var u []*models.StatisticsSummary
	q := s.db.
		NewSelect().
		Model((*models.BillingRequest)(nil)).
		ColumnExpr("billing_request.network AS network").
		ColumnExpr("count(billing_request.network) AS total").
		Where("billing_request.user_id = ?", userID).
		Group("billing_request.network")

	if last != nil && *last {
		q = q.Where("billing_request.last = ?", true)
	}

	if from != nil {
		q = q.Where("billing_request.created_at::date >= ?", *from)
	}
	if to != nil {
		q = q.Where("billing_request.created_at::date <= ?", *to)
	}
	if err := q.Scan(ctx, &u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *storeImpl) BillingRequestFindByID(ctx context.Context, id int64) (*models.BillingRequest, error) {
	u := new(models.BillingRequest)
	err := s.db.
		NewSelect().
		Model(u).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return nil, err
	}
	return u, nil
}

func (s *storeImpl) BillingRisks(ctx context.Context, id int64) ([]*models.BillingRisk, error) {
	req, err := s.BillingRequestFindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var u []*models.BillingRisk
	err = s.db.
		NewSelect().
		Model((*models.BillingRisk)(nil)).
		Where("billing_risk.billing_request_id = ?", id).
		Where("billing_risk.is_calculated = ?", req.IsCalculated).
		Where("billing_risk.is_wallet = ?", req.IsWallet).
		Where("billing_risk.is_reported = ?", req.IsReported).
		Order("billing_risk.percent_raw DESC").
		Scan(ctx, &u)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (s *storeImpl) BillingRegisterRequest(ctx context.Context, userID int64, query string, risk *neo4jstore.Risk, network string) (*models.Billing, error) {
	billing, err := s.BillingGetOrCreate(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "billing get or create")
	}

	if billing.Requests == 0 {
		return nil, ErrInsufficientBalance
	}

	req := s.createBillingRequestFromRisk(userID, query, network, risk)

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	_, err = s.db.NewUpdate().
		Model((*models.BillingRequest)(nil)).
		Where("user_id = ? AND query = ? AND network = ?", userID, query, network).
		Set("last = ?", false).
		Exec(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "update is last request error")
	}

	_, txErr := tx.NewInsert().Model(req).Exec(ctx)
	if txErr != nil {
		if err := tx.Rollback(); err != nil {
			return nil, errors.Wrap(err, "rollback error")
		}
		return nil, errors.Wrap(err, "tx error")
	}

	risks, err := s.createBillingRisks(ctx, req, risk)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}
	for _, r := range risks {
		_, err = tx.NewInsert().Model(r).Exec(ctx)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	billing.Requests -= 1
	_, err = tx.
		NewUpdate().
		Model(billing).
		WherePK().
		Column("requests").
		Exec(ctx)

	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "tx commit")
	}

	return billing, nil
}
