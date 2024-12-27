// Copyright 2024, Northwood Labs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package httptls is a library which contains functionality for testing HTTP
versions, TLS versions, and TLS ciphersuites.

The library is used by the `http` and `tls` subcommands of the `devsec-tools`
CLI.

This package leverages concurrency to test multiple versions of HTTP and TLS
simultaneously. The library also provides a `ParseDomain` function to parse a
domain from a URL-like string.
*/
package httptls