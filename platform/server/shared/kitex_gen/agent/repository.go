// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package agent

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type AddRepositoryReq struct {
	RepositoryType int32  `thrift:"repository_type,1" frugal:"1,default,i32" json:"repository_type"`
	RepositoryUrl  string `thrift:"repository_url,2" frugal:"2,default,string" json:"repository_url"`
	Token          string `thrift:"token,3" frugal:"3,default,string" json:"token"`
}

func NewAddRepositoryReq() *AddRepositoryReq {
	return &AddRepositoryReq{}
}

func (p *AddRepositoryReq) InitDefault() {
	*p = AddRepositoryReq{}
}

func (p *AddRepositoryReq) GetRepositoryType() (v int32) {
	return p.RepositoryType
}

func (p *AddRepositoryReq) GetRepositoryUrl() (v string) {
	return p.RepositoryUrl
}

func (p *AddRepositoryReq) GetToken() (v string) {
	return p.Token
}
func (p *AddRepositoryReq) SetRepositoryType(val int32) {
	p.RepositoryType = val
}
func (p *AddRepositoryReq) SetRepositoryUrl(val string) {
	p.RepositoryUrl = val
}
func (p *AddRepositoryReq) SetToken(val string) {
	p.Token = val
}

var fieldIDToName_AddRepositoryReq = map[int16]string{
	1: "repository_type",
	2: "repository_url",
	3: "token",
}

func (p *AddRepositoryReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AddRepositoryReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AddRepositoryReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.RepositoryType = v
	}
	return nil
}

func (p *AddRepositoryReq) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.RepositoryUrl = v
	}
	return nil
}

func (p *AddRepositoryReq) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Token = v
	}
	return nil
}

func (p *AddRepositoryReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("AddRepositoryReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AddRepositoryReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("repository_type", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.RepositoryType); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *AddRepositoryReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("repository_url", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.RepositoryUrl); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *AddRepositoryReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("token", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Token); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *AddRepositoryReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddRepositoryReq(%+v)", *p)
}

func (p *AddRepositoryReq) DeepEqual(ano *AddRepositoryReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.RepositoryType) {
		return false
	}
	if !p.Field2DeepEqual(ano.RepositoryUrl) {
		return false
	}
	if !p.Field3DeepEqual(ano.Token) {
		return false
	}
	return true
}

func (p *AddRepositoryReq) Field1DeepEqual(src int32) bool {

	if p.RepositoryType != src {
		return false
	}
	return true
}
func (p *AddRepositoryReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.RepositoryUrl, src) != 0 {
		return false
	}
	return true
}
func (p *AddRepositoryReq) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Token, src) != 0 {
		return false
	}
	return true
}

type AddRepositoryRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewAddRepositoryRes() *AddRepositoryRes {
	return &AddRepositoryRes{}
}

func (p *AddRepositoryRes) InitDefault() {
	*p = AddRepositoryRes{}
}

func (p *AddRepositoryRes) GetCode() (v int32) {
	return p.Code
}

func (p *AddRepositoryRes) GetMsg() (v string) {
	return p.Msg
}
func (p *AddRepositoryRes) SetCode(val int32) {
	p.Code = val
}
func (p *AddRepositoryRes) SetMsg(val string) {
	p.Msg = val
}

var fieldIDToName_AddRepositoryRes = map[int16]string{
	1: "code",
	2: "msg",
}

func (p *AddRepositoryRes) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AddRepositoryRes[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AddRepositoryRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}

func (p *AddRepositoryRes) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}

func (p *AddRepositoryRes) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("AddRepositoryRes"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AddRepositoryRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *AddRepositoryRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *AddRepositoryRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddRepositoryRes(%+v)", *p)
}

func (p *AddRepositoryRes) DeepEqual(ano *AddRepositoryRes) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Code) {
		return false
	}
	if !p.Field2DeepEqual(ano.Msg) {
		return false
	}
	return true
}

func (p *AddRepositoryRes) Field1DeepEqual(src int32) bool {

	if p.Code != src {
		return false
	}
	return true
}
func (p *AddRepositoryRes) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Msg, src) != 0 {
		return false
	}
	return true
}

type SyncRepositoryByIdReq struct {
	Ids []int64 `thrift:"ids,1" frugal:"1,default,list<i64>" json:"ids"`
}

func NewSyncRepositoryByIdReq() *SyncRepositoryByIdReq {
	return &SyncRepositoryByIdReq{}
}

func (p *SyncRepositoryByIdReq) InitDefault() {
	*p = SyncRepositoryByIdReq{}
}

func (p *SyncRepositoryByIdReq) GetIds() (v []int64) {
	return p.Ids
}
func (p *SyncRepositoryByIdReq) SetIds(val []int64) {
	p.Ids = val
}

var fieldIDToName_SyncRepositoryByIdReq = map[int16]string{
	1: "ids",
}

