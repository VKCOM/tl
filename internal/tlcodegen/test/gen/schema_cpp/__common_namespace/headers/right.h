#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/right.h"
#include "../../service6/types/service6.findResultRow.h"
#include "../../service6/types/service6.error.h"

namespace tl2 { namespace details { 

void RightService6ErrorVectorService6FindResultRowReset(::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept;

bool RightService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept;
bool RightService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept; 
bool RightService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept;
bool RightService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool RightService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);

}} // namespace tl2::details

