// Copyright (c) 2020 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package fakeapi

import (
	"sync"

	"github.com/ProtonMail/proton-bridge/pkg/pmapi"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	// Internal states.
	lock               *sync.RWMutex
	fakeAPIs           []*FakePMAPI
	calls              []*fakeCall
	labelIDGenerator   idGenerator
	messageIDGenerator idGenerator
	tokenGenerator     idGenerator

	// State controlled by test.
	noInternetConnection bool
	usersByUsername      map[string]*fakeUser
	sessionsByUID        map[string]*fakeSession
	addressesByUsername  map[string]*pmapi.AddressList
	labelsByUsername     map[string][]*pmapi.Label
	messagesByUsername   map[string][]*pmapi.Message

	log *logrus.Entry
}

func NewController() *Controller {
	return &Controller{
		lock:               &sync.RWMutex{},
		fakeAPIs:           []*FakePMAPI{},
		calls:              []*fakeCall{},
		labelIDGenerator:   100, // We cannot use system label IDs.
		messageIDGenerator: 0,
		tokenGenerator:     1000, // No specific reason; 1000 simply feels right.

		noInternetConnection: false,
		usersByUsername:      map[string]*fakeUser{},
		sessionsByUID:        map[string]*fakeSession{},
		addressesByUsername:  map[string]*pmapi.AddressList{},
		labelsByUsername:     map[string][]*pmapi.Label{},
		messagesByUsername:   map[string][]*pmapi.Message{},

		log: logrus.WithField("pkg", "fakeapi-controller"),
	}
}

func (cntrl *Controller) GetClient(userID string) *FakePMAPI {
	fakeAPI := New(cntrl)
	cntrl.fakeAPIs = append(cntrl.fakeAPIs, fakeAPI)
	return fakeAPI
}
