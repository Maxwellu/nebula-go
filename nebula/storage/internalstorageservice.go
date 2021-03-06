// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package storage

import (
	"bytes"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
	nebula0 "github.com/vesoft-inc/nebula-go/nebula"
	meta1 "github.com/vesoft-inc/nebula-go/nebula/meta"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal

var _ = nebula0.GoUnusedProtection__
var _ = meta1.GoUnusedProtection__
type InternalStorageService interface {
  // Parameters:
  //  - Req
  GetValue(req *GetValueRequest) (r *GetValueResponse, err error)
  // Parameters:
  //  - Req
  ForwardTransaction(req *InternalTxnRequest) (r *ExecResponse, err error)
}

type InternalStorageServiceClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
}

func (client *InternalStorageServiceClient) Close() error {
  return client.Transport.Close()
}

func NewInternalStorageServiceClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *InternalStorageServiceClient {
  return &InternalStorageServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewInternalStorageServiceClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *InternalStorageServiceClient {
  return &InternalStorageServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - Req
func (p *InternalStorageServiceClient) GetValue(req *GetValueRequest) (r *GetValueResponse, err error) {
  if err = p.sendGetValue(req); err != nil { return }
  return p.recvGetValue()
}

func (p *InternalStorageServiceClient) sendGetValue(req *GetValueRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getValue", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := InternalStorageServiceGetValueArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *InternalStorageServiceClient) recvGetValue() (value *GetValueResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "getValue" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "getValue failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "getValue failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error397 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error398 error
    error398, err = error397.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error398
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getValue failed: invalid message type")
    return
  }
  result := InternalStorageServiceGetValueResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Req
func (p *InternalStorageServiceClient) ForwardTransaction(req *InternalTxnRequest) (r *ExecResponse, err error) {
  if err = p.sendForwardTransaction(req); err != nil { return }
  return p.recvForwardTransaction()
}

func (p *InternalStorageServiceClient) sendForwardTransaction(req *InternalTxnRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("forwardTransaction", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := InternalStorageServiceForwardTransactionArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *InternalStorageServiceClient) recvForwardTransaction() (value *ExecResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "forwardTransaction" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "forwardTransaction failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "forwardTransaction failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error399 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error400 error
    error400, err = error399.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error400
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "forwardTransaction failed: invalid message type")
    return
  }
  result := InternalStorageServiceForwardTransactionResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type InternalStorageServiceThreadsafeClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
  Mu sync.Mutex
}

func NewInternalStorageServiceThreadsafeClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *InternalStorageServiceThreadsafeClient {
  return &InternalStorageServiceThreadsafeClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewInternalStorageServiceThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *InternalStorageServiceThreadsafeClient {
  return &InternalStorageServiceThreadsafeClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

func (p *InternalStorageServiceThreadsafeClient) Threadsafe() {}

// Parameters:
//  - Req
func (p *InternalStorageServiceThreadsafeClient) GetValue(req *GetValueRequest) (r *GetValueResponse, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendGetValue(req); err != nil { return }
  return p.recvGetValue()
}

func (p *InternalStorageServiceThreadsafeClient) sendGetValue(req *GetValueRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getValue", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := InternalStorageServiceGetValueArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *InternalStorageServiceThreadsafeClient) recvGetValue() (value *GetValueResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "getValue" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "getValue failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "getValue failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error401 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error402 error
    error402, err = error401.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error402
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getValue failed: invalid message type")
    return
  }
  result := InternalStorageServiceGetValueResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Req
func (p *InternalStorageServiceThreadsafeClient) ForwardTransaction(req *InternalTxnRequest) (r *ExecResponse, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendForwardTransaction(req); err != nil { return }
  return p.recvForwardTransaction()
}

func (p *InternalStorageServiceThreadsafeClient) sendForwardTransaction(req *InternalTxnRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("forwardTransaction", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := InternalStorageServiceForwardTransactionArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *InternalStorageServiceThreadsafeClient) recvForwardTransaction() (value *ExecResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "forwardTransaction" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "forwardTransaction failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "forwardTransaction failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error403 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error404 error
    error404, err = error403.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error404
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "forwardTransaction failed: invalid message type")
    return
  }
  result := InternalStorageServiceForwardTransactionResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type InternalStorageServiceProcessor struct {
  processorMap map[string]thrift.ProcessorFunction
  handler InternalStorageService
}

func (p *InternalStorageServiceProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *InternalStorageServiceProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction, err error) {
  if processor, ok := p.processorMap[key]; ok {
    return processor, nil
  }
  return nil, nil // generic error message will be sent
}

func (p *InternalStorageServiceProcessor) ProcessorMap() map[string]thrift.ProcessorFunction {
  return p.processorMap
}

func NewInternalStorageServiceProcessor(handler InternalStorageService) *InternalStorageServiceProcessor {
  self405 := &InternalStorageServiceProcessor{handler:handler, processorMap:make(map[string]thrift.ProcessorFunction)}
  self405.processorMap["getValue"] = &internalStorageServiceProcessorGetValue{handler:handler}
  self405.processorMap["forwardTransaction"] = &internalStorageServiceProcessorForwardTransaction{handler:handler}
  return self405
}

type internalStorageServiceProcessorGetValue struct {
  handler InternalStorageService
}

func (p *internalStorageServiceProcessorGetValue) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := InternalStorageServiceGetValueArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *internalStorageServiceProcessorGetValue) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("getValue", messageType, seqId); err2 != nil {
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
  return err
}

func (p *internalStorageServiceProcessorGetValue) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*InternalStorageServiceGetValueArgs)
  var result InternalStorageServiceGetValueResult
  if retval, err := p.handler.GetValue(args.Req); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getValue: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}

