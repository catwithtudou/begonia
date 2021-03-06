// Code generated by Begonia. DO NOT EDIT.
// versions:
// 	Begonia v1.0.2
// source: server.go
// begonia server file

package main

import (
	"context"
	"errors"
	"github.com/MashiroC/begonia/app"
	"github.com/MashiroC/begonia/app/coding"
	cRegister "github.com/MashiroC/begonia/core/register"
)

var (
	_TestServiceFuncList []cRegister.FunInfo

	_TestServiceEchoInSchema = `
{
			"namespace":"begonia.func.Echo",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":"int","alias":"i1"}
,{"name":"F2","type":"int","alias":"it"}
,{"name":"F3","type":"int","alias":"i2"}
,{"name":"F4","type":"int","alias":"i3"}
,{"name":"F5","type":"int","alias":"i4"}
,{"name":"F6","type":"long","alias":"i5"}
,{"name":"F7","type":"float","alias":"f1"}
,{"name":"F8","type":"double","alias":"f2"}
,{"name":"F9","type":"boolean","alias":"ok"}
,{"name":"F10","type":"string","alias":"str"}
,{"name":"F11","type":{
				"type": "array",
				"items": "int"
			},"alias":"s1"}
,{"name":"F12","type":{
				"type": "array",
				"items": "string"
			},"alias":"s2"}
,{"name":"F13","type":"bytes","alias":"s6"}
,{"name":"F14","type":{
				"type": "record",
				"name": "TestStruct",
				"fields":[{"name":"I1","type":"int"}
,{"name":"I2","type":"int"}
,{"name":"I3","type":"int"}
,{"name":"I4","type":"int"}
,{"name":"I5","type":"long"}
,{"name":"Str","type":"string"}
,{"name":"S1","type":{
				"type": "array",
				"items": "int"
			}}
,{"name":"S2","type":{
				"type": "array",
				"items": "string"
			}}
,{"name":"TestStruct2","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Test3","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Map1","type":{"type":"map","values":"string"}}
,{"name":"Map2","type":{"type":"map","values":{
				"type": "array",
				"items": "int"
			}}}

				]
			},"alias":"st"}
,{"name":"F15","type":{"type":"map","values":"string"},"alias":"m1"}
,{"name":"F16","type":{"type":"map","values":"int"},"alias":"m2"}
,{"name":"F17","type":{"type":"map","values":{
				"type": "record",
				"name": "TestStruct",
				"fields":[{"name":"I1","type":"int"}
,{"name":"I2","type":"int"}
,{"name":"I3","type":"int"}
,{"name":"I4","type":"int"}
,{"name":"I5","type":"long"}
,{"name":"Str","type":"string"}
,{"name":"S1","type":{
				"type": "array",
				"items": "int"
			}}
,{"name":"S2","type":{
				"type": "array",
				"items": "string"
			}}
,{"name":"TestStruct2","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Test3","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Map1","type":{"type":"map","values":"string"}}
,{"name":"Map2","type":{"type":"map","values":{
				"type": "array",
				"items": "int"
			}}}

				]
			}},"alias":"m3"}

			]
		}`
	_TestServiceEchoOutSchema = `
{
			"namespace":"begonia.func.Echo",
			"type":"record",
			"name":"Out",
			"fields":[
				{"name":"F1","type":"string"}

			]
		}`
	_TestServiceEchoInCoder  coding.Coder
	_TestServiceEchoOutCoder coding.Coder

	_TestServiceEcho2InSchema = `
{
			"namespace":"begonia.func.Echo2",
			"type":"record",
			"name":"In",
			"fields":[
				
			]
		}`
	_TestServiceEcho2OutSchema = `
{
			"namespace":"begonia.func.Echo2",
			"type":"record",
			"name":"Out",
			"fields":[
				{"name":"F1","type":"int","alias":"i1"}
,{"name":"F2","type":"int","alias":"i2"}
,{"name":"F3","type":"int","alias":"i3"}
,{"name":"F4","type":"int","alias":"i4"}
,{"name":"F5","type":"long","alias":"i5"}
,{"name":"F6","type":"float","alias":"f1"}
,{"name":"F7","type":"double","alias":"f2"}
,{"name":"F8","type":"boolean","alias":"ok"}
,{"name":"F9","type":"string","alias":"str"}
,{"name":"F10","type":{
				"type": "array",
				"items": "int"
			},"alias":"s1"}
,{"name":"F11","type":{
				"type": "array",
				"items": "string"
			},"alias":"s2"}
,{"name":"F12","type":"bytes","alias":"s6"}
,{"name":"F13","type":{
				"type": "record",
				"name": "TestStruct",
				"fields":[{"name":"I1","type":"int"}
,{"name":"I2","type":"int"}
,{"name":"I3","type":"int"}
,{"name":"I4","type":"int"}
,{"name":"I5","type":"long"}
,{"name":"Str","type":"string"}
,{"name":"S1","type":{
				"type": "array",
				"items": "int"
			}}
,{"name":"S2","type":{
				"type": "array",
				"items": "string"
			}}
,{"name":"TestStruct2","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Test3","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Map1","type":{"type":"map","values":"string"}}
,{"name":"Map2","type":{"type":"map","values":{
				"type": "array",
				"items": "int"
			}}}

				]
			},"alias":"st"}
,{"name":"F14","type":{"type":"map","values":"string"},"alias":"m1"}
,{"name":"F15","type":{"type":"map","values":"int"},"alias":"m2"}
,{"name":"F16","type":{"type":"map","values":{
				"type": "record",
				"name": "TestStruct",
				"fields":[{"name":"I1","type":"int"}
,{"name":"I2","type":"int"}
,{"name":"I3","type":"int"}
,{"name":"I4","type":"int"}
,{"name":"I5","type":"long"}
,{"name":"Str","type":"string"}
,{"name":"S1","type":{
				"type": "array",
				"items": "int"
			}}
,{"name":"S2","type":{
				"type": "array",
				"items": "string"
			}}
,{"name":"TestStruct2","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Test3","type":{
				"type": "record",
				"name": "TestStruct2",
				"fields":[{"name":"B1","type":"bytes"}
,{"name":"B2","type":"bytes"}

				]
			}}
,{"name":"Map1","type":{"type":"map","values":"string"}}
,{"name":"Map2","type":{"type":"map","values":{
				"type": "array",
				"items": "int"
			}}}

				]
			}},"alias":"m3"}

			]
		}`
	_TestServiceEcho2InCoder  coding.Coder
	_TestServiceEcho2OutCoder coding.Coder
)

