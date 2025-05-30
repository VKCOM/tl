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
class cases_replace7plus implements TL\Readable, TL\Writeable {

  /** Field mask for $a field */
  const BIT_A_0 = (1 << 0);

  /** @var int */
  public $n = 0;

  /** @var int */
  public $m = 0;

  /** @var int[][]|null */
  public $a = null;

  /**
   * @param int $n
   * @param int $m
   */
  public function __construct($n = 0, $m = 0) {
    $this->n = $n;
    $this->m = $m;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x197858f5) {
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
    [$this->m, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    if (($this->n & (1 << 0)) != 0) {
      $this->a = [];
      for($i15 = 0; $i15 < $this->n; $i15++) {
        /** @var int[] */
        $obj15 = [];
        for($i9 = 0; $i9 < $this->m; $i9++) {
          /** @var int */
          $obj9 = 0;
          [$obj9, $success] = $stream->read_int32();
          if (!$success) {
            return false;
          }
          $obj15[] = $obj9;
        }
        $this->a[] = $obj15;
      }
    } else {
      $this->a = null;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0x197858f5);
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
    $success = $stream->write_uint32($this->m);
    if (!$success) {
      return false;
    }
    if (($this->n & (1 << 0)) != 0) {
      for($i15 = 0; $i15 < $this->n; $i15++) {
        for($i9 = 0; $i9 < $this->m; $i9++) {
          $success = $stream->write_int32($this->a[$i15][$i9]);
          if (!$success) {
            return false;
          }
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

    return $mask;
  }

}
