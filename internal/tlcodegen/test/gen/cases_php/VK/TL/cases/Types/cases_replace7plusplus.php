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
class cases_replace7plusplus implements TL\Readable, TL\Writeable {

  /** Field mask for $A field */
  const BIT_A_0 = (1 << 0);

  /** @var int */
  public $N = 0;

  /** @var int */
  public $M = 0;

  /** @var int[][]|null */
  public $A = null;

  /**
   * @param int $N
   * @param int $M
   */
  public function __construct($N = 0, $M = 0) {
    $this->N = $N;
    $this->M = $M;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0xabc39b68) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    [$this->N, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    [$this->M, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    if (($this->N & (1 << 0)) != 0) {
      $this->A = [];
      for($i15 = 0; $i15 < $this->N; $i15++) {
        $array_array_int___element = [];
        for($i9 = 0; $i9 < $this->M; $i9++) {
          $array_int___element = 0;
          [$array_int___element, $success] = $stream->read_int32();
          if (!$success) {
            return false;
          }
          $array_array_int___element[] = $array_int___element;
        }
        $this->A[] = $array_array_int___element;
      }
    } else {
      $this->A = null;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0xabc39b68);
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
    $success = $stream->write_uint32($this->N);
    if (!$success) {
      return false;
    }
    $success = $stream->write_uint32($this->M);
    if (!$success) {
      return false;
    }
    if (($this->N & (1 << 0)) != 0) {
      for($i15 = 0; $i15 < $this->N; $i15++) {
        for($i9 = 0; $i9 < $this->M; $i9++) {
          $success = $stream->write_int32($this->A[$i15][$i9]);
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

    if (!is_null($this->A)) {
      $mask |= self::BIT_A_0;
    }

    return $mask;
  }

}
