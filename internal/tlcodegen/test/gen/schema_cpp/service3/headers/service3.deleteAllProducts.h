#pragma once

#include "../../basictl/io_streams.h"
#include "../functions/service3.deleteAllProducts.h"
#include "../../__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service3DeleteAllProductsReset(::tl2::service3::DeleteAllProducts& item);

bool Service3DeleteAllProductsWriteJSON(std::ostream& s, const ::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsRead(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item);
bool Service3DeleteAllProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::DeleteAllProducts& item);

bool Service3DeleteAllProductsReadResult(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item, bool& result);
bool Service3DeleteAllProductsWriteResult(::basictl::tl_ostream & s, ::tl2::service3::DeleteAllProducts& item, bool& result);
		
}} // namespace tl2::details

