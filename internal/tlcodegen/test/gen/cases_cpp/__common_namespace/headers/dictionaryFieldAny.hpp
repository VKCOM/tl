#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/dictionaryFieldAny.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldAnyDoubleIntReset(std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyDoubleIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyDoubleIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldAnyIntIntReset(std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyIntIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyIntIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldAnyDoubleIntReset(::tl2::DictionaryFieldAny<double, int32_t>& item);
bool DictionaryFieldAnyDoubleIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<double, int32_t>& item);
bool DictionaryFieldAnyDoubleIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<double, int32_t>& item);
bool DictionaryFieldAnyDoubleIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<double, int32_t>& item);
bool DictionaryFieldAnyDoubleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<double, int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldAnyIntIntReset(::tl2::DictionaryFieldAny<int32_t, int32_t>& item);
bool DictionaryFieldAnyIntIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<int32_t, int32_t>& item);
bool DictionaryFieldAnyIntIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item);
bool DictionaryFieldAnyIntIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<int32_t, int32_t>& item);
bool DictionaryFieldAnyIntIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item);

}} // namespace tl2::details

