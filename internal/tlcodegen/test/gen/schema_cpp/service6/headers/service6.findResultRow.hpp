#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service6.findResultRow.hpp"

namespace tl2 { namespace details { 

void BuiltinVectorService6FindResultRowReset(std::vector<::tl2::service6::FindResultRow>& item);

bool BuiltinVectorService6FindResultRowWriteJSON(std::ostream & s, const std::vector<::tl2::service6::FindResultRow>& item);
bool BuiltinVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item);
bool BuiltinVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service6FindResultRowReset(::tl2::service6::FindResultRow& item);

bool Service6FindResultRowWriteJSON(std::ostream& s, const ::tl2::service6::FindResultRow& item);
bool Service6FindResultRowRead(::basictl::tl_istream & s, ::tl2::service6::FindResultRow& item);
bool Service6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::service6::FindResultRow& item);
bool Service6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::service6::FindResultRow& item);
bool Service6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::FindResultRow& item);

}} // namespace tl2::details

