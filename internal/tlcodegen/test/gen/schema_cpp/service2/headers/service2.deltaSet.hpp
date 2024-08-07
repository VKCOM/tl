#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service2.deltaSet.hpp"

namespace tl2 { namespace details { 

void BuiltinTupleService2DeltaSetReset(std::vector<::tl2::service2::DeltaSet>& item);

bool BuiltinTupleService2DeltaSetWriteJSON(std::ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_n, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum);
bool BuiltinTupleService2DeltaSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_n, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum);
bool BuiltinTupleService2DeltaSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_n, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service2DeltaSetReset(::tl2::service2::DeltaSet& item);

bool Service2DeltaSetWriteJSON(std::ostream& s, const ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
bool Service2DeltaSetRead(::basictl::tl_istream & s, ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
bool Service2DeltaSetWrite(::basictl::tl_ostream & s, const ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
bool Service2DeltaSetReadBoxed(::basictl::tl_istream & s, ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);
bool Service2DeltaSetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum);

}} // namespace tl2::details

