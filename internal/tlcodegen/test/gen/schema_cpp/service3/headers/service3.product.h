#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/service3.product.h"

namespace tl2 { namespace details { 

void BuiltinVectorService3ProductReset(std::vector<::tl2::service3::Product>& item);

bool BuiltinVectorService3ProductWriteJSON(std::ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t);
bool BuiltinVectorService3ProductRead(::basictl::tl_istream & s, std::vector<::tl2::service3::Product>& item, uint32_t nat_t);
bool BuiltinVectorService3ProductWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorService3Product0Reset(std::vector<::tl2::service3::Productmode<0>>& item);

bool BuiltinVectorService3Product0WriteJSON(std::ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item);
bool BuiltinVectorService3Product0Read(::basictl::tl_istream & s, std::vector<::tl2::service3::Productmode<0>>& item);
bool BuiltinVectorService3Product0Write(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service3ProductReset(::tl2::service3::Product& item);

bool Service3ProductWriteJSON(std::ostream& s, const ::tl2::service3::Product& item, uint32_t nat_mode);
bool Service3ProductRead(::basictl::tl_istream & s, ::tl2::service3::Product& item, uint32_t nat_mode);
bool Service3ProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::Product& item, uint32_t nat_mode);
bool Service3ProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::Product& item, uint32_t nat_mode);
bool Service3ProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::Product& item, uint32_t nat_mode);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void Service3Product0Reset(::tl2::service3::Productmode<0>& item);

bool Service3Product0WriteJSON(std::ostream& s, const ::tl2::service3::Productmode<0>& item);
bool Service3Product0Read(::basictl::tl_istream & s, ::tl2::service3::Productmode<0>& item);
bool Service3Product0Write(::basictl::tl_ostream & s, const ::tl2::service3::Productmode<0>& item);
bool Service3Product0ReadBoxed(::basictl::tl_istream & s, ::tl2::service3::Productmode<0>& item);
bool Service3Product0WriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::Productmode<0>& item);

}} // namespace tl2::details

