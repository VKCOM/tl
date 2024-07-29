#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/right.hpp"
#include "../../../service6/types/service6.findWithBoundsResult.hpp"
#include "../../../service6/types/service6.findResultRow.hpp"
#include "../../../service6/types/service6.error.hpp"

namespace tl2 { namespace details { 

void RightIntVectorService6FindWithBoundsResultReset(::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool RightIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool RightIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool RightIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool RightIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void RightService6ErrorVectorService6FindResultRowReset(::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool RightService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool RightService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool RightService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool RightService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);

}} // namespace tl2::details

