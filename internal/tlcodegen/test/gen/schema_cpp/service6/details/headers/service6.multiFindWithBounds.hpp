#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service6.multiFindWithBounds.hpp"
#include "../../types/service6.findWithBoundsResult.hpp"
#include "../../../__common/types/Either.hpp"

namespace tl2 { namespace details { 

void Service6MultiFindWithBoundsReset(::tl2::service6::MultiFindWithBounds& item);
bool Service6MultiFindWithBoundsRead(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item);
bool Service6MultiFindWithBoundsWrite(::basictl::tl_ostream & s, const ::tl2::service6::MultiFindWithBounds& item);
bool Service6MultiFindWithBoundsReadBoxed(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item);
bool Service6MultiFindWithBoundsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::MultiFindWithBounds& item);

bool Service6MultiFindWithBoundsReadResult(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& result);
bool Service6MultiFindWithBoundsWriteResult(::basictl::tl_ostream & s, ::tl2::service6::MultiFindWithBounds& item, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& result);
		
}} // namespace tl2::details

