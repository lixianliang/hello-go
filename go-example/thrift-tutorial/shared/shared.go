// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package shared

import (
	"bytes"
	"reflect"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - Key
//  - Value
type SharedStruct struct {
  Key int32 `thrift:"key,1" db:"key" json:"key"`
  Value string `thrift:"value,2" db:"value" json:"value"`
}

func NewSharedStruct() *SharedStruct {
  return &SharedStruct{}
}


func (p *SharedStruct) GetKey() int32 {
  return p.Key
}

func (p *SharedStruct) GetValue() string {
  return p.Value
}
func (p *SharedStruct) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SharedStruct)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *SharedStruct)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Value = v
}
  return nil
}

func (p *SharedStruct) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SharedStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedStruct) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteI32(int32(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *SharedStruct) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:value: ", p), err) }
  if err := oprot.WriteString(string(p.Value)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.value (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:value: ", p), err) }
  return err
}

func (p *SharedStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SharedStruct(%+v)", *p)
}

type SharedService interface {
  // Parameters:
  //  - Key
  GetStruct(key int32) (r *SharedStruct, err error)
}

type SharedServiceClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewSharedServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SharedServiceClient {
  return &SharedServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewSharedServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SharedServiceClient {
  return &SharedServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - Key
func (p *SharedServiceClient) GetStruct(key int32) (r *SharedStruct, err error) {
  if err = p.sendGetStruct(key); err != nil { return }
  return p.recvGetStruct()
}

func (p *SharedServiceClient) sendGetStruct(key int32)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getStruct", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := SharedServiceGetStructArgs{
  Key : key,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *SharedServiceClient) recvGetStruct() (value *SharedStruct, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "getStruct" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getStruct failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getStruct failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getStruct failed: invalid message type")
    return
  }
  result := SharedServiceGetStructResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type SharedServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler SharedService
}

func (p *SharedServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *SharedServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *SharedServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction{
  return p.processorMap
}

func NewSharedServiceProcessor(handler SharedService) *SharedServiceProcessor {

  self2 := &SharedServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["getStruct"] = &sharedServiceProcessorGetStruct{handler:handler}
return self2
}

func (p *SharedServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x3

}

type sharedServiceProcessorGetStruct struct {
  handler SharedService
}

func (p *sharedServiceProcessorGetStruct) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := SharedServiceGetStructArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getStruct", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := SharedServiceGetStructResult{}
var retval *SharedStruct
  var err2 error
  if retval, err2 = p.handler.GetStruct(args.Key); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getStruct: " + err2.Error())
    oprot.WriteMessageBegin("getStruct", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getStruct", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Key
type SharedServiceGetStructArgs struct {
  Key int32 `thrift:"key,1" db:"key" json:"key"`
}

func NewSharedServiceGetStructArgs() *SharedServiceGetStructArgs {
  return &SharedServiceGetStructArgs{}
}


func (p *SharedServiceGetStructArgs) GetKey() int32 {
  return p.Key
}
func (p *SharedServiceGetStructArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SharedServiceGetStructArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *SharedServiceGetStructArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getStruct_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedServiceGetStructArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteI32(int32(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *SharedServiceGetStructArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SharedServiceGetStructArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SharedServiceGetStructResult struct {
  Success *SharedStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSharedServiceGetStructResult() *SharedServiceGetStructResult {
  return &SharedServiceGetStructResult{}
}

var SharedServiceGetStructResult_Success_DEFAULT *SharedStruct
func (p *SharedServiceGetStructResult) GetSuccess() *SharedStruct {
  if !p.IsSetSuccess() {
    return SharedServiceGetStructResult_Success_DEFAULT
  }
return p.Success
}
func (p *SharedServiceGetStructResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *SharedServiceGetStructResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SharedServiceGetStructResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &SharedStruct{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *SharedServiceGetStructResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getStruct_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedServiceGetStructResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *SharedServiceGetStructResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SharedServiceGetStructResult(%+v)", *p)
}


