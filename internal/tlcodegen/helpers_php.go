// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

var PHPNamesToIgnoreForLinterCheck = []string{
	// primitives from https://github.com/VKCOM/kphp/blob/ff72b37bda68cf282c37a92d23bc763de6f47795/common/tl2php/combinator-to-php.cpp#L84
	"String",
	"Int",
	"#",
	"Long",
	"Double",
	"Float",
	"Bool",
	//"False", no such type?
	"True",
	// brackets from https://github.com/VKCOM/kphp/blob/ff72b37bda68cf282c37a92d23bc763de6f47795/common/tl2php/combinator-to-php.cpp#L104
	"Vector",
	"Tuple",
	"Dictionary",
	"IntKeyDictionary",
	"LongKeyDictionary",
	// and Maybe
	"Maybe",
}

const BasicTlPathPHP = "tl_streams.php"
const BasicTLCodePHP = `<?php

namespace VK\TL;

class tl_constants {
    const tinyStringLen    = 253;
	const bigStringMarker  = 0xfe;
	const hugeStringMarker = 0xff;
	const bigStringLen     = (1 << 24) - 1;
	const hugeStringLen    = (1 << 56) - 1;
}

class tl_input_stream {
    /** @var string */
    private $data = "";
    /** @var int */
    private $offset = 0;

    public function __construct(string $data) {
        $this->data = $data;
    }

    public function get_offset(): int {
        return $this->offset;
    }

    public function remaining_size(): int {
        return strlen($this->data) - $this->offset;
    }
    
    /** @return tuple(int, bool) */
    public function read_int32() {
        // TODO: 'l' - signed 32 bit SYSTEM DEPENDENT ENDIAN
        $data = unpack('l', $this->data, $this->offset);
        if (!$data) {
            return [0, false];
        } else {
            $this->offset += 4;
            return [$data[1], true];
        }
    }

    /** @return tuple(int, bool) */
    public function read_uint32() {
        $data = unpack('V', $this->data, $this->offset);
        if (!$data) {
            return [0, false];
        } else {
            $this->offset += 4;
            return [$data[1], true];
        }
    }

    /** @return tuple(int, bool) */
    public function read_int64() {
        $data = unpack('q', $this->data, $this->offset);
        if (!$data) {
            return [0, false];
        } else {
            $this->offset += 8;
            return [$data[1], true];
        }
    }

    /** @return tuple(bool, bool) */
    public function read_bool(int $false_tag, $true_tag) {
        [$tag, $success] = $this->read_uint32();
        if (!$success) {
            return [false, false];
        } else if ($tag == $false_tag) {
            return [false, true];
        } else if ($tag == $true_tag) {
            return [true, true];
        }
        return [false, false];
    }

    /** @return tuple(float, bool) */
    public function read_float() {
        $data = unpack('f', $this->data, $this->offset);
        if (!$data) {
            return [0, false];
        } else {
            $this->offset += 4;
            return [$data[1], true];
        }
    }

    /** @return tuple(float, bool) */
    public function read_double() {
        $data = unpack('d', $this->data, $this->offset);
        if (!$data) {
            return [0, false];
        } else {
            $this->offset += 8;
            return [$data[1], true];
        }
    }

    /** @return tuple(string, bool) */
    public function read_string() {
        $size = $this->remaining_size(); 
        if ($size == 0) {
            return ["", false];
        }
        $first_byte = ord($this->data[$this->offset]);
        $l = 0;
        $p = 0;
        if ($first_byte < tl_constants::tinyStringLen) {
            $l = $first_byte;
            $this->offset += 1;
            $p = $l + 1;
        } elseif ($first_byte == tl_constants::bigStringMarker) {
            if ($size < 4) {
                return ["", false];
            }
            $l = (ord($this->data[$this->offset + 3]) << 16) + (ord($this->data[$this->offset + 2]) << 8) + (ord($this->data[$this->offset + 1]) << 0);
            $this->offset += 4;
            $p = $l;
            if ($l <= tl_constants::tinyStringLen) {
                return ["", false];
            }
        } else {
            if ($size < 8) {
                return ["", false];
            }
            $l64 = (ord($this->data[$this->offset + 7]) << 48) + (ord($this->data[$this->offset + 6]) << 40) + (ord($this->data[$this->offset + 5]) << 32) + (ord($this->data[$this->offset + 4]) << 24) + (ord($this->data[$this->offset + 3]) << 16) + (ord($this->data[$this->offset + 2]) << 8) + (ord($this->data[$this->offset + 1]) << 0);
            // TODO: check l64 > max int
            $l = $l64;
            $this->offset += 8;
            $p = $l;
            if ($l <= tl_constants::bigStringLen) {
                return ["", false];
            }
        }
        $start = $this->offset;
        if ($l > 0) {
            if ($this->remaining_size() < $l) {
                return ["", false];
            }
        }
        $padding = $this->paddingLen($p);
        for ($i = 0; $i < $padding; $i++) {
            if (ord($this->data[$this->offset + $l + $i]) != 0) {
                return ["", false];
            }
        }
        $this->offset += $l + $padding;
        return [substr($this->data, $start, $l), true];
    }
    
    function paddingLen(int $l): int {
        return (4 - ($l % 4)) % 4;
    }   
}

class tl_output_stream {
    /** @var string */
    private $data = "";

    public function __construct(string $data) {
        $this->data = $data;
    }

    public function get_data(): string {
        return $this->data;
    }
	
	/** @return bool */
    public function write_uint32(int $value) {
        $this->data .= pack('V', $value);
 		return true;       
    }

    /** @return bool */
    public function write_int32(int $value) {
        $this->data .= pack('l', $value);
 		return true;
    }

    /** @return bool */
    public function write_int64(int $value) {
        $this->data .= pack('q', $value);
 		return true;
    }

    /** @return bool */
    public function write_bool(bool $value, int $false_tag, $true_tag) {
        if ($value) {
            $this->data .= pack('V', $true_tag);
        } else {
            $this->data .= pack('V', $false_tag);
        }
 		return true;
    }

    /** @return bool */
    public function write_float(float $value) {
        $this->data .= pack('f', $value);
 		return true; 
    }

    /** @return bool */
    public function write_double(float $value) {
        $this->data .= pack('d', $value);
 		return true; 
    }

    /** @return bool */
    public function write_string(string $value) {
        $l = strlen($value);
        $p = 0;
        if ($l <= tl_constants::tinyStringLen) {
            $this->data .= chr($l);
            $p = $l + 1;
        } else if ($l <= tl_constants::bigStringLen) {
            $this->data .= chr(tl_constants::bigStringMarker);
            $this->data .= chr(($l & 255));
            $this->data .= chr((($l >> 8) & 255));
            $this->data .= chr((($l >> 16) & 255));
            $p = $l;
        } else {
            if ($l > tl_constants::hugeStringLen) {
                $l = tl_constants::hugeStringLen;
                $value = substr($value, 0, $l);
            }
            $this->data .= chr(tl_constants::hugeStringMarker);
            $this->data .= chr(($l & 255));
            $this->data .= chr((($l >> 8) & 255));
            $this->data .= chr((($l >> 16) & 255));
            $this->data .= chr((($l >> 24) & 255));
            $this->data .= chr((($l >> 32) & 255));
            $this->data .= chr((($l >> 40) & 255));
            $this->data .= chr((($l >> 48) & 255));
            $p = $l;
        }
        $this->data .= $value;
        if ($p % 4 == 1) {
            $this->data .= chr(0);
            $this->data .= chr(0);
            $this->data .= chr(0);
        } else if ($p % 4 == 2) {
            $this->data .= chr(0);
            $this->data .= chr(0);
        } else if ($p % 4 == 3) {
            $this->data .= chr(0);
        }
 		return true; 
    }
}
?>
`

