/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	base2 "github.com/persistenceOne/persistenceSDK/schema/proto/types/base"
	tendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

func CreateTestInput2(t *testing.T) (sdkTypes.Context, helpers.Keeper) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	std.RegisterLegacyAminoCodec(Codec)
	std.RegisterInterfaces(interfaceRegistry)

	encodingConfig := applications.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             Codec,
	}


	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := paramsKeeper.NewKeeper(
		encodingConfig.Marshaler,
		encodingConfig.Amino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, tendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	testQueryKeeper := keeperPrototype().Initialize(mapper, Parameters, []interface{}{})

	return context, testQueryKeeper
}

func Test_Query_Keeper_Identity(t *testing.T) {

	context, keepers := CreateTestInput2(t)
	immutableProperties, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableProperties, Error2 := base.ReadProperties("burn:S|100")
	require.Equal(t, nil, Error2)

	classificationID := base2.NewID("ClassificationID")
	identityID := key.NewIdentityID(classificationID, immutableProperties)
	keepers.(queryKeeper).mapper.NewCollection(context).Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	testQueryRequest := newQueryRequest(classificationID)
	require.Equal(t, queryResponse{Success: true, Error: nil, List: keepers.(queryKeeper).mapper.NewCollection(context).Fetch(key.FromID(identityID)).GetList()}, keepers.(queryKeeper).Enquire(context, testQueryRequest))

}
