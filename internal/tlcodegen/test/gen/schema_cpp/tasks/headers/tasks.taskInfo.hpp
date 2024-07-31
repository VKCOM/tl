#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tasks.taskInfo.hpp"

namespace tl2 { namespace details { 

void TasksTaskInfoReset(::tl2::tasks::TaskInfo& item);

bool TasksTaskInfoWriteJSON(std::ostream& s, const ::tl2::tasks::TaskInfo& item);
bool TasksTaskInfoRead(::basictl::tl_istream & s, ::tl2::tasks::TaskInfo& item);
bool TasksTaskInfoWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskInfo& item);
bool TasksTaskInfoReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskInfo& item);
bool TasksTaskInfoWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskInfo& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool TasksTaskInfoMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::tasks::TaskInfo>& item);

bool TasksTaskInfoMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::tasks::TaskInfo>& item);
bool TasksTaskInfoMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::tasks::TaskInfo>& item);


}} // namespace tl2::details

