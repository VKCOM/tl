#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../functions/service3.restoreGroupedProducts.hpp"
#include "../../../__common/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service3RestoreGroupedProductsReset(::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsRead(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreGroupedProducts& item);

bool Service3RestoreGroupedProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item, bool& result);
bool Service3RestoreGroupedProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::RestoreGroupedProducts& item, bool& result);
		
}} // namespace tl2::details

