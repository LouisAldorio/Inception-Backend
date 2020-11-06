// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package graphql

import (
	"context"
	"github.com/vektah/gqlparser/v2/ast"
	"sync"
)

var (
	lockExecutableSchemaMockComplexity sync.RWMutex
	lockExecutableSchemaMockExec       sync.RWMutex
	lockExecutableSchemaMockSchema     sync.RWMutex
)

// Ensure, that ExecutableSchemaMock does implement ExecutableSchema.
// If this is not the case, regenerate this file with moq.
var _ ExecutableSchema = &ExecutableSchemaMock{}

// ExecutableSchemaMock is a mock implementation of ExecutableSchema.
//
//     func TestSomethingThatUsesExecutableSchema(t *testing.T) {
//
//         // make and configure a mocked ExecutableSchema
//         mockedExecutableSchema := &ExecutableSchemaMock{
//             ComplexityFunc: func(typeName string, fieldName string, childComplexity int, args map[string]interface{}) (int, bool) {
// 	               panic("mock out the Complexity method")
//             },
//             ExecFunc: func(ctx context.Context) ResponseHandler {
// 	               panic("mock out the Exec method")
//             },
//             SchemaFunc: func() *ast.Schema {
// 	               panic("mock out the Schema method")
//             },
//         }
//
//         // use mockedExecutableSchema in code that requires ExecutableSchema
//         // and then make assertions.
//
//     }
type ExecutableSchemaMock struct {
	// ComplexityFunc mocks the Complexity method.
	ComplexityFunc func(typeName string, fieldName string, childComplexity int, args map[string]interface{}) (int, bool)

	// ExecFunc mocks the Exec method.
	ExecFunc func(ctx context.Context) ResponseHandler

	// SchemaFunc mocks the Schema method.
	SchemaFunc func() *ast.Schema

	// calls tracks calls to the methods.
	calls struct {
		// Complexity holds details about calls to the Complexity method.
		Complexity []struct {
			// TypeName is the typeName argument value.
			TypeName string
			// FieldName is the fieldName argument value.
			FieldName string
			// ChildComplexity is the childComplexity argument value.
			ChildComplexity int
			// Args is the args argument value.
			Args map[string]interface{}
		}
		// Exec holds details about calls to the Exec method.
		Exec []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Schema holds details about calls to the Schema method.
		Schema []struct {
		}
	}
}

// Complexity calls ComplexityFunc.
func (mock *ExecutableSchemaMock) Complexity(typeName string, fieldName string, childComplexity int, args map[string]interface{}) (int, bool) {
	if mock.ComplexityFunc == nil {
		panic("ExecutableSchemaMock.ComplexityFunc: method is nil but ExecutableSchema.Complexity was just called")
	}
	callInfo := struct {
		TypeName        string
		FieldName       string
		ChildComplexity int
		Args            map[string]interface{}
	}{
		TypeName:        typeName,
		FieldName:       fieldName,
		ChildComplexity: childComplexity,
		Args:            args,
	}
	lockExecutableSchemaMockComplexity.Lock()
	mock.calls.Complexity = append(mock.calls.Complexity, callInfo)
	lockExecutableSchemaMockComplexity.Unlock()
	return mock.ComplexityFunc(typeName, fieldName, childComplexity, args)
}

// ComplexityCalls gets all the calls that were made to Complexity.
// Check the length with:
//     len(mockedExecutableSchema.ComplexityCalls())
func (mock *ExecutableSchemaMock) ComplexityCalls() []struct {
	TypeName        string
	FieldName       string
	ChildComplexity int
	Args            map[string]interface{}
} {
	var calls []struct {
		TypeName        string
		FieldName       string
		ChildComplexity int
		Args            map[string]interface{}
	}
	lockExecutableSchemaMockComplexity.RLock()
	calls = mock.calls.Complexity
	lockExecutableSchemaMockComplexity.RUnlock()
	return calls
}

// Exec calls ExecFunc.
func (mock *ExecutableSchemaMock) Exec(ctx context.Context) ResponseHandler {
	if mock.ExecFunc == nil {
		panic("ExecutableSchemaMock.ExecFunc: method is nil but ExecutableSchema.Exec was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockExecutableSchemaMockExec.Lock()
	mock.calls.Exec = append(mock.calls.Exec, callInfo)
	lockExecutableSchemaMockExec.Unlock()
	return mock.ExecFunc(ctx)
}

// ExecCalls gets all the calls that were made to Exec.
// Check the length with:
//     len(mockedExecutableSchema.ExecCalls())
func (mock *ExecutableSchemaMock) ExecCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockExecutableSchemaMockExec.RLock()
	calls = mock.calls.Exec
	lockExecutableSchemaMockExec.RUnlock()
	return calls
}

// Schema calls SchemaFunc.
func (mock *ExecutableSchemaMock) Schema() *ast.Schema {
	if mock.SchemaFunc == nil {
		panic("ExecutableSchemaMock.SchemaFunc: method is nil but ExecutableSchema.Schema was just called")
	}
	callInfo := struct {
	}{}
	lockExecutableSchemaMockSchema.Lock()
	mock.calls.Schema = append(mock.calls.Schema, callInfo)
	lockExecutableSchemaMockSchema.Unlock()
	return mock.SchemaFunc()
}

// SchemaCalls gets all the calls that were made to Schema.
// Check the length with:
//     len(mockedExecutableSchema.SchemaCalls())
func (mock *ExecutableSchemaMock) SchemaCalls() []struct {
} {
	var calls []struct {
	}
	lockExecutableSchemaMockSchema.RLock()
	calls = mock.calls.Schema
	lockExecutableSchemaMockSchema.RUnlock()
	return calls
}
