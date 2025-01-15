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
class cases_testArray implements TL\Readable, TL\Writeable {

  /** @var int */
  public $n = 0;

  /** @var int[] */
  public $arr = [];

  /**
   * @param int $n
   * @param int[] $arr
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
    if (!$success || $magic != 0xa888030d) {
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
    for($i9 = 0; $i9 < $this->n; $i9++) {
      $array_int___element = 0;
      [$array_int___element, $success] = $stream->read_int32();
      if (!$success) {
        return false;
      }
      $this->arr[] = $array_int___element;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0xa888030d);
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
    for($i9 = 0; $i9 < $this->n; $i9++) {
      $success = $stream->write_int32($this->arr[$i9]);
      if (!$success) {
        return false;
      }
    }
    return true;
  }

}
