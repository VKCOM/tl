#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service2.addOrIncrMany.hpp"
#include "../../types/service2.counterSet.hpp"

namespace tl2 { namespace details { 

void Service2AddOrIncrManyReset(::tl2::service2::AddOrIncrMany& item);
bool Service2AddOrIncrManyRead(::basictl::tl_istream & s, ::tl2::service2::AddOrIncrMany& item);
bool Service2AddOrIncrManyWrite(::basictl::tl_ostream & s, const ::tl2::service2::AddOrIncrMany& item);
bool Service2AddOrIncrManyReadBoxed(::basictl::tl_istream & s, ::tl2::service2::AddOrIncrMany& item);
bool Service2AddOrIncrManyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::AddOrIncrMany& item);

bool Service2AddOrIncrManyReadResult(::basictl::tl_istream & s, ::tl2::service2::AddOrIncrMany& item, std::vector<::tl2::service2::CounterSet>& result);
bool Service2AddOrIncrManyWriteResult(::basictl::tl_ostream & s, ::tl2::service2::AddOrIncrMany& item, std::vector<::tl2::service2::CounterSet>& result);
		
}} // namespace tl2::details

