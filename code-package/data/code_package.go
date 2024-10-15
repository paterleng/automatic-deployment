package data

import (
	"code-package/data/schema"
	"context"
)

type CodePackageRepo interface {
	BaseRepo

	Create(ctx context.Context, plan schema.GetAndPushPlan) error
	SelectPlanById(ctx context.Context, id int64) (schema.GetAndPushPlan, error)
	UpdatePlanStatus(ctx context.Context, id, status int64) error
}

type codePackageRepo struct {
	*Data
}

func (r *codePackageRepo) Create(ctx context.Context, plan schema.GetAndPushPlan) error {
	return r.DB(ctx).Create(&plan).Error
}

func (r *codePackageRepo) SelectPlanById(ctx context.Context, id int64) (schema.GetAndPushPlan, error) {
	req := schema.GetAndPushPlan{}
	err := r.DB(ctx).Where("id = ?", id).Find(&req).Error
	if err != nil {
		return req, nil
	}
	return req, nil
}

func (r *codePackageRepo) UpdatePlanStatus(ctx context.Context, id, status int64) error {
	return r.DB(ctx).Update("status", status).Where("id = ?", id).Error
}

func NewCodePackageRepo(data *Data) CodePackageRepo {
	repo := &codePackageRepo{
		Data: data,
	}
	return repo
}
