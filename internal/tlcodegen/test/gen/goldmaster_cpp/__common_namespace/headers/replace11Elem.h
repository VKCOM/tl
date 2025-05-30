// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/replace11Elem.h"

namespace tl2 { namespace details { 

void BuiltinTupleReplace11ElemLongReset(std::vector<::tl2::Replace11Elem<int64_t>>& item);

bool BuiltinTupleReplace11ElemLongWriteJSON(std::ostream & s, const std::vector<::tl2::Replace11Elem<int64_t>>& item, uint32_t nat_n, uint32_t nat_tn, uint32_t nat_tk);
bool BuiltinTupleReplace11ElemLongRead(::basictl::tl_istream & s, std::vector<::tl2::Replace11Elem<int64_t>>& item, uint32_t nat_n, uint32_t nat_tn, uint32_t nat_tk);
bool BuiltinTupleReplace11ElemLongWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Replace11Elem<int64_t>>& item, uint32_t nat_n, uint32_t nat_tn, uint32_t nat_tk);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Replace11ElemLongReset(::tl2::Replace11Elem<int64_t>& item) noexcept;

bool Replace11ElemLongWriteJSON(std::ostream& s, const ::tl2::Replace11Elem<int64_t>& item, uint32_t nat_n, uint32_t nat_k) noexcept;
bool Replace11ElemLongRead(::basictl::tl_istream & s, ::tl2::Replace11Elem<int64_t>& item, uint32_t nat_n, uint32_t nat_k) noexcept; 
bool Replace11ElemLongWrite(::basictl::tl_ostream & s, const ::tl2::Replace11Elem<int64_t>& item, uint32_t nat_n, uint32_t nat_k) noexcept;

}} // namespace tl2::details

