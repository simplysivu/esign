// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package signinggroups implements the DocuSign SDK
// category SigningGroups.
//
// Use the SigningGroup category to manage signing groups that allow you anyone in the group to sign a document.
//
// The category allows you create the signing group and manage the users in the group.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/v2/reference/SigningGroups
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/v2/signinggroups"
//       "github.com/jfcote87/esign/v2/model"
//   )
//   ...
//   signinggroupsService := signinggroups.New(esignCredential)
package signinggroups // import "github.com/jfcote87/esign/v2/signinggroups"

import (
	"context"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2/model"
)

// Service implements DocuSign SigningGroups Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a signinggroups service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// UsersDelete deletes  one or more members from a signing group.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroupusers/delete
//
// SDK Method SigningGroups::deleteUsers
func (s *Service) UsersDelete(signingGroupID string, signingGroupUsers *model.SigningGroupUsers) *UsersDeleteOp {
	return &UsersDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"signing_groups", signingGroupID, "users"}, "/"),
		Payload:    signingGroupUsers,
		QueryOpts:  make(url.Values),
	}
}

// UsersDeleteOp implements DocuSign API SDK SigningGroups::deleteUsers
type UsersDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UsersDeleteOp) Do(ctx context.Context) (*model.SigningGroupUsers, error) {
	var res *model.SigningGroupUsers
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UsersList gets a list of members in a Signing Group.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroupusers/list
//
// SDK Method SigningGroups::listUsers
func (s *Service) UsersList(signingGroupID string) *UsersListOp {
	return &UsersListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"signing_groups", signingGroupID, "users"}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// UsersListOp implements DocuSign API SDK SigningGroups::listUsers
type UsersListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UsersListOp) Do(ctx context.Context) (*model.SigningGroupUsers, error) {
	var res *model.SigningGroupUsers
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UsersUpdate adds members to a signing group.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroupusers/update
//
// SDK Method SigningGroups::updateUsers
func (s *Service) UsersUpdate(signingGroupID string, signingGroupUsers *model.SigningGroupUsers) *UsersUpdateOp {
	return &UsersUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"signing_groups", signingGroupID, "users"}, "/"),
		Payload:    signingGroupUsers,
		QueryOpts:  make(url.Values),
	}
}

// UsersUpdateOp implements DocuSign API SDK SigningGroups::updateUsers
type UsersUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UsersUpdateOp) Do(ctx context.Context) (*model.SigningGroupUsers, error) {
	var res *model.SigningGroupUsers
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Create creates a signing group.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroups/create
//
// SDK Method SigningGroups::createList
func (s *Service) Create(signingGroupInformation *model.SigningGroupInformation) *CreateOp {
	return &CreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "signing_groups",
		Payload:    signingGroupInformation,
		QueryOpts:  make(url.Values),
	}
}

// CreateOp implements DocuSign API SDK SigningGroups::createList
type CreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateOp) Do(ctx context.Context) (*model.SigningGroupInformation, error) {
	var res *model.SigningGroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Delete deletes one or more signing groups.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroups/delete
//
// SDK Method SigningGroups::deleteList
func (s *Service) Delete(signingGroupInformation *model.SigningGroupInformation) *DeleteOp {
	return &DeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "signing_groups",
		Payload:    signingGroupInformation,
		QueryOpts:  make(url.Values),
	}
}

// DeleteOp implements DocuSign API SDK SigningGroups::deleteList
type DeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteOp) Do(ctx context.Context) (*model.SigningGroupInformation, error) {
	var res *model.SigningGroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Get gets information about a signing group.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroups/get
//
// SDK Method SigningGroups::get
func (s *Service) Get(signingGroupID string) *GetOp {
	return &GetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"signing_groups", signingGroupID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// GetOp implements DocuSign API SDK SigningGroups::get
type GetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOp) Do(ctx context.Context) (*model.SigningGroup, error) {
	var res *model.SigningGroup
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// List gets a list of the Signing Groups in an account.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroups/list
//
// SDK Method SigningGroups::list
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "signing_groups",
		QueryOpts:  make(url.Values),
	}
}

// ListOp implements DocuSign API SDK SigningGroups::list
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.SigningGroupInformation, error) {
	var res *model.SigningGroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupType set the call query parameter group_type
func (op *ListOp) GroupType(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("group_type", val)
	}
	return op
}

// IncludeUsers when set to **true**, the response includes the signing group members.
func (op *ListOp) IncludeUsers() *ListOp {
	if op != nil {
		op.QueryOpts.Set("include_users", "true")
	}
	return op
}

// Update updates a signing group.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroups/update
//
// SDK Method SigningGroups::update
func (s *Service) Update(signingGroupID string, signingGroups *model.SigningGroup) *UpdateOp {
	return &UpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"signing_groups", signingGroupID}, "/"),
		Payload:    signingGroups,
		QueryOpts:  make(url.Values),
	}
}

// UpdateOp implements DocuSign API SDK SigningGroups::update
type UpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateOp) Do(ctx context.Context) (*model.SigningGroup, error) {
	var res *model.SigningGroup
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UpdateList updates signing group names.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/signinggroups/signinggroups/updatelist
//
// SDK Method SigningGroups::updateList
func (s *Service) UpdateList(signingGroupInformation *model.SigningGroupInformation) *UpdateListOp {
	return &UpdateListOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "signing_groups",
		Payload:    signingGroupInformation,
		QueryOpts:  make(url.Values),
	}
}

// UpdateListOp implements DocuSign API SDK SigningGroups::updateList
type UpdateListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateListOp) Do(ctx context.Context) (*model.SigningGroupInformation, error) {
	var res *model.SigningGroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