type internalStorageServiceProcessorForwardTransaction struct {
  handler InternalStorageService
}

func (p *internalStorageServiceProcessorForwardTransaction) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := InternalStorageServiceForwardTransactionArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *internalStorageServiceProcessorForwardTransaction) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("forwardTransaction", messageType, seqId); err2 != nil {
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
  return err
}

func (p *internalStorageServiceProcessorForwardTransaction) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*InternalStorageServiceForwardTransactionArgs)
  var result InternalStorageServiceForwardTransactionResult
  if retval, err := p.handler.ForwardTransaction(args.Req); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing forwardTransaction: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Req
type InternalStorageServiceGetValueArgs struct {
  Req *GetValueRequest `thrift:"req,1" db:"req" json:"req"`
}

func NewInternalStorageServiceGetValueArgs() *InternalStorageServiceGetValueArgs {
  return &InternalStorageServiceGetValueArgs{}
}

var InternalStorageServiceGetValueArgs_Req_DEFAULT *GetValueRequest
func (p *InternalStorageServiceGetValueArgs) GetReq() *GetValueRequest {
  if !p.IsSetReq() {
    return InternalStorageServiceGetValueArgs_Req_DEFAULT
  }
return p.Req
}
func (p *InternalStorageServiceGetValueArgs) IsSetReq() bool {
  return p.Req != nil
}

func (p *InternalStorageServiceGetValueArgs) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField1(iprot); err != nil {
        return err
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

func (p *InternalStorageServiceGetValueArgs)  ReadField1(iprot thrift.Protocol) error {
  p.Req = NewGetValueRequest()
  if err := p.Req.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
  }
  return nil
}

func (p *InternalStorageServiceGetValueArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("getValue_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *InternalStorageServiceGetValueArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err) }
  if err := p.Req.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err) }
  return err
}

func (p *InternalStorageServiceGetValueArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("InternalStorageServiceGetValueArgs(%+v)", *p)
}

// Attributes:
//  - Success
type InternalStorageServiceGetValueResult struct {
  Success *GetValueResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewInternalStorageServiceGetValueResult() *InternalStorageServiceGetValueResult {
  return &InternalStorageServiceGetValueResult{}
}

var InternalStorageServiceGetValueResult_Success_DEFAULT *GetValueResponse
func (p *InternalStorageServiceGetValueResult) GetSuccess() *GetValueResponse {
  if !p.IsSetSuccess() {
    return InternalStorageServiceGetValueResult_Success_DEFAULT
  }
return p.Success
}
func (p *InternalStorageServiceGetValueResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *InternalStorageServiceGetValueResult) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField0(iprot); err != nil {
        return err
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

func (p *InternalStorageServiceGetValueResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewGetValueResponse()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *InternalStorageServiceGetValueResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("getValue_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *InternalStorageServiceGetValueResult) writeField0(oprot thrift.Protocol) (err error) {
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

func (p *InternalStorageServiceGetValueResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("InternalStorageServiceGetValueResult(%+v)", *p)
}

// Attributes:
//  - Req
type InternalStorageServiceForwardTransactionArgs struct {
  Req *InternalTxnRequest `thrift:"req,1" db:"req" json:"req"`
}

func NewInternalStorageServiceForwardTransactionArgs() *InternalStorageServiceForwardTransactionArgs {
  return &InternalStorageServiceForwardTransactionArgs{}
}

var InternalStorageServiceForwardTransactionArgs_Req_DEFAULT *InternalTxnRequest
func (p *InternalStorageServiceForwardTransactionArgs) GetReq() *InternalTxnRequest {
  if !p.IsSetReq() {
    return InternalStorageServiceForwardTransactionArgs_Req_DEFAULT
  }
return p.Req
}
func (p *InternalStorageServiceForwardTransactionArgs) IsSetReq() bool {
  return p.Req != nil
}

func (p *InternalStorageServiceForwardTransactionArgs) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField1(iprot); err != nil {
        return err
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

func (p *InternalStorageServiceForwardTransactionArgs)  ReadField1(iprot thrift.Protocol) error {
  p.Req = NewInternalTxnRequest()
  if err := p.Req.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
  }
  return nil
}

func (p *InternalStorageServiceForwardTransactionArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("forwardTransaction_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *InternalStorageServiceForwardTransactionArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err) }
  if err := p.Req.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err) }
  return err
}

func (p *InternalStorageServiceForwardTransactionArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("InternalStorageServiceForwardTransactionArgs(%+v)", *p)
}

// Attributes:
//  - Success
type InternalStorageServiceForwardTransactionResult struct {
  Success *ExecResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewInternalStorageServiceForwardTransactionResult() *InternalStorageServiceForwardTransactionResult {
  return &InternalStorageServiceForwardTransactionResult{}
}

var InternalStorageServiceForwardTransactionResult_Success_DEFAULT *ExecResponse
func (p *InternalStorageServiceForwardTransactionResult) GetSuccess() *ExecResponse {
  if !p.IsSetSuccess() {
    return InternalStorageServiceForwardTransactionResult_Success_DEFAULT
  }
return p.Success
}
func (p *InternalStorageServiceForwardTransactionResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *InternalStorageServiceForwardTransactionResult) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField0(iprot); err != nil {
        return err
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

func (p *InternalStorageServiceForwardTransactionResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewExecResponse()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *InternalStorageServiceForwardTransactionResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("forwardTransaction_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *InternalStorageServiceForwardTransactionResult) writeField0(oprot thrift.Protocol) (err error) {
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

func (p *InternalStorageServiceForwardTransactionResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("InternalStorageServiceForwardTransactionResult(%+v)", *p)
}


