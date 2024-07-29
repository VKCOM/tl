#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service6.multiFind.hpp"
#include "../types/service6.findResultRow.hpp"
#include "../types/service6.error.hpp"
#include "../../__common_namespace/types/Either.hpp"

namespace tl2 { namespace details { 

void Service6MultiFindReset(::tl2::service6::MultiFind& item);
bool Service6MultiFindRead(::basictl::tl_istream & s, ::tl2::service6::MultiFind& item);
bool Service6MultiFindWrite(::basictl::tl_ostream & s, const ::tl2::service6::MultiFind& item);
bool Service6MultiFindReadBoxed(::basictl::tl_istream & s, ::tl2::service6::MultiFind& item);
bool Service6MultiFindWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::MultiFind& item);

bool Service6MultiFindReadResult(::basictl::tl_istream & s, ::tl2::service6::MultiFind& item, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& result);
bool Service6MultiFindWriteResult(::basictl::tl_ostream & s, ::tl2::service6::MultiFind& item, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& result);
		
}} // namespace tl2::details

