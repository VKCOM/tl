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
class cases_testAllPossibleFieldConfigsContainer implements TL\Readable, TL\Writeable {

  /** @var int */
  public $outer = 0;

  /** @var TL\cases\Types\cases_testAllPossibleFieldConfigs */
  public $value = null;

  /**
   * @param int $outer
   * @param TL\cases\Types\cases_testAllPossibleFieldConfigs $value
   */
  public function __construct($outer = 0, $value = null) {
    $this->outer = $outer;
    $this->value = $value;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0xe3fae936) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    [$this->outer, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    if (is_null($this->value)) {
      $this->value = new TL\cases\Types\cases_testAllPossibleFieldConfigs();
    }
    $success = $this->value->read($stream, $this->outer);
    if (!$success) {
      return false;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0xe3fae936);
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
    $success = $stream->write_uint32($this->outer);
    if (!$success) {
      return false;
    }
    if (is_null($this->value)) {
      $this->value = new TL\cases\Types\cases_testAllPossibleFieldConfigs();
    }
    $success = $this->value->write($stream, $this->outer);
    if (!$success) {
      return false;
    }
    return true;
  }

}
