#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/tasks.taskInfo.h"

namespace tl2 { namespace details { 

void TasksTaskInfoReset(::tl2::tasks::TaskInfo& item) noexcept;

bool TasksTaskInfoWriteJSON(std::ostream& s, const ::tl2::tasks::TaskInfo& item) noexcept;
bool TasksTaskInfoRead(::basictl::tl_istream & s, ::tl2::tasks::TaskInfo& item) noexcept; 
bool TasksTaskInfoWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskInfo& item) noexcept;
bool TasksTaskInfoReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskInfo& item);
bool TasksTaskInfoWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskInfo& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool TasksTaskInfoMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::tasks::TaskInfo>& item);

bool TasksTaskInfoMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::tasks::TaskInfo>& item);
bool TasksTaskInfoMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::tasks::TaskInfo>& item);


}} // namespace tl2::details

