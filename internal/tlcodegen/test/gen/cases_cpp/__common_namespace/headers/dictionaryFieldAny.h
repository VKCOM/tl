#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/dictionaryFieldAny.h"

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldAnyDoubleIntReset(std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item);

bool BuiltinVectorDictionaryFieldAnyDoubleIntWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyDoubleIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyDoubleIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorDictionaryFieldAnyIntIntReset(std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);

bool BuiltinVectorDictionaryFieldAnyIntIntWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyIntIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);
bool BuiltinVectorDictionaryFieldAnyIntIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldAnyDoubleIntReset(::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept;

bool DictionaryFieldAnyDoubleIntWriteJSON(std::ostream& s, const ::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept;
bool DictionaryFieldAnyDoubleIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept; 
bool DictionaryFieldAnyDoubleIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept;
bool DictionaryFieldAnyDoubleIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<double, int32_t>& item);
bool DictionaryFieldAnyDoubleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<double, int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void DictionaryFieldAnyIntIntReset(::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept;

bool DictionaryFieldAnyIntIntWriteJSON(std::ostream& s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept;
bool DictionaryFieldAnyIntIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept; 
bool DictionaryFieldAnyIntIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept;
bool DictionaryFieldAnyIntIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<int32_t, int32_t>& item);
bool DictionaryFieldAnyIntIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item);

}} // namespace tl2::details

