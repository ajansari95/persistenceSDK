/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package block

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type block struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
}

var _ helpers.Block = (*block)(nil)

func (block block) Begin(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {

}

func (block block) End(context sdkTypes.Context, endBlockRequest abciTypes.RequestEndBlock) {
	orders := block.mapper.NewCollection(context)

	orders.Iterate(
		key.New(base.NewID("")),
		func(order helpers.Mappable) bool {
			metaProperties, Error := supplement.GetMetaPropertiesFromResponse(block.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetExpiry(), order.(mappables.Order).GetMakerOwnableSplit())))
			if Error != nil {
				panic(Error)
			}

			if expiryProperty := metaProperties.GetMetaProperty(base.NewID(properties.Expiry)); expiryProperty != nil {
				expiry, Error := expiryProperty.GetMetaFact().GetData().AsHeight()
				if Error != nil {
					panic(Error)
				} else if !expiry.IsGreaterThan(base.NewHeight(endBlockRequest.Height)) {
					makerOwnableSplitProperty := metaProperties.GetMetaProperty(base.NewID(properties.MakerOwnableSplit))
					if makerOwnableSplitProperty == nil {
						panic(errors.MetaDataError)
					}

					makerOwnableSplit, Error := makerOwnableSplitProperty.GetMetaFact().GetData().AsDec()
					if Error != nil {
						panic(Error)
					}

					if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), order.(mappables.Order).GetMakerID(), order.(mappables.Order).GetMakerOwnableID(), makerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}

					orders.Remove(order)
				}
			}

			return false
		},
	)
}

func (block block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Block {
	block.mapper, block.parameters = mapper, parameters

	for _, auxiliaryKeeper := range auxiliaryKeepers {
		switch value := auxiliaryKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				block.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				block.transferAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return block
}
