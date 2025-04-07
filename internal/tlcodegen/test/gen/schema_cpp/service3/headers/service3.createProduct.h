#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/service3.createProduct.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service3CreateProductReset(::tl2::service3::CreateProduct& item);

bool Service3CreateProductWriteJSON(std::ostream& s, const ::tl2::service3::CreateProduct& item);
bool Service3CreateProductRead(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item);
bool Service3CreateProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::CreateProduct& item);
bool Service3CreateProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item);
bool Service3CreateProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::CreateProduct& item);

bool Service3CreateProductReadResult(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item, bool& result);
bool Service3CreateProductWriteResult(::basictl::tl_ostream & s, ::tl2::service3::CreateProduct& item, bool& result);
		
}} // namespace tl2::details

