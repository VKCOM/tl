// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

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
