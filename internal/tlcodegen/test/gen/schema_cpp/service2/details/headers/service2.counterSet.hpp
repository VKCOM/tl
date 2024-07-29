#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../types/service2.counterSet.hpp"

namespace tl2 { namespace details { 

void BuiltinTupleService2CounterSetReset(std::vector<::tl2::service2::CounterSet>& item);
bool BuiltinTupleService2CounterSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_n, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum);
bool BuiltinTupleService2CounterSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_n, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service2CounterSetReset(::tl2::service2::CounterSet& item);
bool Service2CounterSetRead(::basictl::tl_istream & s, ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
bool Service2CounterSetWrite(::basictl::tl_ostream & s, const ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
bool Service2CounterSetReadBoxed(::basictl::tl_istream & s, ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
bool Service2CounterSetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);

}} // namespace tl2::details

