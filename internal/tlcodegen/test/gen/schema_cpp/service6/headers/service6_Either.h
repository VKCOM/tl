#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/Either.h"
#include "service6/types/service6.findWithBoundsResult.h"

namespace tl2 { namespace details { 

void BuiltinVectorEitherIntVectorService6FindWithBoundsResultReset(std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);

bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void EitherIntVectorService6FindWithBoundsResultReset(::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;

bool EitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;
bool EitherIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;
bool EitherIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept;

}} // namespace tl2::details

