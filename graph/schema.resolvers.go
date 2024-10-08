package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"

	"github.com/kyomel/go-gql-blogs/graph/generated"
	"github.com/kyomel/go-gql-blogs/graph/middleware"
	"github.com/kyomel/go-gql-blogs/graph/model"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (string, error) {
	token := r.userService.Register(input)
	if token == "" {
		return "", errors.New("registration failed")
	}
	return token, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	token := r.userService.Login(input)
	if token == "" {
		return "", errors.New("login failed, invalid email or password")
	}
	return token, nil
}

// NewBlog is the resolver for the newBlog field.
func (r *mutationResolver) NewBlog(ctx context.Context, input model.NewBlog) (*model.Blog, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.Blog{}, errors.New("access denied")
	}

	blog, err := r.blogService.CreateBlog(input, *user)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

// EditBlog is the resolver for the editBlog field.
func (r *mutationResolver) EditBlog(ctx context.Context, input model.EditBlog) (*model.Blog, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.Blog{}, errors.New("access denied")
	}

	blog, err := r.blogService.EditBlog(input, *user)
	if err != nil {
		return &model.Blog{}, err
	}

	return blog, nil
}

// DeleteBlog is the resolver for the deleteBlog field.
func (r *mutationResolver) DeleteBlog(ctx context.Context, input model.DeleteBlog) (bool, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return false, errors.New("access denied")
	}

	result, err := r.blogService.DeleteBlog(input, *user)
	if err != nil {
		return false, err
	}

	return result, nil
}

// Blogs is the resolver for the blogs field.
func (r *queryResolver) Blogs(ctx context.Context) ([]*model.Blog, error) {
	blogs := r.blogService.GetAllBlogs()
	return blogs, nil
}

// Blog is the resolver for the blog field.
func (r *queryResolver) Blog(ctx context.Context, id string) (*model.Blog, error) {
	blog, err := r.blogService.GetBlogByID(id)
	if err != nil {
		return &model.Blog{}, err
	}
	return blog, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
