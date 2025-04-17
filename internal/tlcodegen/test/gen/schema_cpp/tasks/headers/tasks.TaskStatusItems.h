// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.TaskStatusItems.h"

namespace tl2 { namespace details { 

void TasksTaskStatusInProgressReset(::tl2::tasks::TaskStatusInProgress& item) noexcept;

bool TasksTaskStatusInProgressWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusInProgress& item) noexcept;
bool TasksTaskStatusInProgressRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusInProgress& item) noexcept; 
bool TasksTaskStatusInProgressWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusInProgress& item) noexcept;
bool TasksTaskStatusInProgressReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusInProgress& item);
bool TasksTaskStatusInProgressWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusInProgress& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TasksTaskStatusNotCurrentlyInEngineReset(::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) noexcept;

bool TasksTaskStatusNotCurrentlyInEngineWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) noexcept;
bool TasksTaskStatusNotCurrentlyInEngineRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) noexcept; 
bool TasksTaskStatusNotCurrentlyInEngineWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) noexcept;
bool TasksTaskStatusNotCurrentlyInEngineReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);
bool TasksTaskStatusNotCurrentlyInEngineWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TasksTaskStatusScheduledReset(::tl2::tasks::TaskStatusScheduled& item) noexcept;

bool TasksTaskStatusScheduledWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusScheduled& item) noexcept;
bool TasksTaskStatusScheduledRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusScheduled& item) noexcept; 
bool TasksTaskStatusScheduledWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusScheduled& item) noexcept;
bool TasksTaskStatusScheduledReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusScheduled& item);
bool TasksTaskStatusScheduledWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusScheduled& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TasksTaskStatusWaitingReset(::tl2::tasks::TaskStatusWaiting& item) noexcept;

bool TasksTaskStatusWaitingWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusWaiting& item) noexcept;
bool TasksTaskStatusWaitingRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusWaiting& item) noexcept; 
bool TasksTaskStatusWaitingWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusWaiting& item) noexcept;
bool TasksTaskStatusWaitingReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusWaiting& item);
bool TasksTaskStatusWaitingWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusWaiting& item);

}} // namespace tl2::details

