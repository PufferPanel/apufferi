/*
 Copyright 2019 Padduck, LLC
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
  	http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package apufferi

import "fmt"

type Error interface {
	error

	GetMessage() string

	GetHumanMessage() string

	GetCode() int

	Is(Error) bool

	SetData(machine []string, human []string) Error
	Set(machine string, human string) Error
}

type genericError struct {
	message   string
	human     string
	code      int
	msgData   []string
	humanData []string
}

func (ge genericError) GetMessage() string {
	return fmt.Sprintf(ge.message, ge.msgData)
}

func (ge genericError) GetHumanMessage() string {
	return fmt.Sprintf(ge.human, ge.humanData)
}

func (ge genericError) GetCode() int {
	return ge.code
}

func (ge genericError) Error() string {
	if ge.human != "" {
		return fmt.Sprintf(ge.human, ge.humanData)
	} else {
		return fmt.Sprintf(ge.message, ge.msgData)
	}
}

func (ge genericError) Is(err Error) bool {
	return ge.GetCode() == err.GetCode()
}

func (ge genericError) SetData(machine []string, human []string) Error {
	cp := ge
	cp.msgData = machine
	if human == nil {
		cp.humanData = cp.msgData
	} else {
		cp.humanData = human
	}
	return cp
}

func (ge genericError) Set(machine string, human string) Error {
	if human == "" {
		human = machine
	}
	return ge.SetData([]string{machine}, []string{human})
}

func CreateError(msg, humanMsg string, code int) Error {
	if humanMsg == "" {
		humanMsg = msg
	}

	return genericError{
		message: msg,
		human:   humanMsg,
		code:    code,
	}
}

func FromError(err error) Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(Error); e != nil && ok {
		return e
	}
	return genericError{
		message: err.Error(),
		human:   err.Error(),
		code:    0,
	}
}
