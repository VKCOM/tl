// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

const basicCPPTLIOStreamsPath = "basictl/io_streams.h"
const basicCPPTLIOThrowableStreamsPath = "basictl/io_throwable_streams.h"

const basicCPPTLCodeHeader = `%s
#pragma once

#include <ostream>
#include <stddef.h>
#include <cstring>
#include <stdexcept>
#include <array>
#include <string>
#include <vector>
#include <utility>
#include <memory>
#include <optional>
#include <variant>

#define TLGEN2_UNLIKELY(x) (x) // __builtin_expect((x), 0) // could improve performance on your platform
#define TLGEN2_NOINLINE // __attribute__ ((noinline)) // could improve performance on your platform

namespace %s {

`

const basicCPPTLCodeFooter = `
} // namespace %s


#undef TLGEN2_NOINLINE
#undef TLGEN2_UNLIKELY
`

const basicCPPTLCodeBody = `

enum {
	TL_MAX_TINY_STRING_LEN          = 253,
	TL_BIG_STRING_LEN               = 0xffffff,
	TL_BIG_STRING_MARKER            = 0xfe,
};

class tl_istream { // TODO - prohibit copy/move
public:
	virtual ~tl_istream() = default;
	bool nat_read(uint32_t & value) {
       	return fetch_data(&value, 4);
	}
	bool nat_read_exact_tag(uint32_t tag) {
	    uint32_t actual_tag = 0;
       	if (TLGEN2_UNLIKELY(!nat_read(actual_tag))) { return false; }
		if (TLGEN2_UNLIKELY(tag != actual_tag)) { return set_error_expected_tag(); }
		return true;
	}
	bool int_read(int32_t & value) {
	    return fetch_data(&value, 4);
	}
	bool long_read(int64_t & value) {
       	return fetch_data(&value, 8);
	}
	bool float_read(float & value) {
       	return fetch_data(&value, 4);
	}
	bool double_read(double & value) {
       	return fetch_data(&value, 8);
	}
	bool bool_read(bool & value, uint32_t f, uint32_t t) {
	    uint32_t tag = 0;
       	if (TLGEN2_UNLIKELY(!nat_read(tag))) { return false; }
		if (tag == t) { value = true; return true; }
       	if (TLGEN2_UNLIKELY(tag != f)) { set_error_bool_tag(); }
		value = false;
		return true;
	}
	bool string_read(std::string & value) {
		if (TLGEN2_UNLIKELY(!ensure_byte())) { return false; }
		auto len = size_t(static_cast<unsigned char>(*ptr));
		if (TLGEN2_UNLIKELY(len >= TL_BIG_STRING_MARKER)) {
			if (TLGEN2_UNLIKELY(len > TL_BIG_STRING_MARKER)) {
				return set_error("TODO - huge string");
			}
			uint32_t len32 = 0;
		   	if (TLGEN2_UNLIKELY(!nat_read(len32))) { return false; }
			len = len32 >> 8U;
			value.clear();
		   	if (TLGEN2_UNLIKELY(!fetch_data_append(value, len))) { return false; }
		   	if (TLGEN2_UNLIKELY(!fetch_pad((-len) & 3))) { return false; }
			return true;
		}
		auto pad = ((-(len+1)) & 3);
		auto fullLen = 1 + len + pad;
		if (TLGEN2_UNLIKELY(ptr + fullLen > end)) {
			ptr += 1;
			value.clear();
		   	if (TLGEN2_UNLIKELY(!fetch_data_append(value, len))) { return false; }
		   	if (TLGEN2_UNLIKELY(!fetch_pad(pad))) { return false; }
			return true;
		}
		// fast path for short strings that fully fit in buffer
		uint32_t x = 0;
		std::memcpy(&x, ptr + fullLen - 4, 4);
		if (TLGEN2_UNLIKELY((x & ~(0xFFFFFFFFU >> (8*pad))) != 0)) {
				return set_error_string_padding();
		}
		value.assign(ptr + 1, len);
		ptr += fullLen;
		return true;
	}
    bool set_error(const char * e) { return false; } // TODO - set error field
	bool set_error_eof() { return set_error("EOF"); }
	bool set_error_sequence_length() { return set_error("sequence_length"); }
	bool set_error_string_padding() { return set_error("string_padding"); }
	bool set_error_bool_tag() { return set_error("bool_tag"); }
	bool set_error_expected_tag() { return set_error("expected_tag"); }
	bool set_error_union_tag() { return set_error("union_tag"); }
protected:
	const char * ptr{};
	const char * end{};
	virtual void grow_buffer(size_t size_hint) = 0; // after call buffer must contain at least single byte, otherwise error
private:
	bool ensure_byte() {
		if (TLGEN2_UNLIKELY(ptr >= end)) {
		    // assert(ptr <= end)
			grow_buffer(1);
		    // assert(ptr <= end)
			if (TLGEN2_UNLIKELY(ptr == end)) {
				return set_error_eof();
			}
		}
		return true;
	}
	bool fetch_data(void * vdata, size_t size) {
    	char * data = reinterpret_cast<char *>(vdata);
		if (TLGEN2_UNLIKELY(ptr + size > end)) {
    	    return fetch_data2(vdata, size);
	    }
   		std::memcpy(data, ptr, size);
    	ptr += size;
	    return true;
    }
	bool fetch_data2(void * vdata, size_t size) {
    	char * data = reinterpret_cast<char *>(vdata);
		for (;TLGEN2_UNLIKELY(ptr + size > end);) {
		    // assert(ptr <= end)
			std::memcpy(data, ptr, end - ptr);
			data += end - ptr;
			size -= end - ptr;
			ptr = end;
			grow_buffer(size);
		    // assert(ptr <= end)
			if (TLGEN2_UNLIKELY(ptr == end)) {
				return set_error_eof();
			}
		}
   		std::memcpy(data, ptr, size);
		ptr += size;
		return true;
	}
	bool fetch_data_append(std::string & value, size_t size) {
		for (;TLGEN2_UNLIKELY(ptr + size > end);) {
		    // assert(ptr <= end)
			value.append(ptr, end - ptr);
			size -= end - ptr;
			ptr = end;
			grow_buffer(size);
		    // assert(ptr <= end)
			if (TLGEN2_UNLIKELY(ptr == end)) {
				return set_error_eof();
			}
		}
		value.append(ptr, size);
		ptr += size;
		return true;
	}
	bool fetch_pad(size_t len) {
		uint32_t x = 0;
       	if (TLGEN2_UNLIKELY(!fetch_data(&x, len))) { return false; }
		if (TLGEN2_UNLIKELY(x != 0)) { return set_error_string_padding(); }
		return true;
	}
};

class tl_ostream { // TODO - prohibit copy/move
public:
	virtual ~tl_ostream() = default;
	bool nat_write(uint32_t value) {
		return store_data(&value, 4);
	}
	bool int_write(int32_t value) {
		return store_data(&value, 4);
	}
	bool long_write(int64_t value) {
		return store_data(&value, 8);
	}
	bool float_write(float value) {
		return store_data(&value, 4);
	}
	bool double_write(double value) {
		return store_data(&value, 8);
	}
	bool string_write(const std::string & value) {
		auto len = value.size();
		if (TLGEN2_UNLIKELY(len > TL_MAX_TINY_STRING_LEN)) {
			if (TLGEN2_UNLIKELY(len > TL_BIG_STRING_LEN)) {
				return set_error("TODO - huge string");
			}
			uint32_t p = (len << 8U) | TL_BIG_STRING_MARKER;
			if (TLGEN2_UNLIKELY(!store_data(&p, 4))) { return false; }
			if (TLGEN2_UNLIKELY(!store_data(value.data(), value.size()))) { return false; }
			if (TLGEN2_UNLIKELY(!store_pad((-len) & 3))) { return false; }
			return true;
		}
		auto pad = ((-(len+1)) & 3);
		auto fullLen = 1 + len + pad;
		if (TLGEN2_UNLIKELY(ptr + fullLen > end)) {
			auto p = static_cast<unsigned char>(len);
			if (TLGEN2_UNLIKELY(!store_data(&p,1))) { return false; }
			if (TLGEN2_UNLIKELY(!store_data(value.data(), value.size()))) { return false; }
			if (TLGEN2_UNLIKELY(!store_pad(pad))) { return false; }
			return true;
		}
		// fast path for short strings that fully fit in buffer
		uint32_t x = 0;
		std::memcpy(ptr + fullLen - 4, &x, 4); // padding first
		*ptr = static_cast<char>(len);
		std::memcpy(ptr + 1, value.data(), len);
		ptr += fullLen;
		return true;
	}
    bool set_error(const char * e) { return false; } // TODO - set error field
	bool set_error_eof() { return set_error("EOF"); }
	bool set_error_sequence_length() { return set_error("sequence_length"); }
protected:
	char * ptr{};
	char * end{};
	virtual void grow_buffer(size_t size) = 0; // after call buffer must contain at least single bytes, otherwise error
private:
	bool store_data(const void * vdata, size_t size) {
    	const char * data = reinterpret_cast<const char *>(vdata);
		for (;TLGEN2_UNLIKELY(ptr + size > end);) {
    		// assert(ptr <= end)
			std::memcpy(ptr, data, end - ptr);
			data += end - ptr;
			size -= end - ptr;
			ptr = end;
			grow_buffer(size);
    		// assert(ptr <= end)
			if (TLGEN2_UNLIKELY(ptr == end)) {
				return set_error_eof();
			}
		}
		std::memcpy(ptr, data, size);
		ptr += size;
		return true;
	}
	bool store_pad(size_t size) {
		for (;TLGEN2_UNLIKELY(ptr + size > end);) {
    		// assert(ptr <= end)
			std::memset(ptr, 0, end - ptr);
			size -= end - ptr;
			ptr = end;
			grow_buffer(size);
    		// assert(ptr <= end)
			if (TLGEN2_UNLIKELY(ptr == end)) {
				return set_error_eof();
			}
		}
		if (size != 0) {
			ptr[0] = 0;
			ptr[size-1] = 0;
			ptr[size/2] = 0;
			ptr += size;
		}
		return true;
	}
};

class tl_istream_string : public tl_istream { // TODO - custom copy/move
public:
	explicit tl_istream_string(const std::string & buf) {
		ptr = buf.data();
		end = ptr + buf.size();
	}
protected:
	void grow_buffer(size_t size) override {}
};

class tl_ostream_string : public tl_ostream { // TODO - custom copy/move
public:
	explicit tl_ostream_string() {
		resize(INITIAL_SIZE);
	}
	std::string & get_buffer() {
		resize(ptr - buf.data());
		return buf;
	}
protected:
	void grow_buffer(size_t size) override {
		auto pos = ptr - buf.data();
		resize(buf.size()*3/2 + INITIAL_SIZE + size); // some arbitrary strategy
		ptr += pos;
	}
private:
	enum { INITIAL_SIZE = 1024 };
	std::string buf;
	void resize(size_t size) {
		buf.resize(size);
		ptr = const_cast<char *>(buf.data()); // works for all known implementations
		end = ptr + buf.size();
	}
};

`
