// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/tuple.h"
#include "__common_namespace/types/string.h"

namespace tl2 { namespace details { 

void TupleStringReset(std::vector<std::string>& item) noexcept;

bool TupleStringWriteJSON(std::ostream& s, const std::vector<std::string>& item, uint32_t nat_n) noexcept;
bool TupleStringRead(::basictl::tl_istream & s, std::vector<std::string>& item, uint32_t nat_n) noexcept; 
bool TupleStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item, uint32_t nat_n) noexcept;
bool TupleStringReadBoxed(::basictl::tl_istream & s, std::vector<std::string>& item, uint32_t nat_n);
bool TupleStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<std::string>& item, uint32_t nat_n);

}} // namespace tl2::details

