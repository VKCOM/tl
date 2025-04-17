#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service6/functions/service6.multiFindWithBounds.h"
#include "service6/types/service6.findWithBoundsResult.h"
#include "__common_namespace/types/Either.h"

namespace tl2 { namespace details { 

void Service6MultiFindWithBoundsReset(::tl2::service6::MultiFindWithBounds& item) noexcept;

bool Service6MultiFindWithBoundsWriteJSON(std::ostream& s, const ::tl2::service6::MultiFindWithBounds& item) noexcept;
bool Service6MultiFindWithBoundsRead(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item) noexcept; 
bool Service6MultiFindWithBoundsWrite(::basictl::tl_ostream & s, const ::tl2::service6::MultiFindWithBounds& item) noexcept;
bool Service6MultiFindWithBoundsReadBoxed(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item);
bool Service6MultiFindWithBoundsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::MultiFindWithBounds& item);

bool Service6MultiFindWithBoundsReadResult(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& result);
bool Service6MultiFindWithBoundsWriteResult(::basictl::tl_ostream & s, ::tl2::service6::MultiFindWithBounds& item, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& result);
		
}} // namespace tl2::details

