// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service3/functions/service3.restoreProduct.h"
#include "__common_namespace/types/Bool.h"

namespace tl2 { namespace details { 

void Service3RestoreProductReset(::tl2::service3::RestoreProduct& item) noexcept;

bool Service3RestoreProductWriteJSON(std::ostream& s, const ::tl2::service3::RestoreProduct& item) noexcept;
bool Service3RestoreProductRead(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item) noexcept; 
bool Service3RestoreProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreProduct& item) noexcept;
bool Service3RestoreProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item);
bool Service3RestoreProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreProduct& item);

bool Service3RestoreProductReadResult(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item, bool& result);
bool Service3RestoreProductWriteResult(::basictl::tl_ostream & s, ::tl2::service3::RestoreProduct& item, bool& result);
		
}} // namespace tl2::details

