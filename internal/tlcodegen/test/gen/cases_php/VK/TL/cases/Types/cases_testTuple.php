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
class cases_testTuple implements TL\Readable, TL\Writeable {

  /** @var int[] */
  public $tpl = [];

  /**
   * @param int[] $tpl
   */
  public function __construct($tpl = []) {
    $this->tpl = $tpl;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x4b9caf8f) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    $this->tpl = [];
    for($i9 = 0; $i9 < 4; $i9++) {
      /** @var int */
      $obj9 = 0;
      [$obj9, $success] = $stream->read_int32();
      if (!$success) {
        return false;
      }
      $this->tpl[] = $obj9;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0x4b9caf8f);
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
    for($i9 = 0; $i9 < 4; $i9++) {
      $success = $stream->write_int32($this->tpl[$i9]);
      if (!$success) {
        return false;
      }
    }
    return true;
  }

}
