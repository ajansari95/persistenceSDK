/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type ID struct {
	IDString string `json:"idString"`
}

var _ types.ID = (*ID)(nil)

func (id ID) String() string {
	return id.IDString

}
func (id ID) Bytes() []byte {
	return []byte(id.IDString)
}
func (id ID) Equals(compareID types.ID) bool {
	return bytes.Equal(id.Bytes(), compareID.Bytes())
}

func NewID(idString string) types.ID {
	return ID{IDString: idString}
}


