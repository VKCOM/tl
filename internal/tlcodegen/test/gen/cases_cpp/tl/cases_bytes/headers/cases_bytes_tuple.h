#pragma once

#include "../../../basics/basictl.h"
#include "../../__common_namespace/types/tuple.h"
#include "../../__common_namespace/types/string.h"

namespace tl2 { namespace details { 

void TupleString4Reset(std::array<std::string, 4>& item);

bool TupleString4WriteJSON(std::ostream& s, const std::array<std::string, 4>& item);
bool TupleString4Read(::basictl::tl_istream & s, std::array<std::string, 4>& item);
bool TupleString4Write(::basictl::tl_ostream & s, const std::array<std::string, 4>& item);
bool TupleString4ReadBoxed(::basictl::tl_istream & s, std::array<std::string, 4>& item);
bool TupleString4WriteBoxed(::basictl::tl_ostream & s, const std::array<std::string, 4>& item);

}} // namespace tl2::details

