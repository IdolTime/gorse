// Copyright 2022 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	"github.com/stretchr/testify/assert"
	"idolTime-gorse/base/task"
	"testing"
	"time"
)

func TestEncodeDecode(t *testing.T) {
	tk := &task.Task{
		Name:       "a",
		Total:      100,
		Done:       50,
		Status:     task.StatusRunning,
		StartTime:  time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local),
		FinishTime: time.Date(2018, time.January, 2, 0, 0, 0, 0, time.Local),
	}
	pb := EncodeTask(tk)
	assert.Equal(t, tk, DecodeTask(pb))
}
