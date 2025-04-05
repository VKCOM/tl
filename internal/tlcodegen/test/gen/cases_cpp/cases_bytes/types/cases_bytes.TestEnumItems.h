#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace cases_bytes { 
struct TestEnum1 {

	std::string_view tl_name() const { return "cases_bytes.testEnum1"; }
	uint32_t tl_tag() const { return 0x58aad3f5; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestEnum1& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases_bytes

namespace tl2 { namespace cases_bytes { 
struct TestEnum2 {

	std::string_view tl_name() const { return "cases_bytes.testEnum2"; }
	uint32_t tl_tag() const { return 0x00b47add; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestEnum2& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases_bytes

namespace tl2 { namespace cases_bytes { 
struct TestEnum3 {

	std::string_view tl_name() const { return "cases_bytes.testEnum3"; }
	uint32_t tl_tag() const { return 0x81911ffa; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestEnum3& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases_bytes

