#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service3/types/service3.groupCountLimit.h"

namespace tl2 { namespace details { 

void BuiltinVectorService3GroupCountLimitReset(std::vector<::tl2::service3::GroupCountLimit>& item);

bool BuiltinVectorService3GroupCountLimitWriteJSON(std::ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item);
bool BuiltinVectorService3GroupCountLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupCountLimit>& item);
bool BuiltinVectorService3GroupCountLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service3GroupCountLimitReset(::tl2::service3::GroupCountLimit& item) noexcept;

bool Service3GroupCountLimitWriteJSON(std::ostream& s, const ::tl2::service3::GroupCountLimit& item) noexcept;
bool Service3GroupCountLimitRead(::basictl::tl_istream & s, ::tl2::service3::GroupCountLimit& item) noexcept; 
bool Service3GroupCountLimitWrite(::basictl::tl_ostream & s, const ::tl2::service3::GroupCountLimit& item) noexcept;
bool Service3GroupCountLimitReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GroupCountLimit& item);
bool Service3GroupCountLimitWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GroupCountLimit& item);

}} // namespace tl2::details

