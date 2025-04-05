#pragma once

#include "../../basictl/io_streams.h"
#include "../types/tasks.TaskStatusItems.h"

namespace tl2 { namespace details { 

void TasksTaskStatusInProgressReset(::tl2::tasks::TaskStatusInProgress& item);

bool TasksTaskStatusInProgressWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusInProgress& item);
bool TasksTaskStatusInProgressRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusInProgress& item);
bool TasksTaskStatusInProgressWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusInProgress& item);
bool TasksTaskStatusInProgressReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusInProgress& item);
bool TasksTaskStatusInProgressWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusInProgress& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TasksTaskStatusNotCurrentlyInEngineReset(::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);

bool TasksTaskStatusNotCurrentlyInEngineWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);
bool TasksTaskStatusNotCurrentlyInEngineRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);
bool TasksTaskStatusNotCurrentlyInEngineWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);
bool TasksTaskStatusNotCurrentlyInEngineReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);
bool TasksTaskStatusNotCurrentlyInEngineWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TasksTaskStatusScheduledReset(::tl2::tasks::TaskStatusScheduled& item);

bool TasksTaskStatusScheduledWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusScheduled& item);
bool TasksTaskStatusScheduledRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusScheduled& item);
bool TasksTaskStatusScheduledWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusScheduled& item);
bool TasksTaskStatusScheduledReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusScheduled& item);
bool TasksTaskStatusScheduledWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusScheduled& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TasksTaskStatusWaitingReset(::tl2::tasks::TaskStatusWaiting& item);

bool TasksTaskStatusWaitingWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusWaiting& item);
bool TasksTaskStatusWaitingRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusWaiting& item);
bool TasksTaskStatusWaitingWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusWaiting& item);
bool TasksTaskStatusWaitingReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusWaiting& item);
bool TasksTaskStatusWaitingWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusWaiting& item);

}} // namespace tl2::details

