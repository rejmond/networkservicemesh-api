// Copyright (c) 2019-2021 Cisco Systems, Inc and/or its affiliates.
//
// Copyright (c) 2021 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kernel - describe kernel mechanism
package kernel

import (
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/cls"
)

// Mechanism is a kernel mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// New returns *networkservice.Mechanism of type kernel using the given netnsURL (file:///proc/${pid}/ns/net)
func New(netnsURL string) *networkservice.Mechanism {
	return &networkservice.Mechanism{
		Cls:  cls.LOCAL,
		Type: MECHANISM,
		Parameters: map[string]string{
			NetNSURL: netnsURL,
		},
	}
}

// ToMechanism converts unified mechanism to helper
// If Mechanism m is *not* of type kernel.MECHANISM, it returns nil
func ToMechanism(m *networkservice.Mechanism) *Mechanism {
	if m.GetType() == MECHANISM {
		return &Mechanism{
			m,
		}
	}
	return nil
}

// GetParameters returns the map of all parameters to the mechanism
func (m *Mechanism) GetParameters() map[string]string {
	if m == nil {
		return map[string]string{}
	}
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	return m.Parameters
}

// GetNetNSInode returns the NetNS inode
func (m *Mechanism) GetNetNSInode() string {
	return m.GetParameters()[NetNSInodeKey]
}

// SetNetNSInode sets the NetNS inode
func (m *Mechanism) SetNetNSInode(netNSInode string) {
	m.GetParameters()[NetNSInodeKey] = netNSInode
}

// GetPCIAddress returns the PCI address of the device
func (m *Mechanism) GetPCIAddress() string {
	return m.GetParameters()[PCIAddressKey]
}

// SetPCIAddress sets the PCI address of the device
func (m *Mechanism) SetPCIAddress(pciAddress string) {
	m.GetParameters()[PCIAddressKey] = pciAddress
}

// IsPCIDevice returns if this mechanism is for a PCI device
func (m *Mechanism) IsPCIDevice() bool {
	return m.GetPCIAddress() != ""
}

// GetDeviceTokenID returns device token ID
func (m *Mechanism) GetDeviceTokenID() string {
	return m.Parameters[DeviceTokenIDKey]
}

// SetDeviceTokenID sets device token ID
func (m *Mechanism) SetDeviceTokenID(tokenID string) {
	m.Parameters[DeviceTokenIDKey] = tokenID
}

// GetInterfaceName returns the Kernel Interface Name
func (m *Mechanism) GetInterfaceName() string {
	return m.GetParameters()[InterfaceNameKey]
}

// SetInterfaceName sets the Kernel Interface Name
func (m *Mechanism) SetInterfaceName(interfaceName string) {
	m.GetParameters()[InterfaceNameKey] = interfaceName
}

// GetNetNSURL returns the NetNS URL, it can be either:
// * file:///proc/${pid}/ns/net - ${pid} process net NS
// * inode://${dev}/${ino} - while transferring file between processes using grpcfd
func (m *Mechanism) GetNetNSURL() string {
	return m.GetParameters()[NetNSURL]
}

// SetNetNSURL sets the NetNS URL - file:///proc/${pid}/ns/net
func (m *Mechanism) SetNetNSURL(urlString string) {
	m.GetParameters()[NetNSURL] = urlString
}

// SupportsVLAN returns SupportsVLAN flag
func (m *Mechanism) SupportsVLAN() bool {
	boolValue, err := strconv.ParseBool(m.GetParameters()[SupportsVLAN])
	if err != nil {
		return false
	}
	return boolValue
}

// SetSupportsVLAN set SupportsVLAN flag
func (m *Mechanism) SetSupportsVLAN(supportsVlan bool) {
	m.GetParameters()[SupportsVLAN] = strconv.FormatBool(supportsVlan)
}

// GetVLAN - return Vlan value - 0 if unset or invalid
func (m *Mechanism) GetVLAN() uint32 {
	// vlan ID range is 0 to 4,095 - can be stored in 12 bit
	vlan, err := strconv.ParseUint(m.GetParameters()[VLAN], 10, 12)
	if err != nil {
		return 0
	}

	return uint32(vlan)
}

// SetVLAN - set the VLAN value
func (m *Mechanism) SetVLAN(vlan uint32) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[VLAN] = strconv.FormatUint(uint64(vlan), 10)

	return m
}
