// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/insolar/blob/master/LICENSE.md.

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package migrationadmin

import (
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/logicrunner/common"
	"github.com/pkg/errors"
)

const PanicIsLogicalError = false

func INS_META_INFO() []map[string]string {
	result := make([]map[string]string, 0)

	return result
}

func INSMETHOD_GetCode(object []byte, data []byte) ([]byte, []byte, error) {
	ph := common.CurrentProxyCtx
	self := new(MigrationAdmin)

	if len(object) == 0 {
		return nil, nil, &foundation.Error{S: "[ Fake GetCode ] ( Generated Method ) Object is nil"}
	}

	err := ph.Deserialize(object, self)
	if err != nil {
		e := &foundation.Error{S: "[ Fake GetCode ] ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return nil, nil, e
	}

	state := []byte{}
	err = ph.Serialize(self, &state)
	if err != nil {
		return nil, nil, err
	}

	ret := []byte{}
	err = ph.Serialize([]interface{}{self.GetCode().Bytes()}, &ret)

	return state, ret, err
}

func INSMETHOD_GetPrototype(object []byte, data []byte) ([]byte, []byte, error) {
	ph := common.CurrentProxyCtx
	self := new(MigrationAdmin)

	if len(object) == 0 {
		return nil, nil, &foundation.Error{S: "[ Fake GetPrototype ] ( Generated Method ) Object is nil"}
	}

	err := ph.Deserialize(object, self)
	if err != nil {
		e := &foundation.Error{S: "[ Fake GetPrototype ] ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return nil, nil, e
	}

	state := []byte{}
	err = ph.Serialize(self, &state)
	if err != nil {
		return nil, nil, err
	}

	ret := []byte{}
	err = ph.Serialize([]interface{}{self.GetPrototype().Bytes()}, &ret)

	return state, ret, err
}

func INSMETHOD_MigrationAdminCall(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(MigrationAdmin)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeMigrationAdminCall ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeMigrationAdminCall ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 3)
	var args0 map[string]interface{}
	args[0] = &args0
	var args1 string
	args[1] = &args1
	var args2 insolar.Reference
	args[2] = &args2

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeMigrationAdminCall ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 interface{}
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.MigrationAdminCall(args0, args1, args2)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetDepositParameters(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(MigrationAdmin)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetDepositParameters ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetDepositParameters ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetDepositParameters ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *VestingParams
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetDepositParameters()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetMigrationDaemonByMemberRef(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(MigrationAdmin)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetMigrationDaemonByMemberRef ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetMigrationDaemonByMemberRef ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetMigrationDaemonByMemberRef ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 insolar.Reference
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetMigrationDaemonByMemberRef(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetMemberByMigrationAddress(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(MigrationAdmin)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetMemberByMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetMemberByMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetMemberByMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *insolar.Reference
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetMemberByMigrationAddress(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetFreeMigrationAddress(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(MigrationAdmin)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetFreeMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetFreeMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetFreeMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 string
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetFreeMigrationAddress(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_AddNewMigrationAddressToMaps(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(MigrationAdmin)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeAddNewMigrationAddressToMaps ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAddNewMigrationAddressToMaps ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 2)
	var args0 string
	args[0] = &args0
	var args1 insolar.Reference
	args[1] = &args1

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAddNewMigrationAddressToMaps ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret0 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0 = self.AddNewMigrationAddressToMaps(args0, args1)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret0 = ph.MakeErrorSerializable(ret0)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func Initialize() insolar.ContractWrapper {
	return insolar.ContractWrapper{
		GetCode:      INSMETHOD_GetCode,
		GetPrototype: INSMETHOD_GetPrototype,
		Methods: insolar.ContractMethods{
			"MigrationAdminCall":            INSMETHOD_MigrationAdminCall,
			"GetDepositParameters":          INSMETHOD_GetDepositParameters,
			"GetMigrationDaemonByMemberRef": INSMETHOD_GetMigrationDaemonByMemberRef,
			"GetMemberByMigrationAddress":   INSMETHOD_GetMemberByMigrationAddress,
			"GetFreeMigrationAddress":       INSMETHOD_GetFreeMigrationAddress,
			"AddNewMigrationAddressToMaps":  INSMETHOD_AddNewMigrationAddressToMaps,
		},
		Constructors: insolar.ContractConstructors{},
	}
}
