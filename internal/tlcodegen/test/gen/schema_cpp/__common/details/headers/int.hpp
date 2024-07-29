#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/int.hpp"

namespace tl2 { namespace details { 

void BuiltinTupleIntReset(std::vector<int32_t>& item);
bool BuiltinTupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorIntReset(std::vector<int32_t>& item);
bool BuiltinVectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool BuiltinVectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool IntMaybeReadBoxed(::basictl::tl_istream & s, std::optional<int32_t>& item);
bool IntMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<int32_t>& item);


}} // namespace tl2::details

