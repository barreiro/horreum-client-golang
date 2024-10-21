package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UserApikeyApikeyItemRequestBuilder builds and executes requests for operations under \api\user\apikey\{id}
type UserApikeyApikeyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewUserApikeyApikeyItemRequestBuilderInternal instantiates a new UserApikeyApikeyItemRequestBuilder and sets the default values.
func NewUserApikeyApikeyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyApikeyItemRequestBuilder) {
    m := &UserApikeyApikeyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/user/apikey/{id}", pathParameters),
    }
    return m
}
// NewUserApikeyApikeyItemRequestBuilder instantiates a new UserApikeyApikeyItemRequestBuilder and sets the default values.
func NewUserApikeyApikeyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserApikeyApikeyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserApikeyApikeyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Rename the rename property
// returns a *UserApikeyItemRenameRequestBuilder when successful
func (m *UserApikeyApikeyItemRequestBuilder) Rename()(*UserApikeyItemRenameRequestBuilder) {
    return NewUserApikeyItemRenameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Revoke the revoke property
// returns a *UserApikeyItemRevokeRequestBuilder when successful
func (m *UserApikeyApikeyItemRequestBuilder) Revoke()(*UserApikeyItemRevokeRequestBuilder) {
    return NewUserApikeyItemRevokeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
