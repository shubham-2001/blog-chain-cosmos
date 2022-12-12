package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Post
	var post = types.Post{
		Author:  msg.Author,
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// Add a post to the store and get back the ID
	id := k.AppendPost(ctx, post)

	// Return the ID of the post
	return &types.MsgCreatePostResponse{Id: id}, nil
}
