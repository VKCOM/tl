#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace cases { 
struct TestEnum1 {

	std::string_view tl_name() const { return "cases.testEnum1"; }
	uint32_t tl_tag() const { return 0x6c6c55ac; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestEnum1& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

namespace tl2 { namespace cases { 
struct TestEnum2 {

	std::string_view tl_name() const { return "cases.testEnum2"; }
	uint32_t tl_tag() const { return 0x86ea88ce; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestEnum2& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

namespace tl2 { namespace cases { 
struct TestEnum3 {

	std::string_view tl_name() const { return "cases.testEnum3"; }
	uint32_t tl_tag() const { return 0x69b83e2f; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestEnum3& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