const RpcFunctionPHP = `<?php

%s#ifndef KPHP

namespace VK\TL;

/**
 * @kphp-tl-class
 */
interface RpcFunction {

  /**
   * @kphp-inline
   *
   * @return string
   */
  public function getTLFunctionName();
}

/**
 * @kphp-tl-class
 */
interface RpcFunctionReturnResult {

}

#endif
`

const RpcFunctionWithFetchersPHP = `<?php

%s#ifndef KPHP

namespace VK\TL;

/**
 * @kphp-tl-class
 */
interface RpcFunction {

  /**
   * @kphp-inline
   *
   * @return int
   */
  public function getTLFunctionMagic();
  /**
   * @kphp-inline
   *
   * @return string
   */
  public function getTLFunctionName();
  /**
   * @kphp-inline
   *
   * @return TL\RpcFunctionFetcher
   */
  public function typedStore();
  /**
   * @kphp-inline
   *
   * @return TL\RpcFunctionFetcher
   */
  public function typedFetch();
}

/**
 * @kphp-tl-class
 */
interface RpcFunctionReturnResult {

}

#endif
`
const RpcResponsePHP = `<?php

%s#ifndef KPHP

namespace VK\TL;

use VK\TL;

/**
 * @kphp-tl-class
 */
interface RpcResponse {

  /** Allows kphp implicitly load all available constructors */
  const CONSTRUCTORS = [
    TL\_common\Types\rpcResponseError::class,
    TL\_common\Types\rpcResponseHeader::class,
    TL\_common\Types\rpcResponseOk::class
  ];

  /**
   * @return TL\RpcFunctionReturnResult
   */
  public function getResult();

  /**
   * @return TL\_common\Types\rpcResponseHeader
   */
  public function getHeader();

  /**
   * @return bool
   */
  public function isError();

  /**
   * @return TL\_common\Types\rpcResponseError
   */
  public function getError();

}

#endif
`