func (p *SyncRepositoryByIdReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SyncRepositoryByIdReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SyncRepositoryByIdReq) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Ids = make([]int64, 0, size)
	for i := 0; i < size; i++ {
		var _elem int64
		if v, err := iprot.ReadI64(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.Ids = append(p.Ids, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *SyncRepositoryByIdReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SyncRepositoryByIdReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SyncRepositoryByIdReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ids", thrift.LIST, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.I64, len(p.Ids)); err != nil {
		return err
	}
	for _, v := range p.Ids {
		if err := oprot.WriteI64(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SyncRepositoryByIdReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SyncRepositoryByIdReq(%+v)", *p)
}

func (p *SyncRepositoryByIdReq) DeepEqual(ano *SyncRepositoryByIdReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Ids) {
		return false
	}
	return true
}

func (p *SyncRepositoryByIdReq) Field1DeepEqual(src []int64) bool {

	if len(p.Ids) != len(src) {
		return false
	}
	for i, v := range p.Ids {
		_src := src[i]
		if v != _src {
			return false
		}
	}
	return true
}

type SyncRepositoryByIdRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewSyncRepositoryByIdRes() *SyncRepositoryByIdRes {
	return &SyncRepositoryByIdRes{}
}

func (p *SyncRepositoryByIdRes) InitDefault() {
	*p = SyncRepositoryByIdRes{}
}

func (p *SyncRepositoryByIdRes) GetCode() (v int32) {
	return p.Code
}

func (p *SyncRepositoryByIdRes) GetMsg() (v string) {
	return p.Msg
}
func (p *SyncRepositoryByIdRes) SetCode(val int32) {
	p.Code = val
}
func (p *SyncRepositoryByIdRes) SetMsg(val string) {
	p.Msg = val
}

var fieldIDToName_SyncRepositoryByIdRes = map[int16]string{
	1: "code",
	2: "msg",
}

func (p *SyncRepositoryByIdRes) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SyncRepositoryByIdRes[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SyncRepositoryByIdRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}

func (p *SyncRepositoryByIdRes) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}

func (p *SyncRepositoryByIdRes) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SyncRepositoryByIdRes"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SyncRepositoryByIdRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SyncRepositoryByIdRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *SyncRepositoryByIdRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SyncRepositoryByIdRes(%+v)", *p)
}

func (p *SyncRepositoryByIdRes) DeepEqual(ano *SyncRepositoryByIdRes) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Code) {
		return false
	}
	if !p.Field2DeepEqual(ano.Msg) {
		return false
	}
	return true
}

func (p *SyncRepositoryByIdRes) Field1DeepEqual(src int32) bool {

	if p.Code != src {
		return false
	}
	return true
}
func (p *SyncRepositoryByIdRes) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Msg, src) != 0 {
		return false
	}
	return true
}

type UpdateRepositoryStatusReq struct {
	Id     int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Status string `thrift:"status,2" frugal:"2,default,string" json:"status"`
}

func NewUpdateRepositoryStatusReq() *UpdateRepositoryStatusReq {
	return &UpdateRepositoryStatusReq{}
}

func (p *UpdateRepositoryStatusReq) InitDefault() {
	*p = UpdateRepositoryStatusReq{}
}

func (p *UpdateRepositoryStatusReq) GetId() (v int64) {
	return p.Id
}

func (p *UpdateRepositoryStatusReq) GetStatus() (v string) {
	return p.Status
}
func (p *UpdateRepositoryStatusReq) SetId(val int64) {
	p.Id = val
}
func (p *UpdateRepositoryStatusReq) SetStatus(val string) {
	p.Status = val
}

var fieldIDToName_UpdateRepositoryStatusReq = map[int16]string{
	1: "id",
	2: "status",
}

func (p *UpdateRepositoryStatusReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UpdateRepositoryStatusReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UpdateRepositoryStatusReq) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Id = v
	}
	return nil
}

func (p *UpdateRepositoryStatusReq) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Status = v
	}
	return nil
}

func (p *UpdateRepositoryStatusReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UpdateRepositoryStatusReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UpdateRepositoryStatusReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UpdateRepositoryStatusReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Status); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UpdateRepositoryStatusReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateRepositoryStatusReq(%+v)", *p)
}

func (p *UpdateRepositoryStatusReq) DeepEqual(ano *UpdateRepositoryStatusReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.Status) {
		return false
	}
	return true
}

func (p *UpdateRepositoryStatusReq) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *UpdateRepositoryStatusReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Status, src) != 0 {
		return false
	}
	return true
}

type UpdateRepositoryStatusRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewUpdateRepositoryStatusRes() *UpdateRepositoryStatusRes {
	return &UpdateRepositoryStatusRes{}
}

func (p *UpdateRepositoryStatusRes) InitDefault() {
	*p = UpdateRepositoryStatusRes{}
}

func (p *UpdateRepositoryStatusRes) GetCode() (v int32) {
	return p.Code
}

func (p *UpdateRepositoryStatusRes) GetMsg() (v string) {
	return p.Msg
}
func (p *UpdateRepositoryStatusRes) SetCode(val int32) {
	p.Code = val
}
func (p *UpdateRepositoryStatusRes) SetMsg(val string) {
	p.Msg = val
}

var fieldIDToName_UpdateRepositoryStatusRes = map[int16]string{
	1: "code",
	2: "msg",
}

func (p *UpdateRepositoryStatusRes) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UpdateRepositoryStatusRes[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UpdateRepositoryStatusRes) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}

func (p *UpdateRepositoryStatusRes) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}

func (p *UpdateRepositoryStatusRes) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UpdateRepositoryStatusRes"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UpdateRepositoryStatusRes) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UpdateRepositoryStatusRes) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UpdateRepositoryStatusRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateRepositoryStatusRes(%+v)", *p)
}

func (p *UpdateRepositoryStatusRes) DeepEqual(ano *UpdateRepositoryStatusRes) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Code) {
		return false
	}
	if !p.Field2DeepEqual(ano.Msg) {
		return false
	}
	return true
}

func (p *UpdateRepositoryStatusRes) Field1DeepEqual(src int32) bool {

	if p.Code != src {
		return false
	}
	return true
}
func (p *UpdateRepositoryStatusRes) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Msg, src) != 0 {
		return false
	}
	return true
}
