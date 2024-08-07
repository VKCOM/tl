#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.groupCountLimit.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorService3GroupCountLimitReset(std::vector<::tl2::service3::GroupCountLimit>& item);

bool BuiltinVectorService3GroupCountLimitWriteJSON(std::ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item);
bool BuiltinVectorService3GroupCountLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupCountLimit>& item);
bool BuiltinVectorService3GroupCountLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service3GroupCountLimitReset(::tl2::service3::GroupCountLimit& item);

bool Service3GroupCountLimitWriteJSON(std::ostream& s, const ::tl2::service3::GroupCountLimit& item);
bool Service3GroupCountLimitRead(::basictl::tl_istream & s, ::tl2::service3::GroupCountLimit& item);
bool Service3GroupCountLimitWrite(::basictl::tl_ostream & s, const ::tl2::service3::GroupCountLimit& item);
bool Service3GroupCountLimitReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GroupCountLimit& item);
bool Service3GroupCountLimitWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GroupCountLimit& item);

}} // namespace tl2::details

