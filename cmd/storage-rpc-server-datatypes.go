/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
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
 */

package cmd

// GenericVolArgs - generic volume args.
type GenericVolArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string
}

// ListVolsReply represents list of vols RPC reply.
type ListVolsReply struct {
	// List of volumes stat information.
	Vols []VolInfo
}

// ReadAllArgs represents read all RPC arguments.
type ReadAllArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string

	// Name of the path.
	Path string
}

// ReadFileArgs represents read file RPC arguments.
type ReadFileArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string

	// Name of the path.
	Path string

	// Starting offset to start reading into Buffer.
	Offset int64

	// Data buffer read from the path at offset.
	Buffer []byte

	// Algorithm used in bit-rot hash computation.
	Algo BitrotAlgorithm

	// Stored hash value used to compare with computed value.
	ExpectedHash []byte

	// Indicates whether the disk has already been verified
	Verified bool
}

// PrepareFileArgs represents append file RPC arguments.
type PrepareFileArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string

	// Name of the path.
	Path string

	// Size of the file to be prepared
	Size int64
}

// AppendFileArgs represents append file RPC arguments.
type AppendFileArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string

	// Name of the path.
	Path string

	// Data buffer to be saved at path.
	Buffer []byte
}

// StatFileArgs represents stat file RPC arguments.
type StatFileArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string

	// Name of the path.
	Path string
}

// DeleteFileArgs represents delete file RPC arguments.
type DeleteFileArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string

	// Name of the path.
	Path string
}

// ListDirArgs represents list contents RPC arguments.
type ListDirArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of the volume.
	Vol string

	// Name of the path.
	Path string
}

// RenameFileArgs represents rename file RPC arguments.
type RenameFileArgs struct {
	// Authentication token generated by Login.
	AuthRPCArgs

	// Name of source volume.
	SrcVol string

	// Source path to be renamed.
	SrcPath string

	// Name of destination volume.
	DstVol string

	// Destination path of renamed file.
	DstPath string
}
