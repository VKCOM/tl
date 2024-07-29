#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service3.restoreProduct.hpp"
#include "../../../__common/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service3RestoreProductReset(::tl2::service3::RestoreProduct& item);
bool Service3RestoreProductRead(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item);
bool Service3RestoreProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreProduct& item);
bool Service3RestoreProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item);
bool Service3RestoreProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreProduct& item);

bool Service3RestoreProductReadResult(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item, bool& result);
bool Service3RestoreProductWriteResult(::basictl::tl_ostream & s, ::tl2::service3::RestoreProduct& item, bool& result);
		
}} // namespace tl2::details

