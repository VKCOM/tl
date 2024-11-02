package tlcodegen

const basictlPHP = `
<?php
include "run.php";

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
        if ($data == false) {
            return [0, false];
        } else {
            $this->offset += 4;
            return [$data[1], true];
        }
    }

    /** @return tuple(int, bool) */
    public function read_uint32() {
        $data = unpack('V', $this->data, $this->offset);
        if ($data == false) {
            return [0, false];
        } else {
            $this->offset += 4;
            return [$data[1], true];
        }
    }

    /** @return tuple(bool, bool) */
    public function read_bool(int $false_tag, $true_tag) {
        $tag = $this->read_uint32();
        if ($tag == false) {
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
        if ($data == false) {
            return [0, false];
        } else {
            $this->offset += 4;
            return [$data[1], true];
        }
    }

    /** @return tuple(double, bool) */
    public function read_double() {
        $data = unpack('d', $this->data, $this->offset);
        if ($data == false) {
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
            // TODO: check l64 > maxint
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
?>
`
