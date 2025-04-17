#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/left.h"
#include "service6/types/service6.findResultRow.h"
#include "service6/types/service6.error.h"

namespace tl2 { namespace details { 

void LeftService6ErrorVectorService6FindResultRowReset(::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept;

bool LeftService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept;
bool LeftService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept; 
bool LeftService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept;
bool LeftService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool LeftService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);

}} // namespace tl2::details

