// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "service6/types/service6.findWithBoundsResult.h"

namespace tl2 { namespace details { 

void VectorService6FindWithBoundsResultReset(std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept;

bool VectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept;
bool VectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept; 
bool VectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept;
bool VectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item);
bool VectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item);

}} // namespace tl2::details

