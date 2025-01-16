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
class cases_myCycle1 implements TL\Readable, TL\Writeable {

  /** Field mask for $a field */
  const BIT_A_0 = (1 << 0);

  /** @var int */
  public $fields_mask = 0;

  /** @var TL\cases\Types\cases_myCycle2|null */
  public $a = null;

  /**
   * @param int $fields_mask
   */
  public function __construct($fields_mask = 0) {
    $this->fields_mask = $fields_mask;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0xd3ca919d) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    [$this->fields_mask, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    if (($this->fields_mask & (1 << 0)) != 0) {
      if (is_null($this->a)) {
        $this->a = new TL\cases\Types\cases_myCycle2();
      }
      $success = $this->a->read($stream);
      if (!$success) {
        return false;
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
    $success = $stream->write_uint32(0xd3ca919d);
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
    $success = $stream->write_uint32($this->fields_mask);
    if (!$success) {
      return false;
    }
    if (($this->fields_mask & (1 << 0)) != 0) {
      if (is_null($this->a)) {
        $this->a = new TL\cases\Types\cases_myCycle2();
      }
      $success = $this->a->write($stream);
      if (!$success) {
        return false;
      }
    }
    return true;
  }

  /**
   * @return int
   */
  public function calculateFieldsMask() {
    $mask = 0;

    if ($this->a) {
      $mask |= self::BIT_A_0;
    }

    return $mask;
  }

}
