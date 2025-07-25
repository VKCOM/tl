// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service3/functions/service3.restoreAllProducts.h"
#include "__common_namespace/types/Bool.h"

namespace tlgen { namespace details { 

void Service3RestoreAllProductsReset(::tlgen::service3::RestoreAllProducts& item) noexcept;

bool Service3RestoreAllProductsWriteJSON(std::ostream& s, const ::tlgen::service3::RestoreAllProducts& item) noexcept;
bool Service3RestoreAllProductsRead(::tlgen::basictl::tl_istream & s, ::tlgen::service3::RestoreAllProducts& item) noexcept; 
bool Service3RestoreAllProductsWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service3::RestoreAllProducts& item) noexcept;
bool Service3RestoreAllProductsReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service3::RestoreAllProducts& item);
bool Service3RestoreAllProductsWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service3::RestoreAllProducts& item);

bool Service3RestoreAllProductsReadResult(::tlgen::basictl::tl_istream & s, const ::tlgen::service3::RestoreAllProducts& item, bool& result);
bool Service3RestoreAllProductsWriteResult(::tlgen::basictl::tl_ostream & s, const ::tlgen::service3::RestoreAllProducts& item, const bool& result);
    
}} // namespace tlgen::details

