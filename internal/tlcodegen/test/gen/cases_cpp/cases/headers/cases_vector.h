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

