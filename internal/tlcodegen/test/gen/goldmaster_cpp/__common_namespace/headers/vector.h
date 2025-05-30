// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/vector.h"
#include "__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void VectorIntReset(std::vector<int32_t>& item) noexcept;

bool VectorIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item) noexcept;
bool VectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item) noexcept; 
bool VectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) noexcept;
bool VectorIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool VectorIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool VectorIntBoxedMaybeWriteJSON(std::ostream & s, const std::optional<std::vector<int32_t>>& item);

bool VectorIntBoxedMaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<int32_t>>& item);
bool VectorIntBoxedMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<int32_t>>& item);


}} // namespace tl2::details

namespace tl2 { namespace details { 

bool VectorIntMaybeWriteJSON(std::ostream & s, const std::optional<std::vector<int32_t>>& item);

bool VectorIntMaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<int32_t>>& item);
bool VectorIntMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<int32_t>>& item);


}} // namespace tl2::details

