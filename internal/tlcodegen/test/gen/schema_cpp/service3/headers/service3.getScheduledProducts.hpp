#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service3.getScheduledProducts.hpp"
#include "../../__common_namespace/types/vector.hpp"
#include "../types/service3.product.hpp"

namespace tl2 { namespace details { 

void Service3GetScheduledProductsReset(::tl2::service3::GetScheduledProducts& item);
bool Service3GetScheduledProductsRead(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item);
bool Service3GetScheduledProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetScheduledProducts& item);
bool Service3GetScheduledProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item);
bool Service3GetScheduledProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetScheduledProducts& item);

bool Service3GetScheduledProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item, std::optional<std::vector<::tl2::service3::Productmode<0>>>& result);
bool Service3GetScheduledProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::GetScheduledProducts& item, std::optional<std::vector<::tl2::service3::Productmode<0>>>& result);
		
}} // namespace tl2::details

