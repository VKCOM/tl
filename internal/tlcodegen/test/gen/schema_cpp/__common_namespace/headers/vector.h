#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "service6/types/service6.findResultRow.h"
#include "service6/types/service6.error.h"
#include "service1/types/service1.Value.h"
#include "__common_namespace/types/Either.h"
#include "__common_namespace/types/string.h"
#include "__common_namespace/types/dictionaryField.h"
#include "__common_namespace/types/long.h"
#include "__common_namespace/types/int.h"
#include "__common_namespace/types/integer.h"

namespace tl2 { namespace details { 

void VectorDictionaryFieldIntReset(std::map<std::string, int32_t>& item) noexcept;

bool VectorDictionaryFieldIntWriteJSON(std::ostream& s, const std::map<std::string, int32_t>& item) noexcept;
bool VectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) noexcept; 
bool VectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) noexcept;
bool VectorDictionaryFieldIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, int32_t>& item);
bool VectorDictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept;

bool VectorEitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept;
bool VectorEitherService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept; 
bool VectorEitherService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept;
bool VectorEitherService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);
bool VectorEitherService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorIntReset(std::vector<int32_t>& item) noexcept;

bool VectorIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item) noexcept;
bool VectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item) noexcept; 
bool VectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) noexcept;
bool VectorIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool VectorIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorIntBoxedReset(std::vector<int32_t>& item) noexcept;

bool VectorIntBoxedWriteJSON(std::ostream& s, const std::vector<int32_t>& item) noexcept;
bool VectorIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item) noexcept; 
bool VectorIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) noexcept;
bool VectorIntBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool VectorIntBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorIntegerReset(std::vector<::tl2::Integer>& item) noexcept;

bool VectorIntegerWriteJSON(std::ostream& s, const std::vector<::tl2::Integer>& item) noexcept;
bool VectorIntegerRead(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item) noexcept; 
bool VectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item) noexcept;
bool VectorIntegerReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item);
bool VectorIntegerWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorLongBoxedReset(std::vector<int64_t>& item) noexcept;

bool VectorLongBoxedWriteJSON(std::ostream& s, const std::vector<int64_t>& item) noexcept;
bool VectorLongBoxedRead(::basictl::tl_istream & s, std::vector<int64_t>& item) noexcept; 
bool VectorLongBoxedWrite(::basictl::tl_ostream & s, const std::vector<int64_t>& item) noexcept;
bool VectorLongBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int64_t>& item);
bool VectorLongBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int64_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService1ValueReset(std::vector<::tl2::service1::Value>& item) noexcept;

bool VectorService1ValueWriteJSON(std::ostream& s, const std::vector<::tl2::service1::Value>& item) noexcept;
bool VectorService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item) noexcept; 
bool VectorService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item) noexcept;
bool VectorService1ValueReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item);
bool VectorService1ValueWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService6FindResultRowReset(std::vector<::tl2::service6::FindResultRow>& item) noexcept;

bool VectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tl2::service6::FindResultRow>& item) noexcept;
bool VectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item) noexcept; 
bool VectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item) noexcept;
bool VectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item);
bool VectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorStringReset(std::vector<std::string>& item) noexcept;

bool VectorStringWriteJSON(std::ostream& s, const std::vector<std::string>& item) noexcept;
bool VectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item) noexcept; 
bool VectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item) noexcept;
bool VectorStringReadBoxed(::basictl::tl_istream & s, std::vector<std::string>& item);
bool VectorStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tl2::details

