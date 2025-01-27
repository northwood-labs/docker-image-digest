// Copyright 2024-2025, Northwood Labs
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

package httptls

// TLSVersion is a map of TLS versions to their human-readable names.
var TLSVersion = map[uint16]string{
	0x0002: "SSLv2",
	0x0300: "SSLv3",
	0x0301: "TLSv1.0",
	0x0302: "TLSv1.1",
	0x0303: "TLSv1.2",
	0x0304: "TLSv1.3",
}
