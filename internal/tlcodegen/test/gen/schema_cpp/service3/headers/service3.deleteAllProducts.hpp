#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service3.deleteAllProducts.hpp"
#include "../../__common_namespace/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service3DeleteAllProductsReset(::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsRead(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::DeleteAllProducts& item);

bool Service3DeleteAllProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item, bool& result);
bool Service3DeleteAllProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::DeleteAllProducts& item, bool& result);
		
}} // namespace tl2::details

