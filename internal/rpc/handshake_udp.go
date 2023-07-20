// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rpc

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

type HandshakeMsgUdp struct {
	Flags     uint32
	SenderPID NetPID
	PeerPID   NetPID
}

type CryptoKeysUdp struct {
	ReadKey  [32]byte
	WriteKey [32]byte
}

func DeriveCryptoKeysUdp(key string, localPid *NetPID, remotePid *NetPID, generation uint32) (*CryptoKeysUdp, error) {
	w, err := writeCryptoInitMsgUdp(key, localPid, remotePid, generation)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize write key derivation data: %w", err)
	}

	//fmt.Println("init write crypto buf", w)
	var keys CryptoKeysUdp
	w1 := md5.Sum(w[1:])
	w2 := sha1.Sum(w)
	copy(keys.WriteKey[:], w1[:])
	copy(keys.WriteKey[12:], w2[:])

	r, err := writeCryptoInitMsgUdp(key, remotePid, localPid, generation)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize read key derivation data: %w", err)
	}
	//fmt.Println("init read crypto buf", r)

	r1 := md5.Sum(r[1:])
	r2 := sha1.Sum(r)
	copy(keys.ReadKey[:], r1[:])
	copy(keys.ReadKey[12:], r2[:])

	return &keys, nil
}

func writeCryptoInitMsgUdp(key string, localPid *NetPID, remotePid *NetPID, generation uint32) ([]byte, error) {
	buf := &bytes.Buffer{}

	if err := binary.Write(buf, binary.LittleEndian, *localPid); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, []byte(key)); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, *remotePid); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, generation); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
