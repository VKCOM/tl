#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/service3.restoreGroupedProducts.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service3RestoreGroupedProductsReset(::tl2::service3::RestoreGroupedProducts& item);

bool Service3RestoreGroupedProductsWriteJSON(std::ostream& s, const ::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsRead(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item);
bool Service3RestoreGroupedProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreGroupedProducts& item);

bool Service3RestoreGroupedProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item, bool& result);
bool Service3RestoreGroupedProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::RestoreGroupedProducts& item, bool& result);
		
}} // namespace tl2::details

