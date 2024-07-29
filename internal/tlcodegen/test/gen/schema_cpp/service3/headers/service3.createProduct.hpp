#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service3.createProduct.hpp"
#include "../../__common_namespace/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service3CreateProductReset(::tl2::service3::CreateProduct& item);
bool Service3CreateProductRead(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item);
bool Service3CreateProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::CreateProduct& item);
bool Service3CreateProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item);
bool Service3CreateProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::CreateProduct& item);

bool Service3CreateProductReadResult(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item, bool& result);
bool Service3CreateProductWriteResult(::basictl::tl_ostream & s, ::tl2::service3::CreateProduct& item, bool& result);
		
}} // namespace tl2::details

