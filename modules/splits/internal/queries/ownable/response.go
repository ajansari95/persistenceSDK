/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package ownable

import (
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse QueryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse QueryResponse) GetError() error {
	return fmt.Errorf(queryResponse.Error)
}
func (queryResponse QueryResponse) LegacyAminoEncode() ([]byte, error) {
	return common.LegacyAminoCodec.MarshalJSON(queryResponse)
}
func (queryResponse QueryResponse) LegacyAminoDecode(bytes []byte) (helpers.QueryResponse, error) {
	if Error := common.LegacyAminoCodec.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return QueryResponse{}
}
func newQueryResponse(value sdkTypes.Dec, error error) QueryResponse {
	success := true
	if error != nil {
		success = false
	}

	return QueryResponse{
		Success: success,
		Error:   error.Error(),
		Value:   value,
	}
}