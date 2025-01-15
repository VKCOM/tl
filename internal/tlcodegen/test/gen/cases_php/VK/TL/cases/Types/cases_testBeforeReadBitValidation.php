<?php

/**
 * AUTOGENERATED, DO NOT EDIT! If you want to modify it, check tl schema.
 *
 * This autogenerated code represents tl class for typed RPC API.
 */

namespace VK\TL\cases\Types;

use VK\TL;

/**
 * @kphp-tl-class
 */
class cases_testBeforeReadBitValidation implements TL\Readable, TL\Writeable {

  /** Field mask for $a field */
  const BIT_A_0 = (1 << 0);

  /** Field mask for $b field */
  const BIT_B_1 = (1 << 1);

  /** @var int */
  public $n = 0;

  /** @var int[]|null */
  public $a = null;

  /** @var int[]|null */
  public $b = null;

  /**
   * @param int $n
   */
  public function __construct($n = 0) {
    $this->n = $n;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x9b2396db) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    [$this->n, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    if (($this->n & (1 << 0)) != 0) {
      $this->a = [];
      for($i9 = 0; $i9 < $this->n; $i9++) {
        $array_int___element = 0;
        [$array_int___element, $success] = $stream->read_int32();
        if (!$success) {
          return false;
        }
        $this->a[] = $array_int___element;
      }
    } else {
      $this->a = null;
    }
    if (($this->n & (1 << 1)) != 0) {
      $this->b = [];
      for($i9 = 0; $i9 < $this->n; $i9++) {
        $array_int___element = 0;
        [$array_int___element, $success] = $stream->read_int32();
        if (!$success) {
          return false;
        }
        $this->b[] = $array_int___element;
      }
    } else {
      $this->b = null;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0x9b2396db);
    if (!$success) {
      return false;
    }
    return $this->write($stream);
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write($stream) {
    $success = $stream->write_uint32($this->n);
    if (!$success) {
      return false;
    }
    if (($this->n & (1 << 0)) != 0) {
      for($i9 = 0; $i9 < $this->n; $i9++) {
        $success = $stream->write_int32($this->a[$i9]);
        if (!$success) {
          return false;
        }
      }
    }
    if (($this->n & (1 << 1)) != 0) {
      for($i9 = 0; $i9 < $this->n; $i9++) {
        $success = $stream->write_int32($this->b[$i9]);
        if (!$success) {
          return false;
        }
      }
    }
    return true;
  }

  /**
   * @return int
   */
  public function calculateN() {
    $mask = 0;

    if ($this->a !== null) {
      $mask |= self::BIT_A_0;
    }

    if ($this->b !== null) {
      $mask |= self::BIT_B_1;
    }

    return $mask;
  }

}
