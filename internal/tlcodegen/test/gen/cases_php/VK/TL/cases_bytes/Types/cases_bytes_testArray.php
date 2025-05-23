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
class cases_bytes_testArray implements TL\Readable, TL\Writeable {

  /** @var int */
  public $n = 0;

  /** @var string[] */
  public $arr = [];

  /**
   * @param int $n
   * @param string[] $arr
   */
  public function __construct($n = 0, $arr = []) {
    $this->n = $n;
    $this->arr = $arr;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x3762fb81) {
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
    $this->arr = [];
    for($i12 = 0; $i12 < $this->n; $i12++) {
      /** @var string */
      $obj12 = '';
      [$obj12, $success] = $stream->read_string();
      if (!$success) {
        return false;
      }
      $this->arr[] = $obj12;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0x3762fb81);
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
    for($i12 = 0; $i12 < $this->n; $i12++) {
      $success = $stream->write_string($this->arr[$i12]);
      if (!$success) {
        return false;
      }
    }
    return true;
  }

}