type _TestServiceEchoIn struct {
	F1  int
	F2  int
	F3  int8
	F4  int16
	F5  int32
	F6  int64
	F7  float32
	F8  float64
	F9  bool
	F10 string
	F11 []int
	F12 []string
	F13 []byte
	F14 TestStruct
	F15 map[string]string
	F16 map[string]int
	F17 map[string]TestStruct
}

type _TestServiceEchoOut struct {
	F1 string
}

type _TestServiceEcho2In struct {
}

type _TestServiceEcho2Out struct {
	F1  int
	F2  int8
	F3  int16
	F4  int32
	F5  int64
	F6  float32
	F7  float64
	F8  bool
	F9  string
	F10 []int
	F11 []string
	F12 []byte
	F13 TestStruct
	F14 map[string]string
	F15 map[string]int
	F16 map[string]TestStruct
}

func init() {
	app.ServiceAppMode = app.Ast

	var err error

	_TestServiceEchoInCoder, err = coding.NewAvro(_TestServiceEchoInSchema)
	if err != nil {
		panic(err)
	}
	_TestServiceEchoOutCoder, err = coding.NewAvro(_TestServiceEchoOutSchema)
	if err != nil {
		panic(err)
	}

	_TestServiceEcho2InCoder, err = coding.NewAvro(_TestServiceEcho2InSchema)
	if err != nil {
		panic(err)
	}
	_TestServiceEcho2OutCoder, err = coding.NewAvro(_TestServiceEcho2OutSchema)
	if err != nil {
		panic(err)
	}

	_TestServiceFuncList = []cRegister.FunInfo{

		{
			Name:      "Echo",
			InSchema:  _TestServiceEchoInSchema,
			OutSchema: _TestServiceEchoOutSchema},

		{
			Name:      "Echo2",
			InSchema:  _TestServiceEcho2InSchema,
			OutSchema: _TestServiceEcho2OutSchema},
	}
}

func (d *TestService) Do(ctx context.Context, fun string, param []byte) (result []byte, err error) {
	switch fun {

	case "Echo":
		var in _TestServiceEchoIn
		err = _TestServiceEchoInCoder.DecodeIn(param, &in)
		if err != nil {
			panic(err)
		}

		res1 := d.Echo(

			in.F1, in.F2, in.F3, in.F4, in.F5, in.F6, in.F7, in.F8, in.F9, in.F10, in.F11, in.F12, in.F13, in.F14, in.F15, in.F16, in.F17,
		)
		if err != nil {
			return nil, err
		}
		var out _TestServiceEchoOut
		out.F1 = res1

		res, err := _TestServiceEchoOutCoder.Encode(out)
		if err != nil {
			panic(err)
		}
		return res, nil

	case "Echo2":
		var in _TestServiceEcho2In
		err = _TestServiceEcho2InCoder.DecodeIn(param, &in)
		if err != nil {
			panic(err)
		}

		res1, res2, res3, res4, res5, res6, res7, res8, res9, res10, res11, res12, res13, res14, res15, res16 := d.Echo2()
		if err != nil {
			return nil, err
		}
		var out _TestServiceEcho2Out
		out.F1 = res1
		out.F2 = res2
		out.F3 = res3
		out.F4 = res4
		out.F5 = res5
		out.F6 = res6
		out.F7 = res7
		out.F8 = res8
		out.F9 = res9
		out.F10 = res10
		out.F11 = res11
		out.F12 = res12
		out.F13 = res13
		out.F14 = res14
		out.F15 = res15
		out.F16 = res16

		res, err := _TestServiceEcho2OutCoder.Encode(out)
		if err != nil {
			panic(err)
		}
		return res, nil

	default:
		err = errors.New("rpc call error: fun not exist")
	}
	return
}

func (d *TestService) FuncList() []cRegister.FunInfo {
	return _TestServiceFuncList
}
