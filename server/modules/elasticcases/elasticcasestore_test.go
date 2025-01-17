// Copyright 2019 Jason Ertel (jertel). All rights reserved.
// Copyright 2020-2022 Security Onion Solutions, LLC. All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package elasticcases

import (
  "context"
  "github.com/security-onion-solutions/securityonion-soc/model"
  "github.com/security-onion-solutions/securityonion-soc/server"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCreateUnauthorized(tester *testing.T) {
  casestore := NewElasticCasestore(server.NewFakeUnauthorizedServer())
  casestore.Init("some/url", "someusername", "somepassword", true)
  socCase := model.NewCase()
  newCase, err := casestore.Create(context.Background(), socCase)
  assert.Error(tester, err)
  assert.Nil(tester, newCase)
}

func TestCreate(tester *testing.T) {
  casestore := NewElasticCasestore(server.NewFakeAuthorizedServer(nil))
  casestore.Init("some/url", "someusername", "somepassword", true)
  caseResponse := `
    {
      "id": "a123",
      "title": "my title"
    }`
  casestore.client.MockStringResponse(caseResponse, 200, nil)
  socCase := model.NewCase()
  newCase, err := casestore.Create(context.Background(), socCase)
  assert.NoError(tester, err)

  assert.Equal(tester, "my title", newCase.Title)
  assert.Equal(tester, "a123", newCase.Id)
}
