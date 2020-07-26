// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/therecipe/qt/core"
)

const (
	KeyvalKeyRole = int(core.Qt__UserRole) + 1<<iota
	KeyvalValRole
)

//go:generate qtmoc
type Keyval struct {
	core.QObject

	_ string `property:"key"`
	_ string `property:"val"`
}

//go:generate qtmoc
type KeyvalList struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ []*Keyval                `property:"list"`
	_ map[int]*core.QByteArray `property:"roles"`

	_ func(int) string  `slot:"valAt"`
	_ func()            `slot:"addEmpty"`
	_ func(int, string) `slot:"editKeyAt"`
	_ func(int, string) `slot:"editValAt"`
	_ func(int)         `slot:"removeAt"`
}

func (m *KeyvalList) init() {
	m.SetRoles(map[int]*core.QByteArray{
		int(core.Qt__DisplayRole): core.NewQByteArray2("display", -1),
		KeyvalKeyRole:             core.NewQByteArray2("role", -1),
		KeyvalValRole:             core.NewQByteArray2("val", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectValAt(m.valAt)
	m.ConnectAddEmpty(m.addEmpty)
	m.ConnectEditKeyAt(m.editKeyAt)
	m.ConnectEditValAt(m.editValAt)
	m.ConnectRemoveAt(m.removeAt)

}

func (m *KeyvalList) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	kv := m.List()[index.Row()]

	switch role {
	case int(core.Qt__DisplayRole), KeyvalKeyRole:
		return core.NewQVariant1(kv.Key())
	case KeyvalValRole:
		return core.NewQVariant1(kv.Val())
	default:
		return core.NewQVariant()
	}
}

func (m *KeyvalList) rowCount(parent *core.QModelIndex) int {
	return len(m.List())
}

func (m *KeyvalList) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *KeyvalList) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *KeyvalList) valAt(idx int) string {
	return m.List()[idx].Val()
}

func (m *KeyvalList) addEmpty() {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.List()), len(m.List()))
	m.SetList(append(m.List(), NewKeyval(nil)))
	m.EndInsertRows()
}

func (m *KeyvalList) editKeyAt(row int, key string) {
	kv := m.List()[row]
	if kv.Key() == key {
		return
	}
	kv.SetKey(key)
	idx := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(idx, idx, []int{int(core.Qt__DisplayRole), KeyvalKeyRole})
}

func (m *KeyvalList) editValAt(row int, val string) {
	kv := m.List()[row]
	if kv.Val() == val {
		return
	}
	kv.SetVal(val)
	idx := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(idx, idx, []int{KeyvalValRole})
}

func (m *KeyvalList) removeAt(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetList(append(m.List()[:row], m.List()[row+1:]...))
	m.EndRemoveRows()
}
