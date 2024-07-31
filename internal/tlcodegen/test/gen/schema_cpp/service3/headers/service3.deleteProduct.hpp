#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/service3.deleteProduct.hpp"
#include "../../__common_namespace/types/Bool.hpp"

namespace tl2 { namespace details { 

void Service3DeleteProductReset(::tl2::service3::DeleteProduct& item);

bool Service3DeleteProductWriteJSON(std::ostream& s, const ::tl2::service3::DeleteProduct& item);
bool Service3DeleteProductRead(::basictl::tl_istream & s, ::tl2::service3::DeleteProduct& item);
bool Service3DeleteProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::DeleteProduct& item);
bool Service3DeleteProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::DeleteProduct& item);
bool Service3DeleteProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::DeleteProduct& item);

bool Service3DeleteProductReadResult(::basictl::tl_istream & s, ::tl2::service3::DeleteProduct& item, bool& result);
bool Service3DeleteProductWriteResult(::basictl::tl_ostream & s, ::tl2::service3::DeleteProduct& item, bool& result);
		
}} // namespace tl2::details

