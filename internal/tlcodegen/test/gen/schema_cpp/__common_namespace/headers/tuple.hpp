#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/tuple.hpp"
#include "../../service2/types/service2.counterSet.hpp"
#include "../../service1/types/service1.Value.hpp"
#include "../types/int.hpp"

namespace tl2 { namespace details { 

void TupleIntReset(std::vector<int32_t>& item);
bool TupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleInt3Reset(std::array<int32_t, 3>& item);
bool TupleInt3Read(::basictl::tl_istream & s, std::array<int32_t, 3>& item);
bool TupleInt3Write(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item);
bool TupleInt3ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 3>& item);
bool TupleInt3WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleInt5Reset(std::array<int32_t, 5>& item);
bool TupleInt5Read(::basictl::tl_istream & s, std::array<int32_t, 5>& item);
bool TupleInt5Write(::basictl::tl_ostream & s, const std::array<int32_t, 5>& item);
bool TupleInt5ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 5>& item);
bool TupleInt5WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 5>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleIntBoxedReset(std::vector<int32_t>& item);
bool TupleIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleIntBoxed10Reset(std::array<int32_t, 10>& item);
bool TupleIntBoxed10Read(::basictl::tl_istream & s, std::array<int32_t, 10>& item);
bool TupleIntBoxed10Write(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item);
bool TupleIntBoxed10ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 10>& item);
bool TupleIntBoxed10WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleIntBoxed2Reset(std::array<int32_t, 2>& item);
bool TupleIntBoxed2Read(::basictl::tl_istream & s, std::array<int32_t, 2>& item);
bool TupleIntBoxed2Write(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item);
bool TupleIntBoxed2ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 2>& item);
bool TupleIntBoxed2WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleService1Value3Reset(std::array<::tl2::service1::Value, 3>& item);
bool TupleService1Value3Read(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item);
bool TupleService1Value3Write(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item);
bool TupleService1Value3ReadBoxed(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item);
bool TupleService1Value3WriteBoxed(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleService2CounterSetReset(std::vector<::tl2::service2::CounterSet>& item);
bool TupleService2CounterSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);
bool TupleService2CounterSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);
bool TupleService2CounterSetReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);
bool TupleService2CounterSetWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);

}} // namespace tl2::details

