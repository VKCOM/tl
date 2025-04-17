#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/right.h"
#include "service6/types/service6.findWithBoundsResult.h"

namespace tl2 { namespace details { 

void RightIntVectorService6FindWithBoundsResultReset(::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;

bool RightIntVectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;
bool RightIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept; 
bool RightIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;
bool RightIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool RightIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);

}} // namespace tl2::details

