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
class cases_testOutFieldMask {

  /** Field mask for $f1 field */
  const BIT_F1_0 = (1 << 0);

  /** Field mask for $f2 field */
  const BIT_F2_3 = (1 << 3);

  /** @var int|null */
  public $f1 = null;

  /** @var boolean */
  public $f2 = false;

  /** @var int[] */
  public $f3 = [];

  /**
   * @param int[] $f3
   */
  public function __construct($f3 = []) {
    $this->f3 = $f3;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @param int $f
   * @return bool 
   */
  public function read_boxed($stream, $f) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0xbd6b4b3c) {
      return false;
    }
    return $this->read($stream, $f);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @param int $f
   * @return bool 
   */
  public function read($stream, $f) {
    if (($f & (1 << 0)) != 0) {
      [$this->f1, $success] = $stream->read_uint32();
      if (!$success) {
        return false;
      }
    } else {
      $this->f1 = null;
    }
    if (($f & (1 << 3)) != 0) {
      $this->f2 = true;
    } else {
      $this->f2 = false;
    }
    $this->f3 = [];
    for($i9 = 0; $i9 < $f; $i9++) {
      $array_int___element = 0;
      [$array_int___element, $success] = $stream->read_int32();
      if (!$success) {
        return false;
      }
      $this->f3[] = $array_int___element;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @param int $f
   * @return bool 
   */
  public function write_boxed($stream, $f) {
    $success = $stream->write_uint32(0xbd6b4b3c);
    if (!$success) {
      return false;
    }
    return $this->write($stream, $f);
  }

  /**
   * @param TL\tl_output_stream $stream
   * @param int $f
   * @return bool 
   */
  public function write($stream, $f) {
    if (($f & (1 << 0)) != 0) {
      $success = $stream->write_uint32($this->f1);
      if (!$success) {
        return false;
      }
    }
    for($i9 = 0; $i9 < $f; $i9++) {
      $success = $stream->write_int32($this->f3[$i9]);
      if (!$success) {
        return false;
      }
    }
    return true;
  }

  /**
   * @return int
   */
  public function calculateF() {
    $mask = 0;

    if (!is_null($this->f1)) {
      $mask |= self::BIT_F1_0;
    }

    if ($this->f2) {
      $mask |= self::BIT_F2_3;
    }

    return $mask;
  }

}
