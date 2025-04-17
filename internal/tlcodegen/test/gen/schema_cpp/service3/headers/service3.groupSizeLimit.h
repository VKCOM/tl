#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service3/types/service3.groupSizeLimit.h"

namespace tl2 { namespace details { 

void BuiltinVectorService3GroupSizeLimitReset(std::vector<::tl2::service3::GroupSizeLimit>& item);

bool BuiltinVectorService3GroupSizeLimitWriteJSON(std::ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item);
bool BuiltinVectorService3GroupSizeLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupSizeLimit>& item);
bool BuiltinVectorService3GroupSizeLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service3GroupSizeLimitReset(::tl2::service3::GroupSizeLimit& item) noexcept;

bool Service3GroupSizeLimitWriteJSON(std::ostream& s, const ::tl2::service3::GroupSizeLimit& item) noexcept;
bool Service3GroupSizeLimitRead(::basictl::tl_istream & s, ::tl2::service3::GroupSizeLimit& item) noexcept; 
bool Service3GroupSizeLimitWrite(::basictl::tl_ostream & s, const ::tl2::service3::GroupSizeLimit& item) noexcept;
bool Service3GroupSizeLimitReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GroupSizeLimit& item);
bool Service3GroupSizeLimitWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GroupSizeLimit& item);

}} // namespace tl2::details

