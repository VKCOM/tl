#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/tuple.hpp"
#include "../../types/service2.deltaSet.hpp"
#include "../../../__common/types/double.hpp"

namespace tl2 { namespace details { 

void TupleDoubleReset(std::vector<double>& item);
bool TupleDoubleRead(::basictl::tl_istream & s, std::vector<double>& item, uint32_t nat_n);
bool TupleDoubleWrite(::basictl::tl_ostream & s, const std::vector<double>& item, uint32_t nat_n);
bool TupleDoubleReadBoxed(::basictl::tl_istream & s, std::vector<double>& item, uint32_t nat_n);
bool TupleDoubleWriteBoxed(::basictl::tl_ostream & s, const std::vector<double>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleService2DeltaSetReset(std::vector<::tl2::service2::DeltaSet>& item);
bool TupleService2DeltaSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);
bool TupleService2DeltaSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);
bool TupleService2DeltaSetReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);
bool TupleService2DeltaSetWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n);

}} // namespace tl2::details

