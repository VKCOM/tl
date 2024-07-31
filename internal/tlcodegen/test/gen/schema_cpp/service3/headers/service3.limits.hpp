#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.limits.hpp"

namespace tl2 { namespace details { 

void Service3LimitsReset(::tl2::service3::Limits& item);

bool Service3LimitsWriteJSON(std::ostream& s, const ::tl2::service3::Limits& item);
bool Service3LimitsRead(::basictl::tl_istream & s, ::tl2::service3::Limits& item);
bool Service3LimitsWrite(::basictl::tl_ostream & s, const ::tl2::service3::Limits& item);
bool Service3LimitsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::Limits& item);
bool Service3LimitsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::Limits& item);

}} // namespace tl2::details

