// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/Either.h"
#include "service6/types/service6.findWithBoundsResult.h"
#include "service6/types/service6.findResultRow.h"
#include "service6/types/service6.error.h"

namespace tlgen { namespace details { 

void BuiltinVectorEitherIntVectorService6FindWithBoundsResultReset(std::vector<::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>>& item);

bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const std::vector<::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>>& item);
bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultRead(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>>& item);
bool BuiltinVectorEitherIntVectorService6FindWithBoundsResultWrite(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>>& item);

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void BuiltinVectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item);

bool BuiltinVectorEitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream & s, const std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item);
bool BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(::tlgen::basictl::tl_istream & s, std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item);
bool BuiltinVectorEitherService6ErrorVectorService6FindResultRowWrite(::tlgen::basictl::tl_ostream & s, const std::vector<::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>>& item);

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void EitherIntVectorService6FindWithBoundsResultReset(::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>& item) noexcept;

bool EitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const ::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>& item) noexcept;
bool EitherIntVectorService6FindWithBoundsResultReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>& item) noexcept;
bool EitherIntVectorService6FindWithBoundsResultWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Either<int32_t, std::vector<::tlgen::service6::FindWithBoundsResult>>& item) noexcept;

}} // namespace tlgen::details

namespace tlgen { namespace details { 

void EitherService6ErrorVectorService6FindResultRowReset(::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>& item) noexcept;

bool EitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream & s, const ::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>& item) noexcept;
bool EitherService6ErrorVectorService6FindResultRowReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>& item) noexcept;
bool EitherService6ErrorVectorService6FindResultRowWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::Either<::tlgen::service6::Error, std::vector<::tlgen::service6::FindResultRow>>& item) noexcept;

}} // namespace tlgen::details