const RpcFunctionFetchersPHP = `<?php

%s#ifndef KPHP

namespace VK\TL;

interface RpcFunctionFetcher {

  /**
   * @param \VK\TL\RpcFunctionReturnResult $result
   */
  public function typedStore($result);
  /**
   * @return \VK\TL\RpcFunctionReturnResult
   */
  public function typedFetch();
}

#endif
`

const TLInterfacesPathPHP = "tl_interfaces.php"
const TLInterfacesCodePHP = `<?php

namespace VK\TL;

use VK\TL;

interface Readable {
  /**
   * @param TL\tl_input_stream $stream
   * @return bool
   */
  public function read($stream);

  /**
   * @param TL\tl_input_stream $stream
   * @return bool
   */
  public function read_boxed($stream);
}

interface Writeable {
  /**
   * @param TL\tl_output_stream $stream
   * @return bool
   */
  public function write($stream);

  /**
   * @param TL\tl_output_stream $stream
   * @return bool
   */
  public function write_boxed($stream);
}

interface TL_Object extends Readable, Writeable {}
`

const TL2SupportPHP = `class tl2_support {
  const TinyStringLen = 253;
  const BigStringMarker = 254;
  const HugeStringMarker = 255;
  const BigStringLen = 254 + (1 << 16) - 1;
  const HugeStringLen = (1 << 64) - 1;
  
  /**
   * @return int
   */
  public static function fetch_size() {
    $b0 = fetch_byte();
    if ($b0 <= self::TinyStringLen) {
      return $b0;
    } else if ($b0 == self::BigStringMarker) {
      $b1 = fetch_byte() & 0xFF;
      $b2 = fetch_byte() & 0xFF;
      return self::BigStringMarker + ($b2 << 8) + $b1;
    } else {
      return fetch_long();
    }
    return 0;
  }

  /**
   * @param int $size
   */
  public static function store_size($size) {
    if ($size <= self::TinyStringLen) {
      store_byte($size & 0xFF);
    } else if ($size <= self::BigStringLen) {
      store_byte(self::BigStringMarker);
      $size -= self::BigStringMarker;
      store_byte($size & 0xFF);
      store_byte(($size >> 8) & 0xFF);
    } else {
      store_byte(self::HugeStringMarker);
      store_long($size);
    }
  }

  /**
   * @param int $size
   * @return int
   */
  public static function count_used_bytes($size) {
    if ($size <= self::TinyStringLen) {
      return 1;
    } else if ($size <= self::BigStringLen) {
      return 3;
    } else {
      return 9;
    }
  }

  /**
   * @param int $count
   * @return int
   */
  public static function skip_bytes($count) {
    if ($count < 0) {
      throw new \Exception("can't skip negative number of bytes");
    }
    for ($i = 0; $i < $count; $i++) {
      fetch_byte();
    }
    return $count;
  }

  /**
   * @return boolean
   */
  public static function fetch_legacy_bool_tl2() {
    $b = fetch_byte();
    return $b != 0;
  }

  /**
   * @param boolean|null $value
   */
  public static function store_legacy_bool_tl2($value) {
    if ($value === null) {
      return;
    }
    if ($value == 0) {
      store_byte(0);
    } else {
      store_byte(1);
    }
  }
}
`

