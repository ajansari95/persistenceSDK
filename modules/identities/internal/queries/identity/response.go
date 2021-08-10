/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
)


var _ helpers.QueryResponse = (*queryResponse)(nil)

func (queryResponse queryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse queryResponse) GetError() error {
	return queryResponse.Error
}
func (queryResponse queryResponse) Encode() ([]byte, error) {
	return common.Codec.MarshalJSON(queryResponse)
}
func (queryResponse queryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if Error := common.Codec.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return queryResponse{}
}
func newQueryResponse(collection helpers.Collection, error protoTypes.Error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}

	return queryResponse{
		Success: success,
		Error:   error,
		List:    collection.GetList(),
	}
}
