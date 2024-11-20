#pragma once

#include "../../../basics/basictl.h"
#include "../types/Either.h"
#include "../../service6/types/service6.findWithBoundsResult.h"
#include "../../service6/types/service6.findResultRow.h"
#include "../../service6/types/service6.error.h"

namespace tl2 { namespace details { 

void BuiltinVectorEitherIntVectorService6FindWithBoundsResultReset(std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);

bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);

bool BuiltinVectorEitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);
bool BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);
bool BuiltinVectorEitherService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void EitherIntVectorService6FindWithBoundsResultReset(::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);

bool EitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool EitherIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);
bool EitherIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void EitherService6ErrorVectorService6FindResultRowReset(::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);

bool EitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream & s, const ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool EitherService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);
bool EitherService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item);

}} // namespace tl2::details

