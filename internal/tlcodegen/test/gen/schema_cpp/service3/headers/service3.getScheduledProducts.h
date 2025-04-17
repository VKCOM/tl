#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service3/functions/service3.getScheduledProducts.h"
#include "__common_namespace/types/vector.h"
#include "service3/types/service3.product.h"

namespace tl2 { namespace details { 

void Service3GetScheduledProductsReset(::tl2::service3::GetScheduledProducts& item) noexcept;

bool Service3GetScheduledProductsWriteJSON(std::ostream& s, const ::tl2::service3::GetScheduledProducts& item) noexcept;
bool Service3GetScheduledProductsRead(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item) noexcept; 
bool Service3GetScheduledProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetScheduledProducts& item) noexcept;
bool Service3GetScheduledProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item);
bool Service3GetScheduledProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetScheduledProducts& item);

bool Service3GetScheduledProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item, std::optional<std::vector<::tl2::service3::Productmode<0>>>& result);
bool Service3GetScheduledProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::GetScheduledProducts& item, std::optional<std::vector<::tl2::service3::Productmode<0>>>& result);
		
}} // namespace tl2::details

