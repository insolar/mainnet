// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/insolar/blob/master/LICENSE.md.

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package member

import (
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/logicrunner/common"
	"github.com/insolar/mainnet/application/appfoundation"
	"github.com/pkg/errors"
)

const PanicIsLogicalError = false

func INS_META_INFO() []map[string]string {
	result := make([]map[string]string, 0)

	{
		info := make(map[string]string, 3)
		info["Type"] = "SagaInfo"
		info["MethodName"] = "Accept"
		info["RollbackMethodName"] = "INS_FLAG_NO_ROLLBACK_METHOD"
		result = append(result, info)
	}

	{
		info := make(map[string]string, 3)
		info["Type"] = "SagaInfo"
		info["MethodName"] = "AcceptBurn"
		info["RollbackMethodName"] = "INS_FLAG_NO_ROLLBACK_METHOD"
		result = append(result, info)
	}

	return result
}

func INSMETHOD_GetCode(object []byte, data []byte) ([]byte, []byte, error) {
	ph := common.CurrentProxyCtx
	self := new(Member)

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
	self := new(Member)

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

func INSMETHOD_GetWallet(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Member)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetWallet ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetWallet ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetWallet ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
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

	ret0, ret1 = self.GetWallet()

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

func INSMETHOD_GetAccount(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Member)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetAccount ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetAccount ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetAccount ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
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

	ret0, ret1 = self.GetAccount(args0)

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

func INSMETHOD_Call(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Member)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeCall ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeCall ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 []byte
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeCall ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
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

	ret0, ret1 = self.Call(args0)

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

func INSMETHOD_GetMigrationAddress(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Member)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetMigrationAddress ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
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

	ret0, ret1 = self.GetMigrationAddress()

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

func INSMETHOD_Accept(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Member)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeAccept ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAccept ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 appfoundation.SagaAcceptInfo
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAccept ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
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

	ret0 = self.Accept(args0)

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

func INSMETHOD_AcceptBurn(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Member)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeAcceptBurn ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAcceptBurn ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 appfoundation.SagaAcceptInfo
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAcceptBurn ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
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

	ret0 = self.AcceptBurn(args0)

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

func INSCONSTRUCTOR_New(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := make([]interface{}, 3)
	var args0 string
	args[0] = &args0
	var args1 string
	args[1] = &args1
	var args2 insolar.Reference
	args[2] = &args2

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNew ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Member
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute constructor (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				err = serializeResults()
				if err == nil {
					state = data
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = New(args0, args1, args2)

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func Initialize() insolar.ContractWrapper {
	return insolar.ContractWrapper{
		Methods: insolar.ContractMethods{
			"GetWallet":           INSMETHOD_GetWallet,
			"GetAccount":          INSMETHOD_GetAccount,
			"Call":                INSMETHOD_Call,
			"GetMigrationAddress": INSMETHOD_GetMigrationAddress,
			"Accept":              INSMETHOD_Accept,
			"AcceptBurn":          INSMETHOD_AcceptBurn,

			"GetCode":      INSMETHOD_GetCode,
			"GetPrototype": INSMETHOD_GetPrototype,
		},
		Constructors: insolar.ContractConstructors{
			"New": INSCONSTRUCTOR_New,
		},
	}
}
