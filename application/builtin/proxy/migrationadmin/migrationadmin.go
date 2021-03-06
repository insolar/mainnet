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
)

type AddMaResponse struct {
	Count int `json:"totalNumber"`
}
type CheckDaemonResponse struct {
	Status string `json:"status"`
}
type GetAddressCountResponse struct {
	ShardIndex int `json:"shardIndex"`
	FreeCount  int `json:"freeCount"`
}
type VestingParams struct {
	Lockup      int64 `json:"lockupInPulses"`
	Vesting     int64 `json:"vestingInPulses"`
	VestingStep int64 `json:"vestingStepInPulses"`
}

// PrototypeReference to prototype of this contract
// error checking hides in generator
var PrototypeReference, _ = insolar.NewObjectReferenceFromString("insolar:0AAABAv4b40_lF0ivLCNhzPcq1hKkHWpRSaZCfZuPDUU")

// MigrationAdmin holds proxy type
type MigrationAdmin struct {
	Reference insolar.Reference
	Prototype insolar.Reference
	Code      insolar.Reference
}

// ContractConstructorHolder holds logic with object construction
type ContractConstructorHolder struct {
	constructorName string
	argsSerialized  []byte
}

// AsChild saves object as child
func (r *ContractConstructorHolder) AsChild(objRef insolar.Reference) (*MigrationAdmin, error) {
	ret, err := common.CurrentProxyCtx.SaveAsChild(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}

	var ref insolar.Reference
	var constructorError *foundation.Error
	resultContainer := foundation.Result{
		Returns: []interface{}{&ref, &constructorError},
	}
	err = common.CurrentProxyCtx.Deserialize(ret, &resultContainer)
	if err != nil {
		return nil, err
	}

	if resultContainer.Error != nil {
		return nil, resultContainer.Error
	}

	if constructorError != nil {
		return nil, constructorError
	}

	return &MigrationAdmin{Reference: ref}, nil
}

// GetObject returns proxy object
func GetObject(ref insolar.Reference) *MigrationAdmin {
	if !ref.IsObjectReference() {
		return nil
	}
	return &MigrationAdmin{Reference: ref}
}

// GetPrototype returns reference to the prototype
func GetPrototype() insolar.Reference {
	return *PrototypeReference
}

// GetReference returns reference of the object
func (r *MigrationAdmin) GetReference() insolar.Reference {
	return r.Reference
}

// GetPrototype returns reference to the code
func (r *MigrationAdmin) GetPrototype() (insolar.Reference, error) {
	if r.Prototype.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 insolar.Reference
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetPrototype", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = common.CurrentProxyCtx.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Prototype = ret0
	}

	return r.Prototype, nil

}

// GetCode returns reference to the code
func (r *MigrationAdmin) GetCode() (insolar.Reference, error) {
	if r.Code.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 insolar.Reference
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetCode", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = common.CurrentProxyCtx.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Code = ret0
	}

	return r.Code, nil
}

// MigrationAdminCall is proxy generated method
func (r *MigrationAdmin) MigrationAdminCall(params map[string]interface{}, nameMethod string, caller insolar.Reference) (interface{}, error) {
	var args [3]interface{}
	args[0] = params
	args[1] = nameMethod
	args[2] = caller

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 interface{}
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "MigrationAdminCall", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// MigrationAdminCallAsImmutable is proxy generated method
func (r *MigrationAdmin) MigrationAdminCallAsImmutable(params map[string]interface{}, nameMethod string, caller insolar.Reference) (interface{}, error) {
	var args [3]interface{}
	args[0] = params
	args[1] = nameMethod
	args[2] = caller

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 interface{}
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "MigrationAdminCall", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetDepositParameters is proxy generated method
func (r *MigrationAdmin) GetDepositParameters() (*VestingParams, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 *VestingParams
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetDepositParameters", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetDepositParametersAsImmutable is proxy generated method
func (r *MigrationAdmin) GetDepositParametersAsImmutable() (*VestingParams, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 *VestingParams
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetDepositParameters", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetMigrationDaemonByMemberRef is proxy generated method
func (r *MigrationAdmin) GetMigrationDaemonByMemberRefAsMutable(memberRef string) (insolar.Reference, error) {
	var args [1]interface{}
	args[0] = memberRef

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 insolar.Reference
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetMigrationDaemonByMemberRef", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetMigrationDaemonByMemberRefAsImmutable is proxy generated method
func (r *MigrationAdmin) GetMigrationDaemonByMemberRef(memberRef string) (insolar.Reference, error) {
	var args [1]interface{}
	args[0] = memberRef

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 insolar.Reference
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetMigrationDaemonByMemberRef", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetMemberByMigrationAddress is proxy generated method
func (r *MigrationAdmin) GetMemberByMigrationAddressAsMutable(migrationAddress string) (*insolar.Reference, error) {
	var args [1]interface{}
	args[0] = migrationAddress

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 *insolar.Reference
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetMemberByMigrationAddress", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetMemberByMigrationAddressAsImmutable is proxy generated method
func (r *MigrationAdmin) GetMemberByMigrationAddress(migrationAddress string) (*insolar.Reference, error) {
	var args [1]interface{}
	args[0] = migrationAddress

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 *insolar.Reference
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetMemberByMigrationAddress", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetFreeMigrationAddress is proxy generated method
func (r *MigrationAdmin) GetFreeMigrationAddressAsMutable(publicKey string) (string, error) {
	var args [1]interface{}
	args[0] = publicKey

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetFreeMigrationAddress", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetFreeMigrationAddressAsImmutable is proxy generated method
func (r *MigrationAdmin) GetFreeMigrationAddress(publicKey string) (string, error) {
	var args [1]interface{}
	args[0] = publicKey

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetFreeMigrationAddress", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// AddNewMigrationAddressToMaps is proxy generated method
func (r *MigrationAdmin) AddNewMigrationAddressToMaps(migrationAddress string, memberRef insolar.Reference) error {
	var args [2]interface{}
	args[0] = migrationAddress
	args[1] = memberRef

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "AddNewMigrationAddressToMaps", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// AddNewMigrationAddressToMapsAsImmutable is proxy generated method
func (r *MigrationAdmin) AddNewMigrationAddressToMapsAsImmutable(migrationAddress string, memberRef insolar.Reference) error {
	var args [2]interface{}
	args[0] = migrationAddress
	args[1] = memberRef

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "AddNewMigrationAddressToMaps", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}