const TLSwitcherPHP = `<?php

%[1]snamespace VK\TL;

use VK\TL;

class tl_switcher {
  /** @var int[] */
  public static $tl_namespaces_info = [];

  /**
   * @param string $tl_namespace
   * @return int
   */
  public static function tl_get_namespace_methods_mode($tl_namespace) {
    if (array_key_exists($tl_namespace, self::$tl_namespaces_info)) {
      return self::$tl_namespaces_info[$tl_namespace];
    }
    return 0;
  }
}
`

const TL2ContextPHP = `<?php

%[1]snamespace VK\TL;

class tl2_context {
  /** @var int[] */
  private $values = [];

  /** @var int */
  private $current_index = 0;

  /** @var int */
  private $current_size = 0;

  /**
   * @kphp-inline
   */
  public function __construct() {
  }

  /**
   * @param int $value
   * 
   * @return int
   */
  public function push_back($value) {
    $index = $this->current_size;
    if ($index == count($this->values)) {
      $this->values[] = $value;
    } else {
      $this->values[$index] = $value;
    }
    $this->current_size += 1;
    return $index;
  }

  /**
   * @return int
   */
  public function pop_front() {
    if ($this->current_index >= $this->current_size) {
      throw new \Exception("context can't pop front value");
    }
    $value = $this->values[$this->current_index];
    $this->current_index += 1;
    return $value;
  }

  /**
   * @return int
   */
  public function get_current_size() {
    return $this->current_size;
  }

  /**
   * @param int $index
   * @param int $value
   */
  public function set_value($index, $value) {
    if ($index >= $this->current_size || $index < 0) {
      throw new \Exception("invalid index to set for context");
    }
    $this->values[$index] = $value;
  }

  /**
   * @param int $index
   * 
   * @return int
   */
  public function get_value($index) {
    if ($index >= $this->current_size || $index < 0) {
      throw new \Exception("invalid index to get for context");
    }
    return $this->values[$index];
  }

  /**
   * @param int $size
   */
  public function cut_tail($size) {
    if ($size > $this->current_size || $size < 0) {
      throw new \Exception("invalid index to cut tail");
    }
    $this->current_size = $size;
  }

  /**
   * @return string
   */
  public function show_state() {
    $s = "[";
    for ($i = $this->current_index; $i < $this->current_size; $i++) {
      if ($i != $this->current_index) {
        $s .= ", ";
      }
      $s .= $this->values[$i] . "";
    }
    $s .= "]";
    return $s;
  }
}`
