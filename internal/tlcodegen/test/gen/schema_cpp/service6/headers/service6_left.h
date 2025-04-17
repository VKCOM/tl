#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/left.h"
#include "service6/types/service6.findWithBoundsResult.h"

namespace tl2 { namespace details { 

void LeftIntVectorService6FindWithBoundsResultReset(::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;

bool LeftIntVectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;
bool LeftIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept; 
bool LeftIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;
bool LeftIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool LeftIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);

}} // namespace tl2::details

