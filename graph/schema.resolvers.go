package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/suapapa/sharefit-gql-server/graph/generated"
	"github.com/suapapa/sharefit-gql-server/graph/model"
	"github.com/suapapa/sharefit-gql-server/internal/database"
)

func (r *centerResolver) Memberships(ctx context.Context, obj *model.Center) ([]*model.Membership, error) {
	var cards []database.Card
	if err := database.SharefitDB.Where("center_id = ?", obj.ID).Find(&cards).Error; err != nil {
		return nil, err
	}

	var ret []*model.Membership
	for _, c := range cards {
		ret = append(ret, &model.Membership{
			ID:       fmt.Sprint(c.ID),
			Training: c.Training,
			CurrCnt:  c.CurrCnt,
			TotalCnt: c.TotalCnt,
			Expiry:   c.Expiry,
		})
	}
	return ret, nil
}

func (r *membershipResolver) Users(ctx context.Context, obj *model.Membership) ([]*model.User, error) {
	var users []database.User
	if err := database.SharefitDB.Where("card_id = ?", obj.ID).Find(&users).Error; err != nil {
		return nil, err
	}

	var ret []*model.User
	for _, u := range users {
		ret = append(ret, &model.User{
			ID:          fmt.Sprint(u.ID),
			Name:        u.Name,
			PhoneNumber: u.PhoneNumber,
		})
	}

	return ret, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := database.User{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
	}

	database.SharefitDB.Create(&user)

	ret := model.User{
		ID:          fmt.Sprint(user.ID),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}

	return &ret, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID *string, user model.NewUser) (*model.User, error) {
	var u database.User
	if err := database.SharefitDB.Where("id = ?", userID).Find(&u).Error; err != nil {
		return nil, err
	}

	u.Name = user.Name
	u.PhoneNumber = user.PhoneNumber

	database.SharefitDB.Save(&u)

	return &model.User{
		ID:          fmt.Sprint(u.ID),
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}, nil
}

func (r *queryResolver) Memberships(ctx context.Context) ([]*model.Membership, error) {
	var cards []database.Card
	database.SharefitDB.Find(&cards)

	var ret []*model.Membership
	for _, v := range cards {
		ret = append(ret, &model.Membership{
			ID:       fmt.Sprint(v.ID),
			Training: v.Training,
			CurrCnt:  v.CurrCnt,
			TotalCnt: v.TotalCnt,
			Expiry:   v.Expiry,
			// UserIDs: ,
		})
	}

	return ret, nil
}

func (r *queryResolver) Membership(ctx context.Context, membershipID *string) (*model.Membership, error) {
	var card database.Card
	if err := database.SharefitDB.Where("id = ?", membershipID).First(&card).Error; err != nil {
		return nil, err
	}

	return &model.Membership{
		ID:       fmt.Sprint(card.ID),
		Training: card.Training,
		CurrCnt:  card.CurrCnt,
		TotalCnt: card.TotalCnt,
		Expiry:   card.Expiry,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []database.User
	database.SharefitDB.Find(&users)

	var ret []*model.User
	for _, v := range users {
		ret = append(ret, &model.User{
			Name:        v.Name,
			PhoneNumber: v.PhoneNumber,
		})
	}
	return ret, nil
}

func (r *queryResolver) User(ctx context.Context, userID *string) (*model.User, error) {
	var user database.User
	if err := database.SharefitDB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &model.User{
		ID:          fmt.Sprint(user.ID),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (r *queryResolver) Centers(ctx context.Context) ([]*model.Center, error) {
	var centers []database.Center
	if err := database.SharefitDB.Find(&centers).Error; err != nil {
		return nil, err
	}

	var ret []*model.Center
	for _, v := range centers {
		ret = append(ret, &model.Center{
			ID:          fmt.Sprint(v.ID),
			Name:        v.Name,
			PhoneNumber: v.PhoneNumber,
			// Memberships: cards,
		})
	}

	return ret, nil
}

func (r *queryResolver) Center(ctx context.Context, centerID *string) (*model.Center, error) {
	var center database.Center
	if err := database.SharefitDB.Where("id = ?", centerID).First(&center).Error; err != nil {
		return nil, err
	}

	return &model.Center{
		ID:          fmt.Sprint(center.ID),
		Name:        center.Name,
		PhoneNumber: center.PhoneNumber,
	}, nil
}

// Center returns generated.CenterResolver implementation.
func (r *Resolver) Center() generated.CenterResolver { return &centerResolver{r} }

// Membership returns generated.MembershipResolver implementation.
func (r *Resolver) Membership() generated.MembershipResolver { return &membershipResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type centerResolver struct{ *Resolver }
type membershipResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
