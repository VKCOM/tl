#pragma once

#include "../../../basics/basictl.h"
#include "../types/service6.findWithBoundsResult.h"

namespace tl2 { namespace details { 

void BuiltinVectorService6FindWithBoundsResultReset(std::vector<::tl2::service6::FindWithBoundsResult>& item);

bool BuiltinVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item);
bool BuiltinVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item);
bool BuiltinVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service6FindWithBoundsResultReset(::tl2::service6::FindWithBoundsResult& item);

bool Service6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::service6::FindWithBoundsResult& item);
bool Service6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::service6::FindWithBoundsResult& item);
bool Service6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::service6::FindWithBoundsResult& item);
bool Service6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::service6::FindWithBoundsResult& item);
bool Service6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::FindWithBoundsResult& item);

}} // namespace tl2::details

