// Copyright 2023 The gVisor Authors.
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

//go:build amd64
// +build amd64

package linux

// TASK_SIZE can be one of two values, corresponding to 4-level and 5-level
// paging.
//
// The array has to be sorted in decreasing order.
var feasibleTaskSizes = []uintptr{0xfffffffffff000, 0x7ffffffff000}

// Page fault error codes
const (
	X86_PF_PROT = 1 << iota
	X86_PF_WRITE
	X86_PF_USER
	X86_PF_RSVD
	X86_PF_INSTR
)
