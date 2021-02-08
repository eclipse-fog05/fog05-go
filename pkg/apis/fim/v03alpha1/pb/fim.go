/*
* Copyright (c) 2014,2020 ADLINK Technology Inc.
* See the NOTICE file(s) distributed with this work for additional
* information regarding copyright ownership.
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors: Gabriele Baldoni, ADLINK Technology Inc.
* golang API
 */

package pb

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/eclipse-fog05/fog05-go/pkg/types/fdu"
	"google.golang.org/grpc"

	fdu_pb "github.com/eclipse-fog05/fog05-go/pkg/types/fdu/pb"
)

// FDUAPI is a component of FIMAPI
type FDUAPI struct {
	endpoint string
	sysid    string
	tenantid string
}

// OnboardFDU onboards an FDU into the catalog
func (f *FDUAPI) OnboardFDU(descriptor fdu.FDUDescriptor) (*fdu.FDUDescriptor, error) {
	var conn *grpc.ClientConn

	//converting fdu.FDUDescriptor to fdu_pb.FDUDescriptor
	b, err := json.Marshal(descriptor)
	if err != nil {
		return nil, err
	}

	var gDesc = fdu_pb.FDUDescriptor{}
	if err = json.Unmarshal(b, &gDesc); err != nil {
		return nil, err
	}

	// creating grpc connection
	conn, err = grpc.Dial(fmt.Sprintf("%s:9000", f.endpoint), grpc.WithInsecure())
	if err != nil {
		log.Error("DefineFDU, dial error :%s", err)
		return nil, err
	}
	defer conn.Close()

	// creating client
	c := fdu_pb.NewFDUApiClient(conn)

	// calling api
	resp, err := c.OnboardFDU(gDesc)
	if err != nil {
		return nil, err
	}

	if resp.error != "" {
		return nil, errors.New("math: square root of negative number")
	}

	//convert to fdu.FDUDescriptor

	b, err = json.Marshal(resp.ok)
	if err != nil {
		return nil, err
	}

	var res = fdu.FDUDescriptor{}
	if err = json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil

}

// DefineFDU defines an FDU
func (f *FDUAPI) DefineFDU(fduUUID uuid.UUID) (*fdu.FDURecord, error) {
	var conn *grpc.ClientConn

	var gRequest = fdu_pb.DefineFDURequest{
		uuid: fduUUID.String(),
	}

	// creating grpc connection
	conn, err := grpc.Dial(fmt.Sprintf("%s:9000", f.endpoint), grpc.WithInsecure())
	if err != nil {
		log.Error("DefineFDU, dial error :%s", err)
		return nil, err
	}
	defer conn.Close()

	// creating client
	c := fdu_pb.NewFDUApiClient(conn)

	// calling api
	resp, err := c.DefineFDU(gRequest)
	if err != nil {
		return nil, err
	}

	if resp.error != "" {
		return nil, errors.New("math: square root of negative number")
	}

	//convert to fdu.FDURecord

	b, err := json.Marshal(resp.ok)
	if err != nil {
		return nil, err
	}

	var res = fdu.FDURecord{}
	if err = json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil

}
