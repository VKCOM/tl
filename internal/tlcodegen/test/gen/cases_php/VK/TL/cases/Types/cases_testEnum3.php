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
class cases_testEnum3 implements TL\cases\Types\cases_TestEnum, TL\Readable, TL\Writeable {

  /**
   * @kphp-inline
   */
  public function __construct() {
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x69b83e2f) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0x69b83e2f);
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
    return true;
  }

}
