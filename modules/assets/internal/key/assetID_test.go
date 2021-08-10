/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	"strings"
	"testing"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_AssetID_Methods(t *testing.T) {
	classificationID := base2.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base2.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))

	testAssetID := NewAssetID(classificationID, immutableProperties).(assetID)
	require.NotPanics(t, func() {
		require.Equal(t, assetID{ClassificationID: classificationID, HashID: baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID()}, testAssetID)
		require.Equal(t, strings.Join([]string{classificationID.String(), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String()}, constants.FirstOrderCompositeIDSeparator), testAssetID.String())
		require.Equal(t, false, testAssetID.IsPartial())
		require.Equal(t, true, assetID{ClassificationID: classificationID, HashID: base2.NewID("")}.IsPartial())
		require.Equal(t, true, testAssetID.Equals(testAssetID))
		require.Equal(t, false, testAssetID.Equals(assetID{ClassificationID: classificationID, HashID: base2.NewID("")}))
		require.Equal(t, true, testAssetID.Matches(testAssetID))
		require.Equal(t, false, testAssetID.Matches(nil))
		require.Equal(t, false, testAssetID.Matches(assetID{ClassificationID: classificationID, HashID: base2.NewID("")}))
		require.Equal(t, testAssetID, FromID(testAssetID))
		require.Equal(t, assetID{ClassificationID: base2.NewID(""), HashID: base2.NewID("")}, FromID(base2.NewID("")))
		require.Equal(t, testAssetID, readAssetID(testAssetID.String()))
	})

}
