// Copyright 2023-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protovalidate

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"

	pb "buf.build/go/protovalidate/internal/gen/tests/example/v1"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func Example() {
	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Buf Build",
		Home: &pb.Coordinates{
			Lat: 27.380583333333334,
			Lng: 33.631838888888886,
		},
	}

	err := Validate(person)
	fmt.Println("valid:", err)

	person.Email = "not an email"
	err = Validate(person)
	fmt.Println("invalid:", err)

	// output:
	// valid: <nil>
	// invalid: validation error:
	//  - email: value must be a valid email address [string.email]
}

func ExampleWithFailFast() {
	loc := &pb.Coordinates{Lat: 999.999, Lng: -999.999}

	validator, err := New()
	if err != nil {
		log.Fatal(err)
	}
	err = validator.Validate(loc)
	fmt.Println("default:", err)

	validator, err = New(WithFailFast())
	if err != nil {
		log.Fatal(err)
	}
	err = validator.Validate(loc)
	fmt.Println("fail fast:", err)

	// output:
	// default: validation error:
	//  - lat: value must be greater than or equal to -90 and less than or equal to 90 [double.gte_lte]
	//  - lng: value must be greater than or equal to -180 and less than or equal to 180 [double.gte_lte]
	// fail fast: validation error:
	//  - lat: value must be greater than or equal to -90 and less than or equal to 90 [double.gte_lte]
}

func ExampleWithMessages() {
	validator, err := New(
		WithMessages(&pb.Person{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Protocol Buffer",
	}
	err = validator.Validate(person)
	fmt.Println(err)

	// output: <nil>
}

func ExampleWithMessageDescriptors() {
	pbType, err := protoregistry.GlobalTypes.FindMessageByName("tests.example.v1.Person")
	if err != nil {
		log.Fatal(err)
	}

	validator, err := New(
		WithMessageDescriptors(
			pbType.Descriptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Protocol Buffer",
	}
	err = validator.Validate(person)
	fmt.Println(err)

	// output: <nil>
}

func ExampleWithDisableLazy() {
	person := &pb.Person{
		Id:    1234,
		Email: "protovalidate@buf.build",
		Name:  "Buf Build",
		Home: &pb.Coordinates{
			Lat: 27.380583333333334,
			Lng: 33.631838888888886,
		},
	}

	validator, err := New(
		WithMessages(&pb.Coordinates{}),
		WithDisableLazy(),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = validator.Validate(person.GetHome())
	fmt.Println("person.Home:", err)
	err = validator.Validate(person)
	fmt.Println("person:", err)

	// output:
	// person.Home: <nil>
	// person: compilation error: no evaluator available for tests.example.v1.Person
}

func ExampleValidationError() {
	validator, err := New()
	if err != nil {
		log.Fatal(err)
	}

	loc := &pb.Coordinates{Lat: 999.999}
	err = validator.Validate(loc)
	var valErr *ValidationError
	if ok := errors.As(err, &valErr); ok {
		violation := valErr.Violations[0]
		fmt.Println(violation.Proto.GetField().GetElements()[0].GetFieldName(), violation.Proto.GetRuleId())
		fmt.Println(violation.RuleValue, violation.FieldValue)
	}

	// output: lat double.gte_lte
	// -90 999.999
}

func ExampleValidationError_localized() {
	validator, err := New()
	if err != nil {
		log.Fatal(err)
	}

	type ErrorInfo struct {
		FieldName  string
		RuleValue  any
		FieldValue any
	}

	var ruleMessages = map[string]string{
		"string.email_empty": "{{.FieldName}}: メールアドレスは空であってはなりません。\n",
		"string.pattern":     "{{.FieldName}}: 値はパターン「{{.RuleValue}}」一致する必要があります。\n",
		"uint64.gt":          "{{.FieldName}}: 値は{{.RuleValue}}を超える必要があります。（価値：{{.FieldValue}}）\n",
	}

	loc := &pb.Person{Id: 900}
	err = validator.Validate(loc)
	var valErr *ValidationError
	if ok := errors.As(err, &valErr); ok {
		for _, violation := range valErr.Violations {
			_ = template.
				Must(template.New("").Parse(ruleMessages[violation.Proto.GetRuleId()])).
				Execute(os.Stdout, ErrorInfo{
					FieldName:  violation.Proto.GetField().GetElements()[0].GetFieldName(),
					RuleValue:  violation.RuleValue.Interface(),
					FieldValue: violation.FieldValue.Interface(),
				})
		}
	}

	// output:
	// id: 値は999を超える必要があります。（価値：900）
	// email: メールアドレスは空であってはなりません。
	// name: 値はパターン「^[[:alpha:]]+( [[:alpha:]]+)*$」一致する必要があります。
}
