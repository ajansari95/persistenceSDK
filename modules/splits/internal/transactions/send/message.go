 /*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

 import (
	 "encoding/json"
	 "github.com/asaskevich/govalidator"
	 "github.com/cosmos/cosmos-sdk/codec"
	 sdkTypes "github.com/cosmos/cosmos-sdk/types"
	 "github.com/cosmos/cosmos-sdk/types/errors"
	 xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	 "github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	 "github.com/persistenceOne/persistenceSDK/schema/test_types"

	 //"github.com/persistenceOne/persistenceSDK/schema/types"
	 codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
 )



var _ sdkTypes.Msg = (*Message)(nil)

func (message Message) Route() string { return module.Name }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}

	return nil
}
func (message Message) GetSignBytes() []byte {
	js,_ := json.Marshal(message)

	return sdkTypes.MustSortJSON(js)
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{sdkTypes.AccAddress(message.From)}
}
func (Message) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, Message{})
}
func messageFromInterface(msg sdkTypes.Msg) *Message {
	switch value := msg.(type) {
	case *Message:
		return value
	default:
		return &Message{}
	}
}
func messagePrototype() Message {
	return Message{}
}

//TODO:types mismatch

func newMessage(from sdkTypes.AccAddress, fromID test_types.ID, toID test_types.ID, ownableID test_types.ID, value sdkTypes.Dec) *Message {
	return &Message{
		From: from,
		FromID:    fromID,
		ToID:      toID,
		OwnableID: ownableID,
		Value:     value,
	}
}
