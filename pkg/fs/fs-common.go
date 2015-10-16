/*
 * Minio Cloud Storage, (C) 2015 Minio, Inc.
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

package fs

import (
	"bufio"
	"bytes"
	"os"
	"sort"
	"strings"
	"time"
)

// Metadata - carries metadata about object
type Metadata struct {
	Md5sum      []byte
	ContentType string
}

// sortUnique sort a slice in lexical order, removing duplicate elements
func sortUnique(objects []string) []string {
	objectMap := make(map[string]string)
	for _, v := range objects {
		objectMap[v] = v
	}
	var results []string
	for k := range objectMap {
		results = append(results, k)
	}
	sort.Strings(results)
	return results
}

type contentInfo struct {
	os.FileInfo
	Prefix  string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
}

type bucketDir struct {
	files []contentInfo
	root  string
}

func delimiter(object, delimiter string) string {
	readBuffer := bytes.NewBufferString(object)
	reader := bufio.NewReader(readBuffer)
	stringReader := strings.NewReader(delimiter)
	delimited, _ := stringReader.ReadByte()
	delimitedStr, _ := reader.ReadString(delimited)
	return delimitedStr
}

// byObjectMetadataKey is a sortable interface for UploadMetadata slice
type byUploadMetadataKey []*UploadMetadata

func (b byUploadMetadataKey) Len() int           { return len(b) }
func (b byUploadMetadataKey) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byUploadMetadataKey) Less(i, j int) bool { return b[i].Object < b[j].Object }
