#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/service1.Value.hpp"

namespace tl2 { namespace details { 

void BuiltinTuple3Service1ValueReset(std::array<::tl2::service1::Value, 3>& item);
bool BuiltinTuple3Service1ValueRead(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item);
bool BuiltinTuple3Service1ValueWrite(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorService1ValueReset(std::vector<::tl2::service1::Value>& item);
bool BuiltinVectorService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item);
bool BuiltinVectorService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service1ValueReset(::tl2::service1::Value& item);
bool Service1ValueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Value& item);
bool Service1ValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Value& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool Service1ValueBoxedMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::service1::Value>& item);
bool Service1ValueBoxedMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::service1::Value>& item);


}} // namespace tl2::details

