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

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := database.User{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
	}

	database.SharefitDB.Create(&user)
	fmt.Println("user.ID", user.ID)

	ret := model.User{
		ID:          fmt.Sprint(user.ID),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}

	return &ret, nil
}

func (r *queryResolver) Memberships(ctx context.Context) ([]*model.Membership, error) {
	var cards []database.Card
	database.SharefitDB.Find(&cards)

	var ret []*model.Membership
	for _, v := range cards {
		database.SharefitDB.Where("card_id = ?", v.ID).Find(&v.Users)
		var users []*model.User
		for _, u := range v.Users {
			users = append(users, &model.User{
				Name:        u.Name,
				PhoneNumber: u.PhoneNumber,
			})
		}

		ret = append(ret, &model.Membership{
			Training: v.Training,
			CurrCnt:  v.CurrCnt,
			TotalCnt: v.TotalCnt,
			Expiry:   v.Expiry.String(),
			Users:    users,
		})
	}

	return ret, nil
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
	database.SharefitDB.Where("id = ?", userID).First(&user)

	return &model.User{
		ID:          fmt.Sprint(user.ID),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (r *queryResolver) Centers(ctx context.Context) ([]*model.Center, error) {
	var centers []database.Center
	database.SharefitDB.Find(&centers)

	var ret []*model.Center
	for _, v := range centers {
		database.SharefitDB.Where("center_id = ?", v.ID).Find(&v.Cards)
		var cards []*model.Membership
		for _, c := range v.Cards {
			cards = append(cards, &model.Membership{
				Training: c.Training,
				CurrCnt:  c.CurrCnt,
				TotalCnt: c.TotalCnt,
				// TODO: ???? should retrive users ????
			})
		}

		ret = append(ret, &model.Center{
			Name:        v.Name,
			PhoneNumber: v.PhoneNumber,
			Memberships: cards,
		})
	}

	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
