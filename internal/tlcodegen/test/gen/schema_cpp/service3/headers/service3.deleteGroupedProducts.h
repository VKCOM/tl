#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/service3.deleteGroupedProducts.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service3DeleteGroupedProductsReset(::tl2::service3::DeleteGroupedProducts& item);

bool Service3DeleteGroupedProductsWriteJSON(std::ostream& s, const ::tl2::service3::DeleteGroupedProducts& item);
bool Service3DeleteGroupedProductsRead(::basictl::tl_istream & s, ::tl2::service3::DeleteGroupedProducts& item);
bool Service3DeleteGroupedProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::DeleteGroupedProducts& item);
bool Service3DeleteGroupedProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::DeleteGroupedProducts& item);
bool Service3DeleteGroupedProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::DeleteGroupedProducts& item);

bool Service3DeleteGroupedProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::DeleteGroupedProducts& item, bool& result);
bool Service3DeleteGroupedProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::DeleteGroupedProducts& item, bool& result);
		
}} // namespace tl2::details

