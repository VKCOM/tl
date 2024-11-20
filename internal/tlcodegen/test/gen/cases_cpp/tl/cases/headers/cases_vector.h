#pragma once

#include "../../../basics/basictl.h"
#include "../../__common_namespace/types/vector.h"
#include "../../__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void VectorIntReset(std::vector<int32_t>& item);

bool VectorIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item);
bool VectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool VectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item);
bool VectorIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool VectorIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

