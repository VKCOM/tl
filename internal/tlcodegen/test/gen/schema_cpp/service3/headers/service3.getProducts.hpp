#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service3.getProducts.hpp"
#include "../../__common_namespace/types/vector.hpp"
#include "../types/service3.product.hpp"

namespace tl2 { namespace details { 

void Service3GetProductsReset(::tl2::service3::GetProducts& item);
bool Service3GetProductsRead(::basictl::tl_istream & s, ::tl2::service3::GetProducts& item);
bool Service3GetProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetProducts& item);
bool Service3GetProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetProducts& item);
bool Service3GetProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetProducts& item);

bool Service3GetProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::GetProducts& item, std::optional<std::vector<::tl2::service3::Product>>& result);
bool Service3GetProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::GetProducts& item, std::optional<std::vector<::tl2::service3::Product>>& result);
		
}} // namespace tl2::details

