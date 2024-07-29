#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service1.keysStat.hpp"

namespace tl2 { namespace details { 

void Service1KeysStatReset(::tl2::service1::KeysStat& item);
bool Service1KeysStatRead(::basictl::tl_istream & s, ::tl2::service1::KeysStat& item);
bool Service1KeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::KeysStat& item);
bool Service1KeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::KeysStat& item);
bool Service1KeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::KeysStat& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool Service1KeysStatMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::service1::KeysStat>& item);
bool Service1KeysStatMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::service1::KeysStat>& item);


}} // namespace tl2::details

