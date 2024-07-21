#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/dictionaryAny.hpp"
#include "../../types/dictionaryFieldAny.hpp"

namespace tl2 { namespace details { 

void DictionaryAnyDoubleIntReset(::tl2::DictionaryAny<double, int32_t>& item);
bool DictionaryAnyDoubleIntRead(::basictl::tl_istream & s, ::tl2::DictionaryAny<double, int32_t>& item);
bool DictionaryAnyDoubleIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryAny<double, int32_t>& item);
bool DictionaryAnyDoubleIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryAny<double, int32_t>& item);
bool DictionaryAnyDoubleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryAny<double, int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryAnyIntIntReset(std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool DictionaryAnyIntIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool DictionaryAnyIntIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool DictionaryAnyIntIntReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool DictionaryAnyIntIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);

}} // namespace tl2::details

