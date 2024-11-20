#pragma once

#include "../../../basics/basictl.h"
#include "../functions/service3.restoreAllProducts.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service3RestoreAllProductsReset(::tl2::service3::RestoreAllProducts& item);

bool Service3RestoreAllProductsWriteJSON(std::ostream& s, const ::tl2::service3::RestoreAllProducts& item);
bool Service3RestoreAllProductsRead(::basictl::tl_istream & s, ::tl2::service3::RestoreAllProducts& item);
bool Service3RestoreAllProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreAllProducts& item);
bool Service3RestoreAllProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreAllProducts& item);
bool Service3RestoreAllProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreAllProducts& item);

bool Service3RestoreAllProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::RestoreAllProducts& item, bool& result);
bool Service3RestoreAllProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::RestoreAllProducts& item, bool& result);
		
}} // namespace tl2::details

