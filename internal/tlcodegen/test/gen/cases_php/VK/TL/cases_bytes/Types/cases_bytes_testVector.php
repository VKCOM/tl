<?php

/**
 * AUTOGENERATED, DO NOT EDIT! If you want to modify it, check tl schema.
 *
 * This autogenerated code represents tl class for typed RPC API.
 */

namespace VK\TL\cases_bytes\Types;

use VK\TL;

/**
 * @kphp-tl-class
 */
class cases_bytes_testVector implements TL\Readable, TL\Writeable {

  /** @var string[] */
  public $arr = [];

  /**
   * @param string[] $arr
   */
  public function __construct($arr = []) {
    $this->arr = $arr;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x3647c8ae) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    [$vector_size, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    $this->arr = [];
    for($i12 = 0; $i12 < $vector_size; $i12++) {
      $array_string___element = '';
      [$array_string___element, $success] = $stream->read_string();
      if (!$success) {
        return false;
      }
      $this->arr[] = $array_string___element;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0x3647c8ae);
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
    $success = $stream->write_uint32(count($this->arr));
    if (!$success) {
      return false;
    }
    for($i12 = 0; $i12 < count($this->arr); $i12++) {
      $success = $stream->write_string($this->arr[$i12]);
      if (!$success) {
        return false;
      }
    }
    return true;
  }

}
