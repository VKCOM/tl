// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tltasks

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tl/tlTasksTaskInfoMaybe"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksAddTask"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksCronTask"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksCronTaskWithId"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksCronTime"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksGetAnyTask"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksGetQueueSize"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksGetQueueTypes"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksGetTaskFromQueue"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksQueueStats"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksQueueTypeInfo"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksQueueTypeSettings"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksQueueTypeStats"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTask"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTaskInfo"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTaskStatus"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTaskStatusInProgress"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTaskStatusNotCurrentlyInEngine"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTaskStatusScheduled"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTaskStatusWaiting"
)

type (
	AddTask           = tlTasksAddTask.TasksAddTask
	CronTask          = tlTasksCronTask.TasksCronTask
	CronTaskWithId    = tlTasksCronTaskWithId.TasksCronTaskWithId
	CronTime          = tlTasksCronTime.TasksCronTime
	GetAnyTask        = tlTasksGetAnyTask.TasksGetAnyTask
	GetQueueSize      = tlTasksGetQueueSize.TasksGetQueueSize
	GetQueueTypes     = tlTasksGetQueueTypes.TasksGetQueueTypes
	GetTaskFromQueue  = tlTasksGetTaskFromQueue.TasksGetTaskFromQueue
	QueueStats        = tlTasksQueueStats.TasksQueueStats
	QueueTypeInfo     = tlTasksQueueTypeInfo.TasksQueueTypeInfo
	QueueTypeSettings = tlTasksQueueTypeSettings.TasksQueueTypeSettings
	QueueTypeStats    = tlTasksQueueTypeStats.TasksQueueTypeStats
	Task              = tlTasksTask.TasksTask
	TaskInfo          = tlTasksTaskInfo.TasksTaskInfo
	TaskInfoMaybe     = tlTasksTaskInfoMaybe.TasksTaskInfoMaybe
	TaskStatus        = tlTasksTaskStatus.TasksTaskStatus
)

func TaskStatusInProgress() TaskStatus {
	return tlTasksTaskStatusInProgress.TasksTaskStatusInProgress()
}
func TaskStatusNotCurrentlyInEngine() TaskStatus {
	return tlTasksTaskStatusNotCurrentlyInEngine.TasksTaskStatusNotCurrentlyInEngine()
}
func TaskStatusScheduled() TaskStatus { return tlTasksTaskStatusScheduled.TasksTaskStatusScheduled() }
func TaskStatusWaiting() TaskStatus   { return tlTasksTaskStatusWaiting.TasksTaskStatusWaiting() }