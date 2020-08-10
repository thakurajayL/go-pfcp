// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestAssociationReleaseRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Normal",
			Structured: message.NewAssociationReleaseRequest(
				seq, ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
			),
			Serialized: []byte{
				0x20, 0x09, 0x00, 0x25, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseAssociationReleaseRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
