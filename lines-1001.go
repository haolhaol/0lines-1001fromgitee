/*
 *
 * Copyright 2014 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

 package grpc

 import (
	 "testing"
 
	 "google.golang.org/grpc/encoding"
	 "google.golang.org/grpc/encoding/proto"
 )
 
 func (s) TestGetCodecForProtoIsNotNil(t *testing.T) {
	 if encoding.GetCodec(proto.Name) == nil {
		 t.Fatalf("encoding.GetCodec(%q) must not be nil by default", proto.Name)
	 }
 }
 package ssh

import (
	"io/ioutil"

	gossh "golang.org/x/crypto/ssh"
)

// PasswordAuth returns a functional option that sets PasswordHandler on the server.
func PasswordAuth(fn PasswordHandler) Option {
	return func(srv *Server) error {
		srv.PasswordHandler = fn
		return nil
	}
}

// PublicKeyAuth returns a functional option that sets PublicKeyHandler on the server.
func PublicKeyAuth(fn PublicKeyHandler) Option {
	return func(srv *Server) error {
		srv.PublicKeyHandler = fn
		return nil
	}
}

// HostKeyFile returns a functional option that adds HostSigners to the server
// from a PEM file at filepath.
func HostKeyFile(filepath string) Option {
	return func(srv *Server) error {
		pemBytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		signer, err := gossh.ParsePrivateKey(pemBytes)
		if err != nil {
			return err
		}

		srv.AddHostKey(signer)

		return nil
	}
}

// HostKeyPEM returns a functional option that adds HostSigners to the server
// from a PEM file as bytes.
func HostKeyPEM(bytes []byte) Option {
	return func(srv *Server) error {
		signer, err := gossh.ParsePrivateKey(bytes)
		if err != nil {
			return err
		}

		srv.AddHostKey(signer)

		return nil
	}
}

// NoPty returns a functional option that sets PtyCallback to return false,
// denying PTY requests.
func NoPty() Option {
	return func(srv *Server) error {
		srv.PtyCallback = func(ctx Context, pty Pty) bool {
			return false
		}
		return nil
	}
}

// WrapConn returns a functional option that sets ConnCallback on the server.
func WrapConn(fn ConnCallback) Option {
	return func(srv *Server) error {
		srv.ConnCallback = fn
		return nil
	}
}
/*
 *
 * Copyright 2014 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

 package grpc

 import (
	 "context"
 )
 
 // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // All errors returned by Invoke are compatible with the status package.
 func (cc *ClientConn) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...CallOption) error {
	 // allow interceptor to see all applicable call options, which means those
	 // configured as defaults from dial option as well as per-call options
	 opts = combine(cc.dopts.callOptions, opts)
 
	 if cc.dopts.unaryInt != nil {
		 return cc.dopts.unaryInt(ctx, method, args, reply, cc, invoke, opts...)
	 }
	 return invoke(ctx, method, args, reply, cc, opts...)
 }
 
 func combine(o1 []CallOption, o2 []CallOption) []CallOption {
	 // we don't use append because o1 could have extra capacity whose
	 // elements would be overwritten, which could cause inadvertent
	 // sharing (and race conditions) between concurrent calls
	 if len(o1) == 0 {
		 return o2
	 } else if len(o2) == 0 {
		 return o1
	 }
	 ret := make([]CallOption, len(o1)+len(o2))
	 copy(ret, o1)
	 copy(ret[len(o1):], o2)
	 return ret
 }
 
 // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead. // Invoke sends the RPC request on the wire and returns after response is
 // received.  This is typically called by generated code.
 //
 // DEPRECATED: Use ClientConn.Invoke instead.
  // received.  This is typically called by generated code.
 //
 //
 func Invoke(ctx context.Context, method string, args, reply interface{}, cc *ClientConn, opts ...CallOption) error {
	 return cc.Invoke(ctx, method, args, reply, opts...)
 }
 
 var unaryStreamDesc = &StreamDesc{ServerStreams: false, ClientStreams: false}
 
 func invoke(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, opts ...CallOption) error {
	 cs, err := newClientStream(ctx, unaryStreamDesc, cc, method, opts...)
	 if err != nil {
		 return err
	 }
	 if err := cs.SendMsg(req); err != nil {
		 return err
	 }
	 return cs.RecvMsg(reply)
 }