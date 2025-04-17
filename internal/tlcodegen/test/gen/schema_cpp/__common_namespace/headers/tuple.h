#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/tuple.h"
#include "service1/types/service1.Value.h"
#include "__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void TupleIntReset(std::vector<int32_t>& item) noexcept;

bool TupleIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept;
bool TupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) noexcept; 
bool TupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept;
bool TupleIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleInt3Reset(std::array<int32_t, 3>& item) noexcept;

bool TupleInt3WriteJSON(std::ostream& s, const std::array<int32_t, 3>& item) noexcept;
bool TupleInt3Read(::basictl::tl_istream & s, std::array<int32_t, 3>& item) noexcept; 
bool TupleInt3Write(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item) noexcept;
bool TupleInt3ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 3>& item);
bool TupleInt3WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleIntBoxedReset(std::vector<int32_t>& item) noexcept;

bool TupleIntBoxedWriteJSON(std::ostream& s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept;
bool TupleIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) noexcept; 
bool TupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept;
bool TupleIntBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleIntBoxed10Reset(std::array<int32_t, 10>& item) noexcept;

bool TupleIntBoxed10WriteJSON(std::ostream& s, const std::array<int32_t, 10>& item) noexcept;
bool TupleIntBoxed10Read(::basictl::tl_istream & s, std::array<int32_t, 10>& item) noexcept; 
bool TupleIntBoxed10Write(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item) noexcept;
bool TupleIntBoxed10ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 10>& item);
bool TupleIntBoxed10WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleIntBoxed2Reset(std::array<int32_t, 2>& item) noexcept;

bool TupleIntBoxed2WriteJSON(std::ostream& s, const std::array<int32_t, 2>& item) noexcept;
bool TupleIntBoxed2Read(::basictl::tl_istream & s, std::array<int32_t, 2>& item) noexcept; 
bool TupleIntBoxed2Write(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) noexcept;
bool TupleIntBoxed2ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 2>& item);
bool TupleIntBoxed2WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleService1Value3Reset(std::array<::tl2::service1::Value, 3>& item) noexcept;

bool TupleService1Value3WriteJSON(std::ostream& s, const std::array<::tl2::service1::Value, 3>& item) noexcept;
bool TupleService1Value3Read(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item) noexcept; 
bool TupleService1Value3Write(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item) noexcept;
bool TupleService1Value3ReadBoxed(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item);
bool TupleService1Value3WriteBoxed(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item);

}} // namespace tl2::details

