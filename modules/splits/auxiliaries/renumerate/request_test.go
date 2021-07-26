/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package renumerate

import (
	testBase "github.com/persistenceOne/persistenceSDK/schema/test_types/base"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_Burn_Request(t *testing.T) {

	ownerID := testBase.NewID("ownerID")
	ownableID := testBase.NewID("ownableID")
	testValue := sdkTypes.NewDec(10)
	testAuxiliaryRequest := NewAuxiliaryRequest(ownerID, ownableID, testValue)

	require.Equal(t, auxiliaryRequest{OwnerID: ownerID, OwnableID: ownableID, Value: testValue}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
