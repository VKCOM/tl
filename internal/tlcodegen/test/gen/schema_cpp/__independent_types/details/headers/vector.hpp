#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/vector.hpp"
#include "../../../tasks/types/tasks.queueTypeInfo.hpp"
#include "../../../service6/types/service6.findWithBoundsResult.hpp"
#include "../../../__common/types/map.hpp"
#include "../../../service6/types/service6.findResultRow.hpp"
#include "../../../service6/types/service6.error.hpp"
#include "../../../service1/types/service1.Value.hpp"
#include "../../../__common/types/Either.hpp"
#include "../../../__common/types/string.hpp"
#include "../../../__common/types/long.hpp"
#include "../../../__common/types/integer.hpp"
#include "../../../__common/types/int.hpp"

namespace tl2 { namespace details { 

void VectorEitherIntVectorService6FindWithBoundsResultReset(std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool VectorEitherIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool VectorEitherIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool VectorEitherIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);
bool VectorEitherIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);
bool VectorEitherService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);
bool VectorEitherService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);
bool VectorEitherService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);
bool VectorEitherService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorIntBoxedReset(std::vector<int32_t>& item);
bool VectorIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool VectorIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item);
bool VectorIntBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool VectorIntBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorIntegerReset(std::vector<::tl2::Integer>& item);
bool VectorIntegerRead(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item);
bool VectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item);
bool VectorIntegerReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item);
bool VectorIntegerWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorLongBoxedReset(std::vector<int64_t>& item);
bool VectorLongBoxedRead(::basictl::tl_istream & s, std::vector<int64_t>& item);
bool VectorLongBoxedWrite(::basictl::tl_ostream & s, const std::vector<int64_t>& item);
bool VectorLongBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int64_t>& item);
bool VectorLongBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int64_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorMapStringStringReset(std::vector<::tl2::Map<std::string, std::string>>& item);
bool VectorMapStringStringRead(::basictl::tl_istream & s, std::vector<::tl2::Map<std::string, std::string>>& item);
bool VectorMapStringStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item);
bool VectorMapStringStringReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Map<std::string, std::string>>& item);
bool VectorMapStringStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService1ValueReset(std::vector<::tl2::service1::Value>& item);
bool VectorService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item);
bool VectorService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item);
bool VectorService1ValueReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item);
bool VectorService1ValueWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService6FindResultRowReset(std::vector<::tl2::service6::FindResultRow>& item);
bool VectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item);
bool VectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item);
bool VectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item);
bool VectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService6FindWithBoundsResultReset(std::vector<::tl2::service6::FindWithBoundsResult>& item);
bool VectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item);
bool VectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item);
bool VectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item);
bool VectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorStringReset(std::vector<std::string>& item);
bool VectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item);
bool VectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item);
bool VectorStringReadBoxed(::basictl::tl_istream & s, std::vector<std::string>& item);
bool VectorStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<std::string>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorTasksQueueTypeInfoReset(std::vector<::tl2::tasks::QueueTypeInfo>& item);
bool VectorTasksQueueTypeInfoRead(::basictl::tl_istream & s, std::vector<::tl2::tasks::QueueTypeInfo>& item);
bool VectorTasksQueueTypeInfoWrite(::basictl::tl_ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item);
bool VectorTasksQueueTypeInfoReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::tasks::QueueTypeInfo>& item);
bool VectorTasksQueueTypeInfoWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item);

}} // namespace tl2::details

